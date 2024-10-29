package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// CreatePageResource
// @Tags open-apis/core
// @summary 新增页面资源(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreatePageResourceReq true "请求参数体"
// @Router	/open-apis/core/rbac/resource/page [post]
func CreatePageResource(c *gin.Context) {
	var p CreatePageResourceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.PageResource().Create(&entity.PageResource{
		PagePath:      p.PagePath,
		Component:     p.Component,
		ComponentName: p.ComponentName,
		IsCache:       p.IsCache,
		PageName:      p.PageName,
		PageType:      p.PageType,
	})
	utils.R(c, res, err)
}

// UpdatePageResource
// @Tags open-apis/core
// @summary 修改页面资源(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdatePageResourceReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/page [put]
func UpdatePageResource(c *gin.Context) {
	var p UpdatePageResourceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.PageResource().Update(&entity.PageResource{
		Pk:            utils.StringToInt64(p.Pk),
		PagePath:      p.PagePath,
		Component:     p.Component,
		ComponentName: p.ComponentName,
		IsCache:       p.IsCache,
		PageName:      p.PageName,
		PageType:      p.PageType,
	})
	utils.R(c, res, err)
}

// DeletePageResource
// @Tags open-apis/core
// @summary 删除页面资源(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/page [delete]
func DeletePageResource(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
		return
	}
	err := C.S.PageResource().Delete(&entity.PageResource{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// SelectPageInterfaceDetailMessage
// @Tags open-apis/core
// @summary 获取页面及关联接口(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param page_type query int false "页面类型:1平台 2 小程序"
// @Param page_path query string false "路由地址"
// @Param component_name query string false "组件名"
// @Param page_name query string false "页面名称"
// @Router	/open-apis/core/rbac/resource/page [get]
func SelectPageInterfaceDetailMessage(c *gin.Context) {
	var p PageResourceQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	message, page, err := C.S.PageResource().SelectPageInterfaceDetailMessage(
		&entity.PageResource{
			PagePath:      p.PagePath,
			ComponentName: p.ComponentName,
			PageName:      p.PageName,
			PageType:      p.PageType,
		},
		&dao.Pagination{
			Page: p.Page,
			Size: p.Size,
		})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.PageResourceWithInterfaceMessage, 0)
	if message != nil && len(message) > 0 {
		for _, v := range message {
			rows := make([]*pb.PageInterfaceMessage, 0)
			if v.PageInterfaceModel != nil && len(v.PageInterfaceModel) > 0 {
				for _, z := range v.PageInterfaceModel {
					rows = append(rows, &pb.PageInterfaceMessage{
						Pk:            utils.Int64ToString(z.Pk),
						InterfacePk:   utils.Int64ToString(z.InterfaceResource.Pk),
						OperationType: z.OperationType,
						InterfaceKey:  z.InterfaceResource.InterfaceKey,
						InterfaceName: z.InterfaceResource.InterfaceName,
						InterfaceWay:  z.InterfaceResource.InterfaceWay,
						InterfaceUrl:  z.InterfaceResource.InterfaceUrl,
					})
				}
			}
			result = append(result, &pb.PageResourceWithInterfaceMessage{
				Pk:            utils.Int64ToString(v.Pk),
				PagePath:      v.PagePath,
				Component:     v.Component,
				ComponentName: v.ComponentName,
				IsCache:       v.IsCache,
				PageType:      v.PageType,
				PageName:      v.PageName,
				Rows:          rows,
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
