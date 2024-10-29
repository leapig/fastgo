package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type TurnoverRecord interface {
	Create(record *entity.TurnoverRecord) (*entity.TurnoverRecord, error)
	Update(record *entity.TurnoverRecord) (*entity.TurnoverRecord, error)
	SelectDetail(en *model.TurnoverRecord, pg *entity.Pagination) ([]*model.TurnoverRecord, error)
	Count(en *model.TurnoverRecord) (int32, error)
	Select(en *entity.TurnoverRecord, pg *entity.Pagination) ([]*entity.TurnoverRecord, error)
	Find(e *entity.TurnoverRecord) (*entity.TurnoverRecord, error)
}

type turnoverRecord struct {
	db *gorm.DB
}

func NewTurnoverRecord(db *gorm.DB) TurnoverRecord {
	return &turnoverRecord{
		db: db,
	}
}

func (o *turnoverRecord) Create(en *entity.TurnoverRecord) (*entity.TurnoverRecord, error) {
	en.Pk = helper.GetRid(helper.TurnoverRecord)
	tx := o.db.Create(&en)
	return en, tx.Error
}
func (o *turnoverRecord) Update(en *entity.TurnoverRecord) (*entity.TurnoverRecord, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *turnoverRecord) Select(en *entity.TurnoverRecord, pg *entity.Pagination) ([]*entity.TurnoverRecord, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	var rows []*entity.TurnoverRecord
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *turnoverRecord) SelectDetail(en *model.TurnoverRecord, pg *entity.Pagination) ([]*model.TurnoverRecord, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("EnterpriseUser")
	sql.Joins("Enterprise")
	sql.Joins("ApprovalEnterpriseUser")
	var rows []*model.TurnoverRecord
	if en.EnterprisePk != 0 {
		sql.Where("turnover_record.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.EnterpriseUserPk != 0 {
		sql.Where("turnover_record.enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	if en.ApplyStatus != 0 {
		sql.Where("turnover_record.apply_status = ?", en.ApplyStatus)
	}
	tx := sql.Order("turnover_record.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *turnoverRecord) Count(en *model.TurnoverRecord) (int32, error) {
	sql := o.db.Model(&entity.TurnoverRecord{})
	if en.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.EnterpriseUserPk != 0 {
		sql.Where("enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	if en.ApplyStatus != 0 {
		sql.Where("apply_status = ?", en.ApplyStatus)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *turnoverRecord) Find(e *entity.TurnoverRecord) (*entity.TurnoverRecord, error) {
	tx := o.db.Model(&entity.TurnoverRecord{}).Where("pk = ?").Find(&e)
	return e, tx.Error
}
