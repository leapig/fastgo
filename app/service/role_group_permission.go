package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type RoleGroupPermission interface {
	Create(*entity.RoleGroupPermission) (*entity.RoleGroupPermission, error)
	Delete(*entity.RoleGroupPermission) error
	DeleteByGroupPk(in *entity.RoleGroupPermission) error
	Select(*entity.RoleGroupPermission, *dao.Pagination) ([]*entity.RoleGroupPermission, *dao.Pagination, error)
	Update(*entity.RoleGroupPermission) (*entity.RoleGroupPermission, error)
	FindByPk(en *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error)
	SelectRoleGroupPermissionByRoleGroupPk(en *entity.RoleGroupPermission) ([]*model.RoleGroupPermissionModel, error)
	SelectAllRoleGroupPermission(en *entity.RoleGroupPermission) ([]*entity.RoleGroupPermission, error)
	SelectPermissionByRoleGroupPK(roleGroupPk int64) ([]*entity.Permission, error)
}

// Role 接口规范实现类
type roleGroupPermission struct {
	dao dao.Dao
}

// NewRoleGroupPermission 实例化接口规范实现类
func NewRoleGroupPermission(dao dao.Dao) RoleGroupPermission {
	return &roleGroupPermission{dao: dao}
}
func (o *roleGroupPermission) Create(in *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error) {
	//result, err := o.dao.RoleGroup().FindByPk(&entity.RoleGroup{
	//	Pk: in.RoleGroupPk,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result.EnterprisePk, 0); err != nil {
	//	logger.Error(err)
	//}
	in.Pk = helper.GetRid(helper.RoleGroupPermission)
	return o.dao.RoleGroupPermission().Create(in)
}
func (o *roleGroupPermission) Delete(in *entity.RoleGroupPermission) error {
	//result, err := o.dao.RoleGroupPermission().FindByPk(&entity.RoleGroupPermission{
	//	Pk: in.Pk,
	//})
	//if err != nil {
	//	return err
	//}
	//result1, err := o.dao.RoleGroup().FindByPk(&entity.RoleGroup{
	//	Pk: result.RoleGroupPk,
	//})
	//if err != nil {
	//	return err
	//}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result1.EnterprisePk, 0); err != nil {
	//	logger.Error(err)
	//}
	return o.dao.RoleGroupPermission().Delete(in)
}
func (o *roleGroupPermission) DeleteByGroupPk(in *entity.RoleGroupPermission) error {
	return o.dao.RoleGroupPermission().DeleteByGroupPk(in)
}
func (o *roleGroupPermission) Select(in *entity.RoleGroupPermission, pg *dao.Pagination) ([]*entity.RoleGroupPermission, *dao.Pagination, error) {
	if rows, err := o.dao.RoleGroupPermission().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.RoleGroupPermission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *roleGroupPermission) Update(in *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error) {
	return o.dao.RoleGroupPermission().Update(in)
}
func (o *roleGroupPermission) FindByPk(en *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error) {
	return o.dao.RoleGroupPermission().FindByPk(en)
}
func (o *roleGroupPermission) SelectRoleGroupPermissionByRoleGroupPk(en *entity.RoleGroupPermission) ([]*model.RoleGroupPermissionModel, error) {
	return o.dao.RoleGroupPermission().SelectRoleGroupPermissionByRoleGroupPk(en)
}
func (o *roleGroupPermission) SelectAllRoleGroupPermission(in *entity.RoleGroupPermission) ([]*entity.RoleGroupPermission, error) {
	if rows, err := o.dao.RoleGroupPermission().SelectAllRoleGroupPermission(in); err != nil {
		return nil, err
	} else {
		return rows, err
	}
}

// SelectPermissionByRoleGroupPK 通过权限组Pk获取权限列表
func (o *roleGroupPermission) SelectPermissionByRoleGroupPK(roleGroupPk int64) ([]*entity.Permission, error) {
	detail, err := o.dao.RoleGroupPermission().SelectPermissionByRoleGroupPK(roleGroupPk)
	if err != nil {
		return nil, err
	}
	return detail, err
}
