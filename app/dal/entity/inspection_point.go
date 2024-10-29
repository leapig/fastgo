package entity

import "time"

type InspectionPoint struct {
	BaseEntity
	Pk                  int64      `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk        int64      `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	ProjectPk           int64      `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:主键"`
	Name                string     `json:"name" db:"name"  gorm:"column:name;comment:"`
	Address             string     `json:"address" db:"address"  gorm:"column:address;comment:"`
	Level               int32      `json:"level" db:"level"  gorm:"column:level;comment:"`
	Photo               string     `json:"photo" db:"photo"  gorm:"column:photo;comment:"`
	CheckFrequency      int32      `json:"check_frequency" db:"check_frequency"  gorm:"column:check_frequency;comment:"`
	FrequencyType       int32      `json:"frequency_type" db:"frequency_type"  gorm:"column:frequency_type;comment:"`
	FrequencyTypeNumber int32      `json:"frequency_type_number" db:"frequency_type_number"  gorm:"column:frequency_type_number;comment:"`
	IntervalTime        string     `json:"interval_time" db:"interval_time"  gorm:"column:interval_time;comment:"`
	Remark              string     `json:"remark" db:"remark"  gorm:"column:remark;comment:"`
	LastCheckTime       *time.Time `json:"last_check_time"    db:"last_check_time" gorm:"column:last_check_time;comment:"`
	Latitude            string     `json:"latitude" db:"latitude" gorm:"column:latitude;comment:纬度"`
	Longitude           string     `json:"longitude" db:"longitude" gorm:"column:longitude;comment:经度"`
}

func (InspectionPoint) TableName() string {
	return "inspection_point"
}
