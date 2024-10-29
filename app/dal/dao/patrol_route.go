package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PatrolRoute interface {
	Create(route *entity.PatrolRoute) (*entity.PatrolRoute, error)
	Delete(uu *entity.PatrolRoute) error
	Update(uu *entity.PatrolRoute) (*entity.PatrolRoute, error)
	Select(en *entity.PatrolRoute, pg *entity.Pagination) ([]*model.PatrolRoute, error)
	Count(en *entity.PatrolRoute) (int32, error)
	List(en *entity.PatrolRoute) ([]*model.PatrolRoute, error)
}
type patrolRoute struct {
	db *gorm.DB
}

func NewPatrolRoute(db *gorm.DB) PatrolRoute {
	return &patrolRoute{
		db: db,
	}
}

func (o *patrolRoute) Create(en *entity.PatrolRoute) (*entity.PatrolRoute, error) {
	en.Pk = helper.GetRid(helper.PatrolRoute)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *patrolRoute) Delete(uu *entity.PatrolRoute) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *patrolRoute) Update(en *entity.PatrolRoute) (*entity.PatrolRoute, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *patrolRoute) Select(en *entity.PatrolRoute, pg *entity.Pagination) ([]*model.PatrolRoute, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("PatrolRoutePoints")
	sql.Preload("PatrolRoutePoints.PatrolPoint")
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.PatrolType != 0 {
		sql.Where("patrol_type = ?", en.PatrolType)
	}
	var rows []*model.PatrolRoute
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *patrolRoute) Count(en *entity.PatrolRoute) (int32, error) {
	sql := o.db.Model(&entity.PatrolRoute{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.PatrolType != 0 {
		sql.Where("patrol_type = ?", en.PatrolType)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *patrolRoute) List(en *entity.PatrolRoute) ([]*model.PatrolRoute, error) {
	sql := o.db.Model(&model.PatrolRoute{})
	sql.Preload("PatrolRoutePoints")
	sql.Preload("PatrolRoutePoints.PatrolPoint")
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.PatrolType != 0 {
		sql.Where("patrol_type = ?", en.PatrolType)
	}
	var rows []*model.PatrolRoute
	tx := sql.Find(&rows)
	return rows, tx.Error
}
