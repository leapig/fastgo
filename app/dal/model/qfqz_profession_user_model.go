package model

import "github.com/leapig/fastgo/app/dal/entity"

type QfqzProfessionUserModel struct {
	entity.BaseEntity
	Pk              int64                    `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	UserPk          int64                    `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:用户主键"`
	Type            int64                    `json:"type" db:"type" gorm:"column:type;comment:主键"`
	User            *entity.User             `json:"user" gorm:"foreignKey:user_pk;references:pk"`
	UserCredentials []entity.UserCredentials `json:"user_credentials" gorm:"foreignKey:user_pk;references:user_pk"`
	// todo 所对应安保公司字段
}

func (QfqzProfessionUserModel) TableName() string {
	return "user_profession"
}
