package entity

type Sms struct {
	BaseEntity
	Pk      int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	Phone   string `json:"phone" db:"phone"  gorm:"column:phone;comment:手机"`
	Content string `json:"content" db:"content"  gorm:"column:content;comment:短信内容"`
	Status  int8   `json:"status" db:"status"  gorm:"column:status;comment:状态"`
	Type    int8   `json:"type" db:"type"  gorm:"column:type;comment:类型"`
	Reason  string `json:"reason" db:"reason"  gorm:"column:reason;comment:失败原因"`
	CodeNum int32  `json:"code_num" db:"code_num"  gorm:"column:code_num;comment:短信字数"`
	Size    int32  `json:"size" db:"size"  gorm:"column:size;comment:短信条数"`
}
