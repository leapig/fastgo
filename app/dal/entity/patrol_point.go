package entity

type PatrolPoint struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	Name         string `json:"name" db:"name" gorm:"column:name;comment:姓名"`
	Longitude    string `json:"longitude" db:"longitude" gorm:"column:longitude;comment:经度"`
	Latitude     string `json:"latitude" db:"latitude" gorm:"column:latitude;comment:纬度"`
	Address      string `json:"address" db:"address" gorm:"column:address;comment:点位地址"`
	Scope        int32  `json:"scope" db:"scope" gorm:"column:scope;comment:打卡范围"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户pk"`
	ProjectPk    int64  `json:"project_pk" db:"project_pk" gorm:"column:project_pk;comment:项目pk"`
}

func (PatrolPoint) TableName() string {
	return "patrol_point"
}
