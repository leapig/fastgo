package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PageInterface interface {
	Create(*entity.PageInterface) (*entity.PageInterface, error)
	Delete(*entity.PageInterface) error
	DeleteByPagePk(uu *entity.PageInterface) error
	Select(*entity.PageInterface, *Pagination) ([]*entity.PageInterface, error)
	Count(*entity.PageInterface) (int32, error)
	Update(*entity.PageInterface) (*entity.PageInterface, error)
	FindByPk(en *entity.PageInterface) (*entity.PageInterface, error)
	SelectAllPageInterface(en *entity.PageInterface) ([]*entity.PageInterface, error)
	FindByPageInterfacePkAndName(en *entity.PageInterface) (*entity.PageInterface, error)
	FindByPagePkAndOperationType(en *entity.PageInterface) ([]*model.PageInterfaceModel, error)
}
type pageInterface struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewPageInterface(db *gorm.DB, rs *helper.Redis) PageInterface {
	return &pageInterface{
		db: db,
		rs: rs,
	}
}

func (o *pageInterface) Create(en *entity.PageInterface) (*entity.PageInterface, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *pageInterface) Delete(uu *entity.PageInterface) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *pageInterface) DeleteByPagePk(uu *entity.PageInterface) error {
	tx := o.db.Where("page_pk=?", uu.PagePk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *pageInterface) Select(en *entity.PageInterface, pg *Pagination) ([]*entity.PageInterface, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.OperationType != 0 {
		sql = sql.Where("operation_type = ?", en.OperationType)
	}
	if en.PagePk != 0 {
		sql = sql.Where("page_pk = ?", en.PagePk)
	}
	var rows []*entity.PageInterface
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *pageInterface) Count(en *entity.PageInterface) (int32, error) {
	sql := o.db.Model(&entity.PageInterface{})
	if en.OperationType != 0 {
		sql = sql.Where("operation_type = ?", en.OperationType)
	}
	if en.PagePk != 0 {
		sql = sql.Where("page_pk = ?", en.PagePk)
	}
	if en.InterfacePk != 0 {
		sql = sql.Where("interface_pk = ?", en.InterfacePk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *pageInterface) Update(en *entity.PageInterface) (*entity.PageInterface, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *pageInterface) FindByPk(en *entity.PageInterface) (*entity.PageInterface, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

func (o *pageInterface) SelectAllPageInterface(en *entity.PageInterface) ([]*entity.PageInterface, error) {
	sql := o.db.Model(&entity.PageInterface{})
	if en.OperationType != 0 {
		sql = sql.Where("operation_type = ?", en.OperationType)
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	if en.PagePk != 0 {
		sql = sql.Where("page_pk = ?", en.PagePk)
	}
	var rows []*entity.PageInterface
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *pageInterface) FindByPageInterfacePkAndName(en *entity.PageInterface) (*entity.PageInterface, error) {
	sql := o.db.Model(&entity.PageInterface{})
	if en.OperationType != 0 {
		sql = sql.Where("operation_type = ?", en.OperationType)
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	if en.PagePk != 0 {
		sql = sql.Where("page_pk = ?", en.PagePk)
	}
	var result *entity.PageInterface
	res := sql.First(&result)
	return result, res.Error
}
func (o *pageInterface) FindByPagePkAndOperationType(en *entity.PageInterface) ([]*model.PageInterfaceModel, error) {
	sql := o.db.Model(&model.PageInterfaceModel{})
	sql.Preload("InterfaceResource")
	if en.OperationType != 0 {
		sql = sql.Where("page_interface.operation_type = ?", en.OperationType)
	}
	if en.PagePk != 0 {
		sql = sql.Where("page_interface.page_pk = ?", en.PagePk)
	}
	var result []*model.PageInterfaceModel
	res := sql.Find(&result)
	return result, res.Error
}
