package model

import "github.com/leapig/fastgo/app/dal/entity"

type PatrolRouteRecordPoint struct {
	entity.PatrolRouteRecordPoint
	PatrolRoutePoint PatrolRoutePoint `json:"patrol_route"  gorm:"foreignKey:pk;references:patrol_route_point_pk"`
}

func (PatrolRouteRecordPoint) TableName() string {
	return "patrol_route_record_point"
}
