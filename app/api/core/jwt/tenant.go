package jwt

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// GetList
// @Tags jwt模块
// @summary 获取用户单位列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/jwt/tenant/list [get]
func GetList(c *gin.Context) {
	if res, err := C.S.Enterprise().ReadUserEnterprise(utils.StringToInt64(c.GetString("user"))); err == nil {
		rows := make([]*pb.Enterprise, 0)
		for _, value := range res {
			rows = append(rows, &pb.Enterprise{
				Pk:    utils.Int64ToString(value.Pk),
				Name:  value.Name,
				Cover: value.Cover,
				Type:  value.Type,
			})
		}
		utils.R(c, &pb.Resp{Rows: rows}, err)
	} else {
		utils.FR(c, err)
	}
}

// GetMenu
// @Tags jwt模块
// @summary 获取菜单(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/jwt/tenant/menu [get]
func GetMenu(c *gin.Context) {
	res, err := helper.RS.HGet("scope_"+c.GetString("tenant"), c.GetString("user")+"_menu")
	var menu []map[string]interface{}
	_ = json.Unmarshal(res.([]byte), &menu)
	utils.R(c, menu, err)
}

// GetPage
// @Tags jwt模块
// @summary 获取页面(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/jwt/tenant/page [get]
func GetPage(c *gin.Context) {
	res, err := helper.RS.HGet("scope_"+c.GetString("tenant"), c.GetString("user")+"_page")
	var page []map[string]interface{}
	_ = json.Unmarshal(res.([]byte), &page)
	utils.R(c, page, err)
}
