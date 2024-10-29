package entity

type InterfaceResource struct {
	BaseEntity
	Pk            int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	InterfaceKey  string `json:"interface_key" db:"interface_key" gorm:"column:interface_key;comment:接口标识符"`
	InterfaceName string `json:"interface_name" db:"interface_name" gorm:"column:interface_name;comment:接口名"`
	InterfaceWay  string `json:"interface_way" db:"interface_way" gorm:"column:interface_way;comment:接口方法"`
	InterfaceUrl  string `json:"interface_url" db:"interface_url" gorm:"column:interface_url;comment:接口路径"`
}

func (InterfaceResource) TableName() string {
	return "interface_resource"
}
