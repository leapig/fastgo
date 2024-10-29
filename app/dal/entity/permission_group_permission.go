package entity

// TODO 待调整

type PermissionGroupPermission struct {
	BaseEntity
	Pk                int64 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PermissionGroupPk int64 `json:"permission_group_pk"            db:"permission_group_pk"              gorm:"column:permission_group_pk;comment:权限组Pk"`
	PermissionPk      int64 `json:"permission_pk"  db:"permission_pk"   gorm:"column:permission_pk;comment:单一权限Pk"`
}

func (PermissionGroupPermission) TableName() string {
	return "permission_group_permission"
}
