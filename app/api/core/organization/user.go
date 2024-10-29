package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// GetUserBaseInfo
// @Tags 组织架构
// @summary 获取用户基础信息(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/organization/user/base_info [get]
func GetUserBaseInfo(c *gin.Context) {
	if res, err := C.S.User().FindUser(&entity.User{Pk: utils.StringToInt64(c.GetString("user"))}); err == nil {
		idCard := ""
		for _, card := range res.UserCredentials {
			if card.Type == 1 {
				idCard = card.Serial
				break
			}
		}
		lastLiveAt := ""
		if res.UserLiveness != nil && len(res.UserLiveness) > 0 {
			data := res.UserLiveness[0]
			lastLiveAt = utils.TimeToString(data.CreatedAt)
		}
		isRealName := 2
		if res.UserRealNameAuthenticationLog != nil && len(res.UserRealNameAuthenticationLog) > 0 {
			isRealName = 1
		}
		utils.R(c, &pb.User{
			Pk:         utils.Int64ToString(res.Pk),
			Name:       res.Name,
			Phone:      res.Phone,
			IdCard:     idCard,
			LiveCity:   "",
			IsRealName: int32(isRealName),
			CreateAt:   utils.TimeToString(res.CreatedAt),
			LastLiveAt: lastLiveAt,
		}, err)
	} else {
		utils.FR(c, err)
	}
}

// GetUserInfo
// @Tags 组织架构
// @summary 获取用户详细信息(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/organization/user/info [get]
func GetUserInfo(c *gin.Context) {
	if c.GetString("tenant") == "1" {
		GetUserBaseInfo(c)
	} else {
		res, err := C.S.EnterpriseUser().FindByUserPkAndEnterprisePk(&entity.EnterpriseUser{
			UserPk:       utils.StringToInt64(c.GetString("user")),
			EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
		})
		if err != nil {
			utils.FR(c, err)
		}
		departmentPk := make([]string, 0)
		departmentRows := make([]*pb.EnterpriseUserDepartment, 0)
		if dp, dpErr := C.S.Member().List(&entity.Member{UserPk: res.Pk}); dpErr == nil {
			for _, value := range dp {
				departmentPk = append(departmentPk, utils.Int64ToString(value.DepartmentPk))
				row := &pb.EnterpriseUserDepartment{
					Pk:           utils.Int64ToString(value.Pk),
					DepartmentPk: utils.Int64ToString(value.DepartmentPk),
				}
				if departmentMessage, dpErr1 := C.S.Department().FindByPk(&entity.Department{Pk: value.DepartmentPk}); dpErr1 == nil {
					row.Name = departmentMessage.Name
				}
				departmentRows = append(departmentRows, row)
			}
		}
		eu := &pb.EnterpriseUser{
			Pk:                    utils.Int64ToString(res.Pk),
			Name:                  res.Name,
			EnterprisePk:          utils.Int64ToString(res.EnterprisePk),
			UserPk:                utils.Int64ToString(res.UserPk),
			Gender:                res.Gender,
			Phone:                 res.Phone,
			JobTitle:              res.JobTitle,
			JobNumber:             res.JobNumber,
			Nation:                res.Nation,
			Education:             res.Education,
			EmergencyContact:      res.EmergencyContact,
			EmergencyContactPhone: res.EmergencyContactPhone,
			TermOfContract:        res.TermOfContract,
			Address:               res.Address,
			Weight:                utils.Float64ToString(res.Weight),
			Height:                utils.Float64ToString(res.Height),
			DepartmentPk:          departmentPk,
			DepartmentRow:         departmentRows,
		}
		if res.Birthday != nil {
			eu.Birthday = utils.DateTimeToString(*res.Birthday)
		}
		if res.EntryTime != nil {
			eu.EntryTime = utils.TimeToString(*res.EntryTime)
		}
		utils.R(c, eu, err)
	}
}
