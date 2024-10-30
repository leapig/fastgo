package organization

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/boombuler/barcode/code128"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
	"github.com/skip2/go-qrcode"
	"image/png"
)

// PostUserVerify
// @Tags 组织架构
// @summary 实名认证(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body VerifyUserReq true "请求参数体"
// @Router	/open-apis/core/user/verify [post]
func PostUserVerify(c *gin.Context) {
	var p VerifyUserReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
		return
	}
	res, err := C.S.UserCredentials().Find(&entity.UserCredentials{
		Type:   2,
		UserPk: utils.StringToInt64(c.GetString("user")),
	})
	if err != nil && res.Pk != 0 {
		utils.FR(c, errors.New("请勿重复认证"))
		return
	}
	object, err := utils.PutFile(c.Request, "face")
	if err != nil {
		utils.FRP(c)
		return
	}
	face, _ := C.S.Files().SaveFile(c.GetString("tenant"), object)
	if err = C.S.UserRealNameAuthenticationLog().VerifyUser(p.Name, p.IdCard, p.Face); err != nil {
		utils.FR(c, err)
		return
	}
	_, err = C.S.UserCredentials().Create(&entity.UserCredentials{
		UserPk: utils.StringToInt64(c.GetString("user")),
		Serial: p.IdCard,
		Type:   2,
		Cert:   p.IdCard,
		Name:   p.Name,
		Face:   face.Name,
	})
	_, err = C.S.UserCredentials().Create(&entity.UserCredentials{
		UserPk: utils.StringToInt64(c.GetString("user")),
		Type:   1,
		Cert:   face.Name,
		Face:   face.Name,
	})
	// 实名认证
	birthday, gender := utils.ParseIdCard(p.IdCard)
	_ = C.S.User().UpdateBaseInfo(&entity.User{
		Pk:       utils.StringToInt64(c.GetString("user")),
		Birthday: &birthday,
		Gender:   gender,
	})
	_, _ = C.S.EnterpriseUser().Update(&entity.EnterpriseUser{
		UserPk:   utils.StringToInt64(c.GetString("user")),
		Birthday: &birthday,
		Gender:   gender,
	})
	utils.R(c, nil, err)
}

// GetUserVerify
// @Tags 组织架构
// @summary 查询实名认证(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/organization/user/verify [get]
func GetUserVerify(c *gin.Context) {
	res, err := C.S.UserCredentials().Find(&entity.UserCredentials{
		Type:   2,
		UserPk: utils.StringToInt64(c.GetString("user")),
	})
	if err == nil && res.Pk == 0 {
		utils.R(c, nil, nil)
	} else {
		utils.FR(c, errors.New("未实名认证"))
	}
}

// GetUserCode
// @Tags 组织架构
// @summary 获取身份码(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/organization/user/code [get]
func GetUserCode(c *gin.Context) {
	code := utils.RandomStrByNum(9)
	res, err := C.S.UserCredentials().Find(&entity.UserCredentials{
		Type:   3,
		UserPk: utils.StringToInt64(c.GetString("user")),
	})
	if err == nil && res.Pk == 0 {
		res, err = C.S.UserCredentials().Find(&entity.UserCredentials{
			Type:   3,
			UserPk: utils.StringToInt64(c.GetString("user")),
			Cert:   code,
			Serial: code,
		})
	} else {
		code = res.Cert
	}
	q, _ := qrcode.New(code, qrcode.Medium)
	qrCode, _ := q.PNG(600)
	b, _ := code128.Encode(code)
	barcode := new(bytes.Buffer)

	// 将PNG图像编码为字节流并写入缓冲区
	_ = png.Encode(barcode, b)
	utils.R(c, map[string]string{
		"qrcode":  base64.StdEncoding.EncodeToString(qrCode),
		"barcode": base64.StdEncoding.EncodeToString(barcode.Bytes()),
	}, nil)
}
