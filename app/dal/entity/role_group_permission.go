package entity

type RoleGroupPermission struct {
	BaseEntity
	Pk             int64 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RoleGroupPk    int64 `json:"role_group_pk"            db:"role_group_pk"              gorm:"column:role_group_pk;comment:"`
	PermissionPk   int64 `json:"permission_pk" db:"permission_pk" gorm:"column:permission_pk;comment:权限主键"`
	PermissionType int32 `json:"permission_type" db:"permission_type" gorm:"column:permission_type;comment:权限类型：1角色、2权限、3权限组"`
}

func (RoleGroupPermission) TableName() string {
	return "role_group_permission"
}
