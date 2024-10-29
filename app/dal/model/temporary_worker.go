package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type TemporaryWorker struct {
	entity.BaseEntity
	Pk               int64                 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	EnterprisePk     int64                 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	EnterpriseUserPk int64                 `json:"enterprise_user_pk" db:"enterprise_user_pk"  gorm:"column:enterprise_user_pk;comment:租户用户pk"`
	StartTime        *time.Time            `json:"start_time" db:"start_time" gorm:"column:start_time;comment:起始时间"`
	EndTime          *time.Time            `json:"end_time" db:"end_time" gorm:"column:end_time;comment:结束时间"`
	EnterpriseUser   entity.EnterpriseUser `json:"enterprise_user" gorm:"foreignKey:pk;references:enterprise_user_pk"`
}

func (TemporaryWorker) TableName() string {
	return "temporary_worker"
}
