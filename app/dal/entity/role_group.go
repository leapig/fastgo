package entity

type RoleGroup struct {
	BaseEntity
	Pk            int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RoleGroupName string `json:"role_group_name"            db:"role_group_name"              gorm:"column:role_group_name;comment:角色组名称"`
	Remark        string `json:"remark"            db:"remark"              gorm:"column:remark;comment:备注"`
	EnterprisePk  int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
}

func (RoleGroup) TableName() string {
	return "role_group"
}
