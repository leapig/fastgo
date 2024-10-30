package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

func Load(r *gin.RouterGroup) {
	r.GET("sms", utils.TenantAccessTokenJWTMiddleware(), GetCodeSms)
	r.POST("file", utils.TenantAccessTokenJWTMiddleware(), PostFile)
}

type _Phone struct {
	Phone string `json:"phone" form:"phone" query:"phone"`
}

// GetCodeSms
// @Tags open-apis/core
// @summary 发送短信验证码(TenantAccessToken)
// @Accept  json
// @Produce  json
// @Security bearer
// @Param phone query string true "手机号"
// @Router	/open-apis/core/tool/sms/code [get]
func GetCodeSms(c *gin.Context) {
	var p _Phone
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if _, err := C.S.User().Find(&entity.User{Phone: p.Phone}); err != nil {
		utils.FR(c, err)
	}
	utils.R(c, nil, C.S.Sms().SendCode(p.Phone, "0"))
}

// PostFile
// @Tags open-apis/core
// @summary 上传文件(TenantAccessToken)
// @Description 上传文件
// @Accept  multipart/form-data
// @Produce  json
// @Security bearer
// @Param file formData file false "校验文件"
// @Router	/open-apis/core/tool/file [post]
func PostFile(c *gin.Context) {
	if object, err := utils.PutFile(c.Request, "file"); err != nil {
		utils.FRP(c)
	} else {
		res, err := C.S.Files().SaveFile(c.GetString("tenant"), object)
		utils.R(c, res, err)
	}
}
