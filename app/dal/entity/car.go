package entity

type Car struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
	Name         string `json:"name" db:"name" gorm:"column:name;comment:车辆名称"`
	Picture      string `json:"picture" db:"picture" gorm:"column:picture;comment:车辆照片"`
	Color        string `json:"color" db:"color" gorm:"column:color;comment:车辆颜色"`
	Licence      string `json:"licence" db:"license" gorm:"license"`
	Status       int32  `json:"status" db:"status" gorm:"column:status;comment:车辆状态1正常2停用"`
}

func (Car) TableName() string {
	return "car"
}
