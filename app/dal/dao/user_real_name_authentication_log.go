package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type UserRealNameAuthenticationLog interface {
	Create(*entity.UserRealNameAuthenticationLog) (*entity.UserRealNameAuthenticationLog, error)
	Delete(*entity.UserRealNameAuthenticationLog) error
	Select(*entity.UserRealNameAuthenticationLog, *Pagination) ([]*entity.UserRealNameAuthenticationLog, error)
	Count(*entity.UserRealNameAuthenticationLog) (int32, error)
	Update(*entity.UserRealNameAuthenticationLog) (*entity.UserRealNameAuthenticationLog, error)
	FindByPk(en *entity.UserRealNameAuthenticationLog) (*entity.UserRealNameAuthenticationLog, error)
}
type userRealNameAuthenticationLog struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUserRealNameAuthenticationLog(db *gorm.DB, rs *helper.Redis) UserRealNameAuthenticationLog {
	return &userRealNameAuthenticationLog{
		db: db,
		rs: rs,
	}
}
func (o *userRealNameAuthenticationLog) Create(en *entity.UserRealNameAuthenticationLog) (*entity.UserRealNameAuthenticationLog, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *userRealNameAuthenticationLog) Delete(uu *entity.UserRealNameAuthenticationLog) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *userRealNameAuthenticationLog) Select(en *entity.UserRealNameAuthenticationLog, pg *Pagination) ([]*entity.UserRealNameAuthenticationLog, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.IdCard != "" {
		sql = sql.Where("id_card = ?", en.IdCard)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	var rows []*entity.UserRealNameAuthenticationLog
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *userRealNameAuthenticationLog) Count(en *entity.UserRealNameAuthenticationLog) (int32, error) {
	sql := o.db.Model(&entity.UserRealNameAuthenticationLog{})
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.IdCard != "" {
		sql = sql.Where("id_card = ?", en.IdCard)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *userRealNameAuthenticationLog) Update(en *entity.UserRealNameAuthenticationLog) (*entity.UserRealNameAuthenticationLog, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *userRealNameAuthenticationLog) FindByPk(en *entity.UserRealNameAuthenticationLog) (*entity.UserRealNameAuthenticationLog, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
