package model

import "github.com/leapig/fastgo/app/dal/entity"

type RoleGroupPermission struct {
	entity.RoleGroupPermission
	RolePermission            []*RolePermission                 `json:"role_permission" gorm:"foreignKey:role_pk;references:permission_pk"`
	Permission                *entity.Permission                `json:"permission" gorm:"foreignKey:pk;references:permission_pk"`
	PermissionGroupPermission []*PermissionGroupPermissionModel `json:"permission_group_permission" gorm:"foreignKey:permission_group_pk;references:permission_pk"`
}

func (RoleGroupPermission) TableName() string {
	return "role_group_permission"
}
