package entity

import "time"

type Project struct {
	BaseEntity
	Pk                 int64     `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	Name               string    `json:"name" db:"name"  gorm:"column:name;comment:名称"`
	CorporatePk        int64     `json:"corporate_pk" db:"corporate_pk"  gorm:"column:corporate_pk;comment:负责人主键"`
	Address            string    `json:"address" db:"address" gorm:"column:address;comment:详细地址"`
	Latitude           string    `json:"latitude" db:"latitude" gorm:"column:latitude;comment:纬度"`
	Longitude          string    `json:"longitude" db:"longitude" gorm:"column:longitude;comment:经度"`
	CreateEnterprisePk int64     `json:"create_enterprise_pk" db:"create_enterprise_pk"  gorm:"column:create_enterprise_pk;comment:创建企业主键"`
	Province           string    `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City               string    `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District           string    `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County             string    `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	AddressCode        string    `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:"`
	OpenTime           time.Time `json:"open_time" db:"open_time"  gorm:"column:open_time;comment:"`
	FinishTime         time.Time `json:"finish_time" db:"finish_time"  gorm:"column:finish_time;comment:"`
	ProjectStatus      int32     `json:"project_status" db:"project_status"  gorm:"column:project_status;comment:"`
}

func (Project) TableName() string {
	return "project"
}
