package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// CreateInterfaceResource
// @Tags open-apis/core
// @summary 新增接口信息(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateInterfaceResourceReq true "请求参数体"
// @Router	/open-apis/core/rbac/resource/interface [post]
func CreateInterfaceResource(c *gin.Context) {
	var p CreateInterfaceResourceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.InterfaceResource().Create(&entity.InterfaceResource{
		InterfaceUrl:  p.InterfaceUrl,
		InterfaceName: p.InterfaceName,
		InterfaceWay:  p.InterfaceWay,
		InterfaceKey:  p.InterfaceKey,
	})
	utils.R(c, res, err)
}

// UpdateInterfaceResource
// @Tags open-apis/core
// @summary 修改接口信息(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdateInterfaceResourceReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/interface [put]
func UpdateInterfaceResource(c *gin.Context) {
	var p UpdateInterfaceResourceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.InterfaceResource().Update(&entity.InterfaceResource{
		Pk:            utils.StringToInt64(p.Pk),
		InterfaceUrl:  p.InterfaceUrl,
		InterfaceName: p.InterfaceName,
		InterfaceWay:  p.InterfaceWay,
		InterfaceKey:  p.InterfaceKey,
	})
	utils.R(c, res, err)
}

// DeleteInterfaceResource
// @Tags open-apis/core
// @summary 删除接口信息(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/interface [delete]
func DeleteInterfaceResource(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.InterfaceResource().Delete(&entity.InterfaceResource{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// SelectInterfaceResource
// @Tags open-apis/core
// @summary 获取接口信息(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param interface_way query string false "接口方法"
// @Param interface_url query string false "接口路径"
// @Param interface_name query string false "接口名"
// @Param interface_key query string false "接口标识符"
// @Router	/open-apis/core/rbac/resource/interface [get]
func SelectInterfaceResource(c *gin.Context) {
	var p InterfaceResourceQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	message, page, err := C.S.InterfaceResource().Select(
		&entity.InterfaceResource{
			InterfaceUrl:  p.InterfaceUrl,
			InterfaceName: p.InterfaceName,
			InterfaceWay:  p.InterfaceWay,
			InterfaceKey:  p.InterfaceKey,
		},
		&dao.Pagination{
			Page: p.Page,
			Size: p.Size,
		})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.InterfaceResource, 0)
	if message != nil && len(message) > 0 {
		for _, v := range message {
			result = append(result, &pb.InterfaceResource{
				Pk:            utils.Int64ToString(v.Pk),
				InterfaceUrl:  v.InterfaceUrl,
				InterfaceName: v.InterfaceName,
				InterfaceWay:  v.InterfaceWay,
				InterfaceKey:  v.InterfaceKey,
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
