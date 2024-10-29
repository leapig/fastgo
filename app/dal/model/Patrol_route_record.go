package model

import "github.com/leapig/fastgo/app/dal/entity"

type PatrolRouteRecord struct {
	entity.PatrolRouteRecord
	EnterpriseUser entity.EnterpriseUser `json:"enterprise_user"  gorm:"foreignKey:pk;references:enterprise_user_pk"`
	PatrolRoute    entity.PatrolRoute    `json:"patrol_route"  gorm:"foreignKey:pk;references:route_pk"`
}

func (PatrolRouteRecord) TableName() string {
	return "patrol_route_record"
}
