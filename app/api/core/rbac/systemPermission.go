package rbac

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// CreatePermission
// @Tags open-apis/core
// @summary 新增权限项(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreatePermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/permission [post]
func CreatePermission(c *gin.Context) {
	var p CreatePermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if p.ResourceType == 1 {
		//菜单
		if _, err := C.S.MenuResource().FindByPk(&entity.MenuResource{
			Pk: utils.StringToInt64(p.Resource),
		}); err != nil {
			utils.FR(c, err)
		}
	} else if p.ResourceType == 2 {
		//页面
		if _, err := C.S.PageResource().FindByPk(&entity.PageResource{
			Pk: utils.StringToInt64(p.Resource),
		}); err != nil {
			utils.FR(c, err)
		}
	} else {
		utils.FR(c, errors.New("错误的资源类型！！！"))
	}
	res, err := C.S.Permission().Create(&entity.Permission{
		OperationType:  p.OperationType,
		Resource:       utils.StringToInt64(p.Resource),
		ResourceType:   p.ResourceType,
		EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
		PermissionName: p.PermissionName,
		Visibility:     p.Visibility,
	})
	utils.R(c, res, err)
}

// UpdaterPermission
// @Tags open-apis/core
// @summary 修改权限项(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdaterPermissionReq true "请求参数体"
// @Router /open-apis/core/rbac/system/permission [put]
func UpdaterPermission(c *gin.Context) {
	var p UpdaterPermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.Permission().Update(&entity.Permission{
		Pk:             utils.StringToInt64(p.Pk),
		OperationType:  p.OperationType,
		Resource:       utils.StringToInt64(p.Resource),
		ResourceType:   p.ResourceType,
		PermissionName: p.PermissionName,
		Visibility:     p.Visibility,
	})
	utils.R(c, &pb.Permission{
		Pk:             utils.Int64ToString(res.Pk),
		OperationType:  res.OperationType,
		Resource:       utils.Int64ToString(res.Resource),
		ResourceType:   res.ResourceType,
		PermissionName: res.PermissionName,
		Visibility:     res.Visibility,
	}, err)
}

// DeletePermission
// @Tags open-apis/core
// @summary 删除权限项(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/permission [delete]
func DeletePermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.Permission().Delete(&entity.Permission{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// SelectPermission
// @Tags open-apis/core
// @summary 获取权限项列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param operation_type query int false "操作类型"
// @Param permission_name query string false "权限名称"
// @Param resource_type query int false "资源类型"
// @Param visibility query int false "权限类型"
// @Router	/open-apis/core/rbac/system/permission [get]
func SelectPermission(c *gin.Context) {
	var p PermissionListQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	permissions, page, err := C.S.Permission().SelectWithPageAndMenu(&entity.Permission{
		PermissionName: p.PermissionName,
		OperationType:  p.OperationType,
		ResourceType:   p.ResourceType,
		Visibility:     p.Visibility,
		EnterprisePk:   utils.StringToInt64(p.EnterprisePk),
	}, &dao.Pagination{
		Page: p.Page,
		Size: p.Size,
	})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.PermissionWithDetail, 0)
	if permissions != nil && len(permissions) > 0 {
		for _, v := range permissions {
			if v.ResourceType == 1 && v.Resource != 0 && v.MenuResource != nil {
				result = append(result, &pb.PermissionWithDetail{
					Pk:             utils.Int64ToString(v.Pk),
					OperationType:  v.OperationType,
					Resource:       utils.Int64ToString(v.Resource),
					ResourceType:   v.ResourceType,
					EnterprisePk:   utils.Int64ToString(v.EnterprisePk),
					PermissionName: v.PermissionName,
					Visibility:     v.Visibility,
					MenuResource: &pb.MenuResource{
						Pk:          utils.Int64ToString(v.MenuResource.Pk),
						MenuType:    v.MenuResource.MenuType,
						MenuName:    v.MenuResource.MenuName,
						ResourceKey: v.MenuResource.ResourceKey,
						ParentPk:    utils.Int64ToString(v.MenuResource.ParentPk),
						Icon:        v.MenuResource.Icon,
						Sort:        v.MenuResource.Sort,
					},
				})
			} else if v.ResourceType == 2 && v.Resource != 0 && v.PageResource != nil {
				result = append(result, &pb.PermissionWithDetail{
					Pk:             utils.Int64ToString(v.Pk),
					OperationType:  v.OperationType,
					Resource:       utils.Int64ToString(v.Resource),
					ResourceType:   v.ResourceType,
					EnterprisePk:   utils.Int64ToString(v.EnterprisePk),
					PermissionName: v.PermissionName,
					Visibility:     v.Visibility,
					PageResource: &pb.PageResource{
						Pk:            utils.Int64ToString(v.PageResource.Pk),
						PagePath:      v.PageResource.PagePath,
						Component:     v.PageResource.Component,
						ComponentName: v.PageResource.ComponentName,
						IsCache:       v.PageResource.IsCache,
						PageName:      v.PageResource.PageName,
					},
				})
			} else {
				result = append(result, &pb.PermissionWithDetail{
					Pk:             utils.Int64ToString(v.Pk),
					OperationType:  v.OperationType,
					Resource:       utils.Int64ToString(v.Resource),
					ResourceType:   v.ResourceType,
					EnterprisePk:   utils.Int64ToString(v.EnterprisePk),
					PermissionName: v.PermissionName,
					Visibility:     v.Visibility,
				})
			}
		}
	}
	utils.R(c, &pb.Resp{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Rows:  result,
	}, err)
}

// SelectPermissionGroup
// @Tags open-apis/core
// @summary 获取权限组列表接口(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param group_name query string false "组名"
// @Param group_type query int false "页数"
// @Router	/open-apis/core/rbac/system/permission/group [get]
func SelectPermissionGroup(c *gin.Context) {
	var p PermissionGroupQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	permissions, page, err := C.S.PermissionGroup().SelectPermissionGroupWithPermission(&model.PermissionGroupModel{
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
		GroupName:    p.GroupName,
		GroupType:    p.GroupType,
	}, &dao.Pagination{
		Page: p.Page,
		Size: p.Size,
	})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.PermissionGroupWithPermission, 0)
	if permissions != nil && len(permissions) > 0 {
		for _, v := range permissions {
			rows := make([]*pb.PermissionForPermissionGroup, 0)
			if v.PermissionGroupPermissionModel != nil && len(v.PermissionGroupPermissionModel) > 0 {
				for _, z := range v.PermissionGroupPermissionModel {
					rows = append(rows, &pb.PermissionForPermissionGroup{
						Pk:             utils.Int64ToString(z.Permission.Pk),
						OperationType:  z.Permission.OperationType,
						Resource:       utils.Int64ToString(z.Permission.Resource),
						ResourceType:   z.Permission.ResourceType,
						EnterprisePk:   utils.Int64ToString(z.Permission.EnterprisePk),
						PermissionName: z.Permission.PermissionName,
						RelationPk:     utils.Int64ToString(z.Pk),
					})
				}
			}
			result = append(result, &pb.PermissionGroupWithPermission{
				Pk:           utils.Int64ToString(v.Pk),
				EnterprisePk: utils.Int64ToString(v.EnterprisePk),
				GroupName:    v.GroupName,
				Remark:       v.Remark,
				GroupType:    v.GroupType,
				Rows:         rows,
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

// CreatePermissionGroup
// @Tags open-apis/core
// @summary 新增权限组(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreatePermissionGroupReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/permission/group [post]
func CreatePermissionGroup(c *gin.Context) {
	var p CreatePermissionGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if p.EnterprisePk != "1" && p.EnterprisePk != "" {
		enterpriseMessage, err := C.S.Enterprise().FindByPk(&entity.Enterprise{
			Pk: utils.StringToInt64(p.EnterprisePk),
		})
		if err == nil && enterpriseMessage != nil {
			if enterpriseMessage.Type == 2 {
				p.GroupType = 2
			} else if enterpriseMessage.Type == 1 {
				p.GroupType = 4
			} else if enterpriseMessage.Type == 3 {
				p.GroupType = 3
			}
		}
	}
	res, err := C.S.PermissionGroup().Create(&entity.PermissionGroup{
		GroupName:    p.GroupName,
		Remark:       p.Remark,
		GroupType:    p.GroupType,
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
	})
	if err != nil {
		utils.FR(c, err)
	}
	utils.R(c, res, err)
}

// UpdaterPermissionGroup
// @Tags open-apis/core
// @summary 修改权限组(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdaterPermissionGroupReq true "请求参数体"
// @Router /open-apis/core/rbac/system/permission/group [put]
func UpdaterPermissionGroup(c *gin.Context) {
	var p UpdaterPermissionGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.PermissionGroup().Update(&entity.PermissionGroup{
		Pk:        utils.StringToInt64(p.Pk),
		GroupName: p.GroupName,
		Remark:    p.Remark,
		GroupType: p.GroupType,
	})
	utils.R(c, res, err)
}

// DeletePermissionGroup
// @Tags open-apis/core
// @summary 删除权限组(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/permission/group [delete]
func DeletePermissionGroup(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.PermissionGroup().Delete(&entity.PermissionGroup{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// CreatePermissionGroupPermission
// @Tags open-apis/core
// @summary 新增权限组关联(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreatePermissionGroupPermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/system/permission/group/permission [post]
func CreatePermissionGroupPermission(c *gin.Context) {
	var p CreatePermissionGroupPermissionReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if _, err := C.S.PermissionGroup().FindByPk(&entity.PermissionGroup{Pk: utils.StringToInt64(p.PermissionGroupPk)}); err != nil {
		utils.FR(c, errors.New("不存在该权限组！"))
	}
	if _, err := C.S.Permission().FindByPk(&entity.Permission{Pk: utils.StringToInt64(p.PermissionPk)}); err != nil {
		utils.FR(c, errors.New("不存在该权限！"))
	}
	res, err := C.S.PermissionGroupPermission().Create(&entity.PermissionGroupPermission{
		PermissionPk:      utils.StringToInt64(p.PermissionPk),
		PermissionGroupPk: utils.StringToInt64(p.PermissionGroupPk),
	})
	utils.R(c, res, err)
}

// DeletePermissionGroupPermission
// @Tags open-apis/core
// @summary 删除权限组关联关系(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/system/permission/group/permission [delete]
func DeletePermissionGroupPermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.PermissionGroupPermission().Delete(&entity.PermissionGroupPermission{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}
