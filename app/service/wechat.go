package service

import (
	"encoding/base64"
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"github.com/leapig/tpp"
	"github.com/leapig/tpp/mp"
	"github.com/leapig/tpp/oa"
	"github.com/leapig/tpp/util"
	"github.com/skip2/go-qrcode"
	"os"
	"strconv"
	"time"
)

const (
	loginKeyName = "loginKey"
)

type Wechat interface {
	Oa() oa.App
	Mp() mp.App
	OaSaveSubscribe(string, string, int64)
	OaScan(string, string, string) error
	GetWxOaQrCode() (*pb.LoginQrCodeResp, error)
	PlatformLogin(string, string) error
	PlatformLoginLoop(key string) (string, error)
	AppletLogin(code string) (string, error)
	AppletRegister(loginCode, phoneCode, name string) (string, error)
	GetWxMpPhoneNumber(code string) (string, error)
	CreateEnterprise(*entity.User, string, *entity.Enterprise) error
	CreateEnterpriseForRelationProject(user *entity.User, code string, enterprise *entity.Enterprise) (*entity.Enterprise, error)
}

type wechat struct {
	dao dao.Dao
}

func NewWechat(dao dao.Dao) Wechat {
	return &wechat{dao: dao}
}

// Mp 小程序实例
func (o *wechat) Mp() mp.App {
	return tpp.NewTpp().MP(mp.Config{
		Key:     os.Getenv("WX_MP_KEY"),
		AppId:   os.Getenv("WX_MP_APP_ID"),
		Secret:  os.Getenv("WX_MP_APP_SECRET"),
		Version: os.Getenv("WX_MP_VERSION"),
	})
}

// Oa 公众号实例
func (o *wechat) Oa() oa.App {
	return tpp.NewTpp().OA(oa.Config{
		Key:    os.Getenv("WX_MP_KEY"),
		AppId:  os.Getenv("WX_OA_APP_ID"),
		Secret: os.Getenv("WX_OA_APP_SECRET"),
		Token:  os.Getenv("WX_OA_TOKEN"),
		AesKey: os.Getenv("WX_OA_AES_KEY"),
	})
}

// OaSaveSubscribe 方法用于保存或更新用户在微信公众号的订阅状态。
// 参数:
//   - openid: 用户的唯一标识。
//   - clientId: 客户端ID。
//   - subscribe: 用户的订阅状态，1 表示已订阅，0 表示未订阅。
func (o *wechat) OaSaveSubscribe(openid, clientId string, subscribe int64) {
	// 查找数据库中是否已经存在该用户的订阅记录
	old, _ := o.dao.UserClient().Find(&entity.UserClient{
		ClientType: 1,
		OpenId:     openid,
	})

	// 将订阅状态转换为 int8 类型
	subscribeInt8 := int8(subscribe)

	// 通过微信公众号获取用户信息
	user := o.Oa().UserInfo(openid)

	// 创建一个新的 UserClient 对象，用于保存或更新用户的订阅状态
	uc := entity.UserClient{
		ClientId:    clientId,
		ClientType:  1,
		OpenId:      openid,
		WxSubscribe: &subscribeInt8,
	}

	// 如果用户信息中包含 unionid，则将其保存到 UserClient 对象中
	if user["unionid"] != nil && user["unionid"] != "" {
		uc.WxUnionid = user["unionid"].(string)
	}

	// 如果 old 对象的 ClientId、WxSubscribe 和 WxUnionid 与 uc 对象相同，则表示没有任何更改，直接返回
	if old.ClientId == uc.ClientId && *old.WxSubscribe == *uc.WxSubscribe && old.WxUnionid == uc.WxUnionid {
		// 无任何更改
		return
	}

	// 保存或更新用户的订阅状态到数据库
	o.dao.UserClient().SaveUserClientSubscribe(&uc)
}

// GetWxOaQrCode 获取公众号登录二维码
func (o *wechat) GetWxOaQrCode() (*pb.LoginQrCodeResp, error) {
	// 生成一个 32 位的随机字符串作为场景值
	scene := util.GetRandString(32)
	// 调用微信公众号的二维码创建接口，创建一个临时的二维码
	res := o.Oa().QrcodeCreate(scene, false)
	// 根据返回的二维码 URL 生成二维码图片
	q, _ := qrcode.New(res["url"].(string), qrcode.Medium)
	// 将二维码图片编码为 PNG 格式
	qrCode, _ := q.PNG(600)
	// 将场景值作为键，将状态设置为 "unUse"，并设置过期时间为 60 秒，存储到 Redis 中
	_ = o.dao.Redis().SetEx(loginKeyName+"_"+scene, 60, "unUse")
	// 返回包含 Base64 编码的二维码图片和场景值的响应
	return &pb.LoginQrCodeResp{
		QrCode: base64.StdEncoding.EncodeToString(qrCode),
		Key:    scene,
	}, nil
}

// OaScan 公众号扫码登录
func (o *wechat) OaScan(openid, clientId, sense string) error {
	o.OaSaveSubscribe(openid, clientId, 1)
	if loginKey, _ := o.dao.Redis().GetString(loginKeyName + "_" + sense); loginKey == "unUse" {
		_ = o.dao.Redis().SetEx(loginKeyName+"_"+sense, 60, "use")
		if userClientOaMessage, err := o.dao.UserClient().Find(&entity.UserClient{
			OpenId:     openid,
			ClientType: 1,
		}); err == nil && userClientOaMessage.WxUnionid != "" {
			if userClientMpMessage, err := o.dao.UserClient().Find(&entity.UserClient{
				WxUnionid:  userClientOaMessage.WxUnionid,
				ClientType: 2,
			}); err == nil && userClientMpMessage.WxUnionid != "" {
				_ = o.dao.Redis().SetEx(loginKeyName+"_"+sense, 60, strconv.FormatInt(userClientMpMessage.UserPk, 10))
				// 推送模板消息
				go func() {
					_ = o.Oa().MessageTemplateSend(oa.Message{
						Touser:     openid,
						TemplateId: os.Getenv("WX_OA_SCAN_TEMPLATE_ID"),
						Miniprogram: oa.Miniprogram{
							Appid:    o.Mp().Id(),
							Pagepath: "pages/index/index",
						},
						Data: map[string]interface{}{
							"time6":  map[string]interface{}{"value": utils.DateTimeToString(time.Now())},
							"thing9": map[string]interface{}{"value": "登录成功，点击详情进入程序"}},
					})
				}()
				return nil
			}
		}
		// 推送模板消息
		go func() {
			_ = o.Oa().MessageTemplateSend(oa.Message{
				Touser:     openid,
				TemplateId: os.Getenv("WX_OA_SCAN_TEMPLATE_ID"),
				Miniprogram: oa.Miniprogram{
					Appid:    o.Mp().Id(),
					Pagepath: "pages/index/index",
				},
				Data: map[string]interface{}{
					"time6":  map[string]interface{}{"value": utils.DateTimeToString(time.Now())},
					"thing9": map[string]interface{}{"value": "登录失败，点击小程序立即注册"}},
			})
		}()
		return nil
	} else if loginKey == "use" {
		return errors.New("该码已被使用")
	} else {
		return errors.New("请勿重复扫码")
	}
}

// PlatformLogin 后台扫码登录设置
func (o *wechat) PlatformLogin(code, key string) error {
	if loginKey, _ := o.dao.Redis().GetString(loginKeyName + "_" + key); loginKey == "unUse" {
		_ = o.dao.Redis().SetEx(loginKeyName+"_"+key, 60, "use")
		res := o.Mp().JsCode2Session(code)
		if res == nil {
			_ = o.dao.Redis().SetEx(loginKeyName+"_"+key, 30*60, "unUse")
			return errors.New("获取登录信息失败")
		}
		if userClientMessage, err := o.dao.UserClient().Find(&entity.UserClient{
			OpenId:     res["openid"].(string),
			ClientType: 2,
		}); err == nil && userClientMessage.UserPk != 0 {
			_ = o.dao.Redis().SetEx(loginKeyName+"_"+key, 60, strconv.FormatInt(userClientMessage.UserPk, 10))
			return nil
		}
		return errors.New("该账户未注册")
	} else if loginKey == "use" {
		return errors.New("该二维码已使用")
	}
	return nil
}

// PlatformLoginLoop 后台扫码登录状态
func (o *wechat) PlatformLoginLoop(key string) (string, error) {
	if loginKey, err := o.dao.Redis().GetString(loginKeyName + "_" + key); err != nil || loginKey == "" || loginKey == "unUse" || loginKey == "use" {
		return "", nil
	} else {
		_, _ = o.dao.Redis().Del(loginKeyName + "_" + key)
		return loginKey, nil
	}
}

func (o *wechat) CreateEnterprise(user *entity.User, code string, enterprise *entity.Enterprise) error {
	// 查询关联关系
	loginRes := o.Mp().JsCode2Session(code)
	if loginRes == nil || loginRes["openid"] == nil || loginRes["unionid"] == nil {
		return errors.New("获取微信凭证失败")
	}
	// 获取用户
	clientRes, _ := o.dao.UserClient().Find(&entity.UserClient{
		ClientType: 2,
		OpenId:     loginRes["openid"].(string),
	})
	if clientRes.UserPk == 0 {
		// 获取用户
		user, err := o.dao.User().Find(user)
		if user.Pk == 0 {
			user.Pk = helper.GetRid(helper.User)
			user, err = o.dao.User().Create(user)
			if err != nil {
				return err
			}
		}
		// 创建关系
		clientRes = &entity.UserClient{
			UserPk:     user.Pk,
			ClientType: 2,
			OpenId:     loginRes["openid"].(string),
			WxUnionid:  loginRes["unionid"].(string),
		}
		if err := o.dao.UserClient().CreateUserClient(clientRes); err != nil {
			return err
		}
	}
	// 创建企业
	enterprise.Pk = helper.GetRid(helper.Enterprise)
	enterprise.CorporatePk = user.Pk
	if en, err := NewEnterprise(o.dao).Create(enterprise); err != nil {
		return err
	} else {
		// 授予权限
		return NewUserPermission(o.dao).CreateForCorporate(&entity.UserPermission{
			EnterprisePk: en.Pk,
			UserPk:       user.Pk,
		})
	}
}

func (o *wechat) GetWxMpPhoneNumber(code string) (string, error) {
	res := o.Mp().PostWxaBusinessGetUserPhoneNumber(code)
	if res == nil {
		return "", errors.New("获取手机号失败")
	}
	if res["phoneNumber"] == "" {
		return "", errors.New("获取手机号失败：为空")
	}
	return res["phoneNumber"].(string), nil
}

// AppletLogin 小程序
func (o *wechat) AppletLogin(code string) (string, error) {
	res := o.Mp().JsCode2Session(code)
	if res != nil {
		if userClientMessage, err := o.dao.UserClient().Find(&entity.UserClient{
			OpenId:     res["openid"].(string),
			ClientType: 2,
		}); err == nil && userClientMessage.UserPk != 0 {
			return strconv.FormatInt(userClientMessage.UserPk, 10), err
		}
	}
	return "", errors.New("登录失败，请稍后重试")
}

func (o *wechat) AppletRegister(loginCode, phoneCode, name string) (string, error) {
	app := o.Mp()
	phoneRes := app.PostWxaBusinessGetUserPhoneNumber(phoneCode)
	if phoneRes == nil || phoneRes["phoneNumber"] == nil || phoneRes["phoneNumber"].(string) == "" {
		return "", errors.New("获取手机号失败")
	}
	phone := phoneRes["phoneNumber"].(string)
	user, err := o.dao.User().Find(&entity.User{Phone: phone})
	if user.Pk == 0 {
		user, err = o.dao.User().Create(&entity.User{
			Pk:    helper.GetRid(helper.User),
			Phone: phone,
			Name:  name,
		})
		if err != nil {
			return "", err
		}
	}
	loginRes := app.JsCode2Session(loginCode)
	if loginRes == nil || loginRes["openid"] == nil || loginRes["unionid"] == nil {
		return "", errors.New("获取微信凭证失败")
	}
	clientRes, _ := o.dao.UserClient().Find(&entity.UserClient{
		ClientType: 2,
		OpenId:     loginRes["openid"].(string),
	})
	if clientRes.UserPk == user.Pk {
		return strconv.FormatInt(user.Pk, 10), nil
	} else if clientRes.UserPk != 0 {
		return "", errors.New("该微信已绑定其他手机号")
	} else {
		err = o.dao.UserClient().CreateUserClient(&entity.UserClient{
			ClientId:   app.Key(),
			UserPk:     user.Pk,
			ClientType: 2,
			OpenId:     loginRes["openid"].(string),
			WxUnionid:  loginRes["unionid"].(string),
		})
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(user.Pk, 10), nil
	}
}

func (o *wechat) CreateEnterpriseForRelationProject(user *entity.User, code string, enterprise *entity.Enterprise) (*entity.Enterprise, error) {
	// 查询关联关系
	loginRes := o.Mp().JsCode2Session(code)
	if loginRes == nil || loginRes["openid"] == nil || loginRes["unionid"] == nil {
		return enterprise, errors.New("获取微信凭证失败")
	}
	// 获取用户
	clientRes, _ := o.dao.UserClient().Find(&entity.UserClient{
		ClientType: 2,
		OpenId:     loginRes["openid"].(string),
	})
	if clientRes.UserPk == 0 {
		// 获取用户
		user, err := o.dao.User().Find(user)
		if user.Pk == 0 {
			user.Pk = helper.GetRid(helper.User)
			user, err = o.dao.User().Create(user)
			if err != nil {
				return enterprise, err
			}
		}
		// 创建关系
		clientRes = &entity.UserClient{
			UserPk:     user.Pk,
			ClientType: 2,
			OpenId:     loginRes["openid"].(string),
			WxUnionid:  loginRes["unionid"].(string),
		}
		if err := o.dao.UserClient().CreateUserClient(clientRes); err != nil {
			return enterprise, err
		}
	}
	// 创建企业
	enterprise.Pk = helper.GetRid(helper.Enterprise)
	enterprise.CorporatePk = user.Pk
	if en, err := NewEnterprise(o.dao).Create(enterprise); err != nil {
		return enterprise, err
	} else {
		// 授予权限
		return enterprise, NewUserPermission(o.dao).CreateForCorporate(&entity.UserPermission{
			EnterprisePk: en.Pk,
			UserPk:       user.Pk,
		})
	}
}
