package entity

type Department struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户主键"`
	Name         string `json:"name" db:"name"  gorm:"column:name;comment:名称"`
	ParentPk     int64  `json:"parent_pk" db:"parent_pk"  gorm:"column:parent_pk;comment:"`
}

func (Department) TableName() string {
	return "department"
}

type DepartmentList struct {
	BaseEntity
	Pk             int64             `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk   int64             `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户主键"`
	Name           string            `json:"name" db:"name"  gorm:"column:name;comment:名称"`
	ParentPk       int64             `json:"parent_pk" db:"parent_pk"  gorm:"column:parent_pk;comment:"`
	DepartmentList []*DepartmentList `gorm:"-"`
}

func (DepartmentList) TableName() string {
	return "department"
}
