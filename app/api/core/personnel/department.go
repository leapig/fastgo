package personnel

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

// SelectAllDepartment
// @Tags open-apis/core
// @summary 获取租户部门列表
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/personnel/department [get]
func SelectAllDepartment(c *gin.Context) {
	rows, err := C.S.Department().FindAllDepartByEnterprisePk(&entity.Department{
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	})
	utils.R(c, &pb.Resp{
		Rows: rows,
	}, err)
}

// CreateDepartment
// @Tags open-apis/core
// @summary 新增部门
// @Accept json
// @Produce  json
// @Param params body PostDepartmentReq true "请求参数体"
// @Security bearer
// @Router	/open-apis/core/personnel/department [post]
func CreateDepartment(c *gin.Context) {
	var p PostDepartmentReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	res, err := C.S.Department().Create(&entity.Department{
		Name:         p.Name,
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
		ParentPk:     utils.StringToInt64(p.ParentPk),
	})
	utils.R(c, res, err)
}

// UpdateDepartment
// @Tags open-apis/core
// @summary 修改部门名称
// @Accept json
// @Produce json
// @Security bearer
// @Param params body PutDepartmentReq true "请求参数体"
// @Router /open-apis/core/personnel/department [put]
func UpdateDepartment(c *gin.Context) {
	var p PutDepartmentReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	res, err := C.S.Department().Update(&entity.Department{
		Pk:           utils.StringToInt64(p.Pk),
		Name:         p.Name,
		ParentPk:     utils.StringToInt64(p.ParentPk),
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	}, p.Leader)
	utils.R(c, res, err)
}

// GetDepartmentMember
// @Tags open-apis/core
// @summary 获取部门人员列表接口(TenantAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param department_pk query string true "部门pk"
// @Param job_number query string false "工号"
// @Param job_title query string false "职称"
// @Param is_leader query int false "是否主管"
// @Param name query string false "员工姓名"
// @Param phone query string false "员工手机"
// @Router	/open-apis/core/personnel/department/member [get]
func GetDepartmentMember(c *gin.Context) {
	var p GetDepartmentMemberReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	if res, count, err := C.S.Member().Select(&model.Member{
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
		DepartmentPk: utils.StringToInt64(p.DepartmentPk),
		IsLeader:     p.IsLeader,
		JobNumber:    utils.StringToInt64(p.JobNumber),
		JobTitle:     p.JobTitle,
		EnterpriseUser: &entity.EnterpriseUser{
			Name:  p.Name,
			Phone: p.Phone,
		},
	}, &entity.Pagination{
		Page: p.Page,
		Size: p.Size,
	}); err == nil && res != nil && len(res) > 0 {
		rows := make([]*pb.DepartMember, 0)
		for _, value := range res {
			dm := &pb.DepartMember{
				Pk:       utils.Int64ToString(value.Pk),
				UserPk:   utils.Int64ToString(value.UserPk),
				IsLeader: value.IsLeader,
				IsMain:   value.IsMain,
			}
			if value.EnterpriseUser != nil {
				dm.Name = value.EnterpriseUser.Name
				dm.Phone = value.EnterpriseUser.Phone
				dm.JobTitle = value.EnterpriseUser.JobTitle
				dm.JobNumber = value.EnterpriseUser.JobNumber
			}
			rows = append(rows, dm)
		}
		utils.R(c, &pb.Resp{
			Page:  p.Page,
			Size:  p.Size,
			Total: count.Total,
			Rows:  rows,
		}, err)
	} else {
		utils.FR(c, err)
	}
}

// DeleteDepartment
// @Tags open-apis/core
// @summary 删除企业部门(TenantAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param pk query string true "部门pk"
// @Router /open-apis/core/personnel/department [delete]
func DeleteDepartment(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.Department().Delete(&entity.Department{
		Pk:           utils.StringToInt64(p.Pk),
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
	})
	utils.R(c, nil, err)
}

// SelectEnterpriseUser
// @Tags open-apis/core
// @summary 获取租户人员信息_人员档案(tenant_access_token)
// @Accept json
// @Produce  json
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param name query string false "姓名"
// @Param phone query string false "手机号"
// @Param gender query int false "性别"
// @Param job_title query string false "职称"
// @Param job_number query string false "工号"
// @Router	/open-apis/core/personnel/enterprise/user [get]
func SelectEnterpriseUser(c *gin.Context) {
	var p GetEnterpriseUserReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	res, page, err := C.S.EnterpriseUser().Select(&entity.EnterpriseUser{
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
		Name:         p.Name,
		Phone:        p.Phone,
		JobTitle:     p.JobTitle,
		JobNumber:    p.JobNumber,
		Gender:       p.Gender,
		Status:       p.Status,
	}, &dao.Pagination{Page: p.Page, Size: p.Size})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.EnterpriseUser, 0)
	for _, row := range res {
		departmentPk := make([]string, 0)
		if dp, dpErr := C.S.Member().List(&entity.Member{UserPk: row.Pk}); dpErr == nil {
			for _, value := range dp {
				departmentPk = append(departmentPk, utils.Int64ToString(value.DepartmentPk))
			}
		}
		z := &pb.EnterpriseUser{
			Pk:                    utils.Int64ToString(row.Pk),
			EnterprisePk:          utils.Int64ToString(row.EnterprisePk),
			UserPk:                utils.Int64ToString(row.UserPk),
			Name:                  row.Name,
			Phone:                 row.Phone,
			Height:                utils.Float64ToString(row.Height),
			Weight:                utils.Float64ToString(row.Weight),
			Gender:                row.Gender,
			JobTitle:              row.JobTitle,
			JobNumber:             row.JobNumber,
			Nation:                row.Nation,
			Education:             row.Education,
			EmergencyContact:      row.EmergencyContact,
			EmergencyContactPhone: row.EmergencyContactPhone,
			TermOfContract:        row.TermOfContract,
			Address:               row.Address,
			Status:                row.Status,
			DepartmentPk:          departmentPk,
		}
		//查询人员
		if row.EntryTime != nil {
			z.EntryTime = utils.DateTimeToString(*row.EntryTime)
		}
		if row.Birthday != nil {
			z.Birthday = utils.DateTimeToString(*row.Birthday)
		}
		result = append(result, z)
	}
	utils.R(c, &pb.Resp{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Rows:  result,
	}, err)
}

// SelectInServiceEnterpriseUser
// @Tags open-apis/core
// @summary 获取在职租户人员信息_人员管理查询接口(tenant_access_token)
// @Accept json
// @Produce  json
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param name query string false "姓名"
// @Param phone query string false "手机号"
// @Param gender query int false "性别"
// @Param job_title query string false "职称"
// @Param job_number query string false "工号"
// @Router	/open-apis/core/personnel/enterprise/user/inservice [get]
func SelectInServiceEnterpriseUser(c *gin.Context) {
	var p GetEnterpriseUserReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	res, page, err := C.S.EnterpriseUser().Select(&entity.EnterpriseUser{
		EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
		Name:         p.Name,
		Phone:        p.Phone,
		JobTitle:     p.JobTitle,
		JobNumber:    p.JobNumber,
		Gender:       p.Gender,
		Status:       p.Status,
	}, &dao.Pagination{Page: p.Page, Size: p.Size})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.EnterpriseUser, 0)
	for _, row := range res {
		departmentPk := make([]string, 0)
		if dp, dpErr := C.S.Member().List(&entity.Member{UserPk: row.Pk}); dpErr == nil {
			for _, value := range dp {
				departmentPk = append(departmentPk, utils.Int64ToString(value.DepartmentPk))
			}
		}
		z := &pb.EnterpriseUser{
			Pk:                    utils.Int64ToString(row.Pk),
			EnterprisePk:          utils.Int64ToString(row.EnterprisePk),
			UserPk:                utils.Int64ToString(row.UserPk),
			Name:                  row.Name,
			Phone:                 row.Phone,
			Height:                utils.Float64ToString(row.Height),
			Weight:                utils.Float64ToString(row.Weight),
			Gender:                row.Gender,
			JobTitle:              row.JobTitle,
			JobNumber:             row.JobNumber,
			Nation:                row.Nation,
			Education:             row.Education,
			EmergencyContact:      row.EmergencyContact,
			EmergencyContactPhone: row.EmergencyContactPhone,
			TermOfContract:        row.TermOfContract,
			Address:               row.Address,
			Status:                row.Status,
			DepartmentPk:          departmentPk,
		}
		//查询人员
		if row.EntryTime != nil {
			z.EntryTime = utils.DateTimeToString(*row.EntryTime)
		}
		if row.Birthday != nil {
			z.Birthday = utils.DateTimeToString(*row.Birthday)
		}
		result = append(result, z)
	}
	utils.R(c, &pb.Resp{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Rows:  result,
	}, err)
}

// SelectEnterpriseUserDetail
// @Tags open-apis/core
// @summary 获取租户人员信息(tenant_access_token)
// @Accept json
// @Produce  json
// @Security bearer
// @Param user_pk query string true "用户pk"
// @Param enterprise_pk query string true "单位pk"
// @Router	/open-apis/core/personnel/enterprise/user/detail [get]
func SelectEnterpriseUserDetail(c *gin.Context) {
	var p SelectEnterpriseUserDetailReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}

	res, err := C.S.EnterpriseUser().FindByUserPkAndEnterprisePk(&entity.EnterpriseUser{
		UserPk:       utils.StringToInt64(p.UserPk),
		EnterprisePk: utils.StringToInt64(p.EnterprisePk),
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
	cfn := make([]*pb.EnterpriseUserAttachment, 0)
	eu.ContractFileName = cfn
	brfn := make([]*pb.EnterpriseUserAttachment, 0)
	eu.BodyReportFileNames = brfn
	rfn := make([]*pb.EnterpriseUserAttachment, 0)
	eu.ReviewFileNames = rfn
	utils.R(c, eu, err)
}

// CreateEnterpriseUser
// @Tags open-apis/core
// @summary 租户下新增人员接口
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostEnterpriseUserReq true "请求参数体"
// @Router	/open-apis/core/personnel/enterprise/user [post]
func CreateEnterpriseUser(c *gin.Context) {
	var p PostEnterpriseUserReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	enU := &entity.EnterpriseUser{
		EnterprisePk:          utils.StringToInt64(c.GetString("tenant")),
		UserPk:                0,
		Name:                  p.Name,
		Phone:                 p.Phone,
		Gender:                p.Gender,
		Height:                utils.StringToFloat64(p.Height),
		Weight:                utils.StringToFloat64(p.Weight),
		JobTitle:              p.JobTitle,
		JobNumber:             p.JobNumber,
		Nation:                p.Nation,
		Education:             p.Education,
		EmergencyContact:      p.EmergencyContact,
		EmergencyContactPhone: p.EmergencyContactPhone,
		TermOfContract:        p.TermOfContract,
		Address:               p.Address,
		Status:                1,
	}
	if p.Birthday != "" {
		birthday, birthdayErr := utils.StringToDateTime(p.Birthday)
		if birthdayErr == nil {
			enU.Birthday = &birthday
		}
	}
	res, err := C.S.EnterpriseUser().Create(enU)
	if err != nil {
		utils.FR(c, err)
	}
	//给人员分配部门
	for _, v := range p.DepartmentPk {
		if v != "" {
			_, _ = C.S.Member().Create(&entity.Member{
				EnterprisePk: utils.StringToInt64(c.GetString("tenant")),
				DepartmentPk: utils.StringToInt64(v),
				UserPk:       res.Pk,
			})
		}
	}
	utils.R(c, res, err)
}

// UpdateEnterpriseUser
// @Tags open-apis/core
// @summary 租户下人员信息修改
// @Accept json
// @Produce json
// @Security bearer
// @Param params body PutEnterpriseUserReq true "请求参数体"
// @Router /open-apis/core/personnel/enterprise/user [put]
func UpdateEnterpriseUser(c *gin.Context) {
	var p PutEnterpriseUserReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	enterpriseUserMessage, err2 := C.S.EnterpriseUser().FindByPk(&entity.EnterpriseUser{
		Pk: utils.StringToInt64(p.Pk),
	})
	if err2 != nil {
		utils.FR(c, err2)
	}
	eu := &entity.EnterpriseUser{
		Pk:                    utils.StringToInt64(p.Pk),
		Name:                  p.Name,
		Phone:                 p.Phone,
		Gender:                p.Gender,
		Height:                utils.StringToFloat64(p.Height),
		Weight:                utils.StringToFloat64(p.Weight),
		JobTitle:              p.JobTitle,
		JobNumber:             p.JobNumber,
		Nation:                p.Nation,
		Education:             p.Education,
		EmergencyContact:      p.EmergencyContact,
		EmergencyContactPhone: p.EmergencyContactPhone,
		TermOfContract:        p.TermOfContract,
		Address:               p.Address,
	}
	_, err := C.S.EnterpriseUser().Update(eu)
	if err != nil {
		utils.FR(c, err)
	}
	if p.DepartmentPk != nil && len(p.DepartmentPk) > 0 {
		departmentPks := make([]int64, 0)
		for _, v := range p.DepartmentPk {
			if v != "" && v != "0" {
				departmentPks = append(departmentPks, utils.StringToInt64(v))
			}
		}
		if len(departmentPks) > 0 {
			if departmentErr := C.S.Member().UpdateUserDepartmentRelation(&entity.Member{
				UserPk:       enterpriseUserMessage.UserPk,
				EnterprisePk: enterpriseUserMessage.EnterprisePk,
			}, departmentPks); departmentErr != nil {
				utils.FR(c, departmentErr)
			}
		}
	}
	utils.R(c, nil, err)
}

// DeleteEnterpriseUser
// @Tags open-apis/core
// @summary 租户下人员删除 (tenant_access_token)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DeleteReq true "请求参数体"
// @Router /open-apis/core/personnel/enterprise/user [delete]
func DeleteEnterpriseUser(c *gin.Context) {
	var p DeleteReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	enterpriseUserMessage, err3 := C.S.EnterpriseUser().FindByPk(&entity.EnterpriseUser{Pk: utils.StringToInt64(p.Pk)})
	if err3 != nil || enterpriseUserMessage.UserPk == 0 {
		utils.FR(c, err3)
	}
	if enterpriseUserMessage.Status == 1 {
		utils.FR(c, errors.New("在职人员无法删除！"))
	}
	err := C.S.EnterpriseUser().Delete(&entity.EnterpriseUser{
		Pk: utils.StringToInt64(p.Pk),
	})
	utils.R(c, nil, err)
}

// DimissionEnterpriseUser
// @Tags open-apis/core
// @summary 租户人员离职 (tenant_access_token)
// @Accept json
// @Produce json
// @Security bearer
// @Param params body DimissionEnterpriseUserReq true "请求参数体"
// @Router /open-apis/core/personnel/enterprise/user/dimission [put]
func DimissionEnterpriseUser(c *gin.Context) {
	var p DimissionEnterpriseUserReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.EnterpriseUser().DimissionEnterpriseUser(&entity.EnterpriseUser{
		Pk: utils.StringToInt64(p.Pk),
	})
	utils.R(c, nil, err)
}
