package model

import "github.com/leapig/fastgo/app/dal/entity"

type PermissionGroupPermissionModel struct {
	entity.BaseEntity
	Pk                int64              `json:"pk"            db:"pk"              gorm:"column:pk;comment:业务主键"`
	PermissionGroupPk int64              `json:"permission_group_pk"            db:"permission_group_pk"              gorm:"column:permission_group_pk;comment:权限组Pk"`
	PermissionPk      int64              `json:"permission_pk"  db:"permission_pk"   gorm:"column:permission_pk;comment:单一权限Pk"`
	Permission        *entity.Permission `json:"permission" gorm:"foreignKey:pk;references:permission_pk"`
}

func (PermissionGroupPermissionModel) TableName() string {
	return "permission_group_permission"
}
