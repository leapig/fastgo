package entity

type InspectionPlan struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	Name         string `json:"name" db:"name"  gorm:"column:name;comment:"`
	Remark       string `json:"remark" db:"remark"  gorm:"column:remark;comment:"`
	ProjectPk    int64  `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:主键"`
}

func (InspectionPlan) TableName() string {
	return "inspection_plan"
}
