package entity

import "time"

type PatrolRouteRecordPoint struct {
	BaseEntity
	Pk                 int64      `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	RouteRecordPk      int64      `json:"route_record_pk" db:"route_record_pk" gorm:"column:route_record_pk;comment:巡逻记录关联pk"`
	Img                string     `json:"img" db:"img" gorm:"column:img;comment:图片"`
	Remark             string     `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户pk"`
	ClockingTime       *time.Time `json:"clocking_time" db:"clocking_time" gorm:"column:clocking_time;comment:打卡时间"`
	Longitude          string     `json:"longitude" db:"longitude" gorm:"column:longitude;comment:经度"`
	Latitude           string     `json:"latitude" db:"latitude" gorm:"column:latitude;comment:纬度"`
	Address            string     `json:"address" db:"address" gorm:"column:address;comment:点位地址"`
	PatrolRoutePointPk int64      `json:"patrol_route_point_pk" db:"patrol_route_point_pk" gorm:"column:patrol_route_point_pk;comment:巡逻点位关联pk"`
	PatrolPointPk      int64      `json:"patrol_point_pk" db:"patrol_point_pk" gorm:"column:patrol_point_pk;comment:巡逻点位pk"`
}

func (PatrolRouteRecordPoint) TableName() string {
	return "patrol_route_record_point"
}
