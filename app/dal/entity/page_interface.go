package entity

type PageInterface struct {
	BaseEntity
	Pk            int64 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PagePk        int64 `json:"page_pk" db:"page_pk" gorm:"column:page_pk;comment:页面主键"`
	InterfacePk   int64 `json:"interface_pk" db:"interface_pk" gorm:"column:interface_pk;comment:接口主键"`
	OperationType int32 `json:"operation_type" db:"operation_type" gorm:"column:operation_type;comment:操作类型：1:读2:写"`
}

func (PageInterface) TableName() string {
	return "page_interface"
}
