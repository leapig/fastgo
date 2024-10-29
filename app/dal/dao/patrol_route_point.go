package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PatrolRoutePoint interface {
	Create(route *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error)
	Delete(uu *entity.PatrolRoutePoint) error
	Update(uu *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error)
	Select(en *entity.PatrolRoutePoint, pg *entity.Pagination) ([]*entity.PatrolRoutePoint, error)
	DeleteByRoutePk(en *entity.PatrolRoutePoint) error
	Count(e *entity.PatrolRoutePoint) (int32, error)
	FindListByRoutePk(en *entity.PatrolRoutePoint) ([]*model.PatrolRoutePoint, error)
	FindByPk(point *entity.PatrolRoutePoint) (*model.PatrolRoutePoint, error)
	FindNext(e *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error)
	FindFirst(e *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error)
}
type patrolRoutePoint struct {
	db *gorm.DB
}

func NewPatrolRoutePoint(db *gorm.DB) PatrolRoutePoint {
	return &patrolRoutePoint{
		db: db,
	}
}

func (o *patrolRoutePoint) Create(en *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error) {
	en.Pk = helper.GetRid(helper.PatrolRoutePoint)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *patrolRoutePoint) Delete(uu *entity.PatrolRoutePoint) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *patrolRoutePoint) DeleteByRoutePk(en *entity.PatrolRoutePoint) error {
	tx := o.db.Where("route_pk=?", en.RoutePk).Delete(&en)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *patrolRoutePoint) Update(en *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *patrolRoutePoint) Select(en *entity.PatrolRoutePoint, pg *entity.Pagination) ([]*entity.PatrolRoutePoint, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PointPk != 0 {
		sql.Where("point_pk = ?", en.PointPk)
	}
	if en.RoutePk != 0 {
		sql.Where("route_pk = ?", en.RoutePk)
	}
	var rows []*entity.PatrolRoutePoint
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *patrolRoutePoint) FindListByRoutePk(en *entity.PatrolRoutePoint) ([]*model.PatrolRoutePoint, error) {
	sql := o.db.Model(&model.PatrolRoutePoint{})
	sql.Joins("PatrolPoint")
	sql.Where("patrol_route_point.route_pk = ?", en.RoutePk)
	var rows []*model.PatrolRoutePoint
	tx := sql.Order("patrol_route_point.sort desc").Find(&rows)
	return rows, tx.Error
}

func (o *patrolRoutePoint) Count(en *entity.PatrolRoutePoint) (int32, error) {
	sql := o.db.Model(&entity.PatrolRoutePoint{})
	if en.PointPk != 0 {
		sql.Where("point_pk = ?", en.PointPk)
	}
	if en.RoutePk != 0 {
		sql.Where("route_pk = ?", en.RoutePk)
	}
	var res int64
	tx := sql.Find(&res)
	return int32(res), tx.Error
}

func (o *patrolRoutePoint) FindByPk(en *entity.PatrolRoutePoint) (*model.PatrolRoutePoint, error) {
	sql := o.db.Model(&model.PatrolRoutePoint{})
	sql.Joins("PatrolPoint")
	sql.Where("patrol_route_point.pk = ?", en.Pk)
	var data *model.PatrolRoutePoint
	tx := sql.Find(&data)
	return data, tx.Error
}

func (o *patrolRoutePoint) FindNext(e *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error) {
	sql := o.db.Model(&entity.PatrolRoutePoint{})
	sql.Where("route_pk = ?", e.RoutePk)
	sql.Where("sort > (SELECT sort FROM patrol_route_point WHERE pk = ?)", e.Pk)
	sql.Order("sort ASC")
	tx := sql.First(e)
	if tx.RowsAffected == 0 {
		return nil, errors.New("find fatal")
	}
	return e, tx.Error
}
func (o *patrolRoutePoint) FindFirst(e *entity.PatrolRoutePoint) (*entity.PatrolRoutePoint, error) {
	sql := o.db.Model(&entity.PatrolRoutePoint{})
	sql.Where("route_pk = ?", e.RoutePk)
	sql.Order("sort ASC")
	tx := sql.First(e)
	if tx.RowsAffected == 0 {
		return nil, errors.New("find fatal")
	}
	return e, tx.Error
}
