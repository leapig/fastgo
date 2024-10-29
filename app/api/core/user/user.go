package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// GetUserList
// @Tags open-apis/core
// @summary 用户列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param name query string false "姓名"
// @Param phone query string false "手机"
// @Router	/open-apis/core/user/list [get]
func GetUserList(c *gin.Context) {
	var p GetUsersReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if res, count, err := C.S.User().SelectUser(&entity.User{Name: p.Name, Phone: p.Phone}, &entity.Pagination{Page: p.Page, Size: p.Size}); err == nil {
		rows := make([]*pb.User, 0)
		for _, value := range res {
			idCard := ""
			for _, card := range value.UserCredentials {
				if card.Type == 1 {
					idCard = card.Serial
					break
				}
			}
			lastLiveAt := ""
			if value.UserLiveness != nil && len(value.UserLiveness) > 0 {
				data := value.UserLiveness[0]
				lastLiveAt = utils.TimeToString(data.CreatedAt)
			}
			isRealName := 2
			if value.UserRealNameAuthenticationLog != nil && len(value.UserRealNameAuthenticationLog) > 0 {
				isRealName = 1
			}
			rows = append(rows, &pb.User{
				Pk:         utils.Int64ToString(value.Pk),
				Name:       value.Name,
				Phone:      value.Phone,
				IdCard:     idCard,
				LiveCity:   "",
				IsRealName: int32(isRealName),
				CreateAt:   utils.TimeToString(value.CreatedAt),
				LastLiveAt: lastLiveAt,
			})
		}
		utils.R(c, &pb.Resp{
			Size:  p.Size,
			Page:  p.Page,
			Total: int32(count),
			Rows:  rows,
		}, err)
	} else {
		utils.FR(c, err)
	}
}

// GetUserAccount
// @Tags open-apis/core
// @summary 账号列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param name query string false "姓名"
// @Param phone query string false "手机"
// @Router	/open-apis/core/user/account [get]
func GetUserAccount(c *gin.Context) {
	var p GetUsersReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if res, count, err := C.S.User().SelectAccountPage(&entity.User{
		Name:  p.Name,
		Phone: p.Phone,
	}, &entity.Pagination{
		Page: p.Page,
		Size: p.Size,
	}); err == nil {
		rows := make([]*pb.UserAccount, 0)
		for _, value := range res {
			wxOfficialAccountsOpenId := ""
			wxMiniProgramOpenId := ""
			for _, userClient := range value.UserClientList {
				wxOfficialAccountsOpenId = C.S.User().FindOaOpenid(userClient.WxUnionid)
				switch userClient.ClientType {
				case 1:
					wxOfficialAccountsOpenId = userClient.OpenId
				case 2:
					wxMiniProgramOpenId = userClient.OpenId
				}
			}
			rows = append(rows, &pb.UserAccount{
				Pk:                       utils.Int64ToString(value.Pk),
				Name:                     value.Name,
				Phone:                    value.Phone,
				WxOfficialAccountsOpenId: wxOfficialAccountsOpenId,
				WxMiniProgramOpenId:      wxMiniProgramOpenId,
				CreateAt:                 utils.TimeToString(value.CreatedAt),
			})
		}
		utils.R(c, &pb.Resp{
			Page:  p.Page,
			Size:  p.Size,
			Total: int32(count),
			Rows:  rows,
		}, err)
	} else {
		utils.FR(c, err)
	}
}

// GetUserMember
// @Tags open-apis/core
// @summary 租户的账号列表(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param name query string false "姓名"
// @Param phone query string false "手机"
// @Router	/open-apis/core/user/member [get]
func GetUserMember(c *gin.Context) {
	var p GetUsersReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	if res, count, err := C.S.User().SelectEnterpriseAccountPage(&model.QfqzAccountModelForEnterprise{
		Name:         p.Name,
		Phone:        p.Phone,
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	}, &entity.Pagination{
		Page: p.Page,
		Size: p.Size,
	}); err == nil {
		rows := make([]*pb.UserAccount, 0)
		for _, value := range res {
			wxOfficialAccountsOpenId := ""
			wxMiniProgramOpenId := ""
			for _, userClient := range value.UserClientList {
				switch userClient.ClientType {
				case 1:
					wxOfficialAccountsOpenId = userClient.OpenId
				case 2:
					wxMiniProgramOpenId = userClient.OpenId
				}
			}
			rows = append(rows, &pb.UserAccount{
				Pk:                       utils.Int64ToString(value.Pk),
				Name:                     value.Name,
				Phone:                    value.Phone,
				WxOfficialAccountsOpenId: wxOfficialAccountsOpenId,
				WxMiniProgramOpenId:      wxMiniProgramOpenId,
				CreateAt:                 utils.TimeToString(value.CreatedAt),
				UserPk:                   utils.Int64ToString(value.UserPk),
			})
		}
		utils.R(c, &pb.Resp{
			Page:  p.Page,
			Size:  p.Size,
			Total: int32(count),
			Rows:  rows,
		}, err)
	} else {
		utils.FR(c, err)
	}
}

// PostUserAccount
// @Tags open-apis/core
// @summary 新增账号(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body Account true "请求参数体"
// @Router	/open-apis/core/user/account [post]
func PostUserAccount(c *gin.Context) {
	var p Account
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.User().Save(&entity.User{
		Name:  p.Name,
		Phone: p.Phone,
	})
	utils.R(c, res, err)
}

// PutUserPhone
// @Tags open-apis/core
// @summary 更换用户手机号(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PhoneChangeReq true "请求参数体"
// @Router	/open-apis/core/user/phone [put]
func PutUserPhone(c *gin.Context) {
	var p PhoneChangeReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.User().UpdatePhone(&entity.User{
		Phone: p.Phone,
		Pk:    utils.StringToInt64(p.UserPk),
	})
	utils.R(c, nil, err)
}

// PutUserClient
// @Tags open-apis/core
// @summary 解绑用户小程序(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body _UserPk true "请求参数体"
// @Router	/open-apis/core/user/client [put]
func PutUserClient(c *gin.Context) {
	var p _UserPk
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.UserClient().DeleteByUserPkAndClientType(&entity.UserClient{
		UserPk:     utils.StringToInt64(p.UserPk),
		ClientType: 2,
	})
	utils.R(c, nil, err)
}

// SetUserBaseInfo
// @Tags open-apis/core
// @summary 修改用户基础信息(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body SetUserBaseInfoReq true "请求参数体"
// @Router	/open-apis/core/user/base_info [put]
func SetUserBaseInfo(c *gin.Context) {
	var p SetUserBaseInfoReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	if res, err := C.S.User().FindUser(&entity.User{Pk: utils.StringToInt64(c.GetString("user"))}); err == nil {
		flag := 1
		for _, card := range res.UserCredentials {
			if card.Type == 1 {
				if card.Serial == p.IdCard {
					flag = 2
				}
				break
			}
		}
		isRealName := 2
		if res.UserRealNameAuthenticationLog != nil && len(res.UserRealNameAuthenticationLog) > 0 {
			isRealName = 1
		}
		if flag == 1 {
			if isRealName == 1 {
				utils.FR(c, errors.New("实名无法修改身份证信息"))
			} else {
				//修改卡包身份证
				C.S.UserCredentials().Create(&entity.UserCredentials{
					UserPk: utils.StringToInt64(c.GetString("user")),
					Serial: p.IdCard,
					Name:   p.Name,
					Type:   1,
				})
			}
		}
		us := &entity.User{
			Pk:     utils.StringToInt64(c.GetString("user")),
			Name:   p.Name,
			Gender: p.Gender,
		}
		if birthday, birErr := utils.StringToDateTime(p.Birthday); birErr == nil {
			us.Birthday = &birthday
		}
		us.Weight = utils.StringToFloat64(p.Weight)
		us.Height = utils.StringToFloat64(p.Height)
		err := C.S.User().UpdateBaseInfo(us)
		utils.R(c, nil, err)
	} else {
		utils.FR(c, err)
	}
}
