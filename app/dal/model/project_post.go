package model

import "github.com/leapig/fastgo/app/dal/entity"

type ProjectPost struct {
	entity.ProjectPost
	EnterpriseName string `json:"enterprise_name" db:"enterprise_name" gorm:"column:enterprise_name;comment:所属公司名称"`
}

func (ProjectPost) TableName() string {
	return "project_post"
}
