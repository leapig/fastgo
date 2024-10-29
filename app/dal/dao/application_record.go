package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"time"
)

type ApplicationRecord interface {
	Create(*entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	Delete(*entity.ApplicationRecord) error
	Select(en *entity.ApplicationRecord, pg *Pagination, st, et *time.Time) ([]*entity.ApplicationRecord, error)
	CountDetail(en *model.ApplicationRecordModel, st, et *time.Time) (int32, error)
	Count(en *entity.ApplicationRecord, st, et *time.Time) (int32, error)
	Update(*entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	FindByPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	UpdateByUserEntryPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	Find(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error)
	SelectWithDetail(en *model.ApplicationRecordModel, pg *Pagination, st, et *time.Time) ([]*model.ApplicationRecordModel, error)
}
type applicationRecord struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewApplicationRecord(db *gorm.DB, rs *helper.Redis) ApplicationRecord {
	return &applicationRecord{
		db: db,
		rs: rs,
	}
}
func (o *applicationRecord) Create(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *applicationRecord) Delete(uu *entity.ApplicationRecord) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *applicationRecord) Select(en *entity.ApplicationRecord, pg *Pagination, st, et *time.Time) ([]*entity.ApplicationRecord, error) {
	sql := o.db.Model(&entity.ApplicationRecord{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	var rows []*entity.ApplicationRecord
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *applicationRecord) CountDetail(en *model.ApplicationRecordModel, st, et *time.Time) (int32, error) {
	sql := o.db.Model(&model.ApplicationRecordModel{})
	sql.Joins("Enterprise")
	sql.Joins("User")
	if en.UserPk != 0 {
		sql = sql.Where("application_record.user_pk = ?", en.UserPk)
	}
	if en.PositionPk != 0 {
		sql = sql.Where("application_record.position_pk = ?", en.PositionPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("application_record.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.InterviewResult != "" {
		sql = sql.Where("application_record.interview_result = ?", en.InterviewResult)
	}
	if en.SourceType != "" {
		sql = sql.Where("application_record.source_type = ?", en.SourceType)
	}
	if en.User.Name != "" {
		sql = sql.Where("User.name like ?", "%"+en.User.Name+"%")
	}
	if st != nil {
		sql.Where("application_record.create_at > ?", st)
	}
	if et != nil {
		sql.Where("application_record.create_at < ?", et)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *applicationRecord) Count(en *entity.ApplicationRecord, st, et *time.Time) (int32, error) {
	sql := o.db.Model(&entity.ApplicationRecord{})
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.PositionPk != 0 {
		sql = sql.Where("position_pk = ?", en.PositionPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.InterviewResult != "" {
		sql = sql.Where("interview_result = ?", en.InterviewResult)
	}
	if en.SourceType != "" {
		sql = sql.Where("source_type = ?", en.SourceType)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *applicationRecord) Update(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *applicationRecord) UpdateByUserEntryPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	tx := o.db.Where("user_entry_pk = ?", en.UserEntryPk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *applicationRecord) FindByPk(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *applicationRecord) Find(en *entity.ApplicationRecord) (*entity.ApplicationRecord, error) {
	sql := o.db.Model(&entity.ApplicationRecord{})
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	if en.PositionPk != 0 {
		sql = sql.Where("position_pk = ?", en.PositionPk)
	}
	if en.SourceType != "" {
		sql = sql.Where("source_type = ?", en.SourceType)
	}
	if en.UserEntryPk != 0 {
		sql = sql.Where("user_entry_pk = ?", en.UserEntryPk)
	}
	tx := sql.Order("create_at desc").First(&en)
	return en, tx.Error
}
func (o *applicationRecord) SelectWithDetail(en *model.ApplicationRecordModel, pg *Pagination, st, et *time.Time) ([]*model.ApplicationRecordModel, error) {
	sql := o.db.Model(&model.ApplicationRecordModel{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("Enterprise")
	sql.Joins("User")
	sql.Joins("Position")
	if en.UserPk != 0 {
		sql = sql.Where("application_record.user_pk = ?", en.UserPk)
	}
	if en.PositionPk != 0 {
		sql = sql.Where("application_record.position_pk = ?", en.PositionPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("application_record.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.InterviewResult != "" {
		sql = sql.Where("application_record.interview_result = ?", en.InterviewResult)
	}
	if en.SourceType != "" {
		sql = sql.Where("application_record.source_type = ?", en.SourceType)
	}
	var rows []*model.ApplicationRecordModel
	tx := sql.Order("application_record.create_at desc").Find(&rows)
	return rows, tx.Error
}
