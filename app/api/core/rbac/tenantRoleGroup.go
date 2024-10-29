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

// CreateEnterpriseRoleGroup
// @Tags open-apis/core
// @summary 新增角色组(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostRoleGroupReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/role/group [post]
func CreateEnterpriseRoleGroup(c *gin.Context) {
	var p PostRoleGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.RoleGroup().Create(&entity.RoleGroup{
		RoleGroupName: p.RoleGroupName,
		Remark:        p.Remark,
		EnterprisePk:  utils.StringToInt64(c.GetString("tenant")),
	})
	utils.R(c, res, err)
}

// UpdaterEnterpriseRoleGroup
// @Tags open-apis/core
// @summary 修改角色组(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body PutRoleGroupReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/role/group [put]
func UpdaterEnterpriseRoleGroup(c *gin.Context) {
	var p PutRoleGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.RoleGroup().Update(&entity.RoleGroup{
		Pk:            utils.StringToInt64(p.Pk),
		RoleGroupName: p.RoleGroupName,
		Remark:        p.Remark,
	})
	utils.R(c, res, err)
}

// DeleteEnterpriseRoleGroup
// @Tags open-apis/core
// @summary 删除角色组(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/role/group [delete]
func DeleteEnterpriseRoleGroup(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if err := C.S.RoleGroup().Delete(&entity.RoleGroup{
		Pk: utils.StringToInt64(p.Pk),
	}); err == nil {
		err = C.S.RoleGroupPermission().DeleteByGroupPk(&entity.RoleGroupPermission{
			RoleGroupPk: utils.StringToInt64(p.Pk),
		})
		utils.R(c, nil, err)
	} else {
		utils.FR(c, err)
	}
}

// SelectEnterpriseRoleGroup
// @Tags open-apis/core
// @summary 获取角色组列表(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_group_name query string false "组名"
// @Router	/open-apis/core/rbac/tenant/role/group [get]
func SelectEnterpriseRoleGroup(c *gin.Context) {
	var p GetRoleGroupReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	roleGroup, page, err := C.S.RoleGroup().Select(&entity.RoleGroup{
		EnterprisePk:  utils.StringToInt64(c.GetString("tenant")),
		RoleGroupName: p.RoleGroupName,
	}, &dao.Pagination{
		Page: p.Page,
		Size: p.Size,
	})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.RoleGroupWithPermission, 0)
	if roleGroup != nil && len(roleGroup) > 0 {
		for _, v := range roleGroup {
			userRows := make([]*pb.UserForRoleAndRoleGroup, 0)
			roleRows := make([]*pb.RoleForRoleGroup, 0)
			permissionRows := make([]*pb.PermissionForRoleGroup, 0)
			permissionGroupRows := make([]*pb.PermissionGroupForRoleGroup, 0)
			if userPermissions, _, err3 := C.S.UserPermission().SelectWithUserMessage(&entity.UserPermission{
				PermissionType: 1,
				PermissionPk:   v.Pk,
				EnterprisePk:   utils.StringToInt64(c.GetString("tenant")),
			}, &dao.Pagination{
				Page: 1,
				Size: 5,
			}); err3 == nil && userPermissions != nil && len(userPermissions) > 0 {
				for _, z := range userPermissions {
					if z.User != nil {
						userRows = append(userRows, &pb.UserForRoleAndRoleGroup{
							Name: z.User.Name,
						})
					}
				}
			}
			if rolePermissions, err2 := C.S.RoleGroupPermission().SelectRoleGroupPermissionByRoleGroupPk(&entity.RoleGroupPermission{
				RoleGroupPk: v.Pk,
			}); err2 == nil && rolePermissions != nil && len(rolePermissions) > 0 {
				for _, z := range rolePermissions {
					switch z.PermissionType {
					case 1:
						if z.PermissionPk != 0 && z.Role != nil {
							roleRows = append(roleRows, &pb.RoleForRoleGroup{
								Pk:           utils.Int64ToString(z.Role.Pk),
								RoleName:     z.Role.RoleName,
								Remark:       z.Role.Remark,
								EnterprisePk: utils.Int64ToString(z.Role.EnterprisePk),
								RelationPk:   utils.Int64ToString(z.Pk),
							})
						}
					case 2:
						if z.PermissionPk != 0 && z.Permission != nil {
							permissionRows = append(permissionRows, &pb.PermissionForRoleGroup{
								Pk:             utils.Int64ToString(z.Permission.Pk),
								OperationType:  z.Permission.OperationType,
								Resource:       utils.Int64ToString(z.Permission.Resource),
								ResourceType:   z.Permission.ResourceType,
								EnterprisePk:   utils.Int64ToString(z.Permission.EnterprisePk),
								PermissionName: z.Permission.PermissionName,
								RelationPk:     utils.Int64ToString(z.Pk),
							})
						}
					case 3:
						if z.PermissionPk != 0 && z.PermissionGroup != nil {
							permissionGroupRows = append(permissionGroupRows, &pb.PermissionGroupForRoleGroup{
								Pk:           utils.Int64ToString(z.PermissionGroup.Pk),
								GroupName:    z.PermissionGroup.GroupName,
								Remark:       z.PermissionGroup.Remark,
								EnterprisePk: utils.Int64ToString(z.PermissionGroup.EnterprisePk),
								RelationPk:   utils.Int64ToString(z.Pk),
							})
						}
					default:
						logger.Error("错误类型！！！！")
					}
				}
			}
			result = append(result, &pb.RoleGroupWithPermission{
				Pk:                  utils.Int64ToString(v.Pk),
				RoleGroupName:       v.RoleGroupName,
				Remark:              v.Remark,
				EnterprisePk:        utils.Int64ToString(v.EnterprisePk),
				RoleRows:            roleRows,
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

// SelectEnterpriseRoleGroupUser
// @Tags open-apis/core
// @summary 获取角色组人员列表(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_group_pk query string true "角色组主键"
// @Router	/open-apis/core/rbac/tenant/role/group/user [get]
func SelectEnterpriseRoleGroupUser(c *gin.Context) {
	var p GetRoleGroupUserReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if message, page, err := C.S.UserPermission().SelectWithUserMessage(&entity.UserPermission{
		EnterprisePk:   utils.StringToInt64(c.GetString("tenant")),
		PermissionPk:   utils.StringToInt64(p.RoleGroupPk),
		PermissionType: 1,
	}, &dao.Pagination{Page: p.Page, Size: p.Size}); err == nil && len(message) > 0 {
		row := make([]*pb.UserMessageForRoleAndRoleGroup, 0)
		for _, user := range message {
			row = append(row, &pb.UserMessageForRoleAndRoleGroup{
				Name:       user.User.Name,
				Phone:      user.User.Phone,
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

// CreateEnterpriseRoleGroupPermission
// @Tags open-apis/core
// @summary 新增角色组与权限关联关系(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostRoleGroupPermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/role/group/permission [post]
func CreateEnterpriseRoleGroupPermission(c *gin.Context) {
	var p PostRoleGroupPermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	switch p.PermissionType {
	case 1:
		if _, err := C.S.Role().FindByPk(&entity.Role{
			Pk: utils.StringToInt64(p.PermissionPk),
		}); err != nil {
			utils.FR(c, err)
		}
	case 2:
		if _, err := C.S.Permission().FindByPk(&entity.Permission{
			Pk: utils.StringToInt64(p.PermissionPk),
		}); err != nil {
			utils.FR(c, err)
		}
	case 3:
		if _, err := C.S.PermissionGroup().FindByPk(&entity.PermissionGroup{
			Pk: utils.StringToInt64(p.PermissionPk),
		}); err != nil {
			utils.FR(c, err)
		}
	default:
		utils.FR(c, errors.New("错误的类型！！！"))
	}
	if _, err := C.S.RoleGroup().FindByPk(&entity.RoleGroup{
		Pk: utils.StringToInt64(p.RoleGroupPk),
	}); err != nil {
		utils.FR(c, err)
	}
	res, err := C.S.RoleGroupPermission().Create(&entity.RoleGroupPermission{
		PermissionPk:   utils.StringToInt64(p.PermissionPk),
		PermissionType: p.PermissionType,
		RoleGroupPk:    utils.StringToInt64(p.RoleGroupPk),
	})
	utils.R(c, res, err)
}

// DeleteEnterpriseRoleGroupPermission
// @Tags open-apis/core
// @summary 删除角色组与权限关联关系(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/role/group/permission [delete]
func DeleteEnterpriseRoleGroupPermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.RoleGroupPermission().Delete(&entity.RoleGroupPermission{
		Pk: utils.StringToInt64(p.Pk),
	})
	utils.R(c, nil, err)
}
