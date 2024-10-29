package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"time"
)

type ApplicationRecord interface {
	Create(*entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	Delete(*entity.ApplicationRecord) error
	Select(en *entity.ApplicationRecord, pg *dao.Pagination, st, et *time.Time) ([]*entity.ApplicationRecord, *dao.Pagination, error)
	Update(*entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	UpdateByUserEntryPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	FindByPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	Find(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	SelectWithDetail(in *model.ApplicationRecordModel, pg *dao.Pagination, st, et *time.Time) ([]*model.ApplicationRecordModel, *dao.Pagination, error)
	CountDetail(en *model.ApplicationRecordModel, st, et *time.Time) (int32, error)
}

type applicationRecord struct {
	dao dao.Dao
}

func NewApplicationRecord(dao dao.Dao) ApplicationRecord {
	return &applicationRecord{dao: dao}
}
func (o *applicationRecord) Create(in *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	in.Pk = helper.GetRid(helper.ApplicationRecord)
	return o.dao.ApplicationRecord().Create(in)
}
func (o *applicationRecord) Delete(in *entity.ApplicationRecord) error {
	return o.dao.ApplicationRecord().Delete(in)
}
func (o *applicationRecord) Select(in *entity.ApplicationRecord, pg *dao.Pagination, st, et *time.Time) ([]*entity.ApplicationRecord, *dao.Pagination, error) {
	if rows, err := o.dao.ApplicationRecord().Select(in, pg, st, et); err != nil {
		return nil, pg, err
	} else {
		count, err := o.dao.ApplicationRecord().Count(in, st, et)
		pg.Total = int32(count)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *applicationRecord) Update(in *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	return o.dao.ApplicationRecord().Update(in)
}

func (o *applicationRecord) UpdateByUserEntryPk(in *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	return o.dao.ApplicationRecord().UpdateByUserEntryPk(in)
}

func (o *applicationRecord) FindByPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	return o.dao.ApplicationRecord().FindByPk(en)
}
func (o *applicationRecord) Find(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	return o.dao.ApplicationRecord().Find(en)
}
func (o *applicationRecord) SelectWithDetail(in *model.ApplicationRecordModel, pg *dao.Pagination, st, et *time.Time) ([]*model.ApplicationRecordModel, *dao.Pagination, error) {
	if rows, err := o.dao.ApplicationRecord().SelectWithDetail(in, pg, st, et); err != nil {
		return nil, pg, err
	} else {
		count, err := o.dao.ApplicationRecord().CountDetail(in, st, et)
		pg.Total = int32(count)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *applicationRecord) CountDetail(en *model.ApplicationRecordModel, st, et *time.Time) (int32, error) {
	return o.dao.ApplicationRecord().CountDetail(en, st, et)
}
