package model

import "github.com/leapig/fastgo/app/dal/entity"

type RolePermission struct {
	entity.RolePermission
	Permission                *entity.Permission                `json:"permission" gorm:"foreignKey:pk;references:permission_pk"`
	PermissionGroupPermission []*PermissionGroupPermissionModel `json:"permission_group_permission" gorm:"foreignKey:permission_group_pk;references:permission_pk"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
