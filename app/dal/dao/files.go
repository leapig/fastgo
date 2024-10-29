package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Files interface {
	Create(*entity.Files) (*entity.Files, error)
	Delete(*entity.Files) error
}

// files 接口规范实现类
type files struct {
	db *gorm.DB
	rs *helper.Redis
}

// NewFilesDao 实例化接口规范实现类
func NewFilesDao(db *gorm.DB, rs *helper.Redis) Files {
	return &files{
		db: db,
		rs: rs,
	}
}

// Create 创建文件信息数据层具体实现
func (f *files) Create(en *entity.Files) (*entity.Files, error) {
	tx := f.db.Create(&en)
	return en, tx.Error
}

// Delete 删除文件信息数据层具体实现
func (f *files) Delete(en *entity.Files) error {
	tx := f.db.Where("name=? and enterprise_pk=? and type=?", en.Name, en.EnterprisePk, en.Type).Delete(&en)
	if tx.RowsAffected == 0 {
		return errors.New("update fatal")
	}
	return tx.Error
}
