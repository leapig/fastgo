package model

import "github.com/leapig/fastgo/app/dal/entity"

type LogisticsInventory struct {
	entity.LogisticsInventory
	PrincipalEnterpriseUser entity.EnterpriseUser `json:"principal_enterprise_user" gorm:"foreignKey:pk;references:principal"`
}

func (LogisticsInventory) TableName() string {
	return "logistics_inventory"
}
