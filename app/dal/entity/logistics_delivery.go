package entity

import "time"

type LogisticsDelivery struct {
	BaseEntity
	Pk           int64      `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64      `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户主键"`
	Claimant     int64      `json:"claimant" db:"claimant" gorm:"column:claimant;comment:申领人"`
	ProjectPk    int64      `json:"project_pk" db:"project_pk" gorm:"column:project_pk;comment:应用项目"`
	Reason       string     `json:"reason" db:"reason" gorm:"column:reason;comment:申领原因"`
	Manager      int64      `json:"manager" db:"manager" gorm:"column:manager;comment:经办人"`
	InventoryPk  int64      `json:"inventory_pk" db:"inventory_pk" gorm:"column:inventory_pk;comment:库存pk"`
	DeliveryNum  int32      `json:"delivery_num" db:"delivery_num" gorm:"column:delivery_num;comment:出库数量"`
	NeedBack     int32      `json:"need_back" db:"need_back" gorm:"column:need_back;comment:是否需要归还1需要2不需要"`
	BackTime     *time.Time `json:"back_time" db:"back_time" gorm:"column:back_time;comment:归还时间"`
	Status       int32      `json:"status" db:"status" gorm:"column:status;comment:单据状态1已完结2借用中"`
	Remark       string     `json:"remark" db:"remark" gorm:"column:remark;comment:备注"'`
	BackNum      int32      `json:"back_num" db:"back_num" gorm:"column:back_num;comment:归还数量"`
}

func (LogisticsDelivery) TableName() string {
	return "logistics_delivery"
}
