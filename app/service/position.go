package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"time"
)

type Position interface {
	Create(*entity.Position) (*entity.Position, error)
	Delete(*entity.Position) error
	Select(*entity.Position, *dao.Pagination, *time.Time, *time.Time) ([]*model.PositionModel, *dao.Pagination, error)
	Update(*entity.Position) (*entity.Position, error)
	FindByPk(en *entity.Position) (*entity.Position, error)
	Count(position2 *entity.Position, startTime, endTime *time.Time) (int32, error)
}

type position struct {
	dao dao.Dao
}

func NewPosition(dao dao.Dao) Position {
	return &position{dao: dao}
}
func (o *position) Create(in *entity.Position) (*entity.Position, error) {
	return o.dao.Position().Create(in)
}
func (o *position) Delete(in *entity.Position) error {
	return o.dao.Position().Delete(in)
}
func (o *position) Select(in *entity.Position, pg *dao.Pagination, st *time.Time, et *time.Time) ([]*model.PositionModel, *dao.Pagination, error) {
	if rows, err := o.dao.Position().Select(in, pg, st, et); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Position().Count(in, st, et)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *position) Update(in *entity.Position) (*entity.Position, error) {
	return o.dao.Position().Update(in)
}
func (o *position) FindByPk(en *entity.Position) (*entity.Position, error) {
	return o.dao.Position().FindByPk(en)
}

func (o *position) Count(position2 *entity.Position, startTime, endTime *time.Time) (int32, error) {
	return o.dao.Position().Count(position2, startTime, endTime)
}
