package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type TemporaryWorker interface {
	Create(record *entity.TemporaryWorker) (*entity.TemporaryWorker, error)
	Update(record *entity.TemporaryWorker) (*entity.TemporaryWorker, error)
	Count(en *model.TemporaryWorker) (int32, error)
	SelectDetail(en *model.TemporaryWorker, pg *entity.Pagination) ([]*model.TemporaryWorker, error)
	DeleteByEnterpriseUserPk(worker *entity.TemporaryWorker) error
}

type temporaryWorker struct {
	db *gorm.DB
}

func NewTemporaryWorker(db *gorm.DB) TemporaryWorker {
	return &temporaryWorker{
		db: db,
	}
}

func (o *temporaryWorker) Create(en *entity.TemporaryWorker) (*entity.TemporaryWorker, error) {
	en.Pk = helper.GetRid(helper.TemporaryWorker)
	tx := o.db.Create(&en)
	return en, tx.Error
}
func (o *temporaryWorker) Update(en *entity.TemporaryWorker) (*entity.TemporaryWorker, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *temporaryWorker) SelectDetail(en *model.TemporaryWorker, pg *entity.Pagination) ([]*model.TemporaryWorker, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("EnterpriseUser")
	if en.EnterpriseUser.Name != "" {
		sql.Where("EnterpriseUser.name  like ?", "%"+en.EnterpriseUser.Name+"%")
	}
	if en.EnterprisePk != 0 {
		sql.Where("temporary_worker.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.EnterpriseUserPk != 0 {
		sql.Where("temporary_worker.enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	var rows []*model.TemporaryWorker
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *temporaryWorker) Count(en *model.TemporaryWorker) (int32, error) {
	sql := o.db.Model(&model.TemporaryWorker{})
	sql.Joins("EnterpriseUser")
	if en.EnterprisePk != 0 {
		sql.Where("temporary_worker.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.EnterpriseUserPk != 0 {
		sql.Where("temporary_worker.enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *temporaryWorker) DeleteByEnterpriseUserPk(worker *entity.TemporaryWorker) error {
	tx := o.db.Where("enterprise_user_pk=?", worker.EnterpriseUserPk).Delete(&worker)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
