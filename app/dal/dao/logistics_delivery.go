package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"time"
)

type LogisticsDelivery interface {
	Select(*entity.LogisticsDelivery, *entity.Pagination) ([]*model.LogisticsDelivery, error)
	Count(*entity.LogisticsDelivery) (int32, error)
	Create(*entity.LogisticsDelivery) (*entity.LogisticsDelivery, error)
	Update(*entity.LogisticsDelivery) error
	Find(en *entity.LogisticsDelivery) (*model.LogisticsDelivery, error)
	ExpenditureStatistics(startTime, endTime time.Time, pk, status int64) (float64, error)
	ListExpenditureStatistics(startTime time.Time, endTime time.Time, pk int64, i int) ([]*model.LogisticsDelivery, error)
}
type logisticsDelivery struct {
	db *gorm.DB
}

func NewLogisticsDelivery(db *gorm.DB) LogisticsDelivery {
	return &logisticsDelivery{
		db: db,
	}
}

func (o *logisticsDelivery) Select(en *entity.LogisticsDelivery, pg *entity.Pagination) ([]*model.LogisticsDelivery, error) {
	sql := o.db.Model(&model.LogisticsDelivery{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("ClaimantEnterpriseUser")
	sql.Preload("ManagerEnterpriseUser")
	sql.Preload("LogisticsInventory")
	sql.Preload("Project")
	var rows []*model.LogisticsDelivery
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Claimant != 0 {
		sql = sql.Where("claimant = ?", en.Claimant)
	}
	if en.InventoryPk != 0 {
		sql = sql.Where("inventory_pk = ?", en.InventoryPk)
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *logisticsDelivery) Count(en *entity.LogisticsDelivery) (int32, error) {
	sql := o.db.Model(&entity.LogisticsDelivery{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Claimant != 0 {
		sql = sql.Where("claimant = ?", en.Claimant)
	}
	if en.InventoryPk != 0 {
		sql = sql.Where("inventory_pk = ?", en.InventoryPk)
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *logisticsDelivery) Create(en *entity.LogisticsDelivery) (*entity.LogisticsDelivery, error) {
	en.Pk = helper.GetRid(helper.LogisticsDelivery)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *logisticsDelivery) Update(en *entity.LogisticsDelivery) error {
	tx := o.db.Where("pk=? and enterprise_pk=?", en.Pk, en.EnterprisePk).Updates(&en)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *logisticsDelivery) Find(en *entity.LogisticsDelivery) (*model.LogisticsDelivery, error) {
	sql := o.db.Model(&model.LogisticsDelivery{})
	sql.Preload("ClaimantEnterpriseUser")
	sql.Preload("ManagerEnterpriseUser")
	sql.Preload("LogisticsInventory")
	sql.Preload("Project")
	var row *model.LogisticsDelivery
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	tx := sql.First(&row)
	return row, tx.Error
}

func (o *logisticsDelivery) ExpenditureStatistics(startTime, endTime time.Time, pk, status int64) (float64, error) {
	var res float64
	sql := o.db.Model(&model.LogisticsDelivery{}).
		Joins("Left join logistics_inventory li ON li.pk = logistics_delivery.inventory_pk").
		Select("SUM(li.cost_price * logistics_delivery.delivery_num)")
	sql.Where("logistics_delivery.create_at > ? and logistics_delivery.create_at < ? and logistics_delivery.enterprise_pk = ?", startTime, endTime, pk)
	if status != 0 {
		sql.Where("logistics_delivery.status = ?", status)
	}
	tx := sql.Find(&res)
	return res, tx.Error
}

func (o *logisticsDelivery) ListExpenditureStatistics(startTime time.Time, endTime time.Time, pk int64, status int) ([]*model.LogisticsDelivery, error) {
	sql := o.db.Model(&model.LogisticsDelivery{}).Limit(5).Offset(0)
	sql.Preload("ClaimantEnterpriseUser")
	sql.Preload("ManagerEnterpriseUser")
	sql.Preload("LogisticsInventory")
	var rows []*model.LogisticsDelivery
	sql.Where("logistics_delivery.create_at > ? and logistics_delivery.create_at < ? and logistics_delivery.enterprise_pk = ? ", startTime, endTime, pk)
	if status != 0 {
		sql.Where("logistics_delivery.status = ?", status)
	}
	tx := sql.Find(&rows)
	return rows, tx.Error
}
