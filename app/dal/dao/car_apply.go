package dao

import (
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type CarApply interface {
}
type carApply struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewCarApply(db *gorm.DB) CarApply {
	return &car{
		db: db,
	}
}
