package enterprise

import (
	"github.com/gin-gonic/gin"
	utils "github.com/leapig/fastgo/app/library/util"
)

func Load(r *gin.RouterGroup) {
	r.POST("new", PostNew)
	r.POST("create", utils.UserAccessTokenJWTMiddleware(), PostCreate)
	r.GET("list", utils.UserAccessTokenJWTMiddleware(), GetList)
	r.GET("applist", utils.UserAccessTokenJWTMiddleware(), GetAppList)
	r.PUT("corporate", utils.UserAccessTokenJWTMiddleware(), PutCorporate)
	r.GET("security", utils.UserAccessTokenJWTMiddleware(), GetSecurity)
	r.POST("security", utils.UserAccessTokenJWTMiddleware(), PostSecurity)
	r.PUT("security", utils.UserAccessTokenJWTMiddleware(), PutSecurity)
}
