package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type ApplicationRecordModel struct {
	entity.BaseEntity
	Pk              int64             `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	UserPk          int64             `json:"user_pk" db:"user_pk"  gorm:"user_pk:cover;comment:"`
	PositionPk      int64             `json:"position_pk" db:"position_pk"  gorm:"position_pk:cover;comment:"`
	ApplyTime       time.Time         `json:"apply_time" db:"apply_time"  gorm:"apply_time:cover;comment:"`
	InterviewResult string            `json:"interview_result" db:"interview_result"  gorm:"interview_result:cover;comment:"`
	EnterprisePk    int64             `json:"enterprise_pk" db:"enterprise_pk"  gorm:"enterprise_pk:cover;comment:"`
	SourceType      string            `json:"source_type" db:"source_type"  gorm:"source_type:cover;comment:信息来源：0应聘人主动发起 1企业主动推送"`
	Enterprise      entity.Enterprise `json:"enterprise" gorm:"foreignKey:pk;references:enterprise_pk"`
	User            entity.User       `json:"user" gorm:"foreignKey:pk;references:user_pk"`
	Position        entity.Position   `json:"position" gorm:"foreignKey:pk;references:position_pk"`
	UserEntryPk     int64             `json:"user_entry_pk" db:"user_entry_pk"  gorm:"column:user_entry_pk;comment:待入职pk"`
}

func (ApplicationRecordModel) TableName() string {
	return "application_record"
}
