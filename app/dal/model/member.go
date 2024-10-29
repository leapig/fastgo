package model

import "github.com/leapig/fastgo/app/dal/entity"

type MemberModel struct {
	Pk           int64                 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64                 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户组件"`
	DepartmentPk int64                 `json:"department_pk" db:"department_pk"  gorm:"column:department_pk;comment:部门主键"`
	UserPk       int64                 `json:"user_pk" db:"user_pk"  gorm:"column:user_pk;comment:人员主键"`
	IsLeader     int32                 `json:"is_leader" db:"is_leader"  gorm:"column:is_leader;comment:是否是主管"`
	IsMain       int64                 `json:"is_main" db:"is_main"  gorm:"column:is_main;comment:是否是主部门"`
	JobTitle     string                `json:"job_title" db:"job_title"  gorm:"column:job_title;comment:职称"`
	JobNumber    int64                 `json:"job_number" db:"job_number"  gorm:"column:job_number;comment:工号"`
	Enterprise   entity.Enterprise     `json:"enterprise" gorm:"foreignKey:pk;references:enterprise_pk"`
	Department   entity.Department     `json:"department" gorm:"foreignKey:pk;references:department_pk"`
	User         entity.EnterpriseUser `json:"user" gorm:"foreignKey:pk;references:user_pk"`
}

func (MemberModel) TableName() string {
	return "member"
}

type Member struct {
	entity.BaseEntity
	Pk             int64                  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk   int64                  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户组件"`
	DepartmentPk   int64                  `json:"department_pk" db:"department_pk"  gorm:"column:department_pk;comment:部门主键"`
	UserPk         int64                  `json:"references" db:"user_pk"  gorm:"column:user_pk;comment:人员主键"`
	IsLeader       int32                  `json:"is_leader" db:"is_leader"  gorm:"column:is_leader;comment:是否是主管"`
	IsMain         int32                  `json:"is_main" db:"is_main"  gorm:"column:is_main;comment:是否是主部门"`
	JobTitle       string                 `json:"job_title" db:"job_title"  gorm:"column:job_title;comment:职称"`
	JobNumber      int64                  `json:"job_number" db:"job_number"  gorm:"column:job_number;comment:工号"`
	EnterpriseUser *entity.EnterpriseUser `json:"enterprise_user" db:"enterprise_user" gorm:"foreignKey:pk;references:user_pk"`
}

func (Member) TableName() string {
	return "member"
}
