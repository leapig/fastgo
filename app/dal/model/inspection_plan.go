package model

import "github.com/leapig/fastgo/app/dal/entity"

type InspectionPlan struct {
	entity.BaseEntity
	Pk                  int64                  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk        int64                  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	Name                string                 `json:"name" db:"name"  gorm:"column:name;comment:"`
	Remark              string                 `json:"remark" db:"remark"  gorm:"column:remark;comment:"`
	ProjectPk           int64                  `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:主键"`
	InspectionPlanPoint []*InspectionPlanPoint `json:"inspection_plan_point" gorm:"foreignKey:plan_pk;references:pk"`
}

func (InspectionPlan) TableName() string {
	return "inspection_plan"
}
