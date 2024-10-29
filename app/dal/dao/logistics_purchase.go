package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type LogisticsPurchase interface {
	Create(purchase *entity.LogisticsPurchase) (*entity.LogisticsPurchase, error)
	Delete(*entity.LogisticsPurchase) error
	Select(*entity.LogisticsPurchase, *entity.Pagination) ([]*model.LogisticsPurchase, error)
	SelectMyPurchase(*entity.LogisticsPurchase, *entity.Pagination) ([]*model.LogisticsPurchase, error)
	Count(*entity.LogisticsPurchase) (int32, error)
	CountMyPurchase(en *entity.LogisticsPurchase) (int32, error)
	Update(*entity.LogisticsPurchase) (*entity.LogisticsPurchase, error)
	FindByPk(en *entity.LogisticsPurchase) (*entity.LogisticsPurchase, error)
}
type logisticsPurchase struct {
	db *gorm.DB
}

func NewLogisticsPurchase(db *gorm.DB) LogisticsPurchase {
	return &logisticsPurchase{
		db: db,
	}
}

func (o *logisticsPurchase) Create(purchase *entity.LogisticsPurchase) (*entity.LogisticsPurchase, error) {
	purchase.Pk = helper.GetRid(helper.LogisticsPurchase)
	tx := o.db.Create(&purchase)
	return purchase, tx.Error
}

func (o *logisticsPurchase) Delete(purchase *entity.LogisticsPurchase) error {
	tx := o.db.Where("pk=?", purchase.Pk).Delete(&purchase)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *logisticsPurchase) SelectMyPurchase(en *entity.LogisticsPurchase, pg *entity.Pagination) ([]*model.LogisticsPurchase, error) {
	sql := o.db.Model(&model.LogisticsPurchase{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("ApplicantEnterpriseUser")
	sql.Joins("ManagerEnterpriseUser")
	var rows []*model.LogisticsPurchase
	if en.EnterprisePk != 0 {
		sql = sql.Where("logistics_purchase.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Applicant != 0 {
		sql = sql.Where("logistics_purchase.applicant = ?", en.Applicant)
	}
	if en.Status != 0 {
		sql.Where("logistics_purchase.status = ?", en.Status)
	}
	if en.Commodity != "" {
		sql.Where("logistics_purchase.commodity = ?", en.Commodity)
	}
	tx := sql.Order("logistics_purchase.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *logisticsPurchase) Select(en *entity.LogisticsPurchase, pg *entity.Pagination) ([]*model.LogisticsPurchase, error) {
	sql := o.db.Model(&model.LogisticsPurchase{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("ApplicantEnterpriseUser")
	sql.Joins("ManagerEnterpriseUser")
	var rows []*model.LogisticsPurchase
	if en.EnterprisePk != 0 {
		sql = sql.Where("logistics_purchase.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Applicant != 0 {
		sql = sql.Where("logistics_purchase.applicant = ?", en.Applicant)
	}
	if en.Status != 0 {
		sql.Where("logistics_purchase.status = ?", en.Status)
	}
	if en.Commodity != "" {
		sql.Where("logistics_purchase.commodity = ?", en.Commodity)
	}
	sql.Where("logistics_purchase.status != 6 and logistics_purchase.status != 5")
	tx := sql.Order("logistics_purchase.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *logisticsPurchase) Count(en *entity.LogisticsPurchase) (int32, error) {
	sql := o.db.Model(&entity.LogisticsPurchase{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Applicant != 0 {
		sql = sql.Where("applicant = ?", en.Applicant)
	}
	if en.Status != 0 {
		sql.Where("status = ?", en.Status)
	}
	if en.Commodity != "" {
		sql.Where("commodity = ?", en.Commodity)
	}
	sql.Where("status != 6 and status != 5")
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *logisticsPurchase) CountMyPurchase(en *entity.LogisticsPurchase) (int32, error) {
	sql := o.db.Model(&entity.LogisticsPurchase{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Applicant != 0 {
		sql = sql.Where("applicant = ?", en.Applicant)
	}
	if en.Status != 0 {
		sql.Where("status = ?", en.Status)
	}
	if en.Commodity != "" {
		sql.Where("commodity = ?", en.Commodity)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *logisticsPurchase) Update(en *entity.LogisticsPurchase) (*entity.LogisticsPurchase, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *logisticsPurchase) FindByPk(en *entity.LogisticsPurchase) (*entity.LogisticsPurchase, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
