package model

import "github.com/leapig/fastgo/app/dal/entity"

type PositionModel struct {
	entity.BaseEntity
	Pk              int64              `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	Description     string             `json:"description" db:"description"  gorm:"description:cover;comment:职位详细信息"`
	SalaryRange     string             `json:"salary_range" db:"salary_range"  gorm:"salary_range:cover;comment:薪资范围"`
	Title           string             `json:"title" db:"title"  gorm:"title:cover;comment:招聘标题"`
	EnterprisePk    int64              `json:"enterprise_pk" db:"enterprise_pk"  gorm:"enterprise_pk:cover;comment:企业Pk"`
	PositionType    string             `json:"position_type" db:"position_type"  gorm:"position_type:cover;comment:招聘类型：0兼职 1长期"`
	Province        string             `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City            string             `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District        string             `json:"district" db:"district"  gorm:"column:district;comment:区"`
	PositionAddress string             `json:"position_address" db:"position_address"  gorm:"column:position_address;comment:招聘地址"`
	Qualification   string             `json:"qualification" db:"qualification"  gorm:"column:qualification;comment:所需资质"`
	IsIssue         string             `json:"is_issue" db:"is_issue"  gorm:"column:is_issue;comment:是否发布"`
	Longitude       string             `json:"longitude" db:"longitude"  gorm:"column:longitude;comment:经度"`
	Latitude        string             `json:"latitude" db:"latitude"  gorm:"column:latitude;comment:纬度"`
	IsCheck         string             `json:"is_check" db:"is_check"  gorm:"column:is_check;comment:是否需要审核  0需要  1不需要 "`
	Enterprise      *entity.Enterprise `json:"enterprise" gorm:"foreignKey:pk;references:enterprise_pk"`
}

func (PositionModel) TableName() string {
	return "position"
}
