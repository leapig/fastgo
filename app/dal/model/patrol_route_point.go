package model

import "github.com/leapig/fastgo/app/dal/entity"

type PatrolRoutePoint struct {
	entity.PatrolRoutePoint
	PatrolPoint entity.PatrolPoint `json:"patrol_point"  gorm:"foreignKey:pk;references:point_pk"`
}

func (PatrolRoutePoint) TableName() string {
	return "patrol_route_point"
}
