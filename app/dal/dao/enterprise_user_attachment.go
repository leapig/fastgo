package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type EnterpriseUserAttachment interface {
	Create(attachment *entity.EnterpriseUserAttachment) (*entity.EnterpriseUserAttachment, error)
	FindByPk(attachment *entity.EnterpriseUserAttachment) (*entity.EnterpriseUserAttachment, error)
	Select(attachment *entity.EnterpriseUserAttachment, pagination *entity.Pagination) ([]*entity.EnterpriseUserAttachment, error)
	Count(attachment *entity.EnterpriseUserAttachment) (int32, error)
	Delete(attachment *entity.EnterpriseUserAttachment) error
	List(attachment *entity.EnterpriseUserAttachment) ([]*entity.EnterpriseUserAttachment, error)
}

type enterpriseUserAttachment struct {
	db *gorm.DB
}

func NewEnterpriseUserAttachment(db *gorm.DB) EnterpriseUserAttachment {
	return &enterpriseUserAttachment{
		db: db,
	}
}

func (o *enterpriseUserAttachment) Create(en *entity.EnterpriseUserAttachment) (*entity.EnterpriseUserAttachment, error) {
	en.Pk = helper.GetRid(helper.EnterpriseUserAttachment)
	tx := o.db.Create(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("delete fatal")
	}
	return o.FindByPk(en)
}

func (o *enterpriseUserAttachment) FindByPk(attachment *entity.EnterpriseUserAttachment) (*entity.EnterpriseUserAttachment, error) {
	sql := o.db.Model(&entity.EnterpriseUserAttachment{})
	tx := sql.Where("pk = ?", attachment.Pk).Find(attachment)
	return attachment, tx.Error
}

func (o *enterpriseUserAttachment) Select(uu *entity.EnterpriseUserAttachment, pg *entity.Pagination) ([]*entity.EnterpriseUserAttachment, error) {
	sql := o.db.Model(&entity.EnterpriseAreaPermission{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if uu.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", uu.EnterprisePk)
	}
	if uu.UserPk != 0 {
		sql.Where("user_pk = ?", uu.UserPk)
	}
	var rows []*entity.EnterpriseUserAttachment
	tx := sql.Find(&rows)
	return rows, tx.Error
}
func (o *enterpriseUserAttachment) Count(uu *entity.EnterpriseUserAttachment) (int32, error) {
	sql := o.db.Model(&entity.EnterpriseAreaPermission{})
	if uu.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", uu.EnterprisePk)
	}
	if uu.UserPk != 0 {
		sql.Where("user_pk = ?", uu.UserPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *enterpriseUserAttachment) Delete(uu *entity.EnterpriseUserAttachment) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *enterpriseUserAttachment) List(uu *entity.EnterpriseUserAttachment) ([]*entity.EnterpriseUserAttachment, error) {
	sql := o.db.Model(&entity.EnterpriseUserAttachment{})
	if uu.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", uu.EnterprisePk)
	}
	if uu.UserPk != 0 {
		sql.Where("user_pk = ?", uu.UserPk)
	}
	if uu.FileType != "" {
		sql.Where("file_type = ?", uu.FileType)
	}
	var rows []*entity.EnterpriseUserAttachment
	tx := sql.Find(&rows)
	return rows, tx.Error
}
