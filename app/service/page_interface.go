package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
)

type PageInterface interface {
	Create(*entity.PageInterface) (*entity.PageInterface, error)
	Delete(*entity.PageInterface) error
	Select(*entity.PageInterface, *dao.Pagination) ([]*entity.PageInterface, *dao.Pagination, error)
	Update(*entity.PageInterface) (*entity.PageInterface, error)
	FindByPk(en *entity.PageInterface) (*entity.PageInterface, error)
}

type pageInterface struct {
	dao dao.Dao
}

func NewPageInterface(dao dao.Dao) PageInterface {
	return &pageInterface{dao: dao}
}
func (o *pageInterface) Create(in *entity.PageInterface) (*entity.PageInterface, error) {
	in.Pk = helper.GetRid(helper.PageInterface)
	return o.dao.PageInterface().Create(in)
}

// Delete 删除页面
func (o *pageInterface) Delete(in *entity.PageInterface) error {
	return o.dao.PageInterface().Delete(in)
}
func (o *pageInterface) Select(in *entity.PageInterface, pg *dao.Pagination) ([]*entity.PageInterface, *dao.Pagination, error) {
	if rows, err := o.dao.PageInterface().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PageInterface().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *pageInterface) Update(in *entity.PageInterface) (*entity.PageInterface, error) {
	return o.dao.PageInterface().Update(in)
}
func (o *pageInterface) FindByPk(en *entity.PageInterface) (*entity.PageInterface, error) {
	return o.dao.PageInterface().FindByPk(en)
}
