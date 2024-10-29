package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type RoleGroup interface {
	Create(*entity.RoleGroup) (*entity.RoleGroup, error)
	Delete(*entity.RoleGroup) error
	Select(*entity.RoleGroup, *dao.Pagination) ([]*entity.RoleGroup, *dao.Pagination, error)
	Update(*entity.RoleGroup) (*entity.RoleGroup, error)
	FindByPk(en *entity.RoleGroup) (*entity.RoleGroup, error)
	SelectWithRoleMessage(in *entity.RoleGroup, pg *dao.Pagination) ([]*model.RoleGroupWithRoleMessageModel, *dao.Pagination, error)
	SelectAllRoleWithRoleMessage(in *entity.RoleGroup) ([]*model.RoleGroupWithRoleMessageModel, error)
}

// Role 接口规范实现类
type roleGroup struct {
	dao dao.Dao
}

// NewRoleGroup 实例化接口规范实现类
func NewRoleGroup(dao dao.Dao) RoleGroup {
	return &roleGroup{dao: dao}
}
func (o *roleGroup) Create(in *entity.RoleGroup) (*entity.RoleGroup, error) {
	in.Pk = helper.GetRid(helper.RoleGroup)
	return o.dao.RoleGroup().Create(in)
}
func (o *roleGroup) Delete(in *entity.RoleGroup) error {
	// 判断是否存在用户使用
	if count, err := o.dao.UserPermission().Count(&entity.UserPermission{PermissionPk: in.Pk, PermissionType: 1}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	if err := o.dao.RoleGroupPermission().DeleteByGroupPk(&entity.RoleGroupPermission{RoleGroupPk: in.Pk}); err != nil {
		return err
	}
	return o.dao.RoleGroup().Delete(in)
}
func (o *roleGroup) Select(in *entity.RoleGroup, pg *dao.Pagination) ([]*entity.RoleGroup, *dao.Pagination, error) {
	if rows, err := o.dao.RoleGroup().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.RoleGroup().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *roleGroup) Update(in *entity.RoleGroup) (*entity.RoleGroup, error) {
	return o.dao.RoleGroup().Update(in)
}
func (o *roleGroup) FindByPk(en *entity.RoleGroup) (*entity.RoleGroup, error) {
	return o.dao.RoleGroup().FindByPk(en)
}
func (o *roleGroup) SelectWithRoleMessage(in *entity.RoleGroup, pg *dao.Pagination) ([]*model.RoleGroupWithRoleMessageModel, *dao.Pagination, error) {
	if rows, err := o.dao.RoleGroup().SelectWithRoleMessage(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.RoleGroup().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *roleGroup) SelectAllRoleWithRoleMessage(in *entity.RoleGroup) ([]*model.RoleGroupWithRoleMessageModel, error) {
	if rows, err := o.dao.RoleGroup().SelectAllRoleWithRoleMessage(in); err != nil {
		return nil, err
	} else {
		return rows, err
	}
}
