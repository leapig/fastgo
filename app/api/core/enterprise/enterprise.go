package enterprise

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/aliyun"
	utils "github.com/leapig/fastgo/app/library/util"
	C "github.com/leapig/fastgo/app/service"
)

// PostNew
// @Tags open-apis/core
// @summary 小程序新用户新增安保单位接口
// @Accept json
// @Produce  json
// @Param params body PostNewReq true "请求参数体"
// @Router	/open-apis/core/enterprise/new [post]
func PostNew(c *gin.Context) {
	var p PostNewReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	err := C.S.Wechat().CreateEnterprise(&entity.User{
		Name:  p.UserName,
		Phone: p.Phone,
	}, p.Code, &entity.Enterprise{
		Country:     p.Country,
		Province:    p.Province,
		City:        p.City,
		District:    p.District,
		County:      p.County,
		Site:        p.Site,
		Longitude:   p.Longitude,
		Latitude:    p.Latitude,
		Name:        p.EnterpriseName,
		StaffSize:   p.StaffSize,
		Type:        2,
		AddressCode: p.AddressCode,
	})
	utils.R(c, nil, err)
}

// PostCreate
// @Tags open-apis/core
// @summary 小程序老用户新增安保单位(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param params body PostSecurityReq true "请求参数体"
// @Router	/open-apis/core/enterprise/create [post]
func PostCreate(c *gin.Context) {
	var p PostSecurityReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	row, err := C.S.Enterprise().Create(&entity.Enterprise{
		Name:        p.Name,
		Country:     p.Country,
		Province:    p.Province,
		City:        p.City,
		District:    p.District,
		County:      p.County,
		Site:        p.Site,
		Longitude:   p.Longitude,
		Latitude:    p.Latitude,
		Type:        2,
		StaffSize:   p.StaffSize,
		AddressCode: p.AddressCode,
	})
	if err != nil {
		utils.FR(c, err)
	}
	row.CorporatePk = utils.StringToInt64(c.GetString("user"))
	row, err = C.S.Enterprise().Update(&entity.Enterprise{
		Pk:          row.Pk,
		CorporatePk: row.CorporatePk,
	})
	if err != nil {
		utils.FR(c, err)
	}
	//分配默认角色、添加进根部门、添加人员租户关联关系
	if err = C.S.UserPermission().CreateForCorporate(&entity.UserPermission{
		EnterprisePk: row.Pk,
		UserPk:       row.CorporatePk,
	}); err != nil {
		utils.FR(c, err)
	}
	row.Cover = aliyun.Oss().PathToUrl(row.Cover)
	row.License = aliyun.Oss().PathToUrl(row.License)
	utils.R(c, pb.Enterprise{
		Pk:             utils.Int64ToString(row.Pk),
		Cover:          row.Cover,
		Name:           row.Name,
		Serial:         row.Serial,
		License:        row.License,
		Country:        row.Country,
		Province:       row.Province,
		City:           row.City,
		District:       row.District,
		County:         row.County,
		Site:           row.Site,
		Longitude:      row.Longitude,
		Latitude:       row.Latitude,
		Type:           row.Type,
		CorporatePk:    utils.Int64ToString(row.CorporatePk),
		CorporateName:  row.CorporateName,
		CorporatePhone: row.CorporatePhone,
		CreateAt:       utils.TimeToString(row.CreatedAt),
		StaffSize:      row.StaffSize,
		AddressCode:    row.AddressCode,
	}, err)
}

// GetList
// @Tags open-apis/core
// @summary 获取用户单位列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/enterprise/list [get]
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

// GetAppList
// @Tags open-apis/core
// @summary 小程序获取用户单位列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Router	/open-apis/core/enterprise/applist [get]
func GetAppList(c *gin.Context) {
	if res, err := C.S.Enterprise().ReadAppUserEnterprise(utils.StringToInt64(c.GetString("user"))); err == nil {
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

// GetSecurity
// @Tags open-apis/core
// @summary 获取安保单位列表(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param page query int true "页码"
// @Param size query int true "页数"
// @Param name query string false "名称"
// @Param serial query string false "信用代码"
// @Param address_code query string false "区域编码"
// @Router	/open-apis/core/enterprise/security [get]
func GetSecurity(c *gin.Context) {
	var p EnterpriseReq
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.FRP(c)
	}
	res, page, err := C.S.Enterprise().Select(&entity.Enterprise{
		Name:        p.Name,
		Serial:      p.Serial,
		Country:     p.Country,
		Province:    p.Province,
		City:        p.City,
		District:    p.District,
		County:      p.County,
		Type:        2,
		AddressCode: p.AddressCode,
	}, &dao.Pagination{Page: p.Page, Size: p.Size})
	if err != nil {
		utils.FR(c, err)
	}
	result := make([]*pb.Enterprise, 0)
	for _, row := range res {
		row.Cover = aliyun.Oss().PathToUrl(row.Cover)
		row.License = aliyun.Oss().PathToUrl(row.License)
		en := &pb.Enterprise{
			Pk:             utils.Int64ToString(row.Pk),
			Cover:          row.Cover,
			Name:           row.Name,
			Serial:         row.Serial,
			License:        row.License,
			Country:        row.Country,
			Province:       row.Province,
			City:           row.City,
			District:       row.District,
			County:         row.County,
			Site:           row.Site,
			Longitude:      row.Longitude,
			Latitude:       row.Latitude,
			Type:           row.Type,
			CorporatePk:    utils.Int64ToString(row.CorporatePk),
			CorporateName:  row.CorporateName,
			CorporatePhone: row.CorporatePhone,
			CreateAt:       utils.TimeToString(row.CreatedAt),
			StaffSize:      row.StaffSize,
			AddressCode:    row.AddressCode,
		}
		eap := make([]*pb.EnterpriseAreaPermission, 0)
		if list, listErr := C.S.EnterpriseAreaPermission().FindList(&entity.EnterpriseAreaPermission{EnterprisePk: row.Pk}); listErr == nil {
			for _, value := range list {
				eap = append(eap, &pb.EnterpriseAreaPermission{
					Pk:           utils.Int64ToString(value.Pk),
					EnterprisePk: utils.Int64ToString(value.EnterprisePk),
					Province:     value.Province,
					City:         value.City,
					District:     value.District,
					County:       value.County,
				})
			}

		}
		en.EnterpriseAreaPermissions = eap
		result = append(result, en)
	}
	utils.R(c, &pb.Resp{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Rows:  result,
	}, err)
}

// PostSecurity
// @Tags open-apis/core
// @summary 新增安保单位(UserAccessToken)
// @Accept json
// @Produce  json
// @Security bearer
// @Param req body PostSecurityReq true "请求参数"
// @Router	/open-apis/core/enterprise/security [post]
func PostSecurity(c *gin.Context) {
	var p PostSecurityReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}

	row, err := C.S.Enterprise().Create(&entity.Enterprise{
		Name:        p.Name,
		Country:     p.Country,
		Province:    p.Province,
		City:        p.City,
		District:    p.District,
		County:      p.County,
		Site:        p.Site,
		Longitude:   p.Longitude,
		Latitude:    p.Latitude,
		StaffSize:   p.StaffSize,
		Type:        2,
		AddressCode: p.AddressCode,
	})
	if err != nil {
		utils.FR(c, err)
	}
	row.Cover = aliyun.Oss().PathToUrl(row.Cover)
	row.License = aliyun.Oss().PathToUrl(row.License)
	utils.R(c, &pb.Enterprise{
		Pk:             utils.Int64ToString(row.Pk),
		Cover:          row.Cover,
		Name:           row.Name,
		Serial:         row.Serial,
		License:        row.License,
		Country:        row.Country,
		Province:       row.Province,
		City:           row.City,
		District:       row.District,
		County:         row.County,
		Site:           row.Site,
		Longitude:      row.Longitude,
		Latitude:       row.Latitude,
		Type:           row.Type,
		CorporatePk:    utils.Int64ToString(row.CorporatePk),
		CorporateName:  row.CorporateName,
		CorporatePhone: row.CorporatePhone,
		CreateAt:       utils.TimeToString(row.CreatedAt),
		StaffSize:      row.StaffSize,
		AddressCode:    row.AddressCode,
	}, err)
}

// PutSecurity
// @Tags open-apis/core
// @summary 修改单位基本信息(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param req body PutSecurityReq true "请求参数"
// @Router /open-apis/core/enterprise/security [put]
func PutSecurity(c *gin.Context) {
	var p PutSecurityReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	row, err := C.S.Enterprise().Update(&entity.Enterprise{
		Pk:          utils.StringToInt64(p.Pk),
		Name:        p.Name,
		Country:     p.Country,
		Province:    p.Province,
		City:        p.City,
		District:    p.District,
		County:      p.County,
		Site:        p.Site,
		StaffSize:   p.StaffSize,
		Longitude:   p.Longitude,
		Latitude:    p.Latitude,
		AddressCode: p.AddressCode,
	})
	if err != nil {
		utils.FR(c, err)
	}
	row.Cover = aliyun.Oss().PathToUrl(row.Cover)
	row.License = aliyun.Oss().PathToUrl(row.License)
	utils.R(c, &pb.Enterprise{
		Pk:             utils.Int64ToString(row.Pk),
		Cover:          row.Cover,
		Name:           row.Name,
		Serial:         row.Serial,
		License:        row.License,
		Country:        row.Country,
		Province:       row.Province,
		City:           row.City,
		District:       row.District,
		County:         row.County,
		Site:           row.Site,
		Longitude:      row.Longitude,
		Latitude:       row.Latitude,
		Type:           row.Type,
		CorporatePk:    utils.Int64ToString(row.CorporatePk),
		CorporateName:  row.CorporateName,
		CorporatePhone: row.CorporatePhone,
		CreateAt:       utils.TimeToString(row.CreatedAt),
		StaffSize:      row.StaffSize,
		AddressCode:    row.AddressCode,
	}, err)
}

// PutCorporate
// @Tags open-apis/core
// @summary 修改单位负责人(UserAccessToken)
// @Accept json
// @Produce json
// @Security bearer
// @Param req body PutCorporateReq true "请求参数"
// @Router /open-apis/core/enterprise/corporate [put]
func PutCorporate(c *gin.Context) {
	var p PutCorporateReq
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.FRP(c)
	}
	row, err := C.S.Enterprise().Update(&entity.Enterprise{
		Pk:          utils.StringToInt64(p.Pk),
		CorporatePk: utils.StringToInt64(p.CorporatePk),
	})
	if err != nil {
		utils.FR(c, err)
	}
	//分配默认角色、添加进根部门、添加人员租户关联关系
	if err = C.S.UserPermission().CreateForCorporate(&entity.UserPermission{
		EnterprisePk: utils.StringToInt64(p.Pk),
		UserPk:       utils.StringToInt64(p.CorporatePk),
	}); err != nil {
		utils.FR(c, err)
	}
	row.Cover = aliyun.Oss().PathToUrl(row.Cover)
	row.License = aliyun.Oss().PathToUrl(row.License)
	utils.R(c, &pb.Enterprise{
		Pk:             utils.Int64ToString(row.Pk),
		Cover:          row.Cover,
		Name:           row.Name,
		Serial:         row.Serial,
		License:        row.License,
		Country:        row.Country,
		Province:       row.Province,
		City:           row.City,
		District:       row.District,
		County:         row.County,
		Site:           row.Site,
		Longitude:      row.Longitude,
		Latitude:       row.Latitude,
		Type:           row.Type,
		CorporatePk:    utils.Int64ToString(row.CorporatePk),
		CorporateName:  row.CorporateName,
		CorporatePhone: row.CorporatePhone,
		CreateAt:       utils.TimeToString(row.CreatedAt),
		StaffSize:      row.StaffSize,
		AddressCode:    row.AddressCode,
	}, err)
}
