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

// SelectEnterprisePermissionGroup
// @Tags open-apis/core
// @summary 获取权限组列表(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param enterprise_pk query string true "租户Pk"
// @Param group_name query string false "组名"
// @Router	/open-apis/core/rbac/tenant/permission/group [get]
func SelectEnterprisePermissionGroup(c *gin.Context) {
	var p GetPermissionGroupReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	permissions, page, err := C.S.PermissionGroup().SelectPermissionGroupWithPermissionForEnterprise(&model.PermissionGroupModel{
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
		GroupName:    p.GroupName,
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

// CreateEnterprisePermissionGroup
// @Tags open-apis/core
// @summary 新增权限组(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostPermissionGroupReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/permission/group [post]
func CreateEnterprisePermissionGroup(c *gin.Context) {
	var p PostPermissionGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	var GroupType int32
	if c.GetString("tenant") != "1" && c.GetString("tenant") != "" {
		enterpriseMessage, err := C.S.Enterprise().FindByPk(&entity.Enterprise{
			Pk: utils.StringToInt64(c.GetString("tenant")),
		})
		if err == nil && enterpriseMessage != nil {
			if enterpriseMessage.Type == 2 {
				GroupType = 2
			} else if enterpriseMessage.Type == 1 {
				GroupType = 4
			} else if enterpriseMessage.Type == 3 {
				GroupType = 3
			}
		}
	}
	res, err := C.S.PermissionGroup().Create(&entity.PermissionGroup{
		GroupName:    p.GroupName,
		Remark:       p.Remark,
		GroupType:    GroupType,
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	})
	utils.R(c, res, err)
}

// UpdaterEnterprisePermissionGroup
// @Tags open-apis/core
// @summary 修改权限组(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body PutPermissionGroupReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/permission/group [put]
func UpdaterEnterprisePermissionGroup(c *gin.Context) {
	var p PutPermissionGroupReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.PermissionGroup().Update(&entity.PermissionGroup{
		Pk:        utils.StringToInt64(p.Pk),
		GroupName: p.GroupName,
		Remark:    p.Remark,
	})
	utils.R(c, res, err)
}

// DeleteEnterprisePermissionGroup
// @Tags open-apis/core
// @summary 删除权限组(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/permission/group [delete]
func DeleteEnterprisePermissionGroup(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.PermissionGroup().Delete(&entity.PermissionGroup{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// CreateEnterprisePermissionGroupPermission
// @Tags open-apis/core
// @summary 新增权限组关联(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostPermissionGroupPermissionReq true "请求参数体"
// @Router	/open-apis/core/rbac/tenant/permission/group/permission [post]
func CreateEnterprisePermissionGroupPermission(c *gin.Context) {
	var p PostPermissionGroupPermissionReq
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

// DeleteEnterprisePermissionGroupPermission
// @Tags open-apis/core
// @summary 删除权限组关联关系(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/tenant/permission/group/permission [delete]
func DeleteEnterprisePermissionGroupPermission(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.PermissionGroupPermission().Delete(&entity.PermissionGroupPermission{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}
