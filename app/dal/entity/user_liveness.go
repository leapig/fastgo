package entity

type UserLiveness struct {
	BaseEntity
	Pk        int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	UserPk    int64  `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:用户主键"`
	Operation string `json:"operation" db:"operation" gorm:"column:operation;comment:动作"`
}
