package rbac

import (
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// CreateEnterpriseRole
// @Tags open-apis/core
// @summary 新增角色(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostRoleReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/role [post]
func CreateEnterpriseRole(c *gin.Context) {
	var p PostRoleReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.Role().Create(&entity.Role{
		RoleName:     p.RoleName,
		Remark:       p.Remark,
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	})
	utils.R(c, res, err)
}

// UpdaterEnterpriseRole
// @Tags open-apis/core
// @summary 修改角色(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body PutRoleReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/role [put]
func UpdaterEnterpriseRole(c *gin.Context) {
	var p PutRoleReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.Role().Update(&entity.Role{
		Pk:       utils.StringToInt64(p.Pk),
		RoleName: p.RoleName,
		Remark:   p.Remark,
	})
	utils.R(c, res, err)
}

// DeleteEnterpriseRole
// @Tags open-apis/core
// @summary 删除角色(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/role [delete]
func DeleteEnterpriseRole(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.Role().Delete(&entity.Role{Pk: utils.StringToInt64(p.Pk)})
	if err == nil {
		err = C.S.RolePermission().DeleteByRolePk(&entity.RolePermission{
			RolePk: utils.StringToInt64(p.Pk),
		})
	}
	utils.R(c, nil, err)
}

// SelectEnterpriseRole
// @Tags open-apis/core
// @summary 获取角色(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_name query string false "组名"
// @Router	/open-apis/core/rbac/tenant/role [get]
func SelectEnterpriseRole(c *gin.Context) {
	var p GetRoleReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	roles, page, err := C.S.Role().SelectWithDetail(&entity.Role{
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
		RoleName:     p.RoleName,
	}, &dao.Pagination{
		Page: p.Page,
		Size: p.Size,
	})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.RoleWithPermission, 0)
	if roles != nil && len(roles) > 0 {
		for _, v := range roles {
			userRows := make([]*pb.UserForRoleAndRoleGroup, 0)
			permissionRows := make([]*pb.PermissionForRole, 0)
			permissionGroupRows := make([]*pb.PermissionGroupForRole, 0)
			if userPermissions, _, err3 := C.S.UserPermission().SelectWithEnterpriseUserMessage(&entity.UserPermission{
				PermissionType: 2,
				PermissionPk:   v.Pk,
				EnterprisePk:   utils.StringToInt64(c.GetString("tenant")),
			}, &dao.Pagination{
				Page: 1,
				Size: 5,
			}); err3 == nil && userPermissions != nil && len(userPermissions) > 0 {
				for _, z := range userPermissions {
					if z.EnterpriseUser != nil {
						userRows = append(userRows, &pb.UserForRoleAndRoleGroup{
							Name: z.EnterpriseUser.Name,
						})

					}
				}
			}
			if v.RolePermissionModel != nil && len(v.RolePermissionModel) > 0 {
				rolePermissions := v.RolePermissionModel
				for _, p := range rolePermissions {
					switch p.PermissionType {
					case 1:
						if p.PermissionPk != 0 && p.Permission != nil {
							permissionRows = append(permissionRows, &pb.PermissionForRole{
								Pk:             utils.Int64ToString(p.Permission.Pk),
								OperationType:  p.Permission.OperationType,
								Resource:       utils.Int64ToString(p.Permission.Resource),
								ResourceType:   p.Permission.ResourceType,
								EnterprisePk:   utils.Int64ToString(p.Permission.EnterprisePk),
								PermissionName: p.Permission.PermissionName,
								RelationPk:     utils.Int64ToString(p.Pk),
							})
						}
					case 2:
						if p.PermissionPk != 0 && p.PermissionGroup != nil {
							permissionGroupRows = append(permissionGroupRows, &pb.PermissionGroupForRole{
								Pk:           utils.Int64ToString(p.PermissionGroup.Pk),
								GroupName:    p.PermissionGroup.GroupName,
								Remark:       p.PermissionGroup.Remark,
								EnterprisePk: utils.Int64ToString(p.PermissionGroup.EnterprisePk),
								RelationPk:   utils.Int64ToString(p.Pk),
							})
						}
					default:
						logger.Error("错误类型！！！")
					}
				}
			}
			result = append(result, &pb.RoleWithPermission{
				Pk:                  utils.Int64ToString(v.Pk),
				RoleName:            v.RoleName,
				Remark:              v.Remark,
				EnterprisePk:        utils.Int64ToString(v.EnterprisePk),
				PermissionRows:      permissionRows,
				PermissionGroupRows: permissionGroupRows,
				UserRows:            userRows,
			})
		}
	}
	utils.R(c, &pb.Resp{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Rows:  result,
	}, err)
}

// CreateEnterpriseRolePermission
// @Tags open-apis/core
// @summary 新增角色与权限关联关系(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostRolePermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/role/permission [post]
func CreateEnterpriseRolePermission(c *gin.Context) {
	var p PostRolePermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	if _, err := C.S.Role().FindByPk(&entity.Role{Pk: utils.StringToInt64(p.RolePk)}); err != nil {
		utils.FR(c, err)
	}
	switch p.PermissionType {
	case 1:
		if _, err := C.S.Permission().FindByPk(&entity.Permission{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
			utils.FR(c, err)
		}
	case 2:
		if _, err := C.S.PermissionGroup().FindByPk(&entity.PermissionGroup{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
			utils.FR(c, err)
		}
	default:
		utils.FR(c, errors.New("错误类型！！！！！"))
	}
	res, err := C.S.RolePermission().Create(&entity.RolePermission{
		PermissionType: p.PermissionType,
		PermissionPk:   utils.StringToInt64(p.PermissionPk),
		RolePk:         utils.StringToInt64(p.RolePk),
	})
	utils.R(c, res, err)
}

// DeleteEnterpriseRolePermission
// @Tags open-apis/core
// @summary 删除角色与权限关联关系(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/role/permission [delete]
func DeleteEnterpriseRolePermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.RolePermission().Delete(&entity.RolePermission{
		Pk: utils.StringToInt64(p.Pk),
	})
	utils.R(c, nil, err)
}

// SelectEnterpriseRoleUser
// @Tags open-apis/core
// @summary 获取角色人员列表(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_pk query string true "角色主键"
// @Router	/open-apis/core/rbac/tenant/role/user [get]
func SelectEnterpriseRoleUser(c *gin.Context) {
	var p GetRoleUserReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if message, page, err := C.S.UserPermission().SelectWithEnterpriseUserMessage(&entity.UserPermission{
		EnterprisePk:   utils.StringToInt64(c.GetString("tenant")),
		PermissionPk:   utils.StringToInt64(p.RolePk),
		PermissionType: 2,
	}, &dao.Pagination{Page: p.Page, Size: p.Size}); err == nil && len(message) > 0 {
		row := make([]*pb.UserMessageForRoleAndRoleGroup, 0)
		for _, user := range message {
			row = append(row, &pb.UserMessageForRoleAndRoleGroup{
				Name:       user.EnterpriseUser.Name,
				Phone:      user.EnterpriseUser.Phone,
				RelationPk: utils.Int64ToString(user.Pk),
			})
		}
		utils.R(c, &pb.Resp{
			Page:  page.Page,
			Size:  page.Size,
			Total: page.Total,
			Rows:  row,
		}, err)
	} else {
		utils.FR(c, err)
	}
}
