package entity

type EnterpriseAreaPermission struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:单位pk"`
	Province     string `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City         string `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District     string `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County       string `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	AddressCode  string `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:地址编码"`
}

func (EnterpriseAreaPermission) TableName() string {
	return "enterprise_area_permission"
}
