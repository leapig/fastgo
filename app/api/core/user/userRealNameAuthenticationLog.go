package user

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// VerifyUserApplet
// @Tags open-apis/core
// @summary 人证比对(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body VerifyUserReq true "请求参数体"
// @Router	/open-apis/core/user/verify [post]
func VerifyUserApplet(c *gin.Context) {
	var p VerifyUserReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	card, cErr := C.S.UserCredentials().Find(&entity.UserCredentials{
		Name:   p.Name,
		Serial: p.IdCard,
		Type:   1,
		UserPk: utils.StringToInt64(c.GetString("user")),
	})
	if cErr == nil && card.Pk != 0 {
		utils.FR(c, cErr)
	}
	if err := C.S.UserRealNameAuthenticationLog().VerifyUser(c.GetString("user"), p.Name, p.IdCard, p.Face, p.Url); err != nil {
		utils.FR(c, err)
	}

	_, err := C.S.UserCredentials().Create(&entity.UserCredentials{
		UserPk: utils.StringToInt64(c.GetString("user")),
		Serial: p.IdCard,
		Type:   1,
		Name:   p.Name,
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

// CheckRealNameAuthentication
// @Tags open-apis/core
// @summary 查询人证比对(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/user/verify/check [get]
func CheckRealNameAuthentication(c *gin.Context) {
	err := C.S.UserRealNameAuthenticationLog().CheckRealNameAuthenticationLog(c.GetString("user"))
	utils.R(c, nil, err)
}
