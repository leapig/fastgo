package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/library/jwt"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
	"strings"
)

// GetUserToken
// @Tags 登录鉴权
// @summary 后台扫码登录>获取用户UserAccessToken
// @Accept json
// @Produce  json
// @Param key query string true "key"
// @Router	/open-apis/core/jwt/token/user_token [get]
func GetUserToken(c *gin.Context) {
	var p _Key
	if err := c.ShouldBind(&p); err != nil {
		utils.FRP(c)
		return
	}
	if res, err := C.S.Wechat().PlatformLoginLoop(p.Key); res == "" || res == "0" {
		utils.FRA(c)
	} else {
		utils.R(c, jwt.GenerateUserToken(res), err)
	}
}

// PostTenantToken
// @Tags 登录鉴权
// @summary 生成TenantAccessToken(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param param body EnterprisePk true "请求参数体"
// @Router	/open-apis/core/jwt/token/tenant_token [post]
func PostTenantToken(c *gin.Context) {
	var p EnterprisePk
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
		return
	}
	//全量同步人员权限
	if err := C.S.UserPermission().FullRefreshUserPermissionByUserPkAndEnterprisePk(utils.StringToInt64(c.GetString("user")), utils.StringToInt64(p.EnterprisePk)); err != nil {
		utils.FRA(c)
		return
	}
	if p.EnterprisePk == "1" {
		utils.R(c, jwt.GenerateTenantToken(p.EnterprisePk, c.GetString("user"), c.GetString("user")), nil)
	} else {
		// TODO 缺少租户用户
		utils.R(c, jwt.GenerateTenantToken(p.EnterprisePk, c.GetString("user"), ""), nil)
	}
}

// PutUserToken
// @Tags 登录鉴权
// @summary 刷新UserAccessToken（UserRefreshToken）
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/jwt/token/user_token [put]
func PutUserToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		utils.FRA(c)
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
		utils.FRA(c)
		return
	}
	utils.R(c, jwt.RefreshUserToken(parts[1]), nil)
}

// PutTenantToken
// @Tags 登录鉴权
// @summary 刷新TenantAccessToken(TenantRefreshToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/jwt/token/tenant_token [put]
func PutTenantToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		utils.FRA(c)
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
		utils.FRA(c)
		return
	}
	token := jwt.RefreshTenantToken(parts[1])
	if token.TenantAccessToken == "" {
		utils.FR(c, errors.New("生成失败！"))
		return
	} else {
		//检测权限是否有变动
		tenantToken, err := jwt.ParseTenantToken(token.TenantAccessToken)
		if err != nil {
			utils.FR(c, errors.New("token生成失败！"))
		}
		if C.S.UserPermission().CheckFullRefreshUserPermissionByUserPkAndEnterprisePk(utils.StringToInt64(tenantToken.UserPk), utils.StringToInt64(tenantToken.TenantPk)) {
			utils.FR(c, errors.New("权限变动，请重新登录！！"))
			return
		}
		utils.R(c, tenantToken, nil)
	}
}
