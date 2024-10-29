package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type ProjectMember struct {
	entity.ProjectMember
	Phone          string    `json:"phone" db:"phone" gorm:"column:phone;comment:手机 "`
	Name           string    `json:"name" db:"name" gorm:"column:name;comment:姓名"`
	Height         float64   `json:"height" db:"height" gorm:"column:height;comment:身高"`
	Weight         float64   `json:"weight" db:"weight" gorm:"column:weight;comment:体重"`
	Education      string    `json:"education" db:"education"  gorm:"column:education;comment:文化程度"`
	Gender         int32     `json:"gender" db:"gender" gorm:"column:gender;comment:性别"`
	Birthday       time.Time `json:"birthday" db:"birthday" gorm:"column:birthday;comment:生日"`
	EnterpriseName string    `json:"enterprise_name" db:"enterprise_name" gorm:"column:enterprise_name;comment:所属公司名称"`
}

func (ProjectMember) TableName() string {
	return "project_member"
}

type ProjectMemberQueryBody struct {
	EnterpriseUserPk int64  `json:"enterprise_user_pk" db:"enterprise_user_pk"  gorm:"column:enterprise_user_pk;comment:成员主键"`
	ProjectPk        int64  `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:项目主键"`
	EnterprisePk     int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	Name             string `json:"name" db:"name"  gorm:"column:name;comment:人员姓名"`
	Phone            string `json:"phone" db:"phone"  gorm:"column:phone;comment:电话"`
	Gender           int32  `json:"gender" db:"gender"  gorm:"column:gender;comment:性别 1 男 2女 0未知"`
}
