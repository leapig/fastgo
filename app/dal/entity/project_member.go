package entity

import "time"

type ProjectMember struct {
	BaseEntity
	Pk               int64     `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterpriseUserPk int64     `json:"enterprise_user_pk" db:"enterprise_user_pk"  gorm:"column:enterprise_user_pk;comment:成员主键"`
	ProjectPk        int64     `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:项目主键"`
	EnterprisePk     int64     `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	UserPk           int64     `json:"user_pk" db:"user_pk"  gorm:"column:user_pk;comment:主键"`
	UserStatus       int32     `json:"user_status" db:"user_status"  gorm:"column:user_status;comment:"`
	IsLongTerm       int32     `json:"is_long_term" db:"is_long_term"  gorm:"column:is_long_term;comment:"`
	RestDay          string    `json:"rest_day" db:"rest_day"  gorm:"column:rest_day;comment:"`
	RestType         int32     `json:"rest_type" db:"rest_type"  gorm:"column:rest_type;comment:"`
	OutTime          time.Time `json:"out_time"  db:"out_time" gorm:"column:out_time;comment:"`
	WorkType         int32     `json:"work_type" db:"work_type"  gorm:"column:work_type;comment:"`
}

func (ProjectMember) TableName() string {
	return "project_member"
}
