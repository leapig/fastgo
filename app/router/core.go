package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/api/core/enterprise"
	"github.com/leapig/fastgo/app/api/core/jwt"
	"github.com/leapig/fastgo/app/api/core/organization"
	"github.com/leapig/fastgo/app/api/core/personnel"
	"github.com/leapig/fastgo/app/api/core/rbac"
	"github.com/leapig/fastgo/app/api/core/tool"
	"github.com/leapig/fastgo/app/api/core/user"
)

func Core(r *gin.RouterGroup) {
	// 令牌鉴权
	jwt.Load(r.Group("jwt"))
	// 组织架构
	organization.Load(r.Group("organization"))

	// 权限
	rbac.Load(r.Group("rbac"))
	// 租户
	enterprise.Load(r.Group("enterprise"))
	// 工具
	tool.Load(r.Group("tool"))
	// 用户
	user.Load(r.Group("user"))
	// 部门
	personnel.Load(r.Group("personnel"))
}
