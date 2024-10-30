package organization

import (
	"github.com/gin-gonic/gin"
	utils "github.com/leapig/fastgo/app/library/util"
)

func Load(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.GET("base_info", utils.UserAccessTokenJWTMiddleware(), GetUserBaseInfo)
		user.GET("info", utils.TenantAccessTokenJWTMiddleware(), GetUserInfo)
		user.GET("verify", utils.TenantAccessTokenJWTMiddleware(), GetUserVerify)
		user.POST("verify", utils.TenantAccessTokenJWTMiddleware(), PostUserVerify)
		user.GET("code", utils.TenantAccessTokenJWTMiddleware(), GetUserCode)
	}
}

type VerifyUserReq struct {
	Name   string `json:"name"  form:"name" query:"name"`
	IdCard string `json:"id_card"  form:"id_card" query:"id_card"`
	Face   string `json:"face"  form:"face" query:"face"`
}
