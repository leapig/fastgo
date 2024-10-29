package model

import "github.com/leapig/fastgo/app/dal/entity"

type PatrolRoute struct {
	entity.PatrolRoute
	PatrolRoutePoints []*PatrolRoutePoint `json:"patrol_route_points" gorm:"foreignKey:route_pk;references:pk"`
	PointRange        int32               `json:"point_range"`

	//gorm:"foreignKey:plan_pk;references:pk"`
}

func (PatrolRoute) TableName() string {
	return "patrol_route"
}
