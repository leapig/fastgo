package entity

type Files struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:业务主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户主键"`
	Name         string `json:"name" db:"name" gorm:"column:name;comment:文件名称"`
	Size         int64  `json:"size" db:"size" gorm:"column:size;comment:文件大小"`
	Extension    string `json:"extension" db:"extension" gorm:"column:extension;comment:文件扩展"`
	Suffix       string `json:"suffix" db:"suffix" gorm:"column:suffix;comment:文件后缀"`
	Type         int64  `json:"type" db:"type" gorm:"column:type;comment:文件类型"`
	Sn           string `json:"sn" db:"sn" gorm:"column:sn;comment:设备标识"`
}

// TableName 表名称
func (Files) TableName() string {
	return "files"
}
