package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type TurnoverRecord struct {
	Pk                       int64                 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk             int64                 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	EnterpriseUserPk         int64                 `json:"enterprise_user_pk" db:"enterprise_user_pk"  gorm:"column:enterprise_user_pk;comment:租户用户pk"`
	ApplyStatus              int32                 `json:"apply_status" db:"apply_status"  gorm:"column:apply_status;comment:审批状态"`
	Context                  string                `json:"context" db:"context"  gorm:"column:context;comment:申请内容"`
	Feedback                 string                `json:"feedback" db:"feedback"  gorm:"column:feedback;comment:反馈意见"`
	ApplicationTime          *time.Time            `json:"application_time" db:"application_time"  gorm:"column:application_time;comment:申请时间"`
	ApprovalTime             *time.Time            `json:"approval_time" db:"approval_time"  gorm:"column:approval_time;comment:审批时间"`
	ApprovalEnterpriseUserPk int64                 `json:"approval_enterprise_user_pk" db:"approval_enterprise_user_pk"  gorm:"column:approval_enterprise_user_pk;comment:审批人租户人员pk"`
	Enterprise               entity.Enterprise     `json:"enterprise" db:"enterprise_pk" gorm:"foreignKey:pk;references:enterprise_pk"`
	EnterpriseUser           entity.EnterpriseUser `json:"enterprise_user"  gorm:"foreignKey:pk;references:enterprise_user_pk"`
	ApprovalEnterpriseUser   entity.EnterpriseUser `json:"approval_enterprise_user"  gorm:"foreignKey:pk;references:enterprise_user_pk"`
}

func (TurnoverRecord) TableName() string {
	return "turnover_record"
}
