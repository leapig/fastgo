package model

import "github.com/leapig/fastgo/app/dal/entity"

type Project struct {
	entity.Project
	CreateEnterpriseName string `json:"create_enterprise_name" db:"create_enterprise_name" gorm:"column:create_enterprise_name;comment:项目创建企业名称"`
	CorporateName        string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone       string `json:"corporate_phone" db:"corporate_phone" gorm:"column:corporate_phone;comment:负责人手机"`
}

func (Project) TableName() string {
	return "project"
}

type ProjectQuery struct {
	ProjectAddress       string `json:"project_address" db:"project_address" gorm:"column:project_address;comment:详细地址"`
	Province             string `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City                 string `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District             string `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County               string `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	CorporateName        string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone       string `json:"corporate_phone" db:"corporate_phone"  gorm:"column:corporate_phone;comment:负责人手机号"`
	ProjectName          string `json:"project_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CreateEnterpriseName string `json:"create_enterprise_name" db:"create_enterprise_name" gorm:"column:create_enterprise_name;comment:项目创建企业名称"`
	EnterprisePk         int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	AddressCode          string `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:"`
	ProjectStatus        int32  `json:"project_status" db:"project_status"  gorm:"column:project_status;comment:"`
}
type ProjectQueryForUserProfession struct {
	EnterpriseUserPk int64 `json:"enterprise_user_pk" db:"enterprise_user_pk" gorm:"column:enterprise_user_pk;comment:"`
}
