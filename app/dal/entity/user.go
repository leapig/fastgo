package entity

import "time"

type User struct {
	BaseEntity
	Pk       int64      `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	Name     string     `json:"name" db:"name" gorm:"column:name;comment:姓名"`
	Phone    string     `json:"phone" db:"phone" gorm:"column:phone;comment:手机 "`
	Gender   int32      `json:"gender" db:"gender" gorm:"column:gender;comment:性别"`
	Birthday *time.Time `json:"birthday" db:"birthday" gorm:"column:birthday;comment:生日"`
	Height   float64    `json:"height" db:"height" gorm:"column:height;comment:身高"`
	Weight   float64    `json:"weight" db:"weight" gorm:"column:weight;comment:体重"`
}

func (User) TableName() string {
	return "user"
}
