package entity

type PatrolRoutePoint struct {
	BaseEntity
	Pk      int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RoutePk int64  `json:"route_pk" db:"route_pk" gorm:"column:route_pk;comment:线路pk"`
	PointPk int64  `json:"point_pk" db:"point_pk" gorm:"column:point_pk;comment:点位pk"`
	Sort    int32  `json:"sort" db:"sort" gorm:"column:sort;comment:排序"`
	QrCode  string `json:"qr_code" db:"qr_code" gorm:"column:qr_code;comment:二维码"`
}

func (PatrolRoutePoint) TableName() string {
	return "patrol_route_point"
}
