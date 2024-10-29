package entity

type EnterpriseRelation struct {
	BaseEntity
	Pk                   int64 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk         int64 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:主键"`
	RelationEnterprisePk int64 `json:"relation_enterprise_pk" db:"relation_enterprise_pk"  gorm:"column:relation_enterprise_pk;comment:主键"`
	ProjectPk            int64 `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:主键"`
}

func (EnterpriseRelation) TableName() string {
	return "enterprise_relation"
}
