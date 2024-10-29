package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// CreatePageInterface
// @Tags open-apis/core
// @summary 新增页面资源与接口关联(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreatePageInterfaceReq true "请求参数体"
// @Router	/open-apis/core/rbac/resource/page/interface [post]
func CreatePageInterface(c *gin.Context) {
	var p CreatePageInterfaceReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	if _, err := C.S.PageResource().FindByPk(&entity.PageResource{
		Pk: utils.StringToInt64(p.PagePk),
	}); err != nil {
		utils.FR(c, err)
	}
	if _, err := C.S.InterfaceResource().FindByPk(&entity.InterfaceResource{
		Pk: utils.StringToInt64(p.InterfacePk),
	}); err != nil {
		utils.FR(c, err)
	}
	res, err := C.S.PageInterface().Create(&entity.PageInterface{
		PagePk:        utils.StringToInt64(p.PagePk),
		InterfacePk:   utils.StringToInt64(p.InterfacePk),
		OperationType: p.OperationType,
	})
	utils.R(c, res, err)
}

// DeletePageInterface
// @Tags open-apis/core
// @summary 删除页面资源与接口关联关系(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/rbac/resource/page/interface [delete]
func DeletePageInterface(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.PageInterface().Delete(&entity.PageInterface{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}
