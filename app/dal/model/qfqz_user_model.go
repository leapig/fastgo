package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
)

type QfqzUserModel struct {
	entity.User
	UserRealNameAuthenticationLog []*entity.UserRealNameAuthenticationLog `json:"real_name_log"   gorm:"foreignKey:user_pk;references:pk"`
	UserLiveness                  []*entity.UserLiveness                  `json:"user_liveness"   gorm:"foreignKey:user_pk;references:pk"`
	UserCredentials               []*entity.UserCredentials               `json:"user_credentials"   gorm:"foreignKey:user_pk;references:pk"`
}

func (QfqzUserModel) TableName() string {
	return "user"
}
