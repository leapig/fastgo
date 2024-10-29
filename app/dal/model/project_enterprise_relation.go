package model

import "github.com/leapig/fastgo/app/dal/entity"

type ProjectEnterpriseRelation struct {
	entity.ProjectEnterpriseRelation
	Project              *entity.Project `json:"project" gorm:"foreignKey:pk;references:project_pk"`
	CorporateName        string          `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone       string          `json:"corporate_phone" db:"corporate_phone" gorm:"column:corporate_phone;comment:负责人手机"`
	CreateEnterpriseName string          `json:"create_enterprise_name" db:"create_enterprise_name" gorm:"column:create_enterprise_name;comment:项目创建企业名称"`
}

func (ProjectEnterpriseRelation) TableName() string {
	return "project_enterprise_relation"
}

type ProjectEnterpriseRelationForRelationEnterprise struct {
	entity.ProjectEnterpriseRelation
	EnterpriseName string `json:"enterprise_name" db:"enterprise_name" gorm:"column:enterprise_name;comment:项目创建企业名称"`
}

func (ProjectEnterpriseRelationForRelationEnterprise) TableName() string {
	return "project_enterprise_relation"
}

type ProjectEnterpriseRelationQuery struct {
	EnterpriseName string `json:"enterprise_name" db:"enterprise_name"  gorm:"column:enterprise_name;comment:"`
	CorporateName  string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone string `json:"corporate_phone" db:"corporate_phone" gorm:"column:corporate_phone;comment:负责人手机"`
	ProjectPk      int64  `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:项目主键"`
}
type ProjectEnterpriseRelationForProject struct {
	entity.ProjectEnterpriseRelation
	CorporateName  string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone string `json:"corporate_phone" db:"corporate_phone" gorm:"column:corporate_phone;comment:负责人手机"`
	EnterpriseName string `json:"enterprise_name" db:"enterprise_name" gorm:"column:enterprise_name;comment:项目创建企业名称"`
	Cover          string `json:"cover" db:"cover"  gorm:"column:cover;comment:封面"`
	Serial         string `json:"serial" db:"serial"  gorm:"column:serial;comment:统一社会信用代码"`
	StaffSize      string `json:"staff_size" db:"staff_size"  gorm:"column:staff_size;comment:人员规模"`
	License        string `json:"license" db:"license"  gorm:"column:license;comment:营业执照"`
	Country        string `json:"country" db:"country"  gorm:"column:country;comment:国"`
	Province       string `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City           string `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District       string `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County         string `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	Site           string `json:"site" db:"site"  gorm:"column:site;comment:地址"`
	Longitude      string `json:"longitude" db:"longitude"  gorm:"column:longitude;comment:经度"`
	Latitude       string `json:"latitude" db:"latitude"  gorm:"column:latitude;comment:纬度"`
	Type           int32  `json:"type" db:"type"  gorm:"column:type;comment:类型"`
	AddressCode    string `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:地址编码"`
}

func (ProjectEnterpriseRelationForProject) TableName() string {
	return "project_enterprise_relation"
}
