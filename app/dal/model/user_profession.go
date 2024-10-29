package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"time"
)

type UserProfessionQuery struct {
	EnterprisePk       int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:主键"`
	UserName           string `json:"user_name" db:"user_name" gorm:"column:user_name;comment:用户名称"`
	UserPhone          string `json:"user_phone" db:"user_phone" gorm:"column:user_phone;comment:用户手机"`
	UserProfessionType int32  `json:"user_profession_type" db:"user_profession_type" gorm:"column:user_profession_type;comment:"`
	UserStatus         int32  `json:"user_status" db:"user_status" gorm:"column:user_status;comment:人员状态  1在职 2离职"`
	UserGender         int32  `json:"user_gender" db:"user_gender" gorm:"column:user_gender;comment:"`
}
type UserProfessionForPlatformQuery struct {
	UserName   string `json:"user_name" db:"user_name" gorm:"column:user_name;comment:用户名称"`
	UserPhone  string `json:"user_phone" db:"user_phone" gorm:"column:user_phone;comment:用户手机"`
	UserStatus int32  `json:"user_status" db:"user_status" gorm:"column:user_status;comment:人员状态  1在职 2离职"`
	UserGender int32  `json:"user_gender" db:"user_gender" gorm:"column:user_gender;comment:"`
}
type UserProfession struct {
	entity.UserProfession
	UserName     string    `json:"user_name" db:"user_name" gorm:"column:user_name;comment:姓名"`
	UserPhone    string    `json:"user_phone" db:"user_phone" gorm:"column:user_phone;comment:手机 "`
	UserStatus   int32     `json:"user_status" db:"user_status" gorm:"column:user_status;comment:人员状态  1在职 2离职"`
	UserGender   int32     `json:"user_gender" db:"user_gender" gorm:"column:user_gender;comment:"`
	UserBirthday time.Time `json:"user_birthday" db:"user_birthday" gorm:"column:user_birthday;comment:生日"`
	UserHeight   float64   `json:"user_height" db:"user_height" gorm:"column:user_height;comment:身高"`
	UserWeight   float64   `json:"user_weight" db:"user_weight" gorm:"column:user_weight;comment:体重"`
}

func (UserProfession) TableName() string {
	return "user_profession"
}

type UserProfessionForPlatform struct {
	entity.User
}

func (UserProfessionForPlatform) TableName() string {
	return "user"
}

type UserProfessionWithEnterpriseNameAndStatus struct {
	entity.UserProfession
	EnterpriseName string `json:"enterprise_name" gorm:"column:enterprise_name;comment:"`
	Status         int32  `json:"status" gorm:"column:status;comment:"`
}

func (UserProfessionWithEnterpriseNameAndStatus) TableName() string {
	return "user_profession"
}

type UserProfessionForSupervisePlatformQuery struct {
	UserName   string  `json:"user_name" db:"user_name" gorm:"column:user_name;comment:用户名称"`
	UserPhone  string  `json:"user_phone" db:"user_phone" gorm:"column:user_phone;comment:用户手机"`
	UserStatus int32   `json:"user_status" db:"user_status" gorm:"column:user_status;comment:人员状态  1在职 2离职"`
	UserGender int32   `json:"user_gender" db:"user_gender" gorm:"column:user_gender;comment:"`
	ProjectPks []int64 `json:"project_pks" db:"project_pks" gorm:"column:project_pks;comment:"`
}
