package entity

type EnterpriseUserAttachment struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:单位pk"`
	UserPk       int64  `json:"user_pk" db:"user_pk" gorm:"column:user_pk;comment:人员主键"`
	FileName     string `json:"file_name" db:"file_name" gorm:"column:file_name;comment:文件名称"`
	FileType     string `json:"file_type" db:"file_type" gorm:"column:file_type;comment:文件类型"`
}

func (EnterpriseUserAttachment) TableName() string {
	return "enterprise_user_attachment"
}
