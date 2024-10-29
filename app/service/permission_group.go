package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type PermissionGroup interface {
	Create(*entity.PermissionGroup) (*entity.PermissionGroup, error)
	Delete(*entity.PermissionGroup) error
	Select(*entity.PermissionGroup, *dao.Pagination) ([]*entity.PermissionGroup, *dao.Pagination, error)
	Update(*entity.PermissionGroup) (*entity.PermissionGroup, error)
	FindByPk(en *entity.PermissionGroup) (*entity.PermissionGroup, error)
	Find(en *entity.PermissionGroup) (*entity.PermissionGroup, error)
	SelectPermissionGroupWithPermission(in *model.PermissionGroupModel, pg *dao.Pagination) ([]*model.PermissionGroupModel, *dao.Pagination, error)
	SelectPermissionGroupWithPermissionForEnterprise(in *model.PermissionGroupModel, pg *dao.Pagination) ([]*model.PermissionGroupModel, *dao.Pagination, error)
}

// PermissionGroup 接口规范实现类
type permissionGroup struct {
	dao dao.Dao
}

// NewPermissionGroup 实例化接口规范实现类
func NewPermissionGroup(dao dao.Dao) PermissionGroup {
	return &permissionGroup{dao: dao}
}
func (o *permissionGroup) Create(in *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	in.Pk = helper.GetRid(helper.PermissionGroup)
	return o.dao.PermissionGroup().Create(in)
}
func (o *permissionGroup) Delete(in *entity.PermissionGroup) error {
	// 角色组
	if count, err := o.dao.RoleGroupPermission().Count(&entity.RoleGroupPermission{PermissionPk: in.Pk, PermissionType: 3}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	// 角色
	if count, err := o.dao.RolePermission().Count(&entity.RolePermission{PermissionPk: in.Pk, PermissionType: 2}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	if err2 := o.dao.PermissionGroupPermission().DeleteByPermissionGroupPk(&entity.PermissionGroupPermission{
		PermissionGroupPk: in.Pk,
	}); err2 != nil {
		return err2
	}
	err := o.dao.PermissionGroup().Delete(in)
	return err
}
func (o *permissionGroup) Select(in *entity.PermissionGroup, pg *dao.Pagination) ([]*entity.PermissionGroup, *dao.Pagination, error) {
	if rows, err := o.dao.PermissionGroup().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PermissionGroup().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *permissionGroup) Update(in *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	return o.dao.PermissionGroup().Update(in)
}

func (o *permissionGroup) FindByPk(en *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	return o.dao.PermissionGroup().FindByPk(en)
}
func (o *permissionGroup) Find(en *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	return o.dao.PermissionGroup().Find(en)
}
func (o *permissionGroup) SelectPermissionGroupWithPermission(in *model.PermissionGroupModel, pg *dao.Pagination) ([]*model.PermissionGroupModel, *dao.Pagination, error) {
	if rows, err := o.dao.PermissionGroup().SelectPermissionGroupWithPermission(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PermissionGroup().Count(&entity.PermissionGroup{
			GroupName:    in.GroupName,
			EnterprisePk: in.EnterprisePk,
			GroupType:    in.GroupType,
		})
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *permissionGroup) SelectPermissionGroupWithPermissionForEnterprise(in *model.PermissionGroupModel, pg *dao.Pagination) ([]*model.PermissionGroupModel, *dao.Pagination, error) {
	enterpriseMessage, err := o.dao.Enterprise().FindByPk(&entity.Enterprise{
		Pk: in.EnterprisePk,
	})
	if err != nil {
		return nil, nil, err
	}
	var groupType int32
	if enterpriseMessage.Type == 1 {
		groupType = 4
	} else if enterpriseMessage.Type == 2 {
		groupType = 2
	} else if enterpriseMessage.Type == 3 {
		groupType = 3
	}
	in.GroupType = groupType
	if rows, err := o.dao.PermissionGroup().SelectPermissionGroupWithPermission(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PermissionGroup().Count(&entity.PermissionGroup{
			GroupName:    in.GroupName,
			EnterprisePk: in.EnterprisePk,
			GroupType:    groupType,
		})
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
