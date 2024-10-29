package model

import "github.com/leapig/fastgo/app/dal/entity"

type PatrolPoint struct {
	entity.PatrolPoint
}

func (PatrolPoint) TableName() string {
	return "patrol_point"
}
