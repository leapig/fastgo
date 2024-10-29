package entity

type LogisticsWarehousing struct {
	BaseEntity
	Pk           int64   `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk int64   `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户主键"`
	Applicant    int64   `json:"applicant" db:"applicant" gorm:"column:applicant;comment:申请人"`
	Supplier     string  `json:"supplier" db:"supplier" gorm:"column:supplier;comment:供应商"`
	Brand        string  `json:"brand" db:"brand" gorm:"column:brand;comment:品牌"`
	Category     int32   `json:"category" db:"category" gorm:"column:category;comment:品类1消耗品2售卖品3固定资产4试用品"`
	Commodity    string  `json:"commodity" db:"commodity" gorm:"column:commodity;comment:商品"`
	Cover        string  `json:"cover" db:"cover" gorm:"column:cover;comment:商品图"`
	CostPrice    float64 `json:"cost_price" db:"cost_price" gorm:"column:cost_price;comment:成本单价"`
	PurchaseNum  int32   `json:"purchase_num" db:"purchase_num" gorm:"column:purchase_num;comment:入库数量"`
	Source       int32   `json:"source" db:"source" gorm:"column:source;comment:商品来源1采购2租用3试用4其他"`
	Manager      int64   `json:"manager" db:"manager" gorm:"column:manager;comment:经办人"`
	PurchasePk   int64   `json:"purchase_pk" db:"purchase_pk" gorm:"column:purchase_pk;comment:关联申请单"`
	InventoryPk  int64   `json:"inventory_pk" db:"inventory_pk" gorm:"column:inventory_pk;comment:关联库存品"`
	Status       int32   `json:"status" db:"status" gorm:"column:status;comment:入库状态：1待入库2已入库 3作废"`
}

func (LogisticsWarehousing) TableName() string {
	return "logistics_warehousing"
}
