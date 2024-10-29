package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type Permission interface {
	Create(*entity.Permission) (*entity.Permission, error)
	Delete(*entity.Permission) error
	Select(*entity.Permission, *dao.Pagination) ([]*entity.Permission, *dao.Pagination, error)
	SelectWithPageAndMenu(in *entity.Permission, pg *dao.Pagination) ([]*model.PermissionModel, *dao.Pagination, error)
	Update(*entity.Permission) (*entity.Permission, error)
	FindByPk(*entity.Permission) (*entity.Permission, error)
}

// Permission 接口规范实现类
type permission struct {
	dao dao.Dao
}

// NewPermission 实例化接口规范实现类
func NewPermission(dao dao.Dao) Permission {
	return &permission{dao: dao}
}
func (o *permission) Create(in *entity.Permission) (*entity.Permission, error) {
	in.Pk = helper.GetRid(helper.Permission)
	return o.dao.Permission().Create(in)
}
func (o *permission) Delete(in *entity.Permission) error {
	// 角色组
	if count, err := o.dao.RoleGroupPermission().Count(&entity.RoleGroupPermission{PermissionPk: in.Pk, PermissionType: 2}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	// 角色
	if count, err := o.dao.RolePermission().Count(&entity.RolePermission{PermissionPk: in.Pk, PermissionType: 1}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	// 权限组
	if count, err := o.dao.PermissionGroupPermission().Count(&entity.PermissionGroupPermission{PermissionPk: in.Pk}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	return o.dao.Permission().Delete(in)
}
func (o *permission) Select(in *entity.Permission, pg *dao.Pagination) ([]*entity.Permission, *dao.Pagination, error) {
	if rows, err := o.dao.Permission().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Permission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *permission) SelectWithPageAndMenu(in *entity.Permission, pg *dao.Pagination) ([]*model.PermissionModel, *dao.Pagination, error) {
	if in.EnterprisePk != 1 {
		enterprise, _ := o.dao.Enterprise().FindByPk(&entity.Enterprise{Pk: in.EnterprisePk})
		switch enterprise.Type {
		case 1:
			in.Visibility = 4
			break
		case 2:
			in.Visibility = 2
			break
		case 3:
			in.Visibility = 3
		default:
			in.Visibility = 1
			break
		}
	}
	if rows, err := o.dao.Permission().SelectWithPageAndMenu(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Permission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *permission) Update(in *entity.Permission) (*entity.Permission, error) {
	return o.dao.Permission().Update(in)
}
func (o *permission) FindByPk(en *entity.Permission) (*entity.Permission, error) {
	return o.dao.Permission().FindByPk(en)
}
