package entity

type UserClient struct {
	BaseEntity
	Pk          int64  `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	UserPk      int64  `json:"user_pk" db:"user_pk"  gorm:"column:user_pk;comment:用户主键"`
	ClientId    string `json:"client_id" db:"client_id" gorm:"column:client_id;comment:客户端标识"`
	ClientType  int32  `json:"client_type" db:"client_type"  gorm:"column:client_type;comment:客户端类型：1微信公众号、2微信小程序"`
	OpenId      string `json:"open_id" db:"open_id"  gorm:"column:open_id;comment:客户端主键"`
	WxUnionid   string `json:"wx_unionid" db:"wx_unionid"  gorm:"column:wx_unionid;comment:微信开放平台unionId"`
	WxSubscribe *int8  `json:"wx_subscribe" db:"wx_subscribe"  gorm:"column:wx_subscribe;comment:微信公众号关注状态：1已关注、2未关注"`
}

func (UserClient) TableName() string {
	return "user_client"
}
