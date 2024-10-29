package model

import "github.com/leapig/fastgo/app/dal/entity"

type RoleGroupWithRoleMessageModel struct {
	entity.BaseEntity
	Pk                  int64                       `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RoleGroupName       string                      `json:"role_group_name"            db:"role_group_name"              gorm:"column:role_group_name;comment:角色组名称"`
	Remark              string                      `json:"remark"            db:"remark"              gorm:"column:remark;comment:备注"`
	EnterprisePk        int64                       `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
	RoleGroupPermission []*RoleGroupPermissionModel `json:"role_list" gorm:"foreignKey:role_group_pk;references:pk"`
}

func (RoleGroupWithRoleMessageModel) TableName() string {
	return "role_group"
}

type RoleGroupPermissionModel struct {
	entity.BaseEntity
	Pk              int64                   `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RoleGroupPk     int64                   `json:"role_group_pk"            db:"role_group_pk"              gorm:"column:role_group_pk;comment:"`
	PermissionPk    int64                   `json:"permission_pk" db:"permission_pk" gorm:"column:permission_pk;comment:权限主键"`
	PermissionType  int64                   `json:"permission_type" db:"permission_type" gorm:"column:permission_type;comment:权限类型：1角色、2权限、3权限组"`
	Role            *entity.Role            `json:"role" gorm:"foreignKey:pk;references:permission_pk"`
	PermissionGroup *entity.PermissionGroup `json:"permission_group" gorm:"foreignKey:pk;references:permission_pk"`
	Permission      *entity.Permission      `json:"permission" gorm:"foreignKey:pk;references:permission_pk"`
}

func (RoleGroupPermissionModel) TableName() string {
	return "role_group_permission"
}

type RolePermissionModel struct {
	entity.BaseEntity
	Pk              int64                   `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PermissionType  int32                   `json:"permission_type"            db:"permission_type"              gorm:"column:permission_type;comment:权限类型：1单一权限。2权限组"`
	PermissionPk    int64                   `json:"permission_pk"  db:"key"   gorm:"column:permission_pk;comment:权限主键"`
	RolePk          int64                   `json:"role_pk"  db:"role_pk"   gorm:"column:role_pk;comment:角色Pk"`
	PermissionGroup *entity.PermissionGroup `json:"permission_group" gorm:"foreignKey:pk;references:permission_pk"`
	Permission      *entity.Permission      `json:"permission" gorm:"foreignKey:pk;references:permission_pk"`
}

func (RolePermissionModel) TableName() string {
	return "role_permission"
}
