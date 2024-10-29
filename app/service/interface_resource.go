package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
)

type InterfaceResource interface {
	Create(*entity.InterfaceResource) (*entity.InterfaceResource, error)
	Delete(*entity.InterfaceResource) error
	Select(*entity.InterfaceResource, *dao.Pagination) ([]*entity.InterfaceResource, *dao.Pagination, error)
	Update(*entity.InterfaceResource) (*entity.InterfaceResource, error)
	FindByPk(en *entity.InterfaceResource) (*entity.InterfaceResource, error)
}

type interfaceResource struct {
	dao dao.Dao
}

func NewInterfaceResource(dao dao.Dao) InterfaceResource {
	return &interfaceResource{dao: dao}
}
func (o *interfaceResource) Create(in *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	in.Pk = helper.GetRid(helper.InterfaceResource)
	return o.dao.InterfaceResource().Create(in)
}
func (o *interfaceResource) Delete(in *entity.InterfaceResource) error {
	if count, err := o.dao.PageInterface().Count(&entity.PageInterface{InterfacePk: in.Pk}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	return o.dao.InterfaceResource().Delete(in)
}
func (o *interfaceResource) Select(in *entity.InterfaceResource, pg *dao.Pagination) ([]*entity.InterfaceResource, *dao.Pagination, error) {
	if rows, err := o.dao.InterfaceResource().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.InterfaceResource().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *interfaceResource) Update(in *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	return o.dao.InterfaceResource().Update(in)
}
func (o *interfaceResource) FindByPk(en *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	return o.dao.InterfaceResource().FindByPk(en)
}
