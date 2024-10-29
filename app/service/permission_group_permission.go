package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type PermissionGroupPermission interface {
	Create(*entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error)
	Delete(*entity.PermissionGroupPermission) error
	Select(*entity.PermissionGroupPermission, *dao.Pagination) ([]*entity.PermissionGroupPermission, *dao.Pagination, error)
	Update(*entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error)
	SelectAllPermissionGroupPermission(in *model.PermissionGroupPermissionModel) ([]*model.PermissionGroupPermissionModel, error)
}

// PermissionGroupPermission 接口规范实现类
type permissionGroupPermission struct {
	dao dao.Dao
}

// NewPermissionGroupPermission 实例化接口规范实现类
func NewPermissionGroupPermission(dao dao.Dao) PermissionGroupPermission {
	return &permissionGroupPermission{dao: dao}
}
func (o *permissionGroupPermission) Create(in *entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error) {
	//result, err := o.dao.PermissionGroup().FindByPk(&entity.PermissionGroup{
	//	Pk: in.PermissionGroupPk,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result.EnterprisePk, 0); err != nil {
	//	logger.Error(err)
	//}
	in.Pk = helper.GetRid(helper.PermissionGroupPermission)
	return o.dao.PermissionGroupPermission().Create(in)
}
func (o *permissionGroupPermission) Delete(in *entity.PermissionGroupPermission) error {
	//result, err := o.dao.PermissionGroupPermission().FindByPk(&entity.PermissionGroupPermission{
	//	Pk: in.Pk,
	//})
	//if err != nil {
	//	return err
	//}
	//result1, err := o.dao.PermissionGroup().FindByPk(&entity.PermissionGroup{
	//	Pk: result.PermissionGroupPk,
	//})
	//if err != nil {
	//	return err
	//}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result1.EnterprisePk, 0); err != nil {
	//	logger.Error(err)
	//}
	return o.dao.PermissionGroupPermission().Delete(in)
}
func (o *permissionGroupPermission) Select(in *entity.PermissionGroupPermission, pg *dao.Pagination) ([]*entity.PermissionGroupPermission, *dao.Pagination, error) {
	if rows, err := o.dao.PermissionGroupPermission().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.PermissionGroupPermission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *permissionGroupPermission) Update(in *entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error) {
	return o.dao.PermissionGroupPermission().Update(in)
}
func (o *permissionGroupPermission) SelectAllPermissionGroupPermission(in *model.PermissionGroupPermissionModel) ([]*model.PermissionGroupPermissionModel, error) {
	return o.dao.PermissionGroupPermission().SelectAllPermissionGroupPermission(in)
}
