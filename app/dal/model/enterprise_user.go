package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type EnterpriseUser struct {
	entity.BaseEntity
	Pk                    int64              `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	EnterprisePk          int64              `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
	UserPk                int64              `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:人员主键"`
	Name                  string             `json:"name" db:"name" gorm:"column:name;comment:姓名"`
	Phone                 string             `json:"phone" db:"phone" gorm:"column:phone;comment:手机 "`
	Gender                int32              `json:"gender" db:"gender" gorm:"column:gender;comment:性别"`
	Birthday              *time.Time         `json:"birthday" db:"birthday" gorm:"column:birthday;comment:生日"`
	Height                float64            `json:"height" db:"height" gorm:"column:height;comment:身高"`
	Weight                float64            `json:"weight" db:"weight" gorm:"column:weight;comment:体重"`
	JobTitle              string             `json:"job_title" db:"job_title"  gorm:"column:job_title;comment:职称"`
	JobNumber             string             `json:"job_number" db:"job_number"  gorm:"column:job_number;comment:工号"`
	Nation                string             `json:"nation" db:"nation"  gorm:"column:nation;comment:民族"`
	Education             string             `json:"education" db:"education"  gorm:"column:education;comment:文化程度"`
	EmergencyContact      string             `json:"emergency_contact" db:"emergency_contact"  gorm:"column:emergency_contact;comment:紧急联系人"`
	EmergencyContactPhone string             `json:"emergency_contact_phone" db:"emergency_contact_phone"  gorm:"column:emergency_contact_phone;comment:紧急联系人手机号"`
	EntryTime             *time.Time         `json:"entry_time" db:"entry_time" gorm:"column:entry_time;comment:入职时间"`
	TermOfContract        string             `json:"term_of_contract" db:"term_of_contract"  gorm:"column:term_of_contract;comment:合同期限"`
	Address               string             `json:"address" db:"address"  gorm:"column:address;comment:住址"`
	Status                int32              `json:"status" db:"status" gorm:"column:status;comment:人员状态  1在职 2离职"`
	Enterprise            *entity.Enterprise `json:"enterprise" gorm:"foreignKey:pk;references:enterprise_pk"`
}

func (EnterpriseUser) TableName() string {
	return "enterprise_user"
}
