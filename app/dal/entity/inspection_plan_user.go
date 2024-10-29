package entity

type InspectionPlanUser struct {
	BaseEntity
	Pk               int64 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk     int64 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	PlanPk           int64 `json:"plan_pk" db:"plan_pk"  gorm:"column:plan_pk;comment:"`
	EnterpriseUserPk int64 `json:"enterprise_user_pk" db:"enterprise_user_pk"  gorm:"column:enterprise_user_pk;comment:"`
	Status           int32 `json:"status" db:"status"  gorm:"column:status;comment:"`
}

func (InspectionPlanUser) TableName() string {
	return "inspection_plan_user"
}
