package entity

type UserRealNameAuthenticationLog struct {
	BaseEntity
	Pk     int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	UserPk int64  `json:"user_pk" db:"user_pk"  gorm:"column:user_pk;comment:用户主键"`
	Face   string `json:"face" db:"face"  gorm:"column:face;comment:人脸链接"`
	Name   string `json:"name" db:"name"  gorm:"column:name;comment:姓名"`
	IdCard string `json:"id_card" db:"id_card"  gorm:"column:id_card;comment:证件号"`
}

func (UserRealNameAuthenticationLog) TableName() string {
	return "user_real_name_authentication_log"
}
