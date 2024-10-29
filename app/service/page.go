package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"strconv"
)

type PageResource interface {
	Create(*entity.PageResource) (*entity.PageResource, error)
	Delete(*entity.PageResource) error
	Select(*entity.PageResource, *dao.Pagination) ([]*entity.PageResource, *dao.Pagination, error)
	Update(*entity.PageResource) (*entity.PageResource, error)
	FindByPk(en *entity.PageResource) (*entity.PageResource, error)
	SelectPageInterfaceDetailMessage(in *entity.PageResource, pg *dao.Pagination) ([]*model.PageResourceWithInterfaceMessageModel, *dao.Pagination, error)
}

type pageResource struct {
	dao dao.Dao
}

func NewPageResource(dao dao.Dao) PageResource {
	return &pageResource{dao: dao}
}
func (o *pageResource) Create(in *entity.PageResource) (*entity.PageResource, error) {
	in.Pk = helper.GetRid(helper.PageResource)
	return o.dao.PageResource().Create(in)
}
func (o *pageResource) Delete(in *entity.PageResource) error {
	if count, err := o.dao.Permission().Count(&entity.Permission{Resource: in.Pk, ResourceType: 2}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	if count, err := o.dao.MenuResource().Count(&entity.MenuResource{ResourceKey: strconv.FormatInt(in.Pk, 10), MenuType: 1}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	_ = o.dao.PageInterface().DeleteByPagePk(&entity.PageInterface{PagePk: in.Pk})
	return o.dao.PageResource().Delete(in)
}
func (o *pageResource) Select(in *entity.PageResource, pg *dao.Pagination) ([]*entity.PageResource, *dao.Pagination, error) {
	if rows, err := o.dao.PageResource().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PageResource().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *pageResource) Update(in *entity.PageResource) (*entity.PageResource, error) {
	return o.dao.PageResource().Update(in)
}
func (o *pageResource) FindByPk(en *entity.PageResource) (*entity.PageResource, error) {
	return o.dao.PageResource().FindByPk(en)
}
func (o *pageResource) SelectPageInterfaceDetailMessage(in *entity.PageResource, pg *dao.Pagination) ([]*model.PageResourceWithInterfaceMessageModel, *dao.Pagination, error) {
	if rows, err := o.dao.PageResource().SelectPageInterfaceDetailMessage(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PageResource().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
