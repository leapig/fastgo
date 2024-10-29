package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type RolePermission interface {
	Create(*entity.RolePermission) (*entity.RolePermission, error)
	Delete(*entity.RolePermission) error
	DeleteByRolePk(in *entity.RolePermission) error
	Select(*entity.RolePermission, *dao.Pagination) ([]*entity.RolePermission, *dao.Pagination, error)
	Update(*entity.RolePermission) (*entity.RolePermission, error)
	SelectAllRolePermission(in *entity.RolePermission) ([]*entity.RolePermission, error)
	SelectRolePermissionByRolePk(in *entity.RolePermission) ([]*model.RolePermissionModel, error)
	SelectPermissionByRolePK(in *entity.RolePermission) ([]*entity.Permission, error)
}

// RolePermission 接口规范实现类
type rolePermission struct {
	dao dao.Dao
}

// NewRolePermission 实例化接口规范实现类
func NewRolePermission(dao dao.Dao) RolePermission {
	return &rolePermission{dao: dao}
}
func (o *rolePermission) Create(in *entity.RolePermission) (*entity.RolePermission, error) {
	//result, err := o.dao.Role().FindByPk(&entity.Role{
	//	Pk: in.RolePk,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result.EnterprisePk, 0); err != nil {
	//	logger.Error(err)
	//}
	in.Pk = helper.GetRid(helper.RolePermission)
	return o.dao.RolePermission().Create(in)
}
func (o *rolePermission) Delete(in *entity.RolePermission) error {
	//result, err := o.dao.RolePermission().FindByPk(&entity.RolePermission{
	//	Pk: in.Pk,
	//})
	//if err != nil {
	//	return err
	//}
	//result1, err := o.dao.Role().FindByPk(&entity.Role{
	//	Pk: result.RolePk,
	//})
	//if err != nil {
	//	return err
	//}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result1.EnterprisePk, 0); err != nil {
	//	logger.Error(err)
	//}
	return o.dao.RolePermission().Delete(in)
}
func (o *rolePermission) DeleteByRolePk(in *entity.RolePermission) error {
	return o.dao.RolePermission().DeleteByRolePk(in)
}
func (o *rolePermission) Select(in *entity.RolePermission, pg *dao.Pagination) ([]*entity.RolePermission, *dao.Pagination, error) {
	if rows, err := o.dao.RolePermission().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.RolePermission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *rolePermission) Update(in *entity.RolePermission) (*entity.RolePermission, error) {
	return o.dao.RolePermission().Update(in)
}
func (o *rolePermission) SelectAllRolePermission(in *entity.RolePermission) ([]*entity.RolePermission, error) {
	return o.dao.RolePermission().SelectAllRolePermission(in)
}
func (o *rolePermission) SelectPermissionByRolePK(in *entity.RolePermission) ([]*entity.Permission, error) {
	detail, err := o.dao.RolePermission().SelectPermissionByRolePK(in.RolePk)
	if err != nil {
		return nil, err
	}
	return detail, err
}
func (o *rolePermission) SelectRolePermissionByRolePk(in *entity.RolePermission) ([]*model.RolePermissionModel, error) {
	return o.dao.RolePermission().SelectRolePermissionByRolePkTest(in)
}
