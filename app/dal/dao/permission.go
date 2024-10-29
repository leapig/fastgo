package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Permission interface {
	Create(*entity.Permission) (*entity.Permission, error)
	Delete(*entity.Permission) error
	Select(*entity.Permission, *Pagination) ([]*entity.Permission, error)
	SelectWithPageAndMenu(en *entity.Permission, pg *Pagination) ([]*model.PermissionModel, error)
	Count(*entity.Permission) (int32, error)
	Update(*entity.Permission) (*entity.Permission, error)
	FindByPk(en *entity.Permission) (*entity.Permission, error)
	Find(en *entity.Permission) (*entity.Permission, error)
	ListTenantPermission() ([]*entity.Permission, error)
}
type permission struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewPermission(db *gorm.DB, rs *helper.Redis) Permission {
	return &permission{
		db: db,
		rs: rs,
	}
}
func (o *permission) Create(en *entity.Permission) (*entity.Permission, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}
func (o *permission) Delete(uu *entity.Permission) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *permission) Select(en *entity.Permission, pg *Pagination) ([]*entity.Permission, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.OperationType != 0 {
		sql = sql.Where("operation_type = ?", en.OperationType)
	}
	if en.Resource != 0 {
		sql = sql.Where("resource = ?", en.Resource)
	}
	if en.ResourceType != 0 {
		sql = sql.Where("resource_type = ?", en.ResourceType)
	}
	if en.Visibility != 0 {
		sql = sql.Where("visibility = ?", en.Visibility)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ? or enterprise_pk = 1 ", en.EnterprisePk)
	}
	if en.PermissionName != "" {
		sql = sql.Where("permission_name like ?", "%"+en.PermissionName+"%")
	}
	var rows []*entity.Permission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *permission) SelectWithPageAndMenu(en *entity.Permission, pg *Pagination) ([]*model.PermissionModel, error) {
	sql := o.db.Model(&model.PermissionModel{}).Preload("PageResource").Preload("MenuResource").Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.OperationType != 0 {
		sql = sql.Where("permission.operation_type = ?", en.OperationType)
	}
	if en.Resource != 0 {
		sql = sql.Where("permission.resource = ?", en.Resource)
	}
	if en.ResourceType != 0 {
		sql = sql.Where("permission.resource_type = ?", en.ResourceType)
	}
	if en.PermissionName != "" {
		sql = sql.Where("permission.permission_name like ?", "%"+en.PermissionName+"%")
	}
	if en.Visibility != 0 {
		sql = sql.Where("permission.visibility = ?", en.Visibility)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("permission.enterprise_pk = ? or permission.enterprise_pk = 1 ", en.EnterprisePk)
	}
	var rows []*model.PermissionModel
	tx := sql.Order("permission.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *permission) Count(en *entity.Permission) (int32, error) {
	sql := o.db.Model(&entity.Permission{})
	if en.OperationType != 0 {
		sql = sql.Where("operation_type = ?", en.OperationType)
	}
	if en.Resource != 0 {
		sql = sql.Where("resource = ?", en.Resource)
	}
	if en.ResourceType != 0 {
		sql = sql.Where("resource_type = ?", en.ResourceType)
	}
	if en.Visibility != 0 {
		sql = sql.Where("visibility = ?", en.Visibility)
	}
	if en.PermissionName != "" {
		sql = sql.Where("permission_name like ?", "%"+en.PermissionName+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ? or enterprise_pk = 1 ", en.EnterprisePk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *permission) Update(log *entity.Permission) (*entity.Permission, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *permission) FindByPk(en *entity.Permission) (*entity.Permission, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

func (o *permission) Find(en *entity.Permission) (*entity.Permission, error) {
	sql := o.db.Model(&entity.Permission{})
	if en.PermissionName != "" {
		sql.Where("permission_name = ?", en.PermissionName)
	}
	if en.Visibility != 0 {
		sql = sql.Where("visibility = ?", en.Visibility)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	tx := sql.Find(&en)
	return en, tx.Error
}

func (o *permission) ListTenantPermission() ([]*entity.Permission, error) {
	rows := make([]*entity.Permission, 0)
	tx := o.db.Model(&entity.Permission{}).Where("visibility = ?", 2).Find(&rows)
	return rows, tx.Error
}
