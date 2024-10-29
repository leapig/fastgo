package entity

type Enterprise struct {
	BaseEntity
	Pk          int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	Cover       string `json:"cover" db:"cover"  gorm:"column:cover;comment:封面"`
	Name        string `json:"name" db:"name"  gorm:"column:name;comment:名称"`
	Serial      string `json:"serial" db:"serial"  gorm:"column:serial;comment:统一社会信用代码"`
	StaffSize   string `json:"staff_size" db:"staff_size"  gorm:"column:staff_size;comment:人员规模"`
	License     string `json:"license" db:"license"  gorm:"column:license;comment:营业执照"`
	Country     string `json:"country" db:"country"  gorm:"column:country;comment:国"`
	Province    string `json:"province" db:"province"  gorm:"column:province;comment:省"`
	City        string `json:"city" db:"city"  gorm:"column:city;comment:市"`
	District    string `json:"district" db:"district"  gorm:"column:district;comment:区"`
	County      string `json:"county" db:"county"  gorm:"column:county;comment:镇"`
	Site        string `json:"site" db:"site"  gorm:"column:site;comment:地址"`
	Longitude   string `json:"longitude" db:"longitude"  gorm:"column:longitude;comment:经度"`
	Latitude    string `json:"latitude" db:"latitude"  gorm:"column:latitude;comment:纬度"`
	CorporatePk int64  `json:"corporate_pk" db:"corporate_pk"  gorm:"column:corporate_pk;comment:负责人主键"`
	Type        int32  `json:"type" db:"type"  gorm:"column:type;comment:类型"`
	AddressCode string `json:"address_code" db:"address_code"  gorm:"column:address_code;comment:地址编码"`
}

func (Enterprise) TableName() string {
	return "enterprise"
}
