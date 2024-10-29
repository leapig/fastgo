package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"time"
)

type UserEntry interface {
	Create(m *entity.UserEntry) (*entity.UserEntry, error)
	Select(m *model.UserEntryModel, page *entity.Pagination, st, et *time.Time) ([]*model.UserEntryModel, error)
	Count(m *model.UserEntryModel, st, et *time.Time) (int32, error)
	Update(m *entity.UserEntry) (*entity.UserEntry, error)
	FindByPk(en *entity.UserEntry) (*entity.UserEntry, error)
	Find(en *entity.UserEntry) (*entity.UserEntry, error)
}
type userEntry struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUserEntry(db *gorm.DB, rs *helper.Redis) UserEntry {
	return &userEntry{
		db: db,
		rs: rs,
	}
}

func (o *userEntry) Create(en *entity.UserEntry) (*entity.UserEntry, error) {
	en.Pk = helper.GetRid(helper.UserEntry)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *userEntry) Select(m *model.UserEntryModel, pg *entity.Pagination, st, et *time.Time) ([]*model.UserEntryModel, error) {
	sql := o.db.Model(&model.UserEntryModel{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Joins("User")
	sql.Joins("Enterprise")
	if m.Pk != 0 {
		sql.Where("user_entry.pk = ?", m.Pk)
	}
	if m.EnterprisePk != 0 {
		sql.Where("user_entry.enterprise_pk = ?", m.EnterprisePk)
	}
	if m.User.Name != "" {
		sql.Where("User.name like ?", "%"+m.User.Name+"%")
	}
	if m.User.Phone != "" {
		sql.Where("User.phone like ?", "%"+m.User.Phone+"%")
	}
	if m.Status != "" {
		sql.Where("user_entry.status = ?", m.Status)
	}
	if st != nil {
		sql.Where("user_entry.arrival_time > ?", st)
	}
	if st != nil {
		sql.Where("user_entry.arrival_time < ?", et)
	}
	var rows []*model.UserEntryModel
	tx := sql.Order("user_entry.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *userEntry) Count(m *model.UserEntryModel, st, et *time.Time) (int32, error) {
	sql := o.db.Model(&model.UserEntryModel{})
	sql.Joins("User")
	//sql.Joins("Enterprise")
	if m.Pk != 0 {
		sql.Where("user_entry.pk = ?", m.Pk)
	}
	if m.EnterprisePk != 0 {
		sql.Where("user_entry.enterprise_pk = ?", m.EnterprisePk)
	}
	if m.User != nil && m.User.Name != "" {
		sql.Where("User.name like ?", "%"+m.User.Name+"%")
	}
	if m.User != nil && m.User.Phone != "" {
		sql.Where("User.phone like ?", "%"+m.User.Phone+"%")
	}
	if m.Status != "" {
		sql.Where("user_entry.status = ?", m.Status)
	}
	if st != nil {
		sql.Where("user_entry.arrival_time > ?", st)
	}
	if st != nil {
		sql.Where("user_entry.arrival_time < ?", et)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *userEntry) Update(m *entity.UserEntry) (*entity.UserEntry, error) {
	first := make(map[string]interface{}, 0)
	userE := entity.UserEntry{}
	res := o.db.Where("pk = ?", m.Pk).First(&userE)
	if res.Error != nil || userE.Pk == 0 {
		return nil, errors.New("获取设备失败")
	}
	if m.Status != "" {
		first["status"] = m.Status
	}
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if res := tx.Where("pk = ?", m.Pk).Model(&userE).Updates(&first); res.Error != nil {
			return res.Error
		}
		return nil
	})
	return &userE, err
}

func (o *userEntry) FindByPk(en *entity.UserEntry) (*entity.UserEntry, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

func (o *userEntry) Find(p *entity.UserEntry) (*entity.UserEntry, error) {
	sql := o.db.Model(&entity.UserEntry{})
	if p.UserPk != 0 {
		sql.Where("user_pk = ?", p.UserPk)
	}
	if p.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", p.EnterprisePk)
	}
	if p.Status != "" {
		sql.Where("status = ?", p.Status)
	}
	tx := sql.Find(&p)
	if tx.RowsAffected == 0 {
		return nil, errors.New("find fatal")
	}
	return p, tx.Error
}
