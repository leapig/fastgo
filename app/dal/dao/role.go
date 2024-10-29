package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Role interface {
	Create(*entity.Role) (*entity.Role, error)
	Delete(*entity.Role) error
	Select(*entity.Role, *Pagination) ([]*entity.Role, error)
	Count(*entity.Role) (int32, error)
	Update(*entity.Role) (*entity.Role, error)
	FindByPk(en *entity.Role) (*entity.Role, error)
	SelectWithDetail(en *entity.Role, pg *Pagination) ([]*model.RoleModel, error)
}
type role struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewRole(db *gorm.DB, rs *helper.Redis) Role {
	return &role{
		db: db,
		rs: rs,
	}
}
func (o *role) Create(en *entity.Role) (*entity.Role, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *role) Delete(uu *entity.Role) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *role) Select(en *entity.Role, pg *Pagination) ([]*entity.Role, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleName != "" {
		sql = sql.Where("role_name like ?", "%"+en.RoleName+"%")
	}
	var rows []*entity.Role
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *role) Count(en *entity.Role) (int32, error) {
	sql := o.db.Model(&entity.Role{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleName != "" {
		sql = sql.Where("role_name like ?", "%"+en.RoleName+"%")
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *role) Update(log *entity.Role) (*entity.Role, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *role) FindByPk(en *entity.Role) (*entity.Role, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *role) SelectWithDetail(en *entity.Role, pg *Pagination) ([]*model.RoleModel, error) {
	sql := o.db.Model(&model.RoleModel{}).Preload("RolePermissionModel").Preload("RolePermissionModel.PermissionGroup").Preload("RolePermissionModel.Permission").Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("role.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleName != "" {
		sql = sql.Where("role.role_name like ?", "%"+en.RoleName+"%")
	}
	var rows []*model.RoleModel
	tx := sql.Order("role.create_at desc").Find(&rows)
	return rows, tx.Error
}
