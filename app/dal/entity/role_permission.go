package entity

// TODO 待调整

type RolePermission struct {
	BaseEntity
	Pk             int64 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	PermissionType int32 `json:"permission_type"            db:"permission_type"              gorm:"column:permission_type;comment:权限类型：1单一权限。2权限组"`
	PermissionPk   int64 `json:"permission_pk"  db:"key"   gorm:"column:permission_pk;comment:权限主键"`
	RolePk         int64 `json:"role_pk"  db:"role_pk"   gorm:"column:role_pk;comment:角色Pk"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
