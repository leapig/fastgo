package model

import "github.com/leapig/fastgo/app/dal/entity"

type EnterpriseRelation struct {
	entity.EnterpriseRelation
}

func (EnterpriseRelation) TableName() string {
	return "enterprise_relation"
}

type EnterpriseRelationQuery struct {
	EnterprisePk   int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:主键"`
	CustomerName   string `json:"customer_name" db:"customer_name"  gorm:"column:customer_name;comment:客户名称"`
	Country        string `json:"country" db:"country"  gorm:"column:country;comment:国"`
	Province       string `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City           string `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District       string `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County         string `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	Site           string `json:"site" db:"site"  gorm:"column:site;comment:地址"`
	AddressCode    string `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:地址编码"`
	CorporateName  string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone string `json:"corporate_phone" db:"corporate_phone" gorm:"column:corporate_phone;comment:负责人手机"`
}
type EnterpriseRelationProjectQuery struct {
	EnterprisePk         int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:主键"`
	RelationEnterprisePk int64  `json:"relation_enterprise_pk" db:"relation_enterprise_pk"  gorm:"column:relation_enterprise_pk;comment:主键"`
	ProjectAddress       string `json:"project_address" db:"project_address" gorm:"column:project_address;comment:详细地址"`
	Province             string `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City                 string `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District             string `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County               string `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	CorporateName        string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone       string `json:"corporate_phone" db:"corporate_phone"  gorm:"column:corporate_phone;comment:负责人手机号"`
	ProjectName          string `json:"project_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CreateEnterpriseName string `json:"create_enterprise_name" db:"create_enterprise_name" gorm:"column:create_enterprise_name;comment:项目创建企业名称"`
	AddressCode          string `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:"`
	ProjectStatus        int32  `json:"project_status" db:"project_status"  gorm:"column:project_status;comment:"`
}
