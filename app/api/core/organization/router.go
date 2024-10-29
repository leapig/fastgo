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
	}
}
