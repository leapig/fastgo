package user

import (
	"github.com/gin-gonic/gin"
	utils "github.com/leapig/fastgo/app/library/util"
)

func Load(r *gin.RouterGroup) {
	r.GET("list", utils.UserAccessTokenJWTMiddleware(), GetUserList)
	r.GET("account", utils.UserAccessTokenJWTMiddleware(), GetUserAccount)
	r.POST("account", utils.UserAccessTokenJWTMiddleware(), PostUserAccount)
	r.PUT("phone", utils.UserAccessTokenJWTMiddleware(), PutUserPhone)
	r.PUT("client", utils.UserAccessTokenJWTMiddleware(), PutUserClient)
	r.GET("member", utils.TenantAccessTokenJWTMiddleware(), GetUserMember)
	r.PUT("base_info", utils.UserAccessTokenJWTMiddleware(), SetUserBaseInfo)
	//证件
	r.GET("credentials", utils.UserAccessTokenJWTMiddleware(), GetUserCredentials)
	r.POST("credentials", utils.UserAccessTokenJWTMiddleware(), CreateUserCredentials)
	r.DELETE("credentials", utils.UserAccessTokenJWTMiddleware(), DeleteUserCredentials)
}
