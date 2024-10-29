package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type RoleGroup interface {
	Create(*entity.RoleGroup) (*entity.RoleGroup, error)
	Delete(*entity.RoleGroup) error
	Select(*entity.RoleGroup, *Pagination) ([]*entity.RoleGroup, error)
	Count(*entity.RoleGroup) (int32, error)
	Update(*entity.RoleGroup) (*entity.RoleGroup, error)
	FindByPk(*entity.RoleGroup) (*entity.RoleGroup, error)
	SelectWithRoleMessage(en *entity.RoleGroup, pg *Pagination) ([]*model.RoleGroupWithRoleMessageModel, error)
	SelectAllRoleWithRoleMessage(en *entity.RoleGroup) ([]*model.RoleGroupWithRoleMessageModel, error)
}
type roleGroup struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewRoleGroup(db *gorm.DB, rs *helper.Redis) RoleGroup {
	return &roleGroup{
		db: db,
		rs: rs,
	}
}
func (o *roleGroup) Create(en *entity.RoleGroup) (*entity.RoleGroup, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *roleGroup) Delete(uu *entity.RoleGroup) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *roleGroup) Select(en *entity.RoleGroup, pg *Pagination) ([]*entity.RoleGroup, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleGroupName != "" {
		sql = sql.Where("role_group_name like ?", "%"+en.RoleGroupName+"%")
	}
	var rows []*entity.RoleGroup
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *roleGroup) Count(en *entity.RoleGroup) (int32, error) {
	sql := o.db.Model(&entity.RoleGroup{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleGroupName != "" {
		sql = sql.Where("role_group_name like ?", "%"+en.RoleGroupName+"%")
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *roleGroup) Update(log *entity.RoleGroup) (*entity.RoleGroup, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *roleGroup) FindByPk(en *entity.RoleGroup) (*entity.RoleGroup, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *roleGroup) SelectWithRoleMessage(en *entity.RoleGroup, pg *Pagination) ([]*model.RoleGroupWithRoleMessageModel, error) {
	sql := o.db.Model(&model.RoleGroupWithRoleMessageModel{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	//sql.Preload("RoleGroupPermission")
	//sql.Preload("RoleGroupPermission.Role")
	if en.EnterprisePk != 0 {
		sql = sql.Where("role_group.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleGroupName != "" {
		sql = sql.Where("role_group.role_group_name like ?", "%"+en.RoleGroupName+"%")
	}
	var rows []*model.RoleGroupWithRoleMessageModel
	tx := sql.Order("role_group.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *roleGroup) SelectAllRoleWithRoleMessage(en *entity.RoleGroup) ([]*model.RoleGroupWithRoleMessageModel, error) {
	sql := o.db.Model(&model.RoleGroupWithRoleMessageModel{})
	//sql.Preload("RoleGroupPermission")
	//sql.Preload("RoleGroupPermission.Role")
	if en.EnterprisePk != 0 {
		sql = sql.Where("role_group.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RoleGroupName != "" {
		sql = sql.Where("role_group.role_group_name like ?", "%"+en.RoleGroupName+"%")
	}
	var rows []*model.RoleGroupWithRoleMessageModel
	tx := sql.Order("role_group.create_at desc").Find(&rows)
	return rows, tx.Error
}
