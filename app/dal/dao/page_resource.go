package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PageResource interface {
	Create(*entity.PageResource) (*entity.PageResource, error)
	Delete(*entity.PageResource) error
	Select(*entity.PageResource, *Pagination) ([]*entity.PageResource, error)
	Count(*entity.PageResource) (int32, error)
	Update(*entity.PageResource) (*entity.PageResource, error)
	FindByPk(en *entity.PageResource) (*entity.PageResource, error)
	SelectAllPageResource(en *entity.PageResource) ([]*entity.PageResource, error)
	FindByPageResourcePkAndName(en *entity.PageResource) (*entity.PageResource, error)
	FindPageAllInterfaceDetailMessageByPk(en *entity.PageResource) (*model.PageResourceWithInterfaceMessageModel, error)
	SelectPageAllInterfaceDetailMessage(en *entity.PageResource) ([]*model.PageResourceWithInterfaceMessageModel, error)
	SelectPageInterfaceDetailMessage(en *entity.PageResource, pg *Pagination) ([]*model.PageResourceWithInterfaceMessageModel, error)
}
type pageResource struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewPageResource(db *gorm.DB, rs *helper.Redis) PageResource {
	return &pageResource{
		db: db,
		rs: rs,
	}
}

func (o *pageResource) Create(en *entity.PageResource) (*entity.PageResource, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *pageResource) Delete(uu *entity.PageResource) error {
	tx := o.db.Model(&entity.PageResource{}).Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *pageResource) Select(en *entity.PageResource, pg *Pagination) ([]*entity.PageResource, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.ComponentName != "" {
		sql = sql.Where("component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.PageName != "" {
		sql = sql.Where("page_name like ?", "%"+en.PageName+"%")
	}
	if en.PageType != 0 {
		sql = sql.Where("page_type = ?", en.PageType)
	}
	var rows []*entity.PageResource
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *pageResource) Count(en *entity.PageResource) (int32, error) {
	sql := o.db.Model(&entity.PageResource{})
	if en.ComponentName != "" {
		sql = sql.Where("component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.PagePath != "" {
		sql = sql.Where("page_path like ?", "%"+en.PagePath+"%")
	}
	if en.PageName != "" {
		sql = sql.Where("page_name like ?", "%"+en.PageName+"%")
	}
	if en.PageType != 0 {
		sql = sql.Where("page_type = ?", en.PageType)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *pageResource) Update(en *entity.PageResource) (*entity.PageResource, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *pageResource) FindByPk(en *entity.PageResource) (*entity.PageResource, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

func (o *pageResource) SelectAllPageResource(en *entity.PageResource) ([]*entity.PageResource, error) {
	sql := o.db.Model(&entity.PageResource{})
	if en.ComponentName != "" {
		sql = sql.Where("component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	if en.PageName != "" {
		sql = sql.Where("page_name like ?", "%"+en.PageName+"%")
	}
	var rows []*entity.PageResource
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *pageResource) FindByPageResourcePkAndName(en *entity.PageResource) (*entity.PageResource, error) {
	sql := o.db.Model(&entity.PageResource{})
	if en.ComponentName != "" {
		sql = sql.Where("component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	var result *entity.PageResource
	res := sql.First(&result)
	return result, res.Error
}
func (o *pageResource) FindPageAllInterfaceDetailMessageByPk(en *entity.PageResource) (*model.PageResourceWithInterfaceMessageModel, error) {
	sql := o.db.Model(&model.PageResourceWithInterfaceMessageModel{})
	sql.Preload("PageInterfaceModel")
	sql.Preload("PageInterfaceModel.InterfaceResource")
	if en.Pk != 0 {
		sql = sql.Where("page_resource.pk = ?", en.Pk)
	}
	if en.ComponentName != "" {
		sql = sql.Where("page_resource.component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.PagePath != "" {
		sql = sql.Where("page_resource.page_path like ?", "%"+en.PagePath+"%")
	}
	if en.Pk != 0 {
		sql = sql.Where("page_resource.pk = ?", en.Pk)
	}
	if en.PageName != "" {
		sql = sql.Where("page_resource.page_name like ?", "%"+en.PageName+"%")
	}
	if en.PageType != 0 {
		sql = sql.Where("page_resource.page_type = ?", en.PageType)
	}
	var result *model.PageResourceWithInterfaceMessageModel
	tx := sql.Order("page_resource.create_at desc").First(&result)
	return result, tx.Error
}
func (o *pageResource) SelectPageAllInterfaceDetailMessage(en *entity.PageResource) ([]*model.PageResourceWithInterfaceMessageModel, error) {
	sql := o.db.Model(&model.PageResourceWithInterfaceMessageModel{})
	sql.Preload("PageInterfaceModel")
	sql.Preload("PageInterfaceModel.InterfaceResource")
	if en.ComponentName != "" {
		sql = sql.Where("page_resource.component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.PageName != "" {
		sql = sql.Where("page_resource.page_name like ?", "%"+en.PageName+"%")
	}
	if en.PagePath != "" {
		sql = sql.Where("page_resource.page_path like ?", "%"+en.PagePath+"%")
	}
	if en.Pk != 0 {
		sql = sql.Where("page_resource.pk = ?", en.Pk)
	}
	if en.PageType != 0 {
		sql = sql.Where("page_resource.page_type = ?", en.PageType)
	}
	var rows []*model.PageResourceWithInterfaceMessageModel
	tx := sql.Order("page_resource.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *pageResource) SelectPageInterfaceDetailMessage(en *entity.PageResource, pg *Pagination) ([]*model.PageResourceWithInterfaceMessageModel, error) {
	sql := o.db.Model(&model.PageResourceWithInterfaceMessageModel{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("PageInterfaceModel")
	sql.Preload("PageInterfaceModel.InterfaceResource")
	if en.ComponentName != "" {
		sql = sql.Where("page_resource.component_name like ?", "%"+en.ComponentName+"%")
	}
	if en.PagePath != "" {
		sql = sql.Where("page_resource.page_path like ?", "%"+en.PagePath+"%")
	}
	if en.PageName != "" {
		sql = sql.Where("page_resource.page_name like ?", "%"+en.PageName+"%")
	}
	if en.PageType != 0 {
		sql = sql.Where("page_resource.page_type = ?", en.PageType)
	}
	var rows []*model.PageResourceWithInterfaceMessageModel
	tx := sql.Order("page_resource.create_at desc").Find(&rows)
	return rows, tx.Error
}
