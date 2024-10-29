package model

import "github.com/leapig/fastgo/app/dal/entity"

type WorkSchedulePost struct {
	entity.WorkSchedulePost
	ProjectPost  *entity.ProjectPost  `json:"project_post" gorm:"foreignKey:pk;references:post_pk"`
	WorkSchedule *entity.WorkSchedule `json:"work_schedule" gorm:"foreignKey:pk;references:schedule_pk"`
}

func (WorkSchedulePost) TableName() string {
	return "work_schedule_post"
}
