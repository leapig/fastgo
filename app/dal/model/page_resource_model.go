package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
)

type PageResourceWithInterfaceMessageModel struct {
	entity.BaseEntity
	Pk                 int64                 `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	PagePath           string                `json:"page_path" db:"page_path" gorm:"column:page_path;comment:路由地址"`
	Component          string                `json:"component" db:"component" gorm:"column:component;comment:组件路径"`
	ComponentName      string                `json:"component_name" db:"component_name" gorm:"column:component_name;comment:组件名"`
	IsCache            int32                 `json:"is_cache" db:"is_cache" gorm:"column:is_cache;comment:是否缓存：0不缓存 1缓存"`
	PageType           int32                 `json:"page_type" db:"page_type" gorm:"column:page_type;comment:"`
	PageName           string                `json:"page_name" db:"page_name" gorm:"column:page_name;comment:页面名"`
	PageInterfaceModel []*PageInterfaceModel `json:"page_interface_with_interface_model" gorm:"foreignKey:page_pk;references:pk"`
}

func (PageResourceWithInterfaceMessageModel) TableName() string {
	return "page_resource"
}

type PageInterfaceModel struct {
	entity.BaseEntity
	Pk                int64                     `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PagePk            int64                     `json:"page_pk" db:"page_pk" gorm:"column:page_pk;comment:页面主键"`
	InterfacePk       int64                     `json:"interface_pk" db:"interface_pk" gorm:"column:interface_pk;comment:接口主键"`
	OperationType     int32                     `json:"operation_type" db:"operation_type" gorm:"column:operation_type;comment:操作类型： 1:只读2:读写"`
	InterfaceResource *entity.InterfaceResource `json:"interface_resource" gorm:"foreignKey:pk;references:interface_pk"`
}

func (PageInterfaceModel) TableName() string {
	return "page_interface"
}

type PageInterface struct {
	entity.PageInterface
	InterfaceResource *entity.InterfaceResource `json:"interface_resource" gorm:"foreignKey:pk;references:interface_pk"`
}

func (PageInterface) TableName() string {
	return "page_interface"
}

type PageResource struct {
	Pk            int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PagePath      string `json:"page_path" db:"page_path" gorm:"column:page_path;comment:路由地址"`
	Component     string `json:"component" db:"component" gorm:"column:component;comment:组件路径"`
	ComponentName string `json:"component_name" db:"component_name" gorm:"column:component_name;comment:组件名"`
	//PageName      string `json:"page_name" db:"page_name" gorm:"column:page_name;comment:页面名"`
	//IsCache       int32  `json:"is_cache" db:"is_cache" gorm:"column:is_cache;comment:是否缓存：1不缓存 2缓存"`
	//PageType      int32  `json:"page_type" db:"page_type" gorm:"column:page_type;comment:页面类型：1平台 2 小程序"`
	OperationType int32 `json:"operation_type" db:"operation_type" gorm:"column:operation_type;comment:操作类型：1:只读2:读写"`
}

func (PageResource) TableName() string {
	return "page_resource"
}
