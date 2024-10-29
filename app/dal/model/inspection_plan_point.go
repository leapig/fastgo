package model

import "github.com/leapig/fastgo/app/dal/entity"

type InspectionPlanPoint struct {
	entity.BaseEntity
	Pk              int64                   `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	PlanPk          int64                   `json:"plan_pk" db:"plan_pk"  gorm:"column:plan_pk;comment:"`
	PointPk         int64                   `json:"point_pk" db:"point_pk"  gorm:"column:point_pk;comment:"`
	InspectionPoint *entity.InspectionPoint `json:"inspection_point" gorm:"foreignKey:pk;references:point_pk"`
}

func (InspectionPlanPoint) TableName() string {
	return "inspection_plan_point"
}
