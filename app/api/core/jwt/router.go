package jwt

import (
	"github.com/gin-gonic/gin"
	utils "github.com/leapig/fastgo/app/library/util"
)

func Load(r *gin.RouterGroup) {
	wx := r.Group("wechat")
	{
		wx.GET("qrcode", GetQrCode)
		wx.Any("callback/:appid", AnyCallBack)
		wx.POST("sign_up", PostSignUp)
		wx.POST("sign_in", PostSignIn)
		wx.GET("phone", GetPhone)
	}
	tk := r.Group("token")
	{
		tk.GET("user_token", GetUserToken)
		tk.POST("tenant_token", utils.UserAccessTokenJWTMiddleware(), PostTenantToken)
		tk.PUT("user_token", PutUserToken)
		tk.PUT("tenant_token", utils.TenantAccessTokenJWTMiddleware(), PutTenantToken)
	}
	tt := r.Group("tenant")
	{
		tt.GET("list", utils.UserAccessTokenJWTMiddleware(), GetList)
		tt.GET("menu", utils.TenantAccessTokenJWTMiddleware(), GetMenu)
		tt.GET("page", utils.TenantAccessTokenJWTMiddleware(), GetPage)
	}
}

type _Key struct {
	Key string `json:"key" form:"key" query:"key"`
}

type EnterprisePk struct {
	EnterprisePk string `json:"enterprise_pk"`
}

type _Code struct {
	Code string `json:"code" form:"code" query:"code"`
}

type KeyCode struct {
	_Key
	_Code
}
