package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type Role interface {
	Create(*entity.Role) (*entity.Role, error)
	Delete(*entity.Role) error
	Select(*entity.Role, *dao.Pagination) ([]*entity.Role, *dao.Pagination, error)
	Update(*entity.Role) (*entity.Role, error)
	FindByPk(en *entity.Role) (*entity.Role, error)
	SelectWithDetail(in *entity.Role, pg *dao.Pagination) ([]*model.RoleModel, *dao.Pagination, error)
}

// Role 接口规范实现类
type role struct {
	dao dao.Dao
}

// NewRole 实例化接口规范实现类
func NewRole(dao dao.Dao) Role {
	return &role{dao: dao}
}
func (o *role) Create(in *entity.Role) (*entity.Role, error) {
	in.Pk = helper.GetRid(helper.Role)
	return o.dao.Role().Create(in)
}
func (o *role) Delete(in *entity.Role) error {
	if count, err := o.dao.UserPermission().Count(&entity.UserPermission{PermissionPk: in.Pk, PermissionType: 2}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	if count, err := o.dao.RoleGroupPermission().Count(&entity.RoleGroupPermission{PermissionPk: in.Pk, PermissionType: 1}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	return o.dao.Role().Delete(in)
}
func (o *role) Select(in *entity.Role, pg *dao.Pagination) ([]*entity.Role, *dao.Pagination, error) {
	if rows, err := o.dao.Role().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Role().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *role) Update(in *entity.Role) (*entity.Role, error) {
	return o.dao.Role().Update(in)
}
func (o *role) FindByPk(en *entity.Role) (*entity.Role, error) {
	return o.dao.Role().FindByPk(en)
}
func (o *role) SelectWithDetail(in *entity.Role, pg *dao.Pagination) ([]*model.RoleModel, *dao.Pagination, error) {
	if rows, err := o.dao.Role().SelectWithDetail(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Role().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
