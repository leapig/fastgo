package entity

type UserProfession struct {
	BaseEntity
	Pk               int64 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	UserPk           int64 `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:用户主键"`
	EnterprisePk     int64 `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:主键"`
	EnterpriseUserPk int64 `json:"enterprise_user_pk" db:"enterprise_user_pk" gorm:"column:enterprise_user_pk;comment:用户主键"`
	Type             int32 `json:"type" db:"type" gorm:"column:type;comment:类型"`
}

func (UserProfession) TableName() string {
	return "user_profession"
}
