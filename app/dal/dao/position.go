package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"time"
)

type Position interface {
	Create(*entity.Position) (*entity.Position, error)
	Delete(*entity.Position) error
	Select(*entity.Position, *Pagination, *time.Time, *time.Time) ([]*model.PositionModel, error)
	Count(*entity.Position, *time.Time, *time.Time) (int32, error)
	Update(*entity.Position) (*entity.Position, error)
	FindByPk(en *entity.Position) (*entity.Position, error)
}
type position struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewPosition(db *gorm.DB, rs *helper.Redis) Position {
	return &position{
		db: db,
		rs: rs,
	}
}
func (o *position) Create(en *entity.Position) (*entity.Position, error) {
	en.Pk = helper.GetRid(helper.Position)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *position) Delete(uu *entity.Position) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *position) Select(en *entity.Position, pg *Pagination, st *time.Time, et *time.Time) ([]*model.PositionModel, error) {
	sql := o.db.Model(&model.PositionModel{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("Enterprise")
	if en.Pk != 0 {
		sql.Where("position.pk = ?", en.Pk)
	}
	if en.Title != "" {
		sql = sql.Where("position.title like ?", "%"+en.Title+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("position.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Province != "" {
		sql = sql.Where("position.province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("position.city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("position.district = ?", en.District)
	}
	if en.PositionType != "" {
		sql = sql.Where("position.position_type = ?", en.PositionType)
	}
	if en.PositionAddress != "" {
		sql = sql.Where("position.position_address like ?", "%"+en.PositionAddress+"%")
	}
	if en.IsIssue != "" {
		sql = sql.Where("position.is_issue = ?", en.IsIssue)
	}
	if st != nil {
		sql.Where("position.create_at > ?", st)
	}
	if et != nil {
		sql.Where("position.create_at < ?", et)
	}
	var rows []*model.PositionModel
	tx := sql.Order("position.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *position) Count(en *entity.Position, st *time.Time, et *time.Time) (int32, error) {
	sql := o.db.Model(&model.PositionModel{})
	if en.Pk != 0 {
		sql.Where("pk = ?", en.Pk)
	}
	if en.Title != "" {
		sql = sql.Where("title like ?", "%"+en.Title+"%")
	}
	if en.Province != "" {
		sql = sql.Where("province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("district = ?", en.District)
	}
	if en.PositionType != "" {
		sql = sql.Where("position_type = ?", en.PositionType)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.PositionAddress != "" {
		sql = sql.Where("position_address like ?", "%"+en.PositionAddress+"%")
	}
	if en.IsIssue != "" {
		sql = sql.Where("is_issue = ?", en.IsIssue)
	}
	if st != nil {
		sql.Where("create_at > ?", st)
	}
	if et != nil {
		sql.Where("create_at < ?", et)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *position) Update(en *entity.Position) (*entity.Position, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *position) FindByPk(en *entity.Position) (*entity.Position, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
