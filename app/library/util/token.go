package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/library/helper"
	"github.com/leapig/fastgo/app/library/jwt"
	"strings"
)

func UserAccessTokenJWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			FRA(c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
			FRA(c)
			c.Abort()
			return
		}
		res, err := jwt.ParseUserToken(parts[1])
		if err != nil {
			FRA(c)
			c.Abort()
			return
		}
		c.Set("user", res.UserPk)
		if scope, err := helper.RS.SIsmembers("scope_1_"+res.UserPk, c.Request.RequestURI+"_"+strings.ToUpper(c.Request.Method)); err == nil && scope != nil && scope.(int64) == 1 {
			c.Next()
		} else {
			//FRS(c)
			//c.Abort()
			//return
			c.Next()
		}
	}
}

func TenantAccessTokenJWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			FRA(c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
			FRA(c)
			c.Abort()
			return
		}
		res, err := jwt.ParseTenantToken(parts[1])
		if err != nil {
			FRA(c)
			c.Abort()
			return
		}
		c.Set("user", res.UserPk)
		c.Set("tenant", res.TenantPk)
		c.Set("tenant_user", res.TenantUserPk)
		if scope, err := helper.RS.SIsmembers("scope_"+res.TenantPk+"_"+res.UserPk, c.Request.RequestURI+"_"+strings.ToUpper(c.Request.Method)); err == nil && scope != nil && scope.(int64) == 1 {
			c.Next()
		} else {
			//FRS(c)
			//c.Abort()
			//return
			c.Next()
		}
	}
}
