package entity

type PageResource struct {
	BaseEntity
	Pk            int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PagePath      string `json:"page_path" db:"page_path" gorm:"column:page_path;comment:路由地址"`
	Component     string `json:"component" db:"component" gorm:"column:component;comment:组件路径"`
	ComponentName string `json:"component_name" db:"component_name" gorm:"column:component_name;comment:组件名"`
	PageName      string `json:"page_name" db:"page_name" gorm:"column:page_name;comment:页面名"`
	IsCache       int32  `json:"is_cache" db:"is_cache" gorm:"column:is_cache;comment:是否缓存：1不缓存 2缓存"`
	PageType      int32  `json:"page_type" db:"page_type" gorm:"column:page_type;comment:页面类型：1平台 2 小程序"`
}

func (PageResource) TableName() string {
	return "page_resource"
}
