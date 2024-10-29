package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type QfqzAccountModelForEnterprise struct {
	entity.BaseEntity
	Pk                 int64                     `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	UserPk             int64                     `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:人员主键"`
	EnterprisePk       int64                     `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户主键"`
	Name               string                    `json:"name" db:"name" gorm:"column:name;comment:姓名"`
	Phone              string                    `json:"phone" db:"phone" gorm:"column:phone;comment:手机 "`
	CreatedAt          time.Time                 `json:"create_at"    db:"create_at" gorm:"column:create_at;comment:创建时间"`
	UserCredentials    []*entity.UserCredentials `json:"user_credentials"  gorm:"foreignKey:user_pk;references:user_pk"`    //卡片
	UserPermissionList []*entity.UserPermission  `json:"user_permission_list" gorm:"foreignKey:user_pk;references:user_pk"` //权限
	UserClientList     []*entity.UserClient      `json:"user_client_list" gorm:"foreignKey:user_pk;references:user_pk"`     //第三方
}

func (QfqzAccountModelForEnterprise) TableName() string {
	return "enterprise_user"
}

type QfqzAccountModel struct {
	entity.BaseEntity
	Pk                 int64                     `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	UserPk             int64                     `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:人员主键"`
	EnterprisePk       int64                     `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户主键"`
	Name               string                    `json:"name" db:"name" gorm:"column:name;comment:姓名"`
	Phone              string                    `json:"phone" db:"phone" gorm:"column:phone;comment:手机 "`
	CreatedAt          time.Time                 `json:"create_at"    db:"create_at" gorm:"column:create_at;comment:创建时间"`
	UserCredentials    []*entity.UserCredentials `json:"user_credentials"  gorm:"foreignKey:user_pk;references:pk"`    //卡片
	UserPermissionList []*entity.UserPermission  `json:"user_permission_list" gorm:"foreignKey:user_pk;references:pk"` //权限
	UserClientList     []*entity.UserClient      `json:"user_client_list" gorm:"foreignKey:user_pk;references:pk"`     //第三方
}

func (QfqzAccountModel) TableName() string {
	return "user"
}
