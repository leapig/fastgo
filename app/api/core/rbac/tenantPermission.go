package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// GetTenantPermission
// @Tags open-apis/core
// @summary 获取权限项列表方法(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param operation_type query int false "操作类型"
// @Param permission_name query string false "权限名称"
// @Param resource_type query int false "资源类型"
// @Router	/open-apis/core/rbac/tenant/permission [get]
func GetTenantPermission(c *gin.Context) {
	var p GetPermissionReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	permissions, page, err := C.S.Permission().SelectWithPageAndMenu(&entity.Permission{
		PermissionName: p.PermissionName,
		OperationType:  p.OperationType,
		ResourceType:   p.ResourceType,
		EnterprisePk:   utils.StringToInt64(c.GetString("tenant")),
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
