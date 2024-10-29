package model

import "github.com/leapig/fastgo/app/dal/entity"

type LogisticsPurchase struct {
	entity.BaseEntity
	Pk                      int64                 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk            int64                 `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户主键"`
	Commodity               string                `json:"commodity" db:"commodity" gorm:"column:commodity;comment:商品"`
	Cover                   string                `json:"cover" db:"cover" gorm:"column:cover;comment:商品图"`
	Demand                  string                `json:"demand" db:"demand" gorm:"column:demand;comment:商品要求"`
	Num                     int32                 `json:"num" db:"num" gorm:"column:num;comment:需求数量"`
	Applicant               int64                 `json:"applicant" db:"applicant" gorm:"column:applicant;comment:申请人"`
	Reason                  string                `json:"reason" db:"reason" gorm:"column:reason;comment:申请原因"`
	Manager                 int64                 `json:"manager" db:"manager" gorm:"column:manager;comment:经办人"`
	Feedback                string                `json:"feedback" db:"feedback" gorm:"column:feedback;comment:反馈意见"`
	Status                  int32                 `json:"status" db:"status" gorm:"column:status;comment:单据状态1已完结2待处理3采购中4已拒绝"`
	ApplicantEnterpriseUser entity.EnterpriseUser `json:"applicant_enterprise_user" gorm:"foreignKey:pk;references:applicant"`
	ManagerEnterpriseUser   entity.EnterpriseUser `json:"manager_enterprise_user" gorm:"foreignKey:pk;references:manager"`
}

func (LogisticsPurchase) TableName() string {
	return "logistics_purchase"
}
