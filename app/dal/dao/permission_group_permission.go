package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PermissionGroupPermission interface {
	Create(*entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error)
	Delete(*entity.PermissionGroupPermission) error
	DeleteByPermissionGroupPk(uu *entity.PermissionGroupPermission) error
	Select(*entity.PermissionGroupPermission, *Pagination) ([]*entity.PermissionGroupPermission, error)
	Count(*entity.PermissionGroupPermission) (int32, error)
	Update(*entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error)
	SelectAllPermissionGroupPermission(*model.PermissionGroupPermissionModel) ([]*model.PermissionGroupPermissionModel, error)
	FindByPk(en *entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error)
}
type permissionGroupPermission struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewPermissionGroupPermission(db *gorm.DB, rs *helper.Redis) PermissionGroupPermission {
	return &permissionGroupPermission{
		db: db,
		rs: rs,
	}
}
func (o *permissionGroupPermission) Create(en *entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *permissionGroupPermission) Delete(uu *entity.PermissionGroupPermission) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *permissionGroupPermission) DeleteByPermissionGroupPk(uu *entity.PermissionGroupPermission) error {
	tx := o.db.Where("permission_group_pk=?", uu.PermissionGroupPk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *permissionGroupPermission) Select(en *entity.PermissionGroupPermission, pg *Pagination) ([]*entity.PermissionGroupPermission, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_group_pk = ?", en.PermissionPk)
	}
	var rows []*entity.PermissionGroupPermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *permissionGroupPermission) Count(en *entity.PermissionGroupPermission) (int32, error) {
	sql := o.db.Model(&entity.PermissionGroupPermission{})
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk = ?", en.PermissionPk)
	}
	if en.PermissionGroupPk != 0 {
		sql = sql.Where("permission_group_pk = ?", en.PermissionPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *permissionGroupPermission) Update(log *entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *permissionGroupPermission) SelectAllPermissionGroupPermission(en *model.PermissionGroupPermissionModel) ([]*model.PermissionGroupPermissionModel, error) {
	sql := o.db.Model(&model.PermissionGroupPermissionModel{})
	sql.Joins("Permission")
	if en.PermissionGroupPk != 0 {
		sql = sql.Where("permission_group_permission.permission_group_pk = ?", en.PermissionGroupPk)
	}
	var rows []*model.PermissionGroupPermissionModel
	tx := sql.Order("permission_group_permission.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *permissionGroupPermission) FindByPk(en *entity.PermissionGroupPermission) (*entity.PermissionGroupPermission, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
