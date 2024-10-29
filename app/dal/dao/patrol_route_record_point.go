package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PatrolRouteRecordPoint interface {
	Create(route *entity.PatrolRouteRecordPoint) (*entity.PatrolRouteRecordPoint, error)
	Delete(uu *entity.PatrolRouteRecordPoint) error
	Update(uu *entity.PatrolRouteRecordPoint) (*entity.PatrolRouteRecordPoint, error)
	Select(en *entity.PatrolRouteRecordPoint, pg *entity.Pagination) ([]*model.PatrolRouteRecordPoint, error)
	Count(en *entity.PatrolRouteRecordPoint) (int32, error)
	FindLast(en *entity.PatrolRouteRecordPoint) (*entity.PatrolRouteRecordPoint, error)
	FindByPk(e *entity.PatrolRoutePoint) (*model.PatrolRouteRecordPoint, error)
	Find(en *entity.PatrolRoutePoint) (*model.PatrolRouteRecordPoint, error)
}

type patrolRouteRecordPoint struct {
	db *gorm.DB
}

func NewPatrolRouteRecordPoint(db *gorm.DB) PatrolRouteRecordPoint {
	return &patrolRouteRecordPoint{
		db: db,
	}
}

func (o *patrolRouteRecordPoint) Create(en *entity.PatrolRouteRecordPoint) (*entity.PatrolRouteRecordPoint, error) {
	en.Pk = helper.GetRid(helper.PatrolRouteRecordPoint)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *patrolRouteRecordPoint) Delete(uu *entity.PatrolRouteRecordPoint) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *patrolRouteRecordPoint) Update(en *entity.PatrolRouteRecordPoint) (*entity.PatrolRouteRecordPoint, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *patrolRouteRecordPoint) Select(en *entity.PatrolRouteRecordPoint, pg *entity.Pagination) ([]*model.PatrolRouteRecordPoint, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("PatrolRoutePoint")
	sql.Preload("PatrolRoutePoint.PatrolPoint")
	var rows []*model.PatrolRouteRecordPoint
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *patrolRouteRecordPoint) Count(en *entity.PatrolRouteRecordPoint) (int32, error) {
	sql := o.db.Model(&entity.PatrolRouteRecordPoint{})
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *patrolRouteRecordPoint) FindLast(en *entity.PatrolRouteRecordPoint) (*entity.PatrolRouteRecordPoint, error) {
	sql := o.db.Model(&entity.PatrolRouteRecordPoint{})
	sql.Where("route_record_pk = ?", en.RouteRecordPk).Order("clocking_time").Limit(1)
	tx := sql.Find(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("find fatal")
	}
	return en, tx.Error
}

func (o *patrolRouteRecordPoint) FindByPk(en *entity.PatrolRoutePoint) (*model.PatrolRouteRecordPoint, error) {
	sql := o.db.Model(&model.PatrolRoutePoint{})
	sql.Where("pk = ?", en.Pk)
	var data *model.PatrolRouteRecordPoint
	tx := sql.Find(&data)
	if tx.RowsAffected == 0 {
		return nil, errors.New("find fatal")
	}
	return data, tx.Error
}

func (o *patrolRouteRecordPoint) Find(en *entity.PatrolRoutePoint) (*model.PatrolRouteRecordPoint, error) {
	sql := o.db.Model(&model.PatrolRoutePoint{})
	if en.RoutePk != 0 {
		sql.Where("route_pk = ?", en.RoutePk)
	}
	if en.PointPk != 0 {
		sql.Where("point_pk = ?", en.PointPk)
	}
	var data *model.PatrolRouteRecordPoint
	tx := sql.Find(&data)
	if tx.RowsAffected == 0 {
		return nil, errors.New("find fatal")
	}
	return data, tx.Error
}
