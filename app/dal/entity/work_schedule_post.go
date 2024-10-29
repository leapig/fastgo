package entity

type WorkSchedulePost struct {
	BaseEntity
	Pk           int64 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	SchedulePk   int64 `json:"schedule_pk" db:"schedule_pk"  gorm:"column:schedule_pk;comment:"`
	PostPk       int64 `json:"post_pk" db:"post_pk"  gorm:"column:post_pk;comment:"`
	PeopleNumber int32 `json:"people_number"  db:"people_number" gorm:"column:people_number;comment:"`
}

func (WorkSchedulePost) TableName() string {
	return "work_schedule_post"
}
