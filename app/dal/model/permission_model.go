package model

import "github.com/leapig/fastgo/app/dal/entity"

type PermissionModel struct {
	entity.BaseEntity
	Pk             int64                `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	OperationType  int32                `json:"operation_type"            db:"operation_type"              gorm:"column:operation_type;comment:操作类型： 1:只读2:读写"`
	Resource       int64                `json:"resource"  db:"resource"   gorm:"column:resource;comment:资源"`
	ResourceType   int32                `json:"resource_type"  db:"resource_type"   gorm:"column:resource_type;comment:资源类型"`
	PermissionName string               `json:"permission_name"  db:"permission_name"   gorm:"column:permission_name;comment:权限项名称"`
	EnterprisePk   int64                `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
	Visibility     int32                `json:"visibility"  db:"visibility"   gorm:"column:visibility;comment:权限可见性:1可见 2不可见"`
	PageResource   *entity.PageResource `json:"page_resource" gorm:"foreignKey:pk;references:resource"`
	MenuResource   *entity.MenuResource `json:"menu_resource" gorm:"foreignKey:pk;references:resource"`
}

func (PermissionModel) TableName() string {
	return "permission"
}
