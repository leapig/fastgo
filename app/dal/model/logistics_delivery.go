package model

import "github.com/leapig/fastgo/app/dal/entity"

type LogisticsDelivery struct {
	entity.LogisticsDelivery
	ClaimantEnterpriseUser entity.EnterpriseUser     `json:"claimant_enterprise_user" gorm:"foreignKey:pk;references:claimant"`
	ManagerEnterpriseUser  entity.EnterpriseUser     `json:"manager_enterprise_user" gorm:"foreignKey:pk;references:manager"`
	LogisticsInventory     entity.LogisticsInventory `json:"logistics_inventory" gorm:"foreignKey:pk;references:inventory_pk"`
	Project                entity.Project            `json:"project" gorm:"foreignKey:pk;references:project_pk"`
}

func (LogisticsDelivery) TableName() string {
	return "logistics_delivery"
}
