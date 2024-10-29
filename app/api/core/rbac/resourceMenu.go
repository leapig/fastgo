package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// CreateMenuResource
// @Tags open-apis/core
// @summary 新增菜单信息(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateMenuResourceReq true "请求参数体"
// @Router	/open-apis/core/rbac/resource/menu [post]
func CreateMenuResource(c *gin.Context) {
	var p CreateMenuResourceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.MenuResource().Create(&entity.MenuResource{
		MenuType:    p.MenuType,
		MenuName:    p.MenuName,
		ResourceKey: p.ResourceKey,
		ParentPk:    utils.StringToInt64(p.ParentPk),
		Icon:        p.Icon,
		Sort:        p.Sort,
		Path:        p.Path,
	})
	utils.R(c, res, err)
}

// UpdateMenuResource
// @Tags open-apis/core
// @summary 修改菜单信息(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body UpdateMenuResourceReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/menu [put]
func UpdateMenuResource(c *gin.Context) {
	var p UpdateMenuResourceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.MenuResource().Update(&entity.MenuResource{
		Pk:          utils.StringToInt64(p.Pk),
		MenuType:    p.MenuType,
		MenuName:    p.MenuName,
		ResourceKey: p.ResourceKey,
		Icon:        p.Icon,
		Sort:        p.Sort,
		Path:        p.Path,
	})
	utils.R(c, res, err)
}

// DeleteMenuResource
// @Tags open-apis/core
// @summary 删除菜单信息(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/menu [delete]
func DeleteMenuResource(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.MenuResource().Delete(&entity.MenuResource{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}

// SelectAllMenuWithDetail
// @Tags open-apis/core
// @summary 获取菜单信息列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param menu_name query string false "菜单名称"
// @Router	/open-apis/core/rbac/resource/menu [get]
func SelectAllMenuWithDetail(c *gin.Context) {
	var p MenuResourceQueryBody
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.MenuResource().SelectAllMenuWithDetail(&entity.MenuResource{MenuName: p.MenuName})
	utils.R(c, &pb.Resp{
		Rows: res,
	}, err)
}
