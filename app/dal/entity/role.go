package entity

// TODO 待调整

type Role struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RoleName     string `json:"role_name"            db:"role_name"              gorm:"column:role_name;comment:角色名称"`
	Remark       string `json:"remark"            db:"remark"              gorm:"column:remark;comment:备注"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
}

func (Role) TableName() string {
	return "role"
}
