package rbac

import (
	"github.com/gin-gonic/gin"
	utils "github.com/leapig/fastgo/app/library/util"
)

func Load(r *gin.RouterGroup) {
	// 系统级权限
	s := r.Group("system")
	{
		//获取人员权限
		s.GET("user/permission", utils.UserAccessTokenJWTMiddleware(), GetUserAllPermissionByUserPk)
		s.GET("user/role/group", utils.UserAccessTokenJWTMiddleware(), GetUserRoleAndRoleGroupByUserPKAndEnterprisePk)
		//权限：人员权限关联
		s.POST("user/permission", utils.UserAccessTokenJWTMiddleware(), CreateUserPermission)
		s.DELETE("user/permission", utils.UserAccessTokenJWTMiddleware(), DeleteUserPermission)
		//权限：角色
		s.POST("role", utils.UserAccessTokenJWTMiddleware(), CreateRole)
		s.PUT("role", utils.UserAccessTokenJWTMiddleware(), UpdaterRole)
		s.DELETE("role", utils.UserAccessTokenJWTMiddleware(), DeleteRole)
		s.GET("role", utils.UserAccessTokenJWTMiddleware(), SelectRole)
		s.GET("role/user", utils.UserAccessTokenJWTMiddleware(), SelectRoleUser)
		//权限：角色组
		s.POST("role/group", utils.UserAccessTokenJWTMiddleware(), CreateRoleGroup)
		s.PUT("role/group", utils.UserAccessTokenJWTMiddleware(), UpdaterRoleGroup)
		s.DELETE("role/group", utils.UserAccessTokenJWTMiddleware(), DeleteRoleGroup)
		s.GET("role/group", utils.UserAccessTokenJWTMiddleware(), SelectRoleGroup)
		s.GET("role/group/user", utils.UserAccessTokenJWTMiddleware(), SelectRoleGroupUser)
		//权限：角色组关联
		s.POST("role/group/permission", utils.UserAccessTokenJWTMiddleware(), CreateRoleGroupPermission)
		s.DELETE("role/group/permission", utils.UserAccessTokenJWTMiddleware(), DeleteRoleGroupPermission)
		//权限：角色关联
		s.POST("role/permission", utils.UserAccessTokenJWTMiddleware(), CreateRolePermission)
		s.DELETE("role/permission", utils.UserAccessTokenJWTMiddleware(), DeleteRolePermission)
		//权限：单一权限
		s.POST("permission", utils.UserAccessTokenJWTMiddleware(), CreatePermission)
		s.PUT("permission", utils.UserAccessTokenJWTMiddleware(), UpdaterPermission)
		s.DELETE("permission", utils.UserAccessTokenJWTMiddleware(), DeletePermission)
		s.GET("permission", utils.UserAccessTokenJWTMiddleware(), SelectPermission)
		//权限：权限组
		s.GET("permission/group", utils.UserAccessTokenJWTMiddleware(), SelectPermissionGroup)
		s.POST("permission/group", utils.UserAccessTokenJWTMiddleware(), CreatePermissionGroup)
		s.PUT("permission/group", utils.UserAccessTokenJWTMiddleware(), UpdaterPermissionGroup)
		s.DELETE("permission/group", utils.UserAccessTokenJWTMiddleware(), DeletePermissionGroup)
		//权限：权限组关联权限
		s.POST("permission/group/permission", utils.UserAccessTokenJWTMiddleware(), CreatePermissionGroupPermission)
		s.DELETE("permission/group/permission", utils.UserAccessTokenJWTMiddleware(), DeletePermissionGroupPermission)
	}
	// 资源
	p := r.Group("resource")
	{
		//权限：菜单
		p.GET("menu", utils.UserAccessTokenJWTMiddleware(), SelectAllMenuWithDetail)
		p.POST("menu", utils.UserAccessTokenJWTMiddleware(), CreateMenuResource)
		p.PUT("menu", utils.UserAccessTokenJWTMiddleware(), UpdateMenuResource)
		p.DELETE("menu", utils.UserAccessTokenJWTMiddleware(), DeleteMenuResource)
		//权限：页面资源
		p.GET("page", utils.UserAccessTokenJWTMiddleware(), SelectPageInterfaceDetailMessage)
		p.POST("page", utils.UserAccessTokenJWTMiddleware(), CreatePageResource)
		p.PUT("page", utils.UserAccessTokenJWTMiddleware(), UpdatePageResource)
		p.DELETE("page", utils.UserAccessTokenJWTMiddleware(), DeletePageResource)
		//权限：接口资源
		p.GET("interface", utils.UserAccessTokenJWTMiddleware(), SelectInterfaceResource)
		p.POST("interface", utils.UserAccessTokenJWTMiddleware(), CreateInterfaceResource)
		p.PUT("interface", utils.UserAccessTokenJWTMiddleware(), UpdateInterfaceResource)
		p.DELETE("interface", utils.UserAccessTokenJWTMiddleware(), DeleteInterfaceResource)
		//权限：页面资源与接口关联
		p.POST("page/interface", utils.UserAccessTokenJWTMiddleware(), CreatePageInterface)
		p.DELETE("page/interface", utils.UserAccessTokenJWTMiddleware(), DeletePageInterface)
	}
	// 租户级权限
	t := r.Group("tenant")
	{
		//获取人员权限
		t.GET("user/permission", utils.TenantAccessTokenJWTMiddleware(), GetEnterpriseUserAllPermissionByUserPk)
		t.GET("user", utils.TenantAccessTokenJWTMiddleware(), GetUserPermissionByUserPkAndEnterprisePkForRedis)
		t.GET("user/role/group", utils.TenantAccessTokenJWTMiddleware(), GetEnterpriseUserRoleAndRoleGroupByUserPKAndEnterprisePk)
		//权限：人员权限关联
		t.POST("user/permission", utils.TenantAccessTokenJWTMiddleware(), CreateEnterpriseUserPermission)
		t.DELETE("user/permission", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterpriseUserPermission)
		//权限：角色
		t.POST("role", utils.TenantAccessTokenJWTMiddleware(), CreateEnterpriseRole)
		t.PUT("role", utils.TenantAccessTokenJWTMiddleware(), UpdaterEnterpriseRole)
		t.DELETE("role", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterpriseRole)
		t.GET("role", utils.TenantAccessTokenJWTMiddleware(), SelectEnterpriseRole)
		t.GET("role/user", utils.TenantAccessTokenJWTMiddleware(), SelectEnterpriseRoleUser)
		//权限：角色组
		t.POST("role/group", utils.TenantAccessTokenJWTMiddleware(), CreateEnterpriseRoleGroup)
		t.PUT("role/group", utils.TenantAccessTokenJWTMiddleware(), UpdaterEnterpriseRoleGroup)
		t.DELETE("role/group", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterpriseRoleGroup)
		t.GET("role/group", utils.TenantAccessTokenJWTMiddleware(), SelectEnterpriseRoleGroup)
		t.GET("role/group/user", utils.TenantAccessTokenJWTMiddleware(), SelectEnterpriseRoleGroupUser)
		//权限：角色组关联
		t.POST("role/group/permission", utils.TenantAccessTokenJWTMiddleware(), CreateEnterpriseRoleGroupPermission)
		r.DELETE("role/group/permission", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterpriseRoleGroupPermission)
		//权限：角色关联
		t.POST("role/permission", utils.TenantAccessTokenJWTMiddleware(), CreateEnterpriseRolePermission)
		t.DELETE("role/permission", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterpriseRolePermission)
		//权限：权限组
		t.GET("permission/group", utils.TenantAccessTokenJWTMiddleware(), SelectEnterprisePermissionGroup)
		t.POST("permission/group", utils.TenantAccessTokenJWTMiddleware(), CreateEnterprisePermissionGroup)
		t.PUT("permission/group", utils.TenantAccessTokenJWTMiddleware(), UpdaterEnterprisePermissionGroup)
		t.DELETE("permission/group", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterprisePermissionGroup)
		//权限：权限组关联权限
		t.POST("permission/group/permission", utils.TenantAccessTokenJWTMiddleware(), CreateEnterprisePermissionGroupPermission)
		t.DELETE("permission/group/permission", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterprisePermissionGroupPermission)
		//权限：单一权限
		t.GET("permission", utils.TenantAccessTokenJWTMiddleware(), GetTenantPermission)
	}
}
