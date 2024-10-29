package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type UserCredentials interface {
	FindList(client *entity.UserCredentials) ([]*entity.UserCredentials, error)
	FindCardByUserPk(credentials *entity.UserCredentials) (*entity.UserCredentials, error)
	Create(c *entity.UserCredentials) (*entity.UserCredentials, error)
	Find(p *entity.UserCredentials) (*entity.UserCredentials, error)
	Delete(p *entity.UserCredentials) error
	Update(*entity.UserCredentials) error
}

type userCredentials struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUserCredentials(db *gorm.DB, rs *helper.Redis) UserCredentials {
	return &userCredentials{
		db: db,
		rs: rs,
	}
}

func (o *userCredentials) FindList(p *entity.UserCredentials) ([]*entity.UserCredentials, error) {
	sql := o.db.Model(&entity.UserCredentials{})
	if p.UserPk != 0 {
		sql.Where("user_pk = ?", p.UserPk)
	}
	var rows []*entity.UserCredentials
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *userCredentials) FindCardByUserPk(p *entity.UserCredentials) (*entity.UserCredentials, error) {
	sql := o.db.Model(&entity.UserCredentials{})
	if p.UserPk != 0 {
		sql.Where("user_pk = ?", p.UserPk)
	}
	if p.Type != 0 {
		sql.Where("type = ?", p.Type)
	}
	tx := sql.Find(&p)
	return p, tx.Error
}

func (o *userCredentials) Create(c *entity.UserCredentials) (*entity.UserCredentials, error) {
	c.Pk = helper.GetRid(helper.UserCredentials)
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if rs := tx.Create(&c); rs.Error != nil {
			return rs.Error
		}
		return nil
	})
	return c, err
}

func (o *userCredentials) Find(p *entity.UserCredentials) (*entity.UserCredentials, error) {
	sql := o.db.Model(&entity.UserCredentials{})
	if p.UserPk != 0 {
		sql.Where("user_pk = ?", p.UserPk)
	}
	if p.Serial != "" {
		sql.Where("serial = ?", p.Serial)
	}
	if p.Type != 0 {
		sql.Where("type = ?", p.Type)
	}
	tx := sql.Find(&p)
	return p, tx.Error
}

func (o *userCredentials) Delete(uu *entity.UserCredentials) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *userCredentials) Update(uu *entity.UserCredentials) error {
	return o.db.Transaction(func(tx *gorm.DB) error {
		if res := tx.Where("pk = ?", uu.Pk).Updates(uu); res.Error != nil {
			return res.Error
		}
		return nil
	})
}
