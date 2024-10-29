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

// CreateRole
// @Tags open-apis/core
// @summary 新增角色(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateRoleReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/role [post]
func CreateRole(c *gin.Context) {
	var p CreateRoleReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.Role().Create(&entity.Role{
		RoleName:     p.RoleName,
		Remark:       p.Remark,
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
	})
	utils.R(c, res, err)
}

// UpdaterRole
// @Tags open-apis/core
// @summary 修改角色(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdaterRoleReq true "请求参数体"
// @Router /open-apis/core/rbac/system/role [put]
func UpdaterRole(c *gin.Context) {
	var p UpdaterRoleReq
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

// DeleteRole
// @Tags open-apis/core
// @summary 删除角色(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/role [delete]
func DeleteRole(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.Role().Delete(&entity.Role{Pk: utils.StringToInt64(p.Pk)})
	if err == nil {
		_ = C.S.RolePermission().DeleteByRolePk(&entity.RolePermission{
			RolePk: utils.StringToInt64(p.Pk),
		})
	}
	utils.R(c, nil, err)
}

// SelectRole
// @Tags open-apis/core
// @summary 获取角色(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_name query string false "组名"
// @Router	/open-apis/core/rbac/system/role [get]
func SelectRole(c *gin.Context) {
	var p RoleQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	roles, page, err := C.S.Role().SelectWithDetail(&entity.Role{
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
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
			if userPermissions, _, err3 := C.S.UserPermission().SelectWithUserMessage(&entity.UserPermission{
				PermissionType: 2,
				PermissionPk:   v.Pk,
				EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
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

// CreateRoleGroup
// @Tags open-apis/core
// @summary 新增角色组(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateRoleGroupReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/role/group [post]
func CreateRoleGroup(c *gin.Context) {
	var p CreateRoleGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.RoleGroup().Create(&entity.RoleGroup{
		RoleGroupName: p.RoleGroupName,
		Remark:        p.Remark,
		EnterprisePk:  utils.StringToInt64(p.EnterprisePk),
	})
	utils.R(c, res, err)
}

// UpdaterRoleGroup
// @Tags open-apis/core
// @summary 修改角色组(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdaterRoleReq true "请求参数体"
// @Router /open-apis/core/rbac/system/role/group [put]
func UpdaterRoleGroup(c *gin.Context) {
	var p UpdaterRoleGroupReq
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

// DeleteRoleGroup
// @Tags open-apis/core
// @summary 删除角色组(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/role/group [delete]
func DeleteRoleGroup(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if err := C.S.RoleGroup().Delete(&entity.RoleGroup{
		Pk: utils.StringToInt64(p.Pk),
	}); err == nil {
		err := C.S.RoleGroupPermission().DeleteByGroupPk(&entity.RoleGroupPermission{
			RoleGroupPk: utils.StringToInt64(p.Pk),
		})
		utils.R(c, nil, err)
	} else {
		utils.FR(c, err)
	}
}

// SelectRoleGroup
// @Tags open-apis/core
// @summary 获取角色组列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_group_name query string false "组名"
// @Router	/open-apis/core/rbac/system/role/group [get]
func SelectRoleGroup(c *gin.Context) {
	var p RoleGroupQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	roleGroup, page, err := C.S.RoleGroup().Select(&entity.RoleGroup{
		EnterprisePk:  utils.StringToInt64(p.EnterprisePk),
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
				EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
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

// CreateRoleGroupPermission
// @Tags open-apis/core
// @summary 新增角色组与权限关联关系(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateRoleGroupPermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/role/group/permission [post]
func CreateRoleGroupPermission(c *gin.Context) {
	var p CreateRoleGroupPermissionReq
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
	if err != nil {
		utils.FR(c, err)
	}
	utils.R(c, res, err)
}

// DeleteRoleGroupPermission
// @Tags open-apis/core
// @summary 删除角色组与权限关联关系(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/role/group/permission [delete]
func DeleteRoleGroupPermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.RoleGroupPermission().Delete(&entity.RoleGroupPermission{
		Pk: utils.StringToInt64(p.Pk),
	})
	utils.R(c, nil, err)
}

// CreateRolePermission
// @Tags open-apis/core
// @summary 新增角色与权限关联关系(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateRolePermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/role/permission [post]
func CreateRolePermission(c *gin.Context) {
	var p CreateRolePermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if _, err := C.S.Role().FindByPk(&entity.Role{Pk: utils.StringToInt64(p.RolePk)}); err != nil {
		utils.FR(c, errors.New("缺少角色信息！！！！！"))
	}
	switch p.PermissionType {
	case 1:
		if _, err := C.S.Permission().FindByPk(&entity.Permission{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
			utils.FR(c, errors.New("不存在该权限！！！！！"))
		}
	case 2:
		if _, err := C.S.PermissionGroup().FindByPk(&entity.PermissionGroup{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
			utils.FR(c, errors.New("不存在该权限组！！！！！"))
		}
	default:
		utils.FR(c, errors.New("错误类型！！！！！"))
	}
	res, err := C.S.RolePermission().Create(&entity.RolePermission{
		PermissionType: p.PermissionType,
		PermissionPk:   utils.StringToInt64(p.PermissionPk),
		RolePk:         utils.StringToInt64(p.RolePk),
	})
	if err != nil {
		utils.FR(c, err)
	}
	utils.R(c, res, err)
}

// DeleteRolePermission
// @Tags open-apis/core
// @summary 删除角色与权限关联关系(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/role/permission [delete]
func DeleteRolePermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.RolePermission().Delete(&entity.RolePermission{
		Pk: utils.StringToInt64(p.Pk),
	})
	utils.R(c, nil, err)
}

// SelectRoleUser
// @Tags open-apis/core
// @summary 获取角色人员列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_pk query string true "角色主键"
// @Router	/open-apis/core/rbac/system/role/user [get]
func SelectRoleUser(c *gin.Context) {
	var p RoleUserQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if message, page, err := C.S.UserPermission().SelectWithUserMessage(&entity.UserPermission{
		EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
		PermissionPk:   utils.StringToInt64(p.RolePk),
		PermissionType: 2,
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

// SelectRoleGroupUser
// @Tags open-apis/core
// @summary 获取角色组人员列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param role_group_pk query string true "角色组主键"
// @Router	/open-apis/core/rbac/system/role/group/user [get]
func SelectRoleGroupUser(c *gin.Context) {
	var p RoleGroupUserQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if message, page, err := C.S.UserPermission().SelectWithUserMessage(&entity.UserPermission{
		EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
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
