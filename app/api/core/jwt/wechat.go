package jwt

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/library/jwt"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
	"github.com/leapig/tpp/util"
	"io"
	"net/http"
	"os"
	"strings"
)

// AnyCallBack 服务号事件订阅消息
func AnyCallBack(c *gin.Context) {
	var query struct {
		MsgSignature string `json:"msg_signature" form:"msg_signature" query:"msg_signature"`
		Signature    string `json:"signature" form:"signature" query:"signature"`
		Timestamp    string `json:"timestamp"     form:"timestamp"     query:"timestamp"`
		Nonce        string `json:"nonce"         form:"nonce"         query:"nonce"`
		EchoStr      string `json:"echostr"       form:"echostr"       query:"echostr"`
		EncryptType  string `json:"encrypt_type"    form:"encrypt_type"       query:"encrypt_type"`
		Openid       string `json:"openid"    form:"openid"       query:"openid"`
	}
	_ = c.ShouldBindQuery(&query)
	if c.Request.Method == "GET" {
		c.String(http.StatusOK, query.EchoStr)
	} else {
		res := &struct {
			ToUserName   string `xml:"ToUserName"`
			FromUserName string `xml:"FromUserName"`
			CreateTime   uint32 `xml:"CreateTime"`
			MsgType      string `xml:"MsgType"`
			Event        string `xml:"Event"`
			EventKey     string `xml:"EventKey"`
			Content      string `xml:"Content"`
			MsgId        string `xml:"MsgId"`
		}{}
		_ = c.ShouldBindQuery(&query)
		body, _ := io.ReadAll(c.Request.Body)
		cpt := util.NewWXBizMsgCrypt(c.Param("appid"), os.Getenv("WX_OA_TOKEN"), os.Getenv("WX_OA_AES_KEY"))
		if message, _, err := cpt.DecryptMsg(query.MsgSignature, query.Timestamp, query.Nonce, body); err == nil {
			_ = xml.Unmarshal(message, res)
			switch strings.ToLower(res.Event) {
			case "subscribe":
				C.S.Wechat().OaSaveSubscribe(query.Openid, res.ToUserName, 1)
				if res.EventKey != "" {
					if err := C.S.Wechat().OaScan(query.Openid, res.ToUserName, res.EventKey); err != nil {
						c.String(http.StatusOK, `<xml>
  <ToUserName><![CDATA[`+res.FromUserName+`]]></ToUserName>
  <FromUserName><![CDATA[`+res.ToUserName+`]]></FromUserName>
  <CreateTime>`+utils.Int64ToString(int64(res.CreateTime+1))+`</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[`+err.Error()+`]]></Content>
</xml>`)
					}
				}
				break
			case "unsubscribe":
				C.S.Wechat().OaSaveSubscribe(query.Openid, res.ToUserName, 2)
				break
			case "scan":
				if err := C.S.Wechat().OaScan(query.Openid, res.ToUserName, res.EventKey); err != nil {
					c.String(http.StatusOK, `<xml>
  <ToUserName><![CDATA[`+res.FromUserName+`]]></ToUserName>
  <FromUserName><![CDATA[`+res.ToUserName+`]]></FromUserName>
  <CreateTime>`+utils.Int64ToString(int64(res.CreateTime+1))+`</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[`+err.Error()+`]]></Content>
</xml>`)
				}
				break
			default:
				break
			}
			c.String(http.StatusOK, "success")
		}
	}
}

// GetQrCode
// @Tags jwt模块
// @summary 后台扫码登录>获取微信登录二维码
// @Accept json
// @Produce  json
// @Router	/open-apis/core/jwt/wechat/qrcode [get]
func GetQrCode(r *gin.Context) {
	res, err := C.S.Wechat().GetWxOaQrCode()
	utils.R(r, res, err)
}

// GetPhone
// @Tags open-apis/core
// @summary 小程序获取用户绑定手机号
// @Accept json
// @Produce  json
// @Param key query string true "小程序code"
// @Router	/open-apis/core/auth/wechat/phone [get]
func GetPhone(c *gin.Context) {
	var p _Key
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.Wechat().GetWxMpPhoneNumber(p.Key)
	utils.R(c, res, err)
}

// PostSignIn
// @Tags open-apis/core
// @summary 小程序登录获取UserAccessToken
// @Accept json
// @Produce  json
// @Param params body _Code true "请求参数体"
// @Router	/open-apis/core/jwt/wechat/sign_in [post]
func PostSignIn(c *gin.Context) {
	var p _Code
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if res, err := C.S.Wechat().AppletLogin(p.Code); err != nil {
		utils.FRA(c)
	} else {
		utils.R(c, jwt.GenerateUserToken(res), err)
	}
}

// PostSignUp
// @Tags open-apis/core
// @summary 小程序注册获取UserAccessToken
// @Accept json
// @Produce  json
// @Param params body KeyCode true "请求参数体"
// @Router	/open-apis/core/jwt/wechat/sign_up [post]
func PostSignUp(c *gin.Context) {
	var p KeyCode
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if res, err := C.S.Wechat().AppletRegister(p.Code, p.Key, "游客_"+p.Code[0:10]); err != nil {
		utils.FR(c, err)
	} else {
		utils.R(c, jwt.GenerateUserToken(res), err)
	}
}
