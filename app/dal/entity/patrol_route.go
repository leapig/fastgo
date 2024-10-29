package entity

type PatrolRoute struct {
	BaseEntity
	Pk           int64  `json:"pk" db:"pk" gorm:"column:pk;comment:主键"`
	EnterprisePk int64  `json:"enterprise_pk" db:"enterprise_pk" gorm:"column:enterprise_pk;comment:租户pk"`
	Remark       string `json:"remark" db:"remark" gorm:"column:remark;comment:备注"`
	RouteName    string `json:"route_name" db:"route_name" gorm:"column:route_name;comment:巡逻路线名称"`
	ProjectPk    int64  `json:"project_pk" db:"project_pk" gorm:"column:project_pk;comment:项目pk"`
	PatrolType   int32  `json:"patrol_type" db:"patrol_type" gorm:"column:patrol_type;comment:巡逻类型 1线路巡逻 2区域巡逻"`
}

func (PatrolRoute) TableName() string {
	return "patrol_route"
}
