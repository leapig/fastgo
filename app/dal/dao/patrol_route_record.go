package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PatrolRouteRecord interface {
	Create(route *entity.PatrolRouteRecord) (*entity.PatrolRouteRecord, error)
	Delete(uu *entity.PatrolRouteRecord) error
	Update(uu *entity.PatrolRouteRecord) (*entity.PatrolRouteRecord, error)
	Select(en *entity.PatrolRouteRecord, pg *entity.Pagination) ([]*model.PatrolRouteRecord, error)
	Count(en *entity.PatrolRouteRecord) (int32, error)
	Find(en *entity.PatrolRouteRecord) (*entity.PatrolRouteRecord, error)
}

type patrolRouteRecord struct {
	db *gorm.DB
}

func NewPatrolRouteRecord(db *gorm.DB) PatrolRouteRecord {
	return &patrolRouteRecord{
		db: db,
	}
}

func (o *patrolRouteRecord) Create(en *entity.PatrolRouteRecord) (*entity.PatrolRouteRecord, error) {
	en.Pk = helper.GetRid(helper.PatrolRouteRecord)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *patrolRouteRecord) Delete(uu *entity.PatrolRouteRecord) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *patrolRouteRecord) Update(en *entity.PatrolRouteRecord) (*entity.PatrolRouteRecord, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *patrolRouteRecord) Select(en *entity.PatrolRouteRecord, pg *entity.Pagination) ([]*model.PatrolRouteRecord, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("EnterpriseUser")
	sql.Joins("PatrolRoute")
	if en.EnterprisePk != 0 {
		sql = sql.Where("patrol_route_record.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("patrol_route_record.project_pk = ?", en.ProjectPk)
	}
	if en.PatrolType != 0 {
		sql.Where("patrol_route_record.patrol_type = ?", en.PatrolType)
	}
	var rows []*model.PatrolRouteRecord
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *patrolRouteRecord) Count(en *entity.PatrolRouteRecord) (int32, error) {
	sql := o.db.Model(&entity.PatrolRouteRecord{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *patrolRouteRecord) Find(en *entity.PatrolRouteRecord) (*entity.PatrolRouteRecord, error) {
	sql := o.db.Model(&entity.PatrolRouteRecord{})
	if en.Pk != 0 {
		sql.Where("pk = ?", en.Pk)
	}
	if en.EnterpriseUserPk != 0 {
		sql.Where("enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	if en.Status != 0 {
		sql.Where("status = ?", en.Status)
	}
	tx := sql.Find(&en)
	return en, tx.Error
}
