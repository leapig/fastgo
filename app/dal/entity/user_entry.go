package entity

import "time"

type UserEntry struct {
	BaseEntity
	Pk              int64      `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	UserPk          int64      `json:"user_pk" db:"user_pk"  gorm:"column:user_pk;comment:用户主键"`
	EnterprisePk    int64      `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:单位主键"`
	Status          string     `json:"status" db:"status"  gorm:"column:status;comment:状态"`
	ArrivalTime     *time.Time `json:"arrival_time" db:"arrival_time"  gorm:"column:arrival_time;comment:到岗时间"`
	PositionAddress string     `json:"position_address" db:"position_address"  gorm:"column:position_address;comment:地址"`
	Message         string     `json:"message" db:"message"  gorm:"column:message;comment:备注信息"`
	Longitude       string     `json:"longitude" db:"longitude"  gorm:"column:longitude;comment:经度"`
	Latitude        string     `json:"latitude" db:"latitude"  gorm:"column:latitude;comment:纬度"`
	UserType        string     `json:"user_type" db:"user_type"  gorm:"column:user_type;comment:入职人员类型 0 短期  1 长期 "`
}

func (UserEntry) TableName() string {
	return "user_entry"
}
