package entity

import "time"

type WorkSchedule struct {
	BaseEntity
	Pk              int64     `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk    int64     `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	ProjectPk       int64     `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:"`
	Name            string    `json:"name" db:"name"  gorm:"column:name;comment:"`
	BeginTime       time.Time `json:"begin_time"  db:"begin_time" gorm:"column:begin_time;comment:"`
	BeginEarlyTime  time.Time `json:"begin_early_time"  db:"begin_early_time" gorm:"column:begin_early_time;comment:"`
	BeginLateTime   time.Time `json:"begin_late_time"  db:"begin_late_time" gorm:"column:begin_late_time;comment:"`
	FinishTime      time.Time `json:"finish_time"  db:"finish_time" gorm:"column:finish_time;comment:"`
	FinishEarlyTime time.Time `json:"finish_early_time"  db:"finish_early_time" gorm:"column:finish_early_time;comment:"`
	FinishLateTime  time.Time `json:"finish_late_time"  db:"finish_late_time" gorm:"column:finish_late_time;comment:"`
	IsNextDay       int32     `json:"is_next_day"  db:"is_next_day" gorm:"column:is_next_day;comment:"`
}

func (WorkSchedule) TableName() string {
	return "work_schedule"
}
