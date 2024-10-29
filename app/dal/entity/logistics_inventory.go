package entity

type LogisticsInventory struct {
	BaseEntity
	Pk                  int64   `json:"pk" db:"pk"  gorm:"column:pk;comment:主键"`
	EnterprisePk        int64   `json:"enterprise_pk" db:"enterprise_pk"  gorm:"column:enterprise_pk;comment:租户组件"`
	Supplier            string  `json:"supplier" db:"supplier" gorm:"column:supplier;comment:供应商"`
	Brand               string  `json:"brand" db:"brand" gorm:"column:brand;comment:品牌"`
	Category            int32   `json:"category" db:"category" gorm:"column:category;comment:品类1消耗品2售卖品3固定资产4试用品"`
	Commodity           string  `json:"commodity" db:"commodity" gorm:"column:commodity;comment:商品"`
	Cover               string  `json:"cover" db:"cover" gorm:"column:cover;comment:商品图"`
	CostPrice           float64 `json:"cost_price" db:"cost_price" gorm:"column:cost_price;comment:成本单价"`
	SellingPrice        float64 `json:"selling_price" db:"selling_price" gorm:"column:selling_price;comment:销售单价"`
	PurchaseNum         int32   `json:"purchase_num" db:"purchase_num" gorm:"column:purchase_num;comment:入库数量"`
	InventoryNum        int32   `json:"inventory_num" db:"inventory_num" gorm:"column:inventory_num;comment:库存数量"`
	DeliveryNum         int32   `json:"delivery_num" db:"delivery_num" gorm:"column:delivery_num;comment:出库数量"`
	InventoryWarningNum int32   `json:"inventory_warning_num" db:"inventory_warning_num" gorm:"column:inventory_warning_num;comment:库存预警数量"`
	Principal           int64   `json:"principal" db:"principal" gorm:"column:principal;comment:出入库负责人"`
}

func (LogisticsInventory) TableName() string {
	return "logistics_inventory"
}
