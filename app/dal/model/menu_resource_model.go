package model

import (
	"github.com/leapig/fastgo/app/dal/entity"
)

type MenuResourceModel struct {
	entity.BaseEntity
	Pk                                    int64                                  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	MenuType                              int32                                  `json:"menu_type" db:"menu_type" gorm:"column:menu_type;comment:菜单类型：1页面 2URL 3空"`
	MenuName                              string                                 `json:"menu_name" db:"menu_name" gorm:"column:menu_name;comment:菜单名称"`
	ResourceKey                           string                                 `json:"resource_key" db:"resource_key" gorm:"column:resource_key;comment:菜单资源"`
	ParentPk                              int64                                  `json:"parent_pk" db:"parent_pk" gorm:"column:parent_pk;comment:父级菜单PK（顶层菜单等于自身PK）"`
	Icon                                  string                                 `json:"icon" db:"icon" gorm:"column:icon;comment:图标"`
	Sort                                  int32                                  `json:"sort" db:"sort" gorm:"column:sort;comment:排序"`
	Path                                  string                                 `json:"path" db:"path" gorm:"column:path;comment:路由"`
	PageResourceWithInterfaceMessageModel *PageResourceWithInterfaceMessageModel `json:"page_resource_with_interface_message_model" gorm:"foreignKey:pk;references:resource_key"`
}

func (MenuResourceModel) TableName() string {
	return "menu_resource"
}

type MenuResource struct {
	entity.MenuResource
	PageResource *entity.PageResource `json:"page_resource" gorm:"foreignKey:pk;references:resource_key"`
}

func (MenuResource) TableName() string {
	return "menu_resource"
}

type MenuResourceForRedis struct {
	Pk            int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	MenuType      int32  `json:"menu_type" db:"menu_type" gorm:"column:menu_type;comment:菜单类型：1页面 2URL 3空"`
	MenuName      string `json:"menu_name" db:"menu_name" gorm:"column:menu_name;comment:菜单名称"`
	ResourceKey   string `json:"resource_key" db:"resource_key" gorm:"column:resource_key;comment:菜单资源"`
	ParentPk      int64  `json:"parent_pk" db:"parent_pk" gorm:"column:parent_pk;comment:父级菜单PK（顶层菜单等于自身PK）"`
	Icon          string `json:"icon" db:"icon" gorm:"column:icon;comment:图标"`
	Sort          int32  `json:"sort" db:"sort" gorm:"column:sort;comment:排序"`
	Path          string `json:"path" db:"path" gorm:"column:path;comment:路由"`
	IsCache       int32  `json:"is_cache" db:"is_cache" gorm:"column:is_cache;comment:是否缓存：1不缓存 2缓存"`
	Component     string `json:"component" db:"component" gorm:"column:component;comment:组件路径"`
	ComponentName string `json:"component_name" db:"component_name" gorm:"column:component_name;comment:组件名"`
}

func (MenuResourceForRedis) TableName() string {
	return "menu_resource"
}
