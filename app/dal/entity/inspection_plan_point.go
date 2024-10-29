package entity

type InspectionPlanPoint struct {
	BaseEntity
	Pk      int64 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	PlanPk  int64 `json:"plan_pk" db:"plan_pk"  gorm:"column:plan_pk;comment:"`
	PointPk int64 `json:"point_pk" db:"point_pk"  gorm:"column:point_pk;comment:"`
}

func (InspectionPlanPoint) TableName() string {
	return "inspection_plan_point"
}
