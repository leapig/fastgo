package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type LogisticsInventory interface {
	Select(*entity.LogisticsInventory, int32, *entity.Pagination) ([]*model.LogisticsInventory, error)
	Count(*entity.LogisticsInventory, int32) (int32, error)
	Create(*entity.LogisticsInventory) (*entity.LogisticsInventory, error)
	Delete(*entity.LogisticsInventory) error
	Update(*entity.LogisticsInventory) error
	Find(*entity.LogisticsInventory) (*model.LogisticsInventory, error)
}
type logisticsInventory struct {
	db *gorm.DB
}

func NewLogisticsInventory(db *gorm.DB) LogisticsInventory {
	return &logisticsInventory{
		db: db,
	}
}

func (o *logisticsInventory) Select(en *entity.LogisticsInventory, warning int32, pg *entity.Pagination) ([]*model.LogisticsInventory, error) {
	sql := o.db.Model(&model.LogisticsInventory{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("PrincipalEnterpriseUser")
	var rows []*model.LogisticsInventory
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Supplier != "" {
		sql = sql.Where("supplier like ?", "%"+en.Supplier+"%")
	}
	if en.Brand != "" {
		sql = sql.Where("brand like ?", "%"+en.Brand+"%")
	}
	if en.Commodity != "" {
		sql = sql.Where("commodity like ?", "%"+en.Commodity+"%")
	}
	if warning == 1 {
		sql = sql.Where("inventory_num<=inventory_warning_num")
	} else if warning == 2 {
		sql = sql.Where("inventory_num>inventory_warning_num")
	}
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *logisticsInventory) Count(en *entity.LogisticsInventory, warning int32) (int32, error) {
	sql := o.db.Model(&entity.LogisticsInventory{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Supplier != "" {
		sql = sql.Where("supplier like ?", "%"+en.Supplier+"%")
	}
	if en.Brand != "" {
		sql = sql.Where("brand like ?", "%"+en.Brand+"%")
	}
	if en.Commodity != "" {
		sql = sql.Where("commodity like ?", "%"+en.Commodity+"%")
	}
	if warning == 1 {
		sql = sql.Where("inventory_num<=inventory_warning_num")
	} else if warning == 2 {
		sql = sql.Where("inventory_num>inventory_warning_num")
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *logisticsInventory) Create(inventory *entity.LogisticsInventory) (*entity.LogisticsInventory, error) {
	inventory.Pk = helper.GetRid(helper.LogisticsPurchase)
	tx := o.db.Create(&inventory)
	return inventory, tx.Error
}

func (o *logisticsInventory) Delete(inventory *entity.LogisticsInventory) error {
	tx := o.db.Where("pk=?", inventory.Pk).Delete(&inventory)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *logisticsInventory) Update(inventory *entity.LogisticsInventory) error {
	tx := o.db.Where("pk=? and enterprise_pk=?", inventory.Pk, inventory.EnterprisePk).Updates(&inventory)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *logisticsInventory) Find(inventory *entity.LogisticsInventory) (*model.LogisticsInventory, error) {
	sql := o.db.Model(&model.LogisticsInventory{})
	sql.Preload("PrincipalEnterpriseUser")
	var row *model.LogisticsInventory
	if inventory.Pk != 0 {
		sql = sql.Where("pk = ?", inventory.Pk)
	}
	if inventory.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", inventory.EnterprisePk)
	}
	tx := sql.First(&row)
	return row, tx.Error
}
