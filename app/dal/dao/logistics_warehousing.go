package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"time"
)

type LogisticsWarehousing interface {
	Create(warehousing *entity.LogisticsWarehousing) (*entity.LogisticsWarehousing, error)
	Select(warehousing *entity.LogisticsWarehousing, pg *entity.Pagination) ([]*model.LogisticsWarehousing, error)
	Count(warehousing *entity.LogisticsWarehousing) (int32, error)
	FindByPk(warehousing *entity.LogisticsWarehousing) (*entity.LogisticsWarehousing, error)
	Update(warehousing *entity.LogisticsWarehousing) (*entity.LogisticsWarehousing, error)
	InputStatistics(time time.Time, time2 time.Time, pk int64) (float64, error)
	ListInputStatistics(time time.Time, time2 time.Time, pk int64) ([]*model.LogisticsWarehousing, error)
}
type logisticsWarehousing struct {
	db *gorm.DB
}

func NewLogisticsWarehousing(db *gorm.DB) LogisticsWarehousing {
	return &logisticsWarehousing{
		db: db,
	}
}

func (o *logisticsWarehousing) Create(warehousing *entity.LogisticsWarehousing) (*entity.LogisticsWarehousing, error) {
	warehousing.Pk = helper.GetRid(helper.LogisticsWarehousing)
	tx := o.db.Create(&warehousing)
	return warehousing, tx.Error
}

func (o *logisticsWarehousing) Select(en *entity.LogisticsWarehousing, pg *entity.Pagination) ([]*model.LogisticsWarehousing, error) {
	sql := o.db.Model(&model.LogisticsWarehousing{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("ApplicantEnterpriseUser")
	sql.Joins("ManagerEnterpriseUser")
	var rows []*model.LogisticsWarehousing
	if en.EnterprisePk != 0 {
		sql = sql.Where("logistics_warehousing.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Applicant != 0 {
		sql = sql.Where("logistics_warehousing.applicant = ?", en.Applicant)
	}
	if en.Status != 0 {
		sql.Where("logistics_warehousing.status = ?", en.Status)
	}
	if en.Supplier != "" {
		sql.Where("logistics_warehousing.supplier = ?", "%"+en.Supplier+"%")
	}
	if en.Brand != "" {
		sql.Where("logistics_warehousing.brand like ?", "%"+en.Brand+"%")
	}
	if en.Commodity != "" {
		sql.Where("logistics_warehousing.commodity like ?", "%"+en.Commodity+"%")
	}
	if en.Category != 0 {
		sql.Where("logistics_warehousing.category = ?", en.Category)
	}
	tx := sql.Order("logistics_warehousing.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *logisticsWarehousing) Count(en *entity.LogisticsWarehousing) (int32, error) {
	sql := o.db.Model(&entity.LogisticsWarehousing{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("logistics_warehousing.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Applicant != 0 {
		sql = sql.Where("logistics_warehousing.applicant = ?", en.Applicant)
	}
	if en.Status != 0 {
		sql.Where("logistics_warehousing.status = ?", en.Status)
	}
	if en.Supplier != "" {
		sql.Where("logistics_warehousing.supplier = ?", "%"+en.Supplier+"%")
	}
	if en.Brand != "" {
		sql.Where("logistics_warehousing.brand like ?", "%"+en.Brand+"%")
	}
	if en.Commodity != "" {
		sql.Where("logistics_warehousing.commodity like ?", "%"+en.Commodity+"%")
	}
	if en.Category != 0 {
		sql.Where("logistics_warehousing.category = ?", en.Category)
	}
	if en.InventoryPk != 0 {
		sql.Where("logistics_warehousing.inventory_pk = ?", en.InventoryPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *logisticsWarehousing) FindByPk(warehousing *entity.LogisticsWarehousing) (*entity.LogisticsWarehousing, error) {
	tx := o.db.Model(&entity.LogisticsWarehousing{}).Where("pk = ?", warehousing.Pk).Find(&warehousing)
	return warehousing, tx.Error
}

func (o *logisticsWarehousing) Update(en *entity.LogisticsWarehousing) (*entity.LogisticsWarehousing, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *logisticsWarehousing) InputStatistics(startTime, endTime time.Time, pk int64) (float64, error) {
	var res float64
	sql := o.db.Model(&entity.LogisticsWarehousing{}).Select("SUM(cost_price * purchase_num)")
	tx := sql.Where("create_at > ? and create_at < ? and enterprise_pk = ? and status = 2", startTime, endTime, pk).Find(&res)
	return res, tx.Error
}

func (o *logisticsWarehousing) ListInputStatistics(startTime time.Time, endTime time.Time, pk int64) ([]*model.LogisticsWarehousing, error) {
	sql := o.db.Model(&model.LogisticsWarehousing{}).Limit(int(5)).Offset(0)
	sql.Joins("ApplicantEnterpriseUser")
	sql.Joins("ManagerEnterpriseUser")
	var rows []*model.LogisticsWarehousing
	tx := sql.Where("logistics_warehousing.create_at > ? and logistics_warehousing.create_at < ? and logistics_warehousing.enterprise_pk = ? and logistics_warehousing.status = 2", startTime, endTime, pk).
		Order("(logistics_warehousing.cost_price * logistics_warehousing.purchase_num) desc").Find(&rows)
	return rows, tx.Error
}
