package model

import "github.com/leapig/fastgo/app/dal/entity"

type PermissionGroupModel struct {
	entity.BaseEntity
	Pk                             int64                             `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	GroupName                      string                            `json:"group_name"            db:"group_name"              gorm:"column:group_name;comment:权限组名称"`
	Remark                         string                            `json:"remark"  db:"remark"   gorm:"column:remark;comment:备注"`
	EnterprisePk                   int64                             `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:企业主键"`
	GroupType                      int32                             `json:"group_type" db:"group_type" gorm:"column:group_type;comment:"`
	PermissionGroupPermissionModel []*PermissionGroupPermissionModel `json:"permission_group_permission_model" gorm:"foreignKey:permission_group_pk;references:pk"`
}

func (PermissionGroupModel) TableName() string {
	return "permission_group"
}
