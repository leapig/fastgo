package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type InterfaceResource interface {
	Create(*entity.InterfaceResource) (*entity.InterfaceResource, error)
	Delete(*entity.InterfaceResource) error
	Select(*entity.InterfaceResource, *Pagination) ([]*entity.InterfaceResource, error)
	Count(*entity.InterfaceResource) (int32, error)
	Update(*entity.InterfaceResource) (*entity.InterfaceResource, error)
	FindByPk(en *entity.InterfaceResource) (*entity.InterfaceResource, error)
	SelectAllInterfaceResource(en *entity.InterfaceResource) ([]*entity.InterfaceResource, error)
	FindByInterfaceResourcePkAndName(en *entity.InterfaceResource) (*entity.InterfaceResource, error)
}
type interfaceResource struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewInterfaceResource(db *gorm.DB, rs *helper.Redis) InterfaceResource {
	return &interfaceResource{
		db: db,
		rs: rs,
	}
}

func (o *interfaceResource) Create(en *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *interfaceResource) Delete(uu *entity.InterfaceResource) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *interfaceResource) Select(en *entity.InterfaceResource, pg *Pagination) ([]*entity.InterfaceResource, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.InterfaceName != "" {
		sql = sql.Where("interface_name like ?", "%"+en.InterfaceName+"%")
	}
	if en.InterfaceKey != "" {
		sql = sql.Where("interface_key like ?", "%"+en.InterfaceKey+"%")
	}
	if en.InterfaceUrl != "" {
		sql = sql.Where("interface_url like ?", "%"+en.InterfaceUrl+"%")
	}
	if en.InterfaceWay != "" {
		sql = sql.Where("interface_way = ?", en.InterfaceWay)
	}
	var rows []*entity.InterfaceResource
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *interfaceResource) Count(en *entity.InterfaceResource) (int32, error) {
	sql := o.db.Model(&entity.InterfaceResource{})
	if en.InterfaceName != "" {
		sql = sql.Where("interface_name like ?", "%"+en.InterfaceName+"%")
	}
	if en.InterfaceKey != "" {
		sql = sql.Where("interface_key like ?", "%"+en.InterfaceKey+"%")
	}
	if en.InterfaceUrl != "" {
		sql = sql.Where("interface_url like ?", "%"+en.InterfaceUrl+"%")
	}
	if en.InterfaceWay != "" {
		sql = sql.Where("interface_way = ?", en.InterfaceWay)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *interfaceResource) Update(en *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *interfaceResource) FindByPk(en *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

func (o *interfaceResource) SelectAllInterfaceResource(en *entity.InterfaceResource) ([]*entity.InterfaceResource, error) {
	sql := o.db.Model(&entity.InterfaceResource{})
	if en.InterfaceName != "" {
		sql = sql.Where("interface_name like ?", "%"+en.InterfaceName+"%")
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	var rows []*entity.InterfaceResource
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *interfaceResource) FindByInterfaceResourcePkAndName(en *entity.InterfaceResource) (*entity.InterfaceResource, error) {
	sql := o.db.Model(&entity.InterfaceResource{})
	if en.InterfaceName != "" {
		sql = sql.Where("interface_name like ?", "%"+en.InterfaceName+"%")
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	var result *entity.InterfaceResource
	res := sql.First(&result)
	return result, res.Error
}
