package entity

type ProjectEnterpriseRelation struct {
	BaseEntity
	Pk           int64 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	ProjectPk    int64 `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:项目主键"`
}

func (ProjectEnterpriseRelation) TableName() string {
	return "project_enterprise_relation"
}
