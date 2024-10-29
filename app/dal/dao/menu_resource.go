package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type MenuResource interface {
	Create(*entity.MenuResource) (*entity.MenuResource, error)
	Delete(*entity.MenuResource) error
	Select(*entity.MenuResource, *Pagination) ([]*model.MenuResourceModel, error)
	Count(*entity.MenuResource) (int32, error)
	Update(*entity.MenuResource) (*entity.MenuResource, error)
	FindByPk(en *entity.MenuResource) (*entity.MenuResource, error)
	SelectAllMenu(en *entity.MenuResource) ([]*model.MenuResourceModel, error)
}
type menuResource struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewMenuResource(db *gorm.DB, rs *helper.Redis) MenuResource {
	return &menuResource{
		db: db,
		rs: rs,
	}
}
func (o *menuResource) Create(en *entity.MenuResource) (*entity.MenuResource, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *menuResource) Delete(uu *entity.MenuResource) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *menuResource) Select(en *entity.MenuResource, pg *Pagination) ([]*model.MenuResourceModel, error) {
	sql := o.db.Model(&model.MenuResourceModel{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("PageResourceWithInterfaceMessageModel").Where("menu_resource.menu_type =?", 1)
	if en.MenuName != "" {
		sql = sql.Where("menu_resource.menu_name like ?", "%"+en.MenuName+"%")
	}
	if en.MenuType != 0 {
		sql = sql.Where("menu_resource.menu_type = ?", en.MenuType)
	}
	if en.ParentPk != 0 {
		sql = sql.Where("menu_resource.parent_pk = ?", en.ParentPk)
	}
	var rows []*model.MenuResourceModel
	tx := sql.Order("menu_resource.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *menuResource) Count(en *entity.MenuResource) (int32, error) {
	sql := o.db.Model(&entity.MenuResource{})
	if en.MenuName != "" {
		sql = sql.Where("menu_name like ?", "%"+en.MenuName+"%")
	}
	if en.MenuType != 0 {
		sql = sql.Where("menu_type = ?", en.MenuType)
	}
	if en.ParentPk != 0 {
		sql = sql.Where("parent_pk = ?", en.ParentPk)
	}
	if en.ResourceKey != "" {
		sql = sql.Where("resource_key = ?", en.ResourceKey)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *menuResource) Update(en *entity.MenuResource) (*entity.MenuResource, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *menuResource) FindByPk(en *entity.MenuResource) (*entity.MenuResource, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *menuResource) SelectAllMenu(en *entity.MenuResource) ([]*model.MenuResourceModel, error) {
	sql := o.db.Model(&model.MenuResourceModel{})
	if en.MenuName != "" {
		sql = sql.Where("menu_name like ?", "%"+en.MenuName+"%")
	}
	if en.MenuType != 0 {
		sql = sql.Where("menu_type = ?", en.MenuType)
	}
	var rows []*model.MenuResourceModel
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
