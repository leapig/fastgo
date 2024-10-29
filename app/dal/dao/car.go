package dao

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Car interface {
	Select(*entity.Car, *entity.Pagination) ([]*entity.Car, error)
	Count(en *entity.Car) (int32, error)
}
type car struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewCar(db *gorm.DB) Car {
	return &car{
		db: db,
	}
}

func (o *car) Select(en *entity.Car, pg *entity.Pagination) ([]*entity.Car, error) {
	sql := o.db.Model(&entity.Car{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	var rows []*entity.Car
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Licence != "" {
		sql = sql.Where("licence like ?", "%"+en.Licence+"%")
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *car) Count(en *entity.Car) (int32, error) {
	sql := o.db.Model(&entity.Car{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Licence != "" {
		sql = sql.Where("licence like ?", "%"+en.Licence+"%")
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}
