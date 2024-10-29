package model

import "github.com/leapig/fastgo/app/dal/entity"

type Enterprise struct {
	entity.Enterprise
	CorporateName  string `json:"corporate_name" db:"corporate_name"  gorm:"column:corporate_name;comment:负责人姓名"`
	CorporatePhone string `json:"corporate_phone" db:"corporate_phone" gorm:"column:corporate_phone;comment:负责人手机"`
}

func (Enterprise) TableName() string {
	return "enterprise"
}

type EnterpriseModel struct {
	entity.Enterprise
	EnterpriseUser entity.EnterpriseUser `json:"user"  gorm:"foreignKey:pk;references:enterprise_pk"`
}

func (EnterpriseModel) TableName() string {
	return "enterprise"
}
