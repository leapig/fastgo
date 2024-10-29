package rbac

import (
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// GetEnterpriseUserAllPermissionByUserPk
// @Tags open-apis/core
// @summary 获取账号权限(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param user_pk query string true "用户Pk"
// @Router	/open-apis/core/rbac/tenant/user/permission [get]
func GetEnterpriseUserAllPermissionByUserPk(c *gin.Context) {
	var p EnterpriseUserPermissionQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	permissions, err := C.S.UserPermission().SelectPermissionByUserPkAndEnterprisePk(&entity.UserPermission{
		UserPk:       utils.StringToInt64(p.UserPk),
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	})
	if err != nil || permissions == nil || len(permissions) <= 0 {
		utils.FR(c, err)
	}
	result := make([]*pb.Permission, 0)
	for _, v := range permissions {
		result = append(result, &pb.Permission{
			Pk:             utils.Int64ToString(v.Pk),
			OperationType:  v.OperationType,
			Resource:       utils.Int64ToString(v.Resource),
			ResourceType:   v.ResourceType,
			EnterprisePk:   utils.Int64ToString(v.EnterprisePk),
			PermissionName: v.PermissionName,
		})
	}
	utils.R(c, &pb.Resp{Rows: result}, nil)
}

// GetEnterpriseUserRoleAndRoleGroupByUserPKAndEnterprisePk
// @Tags open-apis/core
// @summary 获取账号角色及角色组(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param user_pk query string true "用户Pk"
// @Router	/open-apis/core/rbac/tenant/user/role/group [get]
func GetEnterpriseUserRoleAndRoleGroupByUserPKAndEnterprisePk(c *gin.Context) {
	var p EnterpriseUserPermissionQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if _, err := C.S.User().Find(&entity.User{Pk: utils.StringToInt64(p.UserPk)}); err != nil {
		utils.FR(c, err)
	}
	userPermissionList, err := C.S.UserPermission().SelectAllUserPermission(&entity.UserPermission{
		UserPk:       utils.StringToInt64(p.UserPk),
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	})
	if err != nil {
		utils.FR(c, err)
	}
	var roleRows = make([]*pb.RoleForUserPermission, 0)
	var roleGroupRows = make([]*pb.RoleGroupForUserPermission, 0)
	for _, userPermission := range userPermissionList {
		switch userPermission.PermissionType {
		case 1: //角色组
			func() {
				if roleGroup, err2 := C.S.RoleGroup().FindByPk(&entity.RoleGroup{
					Pk: userPermission.PermissionPk,
				}); err2 == nil && roleGroup != nil {
					roleGroupRows = append(roleGroupRows, &pb.RoleGroupForUserPermission{
						Pk:            utils.Int64ToString(roleGroup.Pk),
						RoleGroupName: roleGroup.RoleGroupName,
						Remark:        roleGroup.Remark,
						EnterprisePk:  utils.Int64ToString(roleGroup.EnterprisePk),
						RelationPk:    utils.Int64ToString(userPermission.Pk),
					})
				}
			}()
		case 2: //角色
			func() {
				if role, err2 := C.S.Role().FindByPk(&entity.Role{
					Pk: userPermission.PermissionPk,
				}); err2 == nil && role != nil {
					roleRows = append(roleRows, &pb.RoleForUserPermission{
						Pk:           utils.Int64ToString(role.Pk),
						RoleName:     role.RoleName,
						Remark:       role.Remark,
						EnterprisePk: utils.Int64ToString(role.EnterprisePk),
						RelationPk:   utils.Int64ToString(userPermission.Pk),
					})
				}
			}()
		default:
			logger.Error("账号权限关联表存在非法类型！！！！！！")
		}
	}
	utils.R(c, &pb.UserPermissionRoleAndRoleGroupResp{
		RoleGroupRows: roleGroupRows,
		RoleRows:      roleRows,
	}, nil)
}

// CreateEnterpriseUserPermission
// @Tags open-apis/core
// @summary 新增人与权限关联(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostUserPermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/user/permission [post]
func CreateEnterpriseUserPermission(c *gin.Context) {
	var p PostUserPermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if _, err := C.S.User().Find(&entity.User{Pk: utils.StringToInt64(p.UserPk)}); err != nil {
		utils.FR(c, errors.New("不存在账号信息！！！"))
	}
	switch p.PermissionType {
	case 2:
		if _, err := C.S.Role().FindByPk(&entity.Role{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
			utils.FR(c, errors.New("缺少角色信息！！！"))
		}
	case 1:
		if _, err := C.S.RoleGroup().FindByPk(&entity.RoleGroup{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
			utils.FR(c, errors.New("不存在该角色组！！！"))
		}
	default:
		utils.FR(c, errors.New("错误类型！！！"))
	}
	res, err := C.S.UserPermission().Create(&entity.UserPermission{
		PermissionType: p.PermissionType,
		PermissionPk:   utils.StringToInt64(p.PermissionPk),
		UserPk:         utils.StringToInt64(p.UserPk),
		EnterprisePk:   utils.StringToInt64(c.GetString("tenant")),
	})
	utils.R(c, res, err)
}

// DeleteEnterpriseUserPermission
// @Tags open-apis/core
// @summary 删除人与权限关联(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/user/permission [delete]
func DeleteEnterpriseUserPermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.UserPermission().Delete(&entity.UserPermission{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// GetUserPermissionByUserPkAndEnterprisePkForRedis
// @Tags open-apis/core
// @summary 获取缓存中的账号权限数据(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/rbac/tenant/user [get]
func GetUserPermissionByUserPkAndEnterprisePkForRedis(c *gin.Context) {
	res, err := C.S.UserPermission().GetUserPermissionByUserPkAndEnterprisePkForRedis(utils.StringToInt64(c.GetString("user")), utils.StringToInt64(c.GetString("tenant")))
	utils.R(c, res, err)
}
