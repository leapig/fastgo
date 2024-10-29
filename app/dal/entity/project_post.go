package entity

type ProjectPost struct {
	BaseEntity
	Pk                   int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk         int64  `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:企业主键"`
	ProjectPk            int64  `json:"project_pk" db:"project_pk"  gorm:"column:project_pk;comment:项目主键"`
	Name                 string `json:"name" db:"name"  gorm:"column:name;comment:名称"`
	Address              string `json:"address" db:"address" gorm:"column:address;comment:详细地址"`
	Latitude             string `json:"latitude" db:"latitude" gorm:"column:latitude;comment:纬度"`
	Longitude            string `json:"longitude" db:"longitude" gorm:"column:longitude;comment:经度"`
	AddressCode          string `json:"address_code" db:"address_code" gorm:"column:address_code;comment:地址编码"`
	Status               int32  `json:"status" db:"status" gorm:"column:status;comment:状态"`
	IsAlarm              int32  `json:"is_alarm" db:"is_alarm" gorm:"column:is_alarm;comment:"`
	CheckInType          int32  `json:"check_in_type" db:"check_in_type" gorm:"column:check_in_type;comment:"`
	CheckInRange         int32  `json:"check_in_range" db:"check_in_range"  gorm:"column:check_in_range;comment:"`
	FirstScope           int32  `json:"first_scope" db:"first_scope"  gorm:"column:first_scope;comment:"`
	SecondScope          int32  `json:"second_scope" db:"second_scope"  gorm:"column:second_scope;comment:"`
	SupportType          int32  `json:"support_type" db:"support_type" gorm:"column:support_type;comment:"`
	PostType             int32  `json:"post_type" db:"post_type" gorm:"column:post_type;comment:"`
	StillnessType        int32  `json:"stillness_type" db:"stillness_type" gorm:"column:stillness_type;comment:"`
	StillnessSwitch      int32  `json:"stillness_switch" db:"stillness_switch" gorm:"column:stillness_switch;comment:"`
	StillnessTime        string `json:"stillness_time" db:"stillness_time" gorm:"column:stillness_time;comment:"`
	StillnessReport      string `json:"stillness_report" db:"stillness_report" gorm:"column:stillness_report;comment:"`
	OutFenceType         int32  `json:"out_fence_type" db:"out_fence_type" gorm:"column:out_fence_type;comment:"`
	OutFenceSwitch       int32  `json:"out_fence_switch" db:"out_fence_switch" gorm:"column:out_fence_switch;comment:"`
	OutFirstFenceReport  string `json:"out_first_fence_report" db:"out_first_fence_report" gorm:"column:out_first_fence_report;comment:"`
	OutSecondFenceReport string `json:"out_second_fence_report" db:"out_second_fence_report" gorm:"column:out_second_fence_report;comment:"`
}

func (ProjectPost) TableName() string {
	return "project_post"
}
