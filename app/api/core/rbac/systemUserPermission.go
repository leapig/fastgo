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

// GetUserAllPermissionByUserPk
// @Tags open-apis/core
// @summary 获取账号权限(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param user_pk query string true "用户Pk"
// @Param enterprise_pk query string true "租户Pk"
// @Router	/open-apis/core/rbac/system/user/permission [get]
func GetUserAllPermissionByUserPk(c *gin.Context) {
	var p PermissionQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	permissions, err := C.S.UserPermission().SelectPermissionByUserPkAndEnterprisePk(&entity.UserPermission{
		UserPk:       utils.StringToInt64(p.UserPk),
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
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
	utils.R(c, &pb.Resp{
		Rows: result,
	}, err)
}

// GetUserRoleAndRoleGroupByUserPKAndEnterprisePk
// @Tags open-apis/core
// @summary 获取账号角色及角色组接口(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param user_pk query string true "用户Pk"
// @Param enterprise_pk query string true "租户Pk"
// @Router	/open-apis/core/rbac/system/user/role/group [get]
func GetUserRoleAndRoleGroupByUserPKAndEnterprisePk(c *gin.Context) {
	var p PermissionQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if _, err := C.S.User().Find(&entity.User{Pk: utils.StringToInt64(p.UserPk)}); err != nil {
		utils.FR(c, err)
	}
	userPermissionList, err := C.S.UserPermission().SelectAllUserPermission(&entity.UserPermission{
		UserPk:       utils.StringToInt64(p.UserPk),
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
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

// CreateUserPermission
// @Tags open-apis/core
// @summary 新增人与权限关联(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateUserPermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/user/permission [post]
func CreateUserPermission(c *gin.Context) {
	var p CreateUserPermissionReq
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
		EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
	})
	if err != nil {
		utils.FR(c, err)
	}
	utils.R(c, res, err)
}

// DeleteUserPermission
// @Tags open-apis/core
// @summary 删除人与权限关联(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/user/permission [delete]
func DeleteUserPermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.UserPermission().Delete(&entity.UserPermission{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}
