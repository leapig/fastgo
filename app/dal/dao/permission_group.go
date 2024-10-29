package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PermissionGroup interface {
	Create(*entity.PermissionGroup) (*entity.PermissionGroup, error)
	Delete(*entity.PermissionGroup) error
	Select(*entity.PermissionGroup, *Pagination) ([]*entity.PermissionGroup, error)
	Count(*entity.PermissionGroup) (int32, error)
	Update(*entity.PermissionGroup) (*entity.PermissionGroup, error)
	FindByPk(en *entity.PermissionGroup) (*entity.PermissionGroup, error)
	Find(en *entity.PermissionGroup) (*entity.PermissionGroup, error)
	SelectPermissionGroupWithPermission(en *model.PermissionGroupModel, pg *Pagination) ([]*model.PermissionGroupModel, error)
	ListTenantPermissionGroup() ([]*entity.PermissionGroup, error)
	ListSupervisePermissionGroup() ([]*entity.PermissionGroup, error)
	ListUserPlatformPermissionGroup() ([]*entity.PermissionGroup, error)
}
type permissionGroup struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewPermissionGroup(db *gorm.DB, rs *helper.Redis) PermissionGroup {
	return &permissionGroup{
		db: db,
		rs: rs,
	}
}
func (o *permissionGroup) Create(en *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *permissionGroup) Delete(uu *entity.PermissionGroup) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *permissionGroup) Select(en *entity.PermissionGroup, pg *Pagination) ([]*entity.PermissionGroup, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.GroupName != "" {
		sql = sql.Where("group_name like ?", "%"+en.GroupName+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ? or enterprise_pk = 1 ", en.EnterprisePk)
	}
	if en.GroupType != 0 {
		sql = sql.Where("group_type = ?", en.GroupType)
	}
	var rows []*entity.PermissionGroup
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *permissionGroup) Count(en *entity.PermissionGroup) (int32, error) {
	sql := o.db.Model(&entity.PermissionGroup{})
	if en.GroupName != "" {
		sql = sql.Where("group_name like ?", "%"+en.GroupName+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ? or enterprise_pk = 1 ", en.EnterprisePk)
	}
	if en.GroupType != 0 {
		sql = sql.Where("group_type = ?", en.GroupType)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *permissionGroup) Update(log *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}

func (o *permissionGroup) Find(en *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	sql := o.db.Model(&entity.PermissionGroup{})
	if en.GroupName != "" {
		sql.Where("group_name = ?", en.GroupName)
	}
	tx := sql.First(&en)
	return en, tx.Error
}
func (o *permissionGroup) FindByPk(en *entity.PermissionGroup) (*entity.PermissionGroup, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *permissionGroup) SelectPermissionGroupWithPermission(en *model.PermissionGroupModel, pg *Pagination) ([]*model.PermissionGroupModel, error) {
	if en.EnterprisePk == 1 {
		sql := o.db.Model(&model.PermissionGroupModel{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
		sql.Preload("PermissionGroupPermissionModel")
		sql.Preload("PermissionGroupPermissionModel.Permission")
		sql = sql.Where("permission_group.enterprise_pk = ?", 1)
		if en.GroupName != "" {
			sql = sql.Where("permission_group.group_name like ?", "%"+en.GroupName+"%")
		}
		if en.GroupType != 0 {
			sql = sql.Where("permission_group.group_type = ?", en.GroupType)
		}
		var rows []*model.PermissionGroupModel
		tx := sql.Order("permission_group.create_at desc").Find(&rows)
		return rows, tx.Error
	} else {
		enterpriseMessage := &entity.Enterprise{}
		enterpriseSql := o.db.Model(&entity.Enterprise{}).Where("pk = ?", en.EnterprisePk).First(&enterpriseMessage)
		if enterpriseSql.Error != nil || enterpriseMessage.Type == 0 {
			return nil, errors.New("系统错误！！")
		}
		sql := o.db.Model(&model.PermissionGroupModel{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
		sql.Preload("PermissionGroupPermissionModel")
		sql.Preload("PermissionGroupPermissionModel.Permission")
		if en.EnterprisePk != 0 {
			sql = sql.Where("permission_group.enterprise_pk = ? or permission_group.enterprise_pk = 1", en.EnterprisePk)
		}
		if en.GroupName != "" {
			sql = sql.Where("permission_group.group_name like ?", "%"+en.GroupName+"%")
		}
		sql = sql.Where("permission_group.group_type = ?", en.GroupType)
		var rows []*model.PermissionGroupModel
		tx := sql.Order("permission_group.create_at desc").Find(&rows)
		return rows, tx.Error
	}
}
func (o *permissionGroup) ListTenantPermissionGroup() ([]*entity.PermissionGroup, error) {
	rows := make([]*entity.PermissionGroup, 0)
	tx := o.db.Model(&entity.PermissionGroup{}).Where("group_type = ?", 2).Find(&rows)
	return rows, tx.Error
}
func (o *permissionGroup) ListSupervisePermissionGroup() ([]*entity.PermissionGroup, error) {
	rows := make([]*entity.PermissionGroup, 0)
	tx := o.db.Model(&entity.PermissionGroup{}).Where("group_type = ?", 4).Find(&rows)
	return rows, tx.Error
}
func (o *permissionGroup) ListUserPlatformPermissionGroup() ([]*entity.PermissionGroup, error) {
	rows := make([]*entity.PermissionGroup, 0)
	tx := o.db.Model(&entity.PermissionGroup{}).Where("group_type = ?", 3).Find(&rows)
	return rows, tx.Error
}
