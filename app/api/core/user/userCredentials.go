package user

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/aliyun"
	utils "github.com/leapig/fastgo/app/library/util"

	C "github.com/leapig/fastgo/app/service"
)

// GetUserCredentials
// @Tags open-apis/core
// @summary 获取用户证件信息(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/user/credentials [get]
func GetUserCredentials(c *gin.Context) {
	res, err := C.S.UserCredentials().FindListCardByUserPk(&entity.UserCredentials{UserPk: utils.StringToInt64(c.GetString("user"))})
	if err != nil {
		utils.FR(c, err)
	}
	rows := make([]*pb.UserCredentials, 0)
	for _, value := range res {
		if value.FrontFileName != "" {
			value.FrontFileName = aliyun.Oss().PathToUrl(value.FrontFileName)
		}
		if value.BackFileName != "" {
			value.BackFileName = aliyun.Oss().PathToUrl(value.BackFileName)
		}
		rows = append(rows, &pb.UserCredentials{
			Pk:            utils.Int64ToString(value.Pk),
			Type:          value.Type,
			Serial:        value.Serial,
			Name:          value.Name,
			FrontFileName: value.FrontFileName,
			BackFileName:  value.BackFileName,
		})
	}
	utils.R(c, &pb.Resp{
		Rows: rows,
	}, err)
}

// CreateUserCredentials
// @Tags open-apis/core
// @summary 新增用户证件(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body CreateUserCredentialsReq true "请求参数体"
// @Router	/open-apis/core/user/credentials [post]
func CreateUserCredentials(c *gin.Context) {
	var p CreateUserCredentialsReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	_, err := C.S.UserCredentials().Create(&entity.UserCredentials{
		UserPk:        utils.StringToInt64(c.GetString("user")),
		Serial:        p.Serial,
		Type:          p.Type,
		Name:          p.Name,
		FrontFileName: p.FrontFileName,
		BackFileName:  p.BackFileName,
	})
	// 实名认证
	if p.Type == 1 {
		birthday, gender := utils.ParseIdCard(p.Serial)
		_ = C.S.User().UpdateBaseInfo(&entity.User{
			Pk:       utils.StringToInt64(c.GetString("user")),
			Birthday: &birthday,
			Gender:   gender,
		})
		_, _ = C.S.EnterpriseUser().Update(&entity.EnterpriseUser{
			UserPk:   utils.StringToInt64(c.GetString("user")),
			Birthday: &birthday,
			Gender:   gender,
		})
	}
	utils.R(c, nil, err)
}

// DeleteUserCredentials
// @Tags open-apis/core
// @summary 删除用户证件(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body DeleteUserCredentialsReq true "请求参数体"
// @Router	/open-apis/core/user/credentials [delete]
func DeleteUserCredentials(c *gin.Context) {
	var p DeleteUserCredentialsReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.UserCredentials().Delete(&entity.UserCredentials{Pk: utils.StringToInt64(p.Pk)})
	utils.R(c, nil, err)
}
