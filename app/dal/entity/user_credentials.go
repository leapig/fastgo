package entity

type UserCredentials struct {
	BaseEntity
	Pk            int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	UserPk        int64  `json:"user_pk" db:"user_pk"  gorm:"column:user_pk;comment:用户主键"`
	Name          string `json:"name" db:"name"  gorm:"column:name;comment:姓名"`
	Serial        string `json:"serial" db:"serial"  gorm:"column:serial;comment:编号"`
	Type          int32  `json:"type" db:"type"  gorm:"column:type;comment:类型"`
	FrontFileName string `json:"front_file_name" db:"front_file_name"  gorm:"column:front_file_name;comment:证件正面"`
	BackFileName  string `json:"back_file_name" db:"back_file_name"  gorm:"column:back_file_name;comment:证件反面"`
}

func (UserCredentials) TableName() string {
	return "user_credentials"
}
