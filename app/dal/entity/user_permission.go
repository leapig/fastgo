package entity

type UserPermission struct {
	BaseEntity
	Pk             int64 `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	UserPk         int64 `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:用户主键"`
	PermissionPk   int64 `json:"permission_pk" db:"permission_pk" gorm:"column:permission_pk;comment:权限主键"`
	PermissionType int32 `json:"permission_type" db:"permission_type" gorm:"column:permission_type;comment:权限类型1角色组、2角色"`
	EnterprisePk   int64 `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
}

func (UserPermission) TableName() string {
	return "user_permission"
}
