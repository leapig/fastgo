package entity

import "time"

type CarApply struct {
	BaseEntity
	Pk           int64     `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	EnterprisePk int64     `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
	CarPk        int64     `json:"car_pk" db:"car_pk" gorm:"column:car_pk;comment:车辆主键"`
	Applicant    int64     `json:"applicant" db:"applicant" gorm:"column:applicant;comment:申请人"`
	Reason       string    `json:"reason" db:"reason" gorm:"column:reason;comment:用车事由"`
	StartTime    time.Time `json:"start_time" db:"start_time" gorm:"column:start_time;comment:借用时间"`
	OverTime     time.Time `json:"over_time" db:"over_time" gorm:"column:over_time;comment:结束时间"`
	BackTime     time.Time `json:"back_time" db:"back_time" gorm:"column:back_time;comment:归还时间"`
	Gas          int32     `json:"gas" db:"gas" gorm:"column:gas;comment:是否加油"`
	Expense      float64   `json:"expense" db:"expense" gorm:"column:expense;comment:加油消费"`
}

func (CarApply) TableName() string {
	return "car_apply"
}
