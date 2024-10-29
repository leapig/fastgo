package entity

import "time"

type PatrolRouteRecord struct {
	BaseEntity
	Pk               int64      `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	EnterpriseUserPk int64      `json:"enterprise_user_pk" db:"enterprise_user_pk" gorm:"column:enterprise_user_pk;comment:租户人员pk"`
	ProjectPk        int64      `json:"project_pk" db:"project_pk" gorm:"column:project_pk;comment:项目pk"`
	EnterprisePk     int64      `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户pk"`
	PatrolType       int32      `json:"patrol_type" db:"patrol_type" gorm:"column:patrol_type;comment:巡逻类型 1线路巡逻 2区域巡逻 3自由巡逻"`
	RoutePk          int64      `json:"route_pk" db:"route_pk" gorm:"column:route_pk;comment:线路pk"`
	StartTime        *time.Time `json:"start_time" db:"start_time" gorm:"column:start_time;comment:巡逻起始时间"`
	EndTime          *time.Time `json:"end_time" db:"end_time" gorm:"column:end_time;comment:巡逻结束时间"`
	Status           int32      `json:"status" db:"status" gorm:"column:status;comment:巡逻状态 1 正在巡逻 2巡逻完成 3终止"`
}

func (PatrolRouteRecord) TableName() string {
	return "patrol_route_record"
}
