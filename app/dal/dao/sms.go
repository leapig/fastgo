package dao

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Sms interface {
	Create(sms2 *entity.Sms) (*entity.Sms, error)
}

type sms struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewSms(db *gorm.DB, rs *helper.Redis) Sms {
	return &sms{
		db: db,
		rs: rs,
	}
}

func (o sms) Create(sms *entity.Sms) (*entity.Sms, error) {
	tx := o.db.Create(&sms)
	return sms, tx.Error
}
