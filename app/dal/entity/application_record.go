package entity

import "time"

type ApplicationRecord struct {
	BaseEntity
	Pk              int64      `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	UserPk          int64      `json:"user_pk" db:"user_pk"  gorm:"user_pk:cover;comment:应聘人Pk"`
	PositionPk      int64      `json:"position_pk" db:"position_pk"  gorm:"position_pk:cover;comment:职位Pk"`
	ApplyTime       *time.Time `json:"apply_time" db:"apply_time"  gorm:"apply_time:cover;comment:应聘时间"`
	InterviewResult string     `json:"interview_result" db:"interview_result"  gorm:"interview_result:cover;comment:申请结果：0已申请但企业未处理  1 企业通过申请邀请入职   9企业主动推送但人员未处理  8企业主动推送但被应聘人忽略"`
	EnterprisePk    int64      `json:"enterprise_pk" db:"enterprise_pk"  gorm:"enterprise_pk:cover;comment:"`
	SourceType      string     `json:"source_type" db:"source_type"  gorm:"source_type:cover;comment:信息来源：0应聘人主动发起 1企业主动推送"`
	ApprovalUserPk  int64      `json:"approval_user_pk" db:"approval_user_pk"  gorm:"approval_user_pk:cover;comment:审批人pk"`
	ApprovalTime    *time.Time `json:"approval_time" db:"approval_time"  gorm:"approval_time:cover;comment:审批时间"`
	UserEntryPk     int64      `json:"user_entry_pk" db:"user_entry_pk"  gorm:"column:user_entry_pk;comment:待入职pk"`
}

func (ApplicationRecord) TableName() string {
	return "application_record"
}
