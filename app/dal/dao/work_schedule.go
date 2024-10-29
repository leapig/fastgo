package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"gorm.io/gorm"
)

type WorkSchedule interface {
	Create(*entity.WorkSchedule) (*entity.WorkSchedule, error)
	Delete(*entity.WorkSchedule) error
	Select(*entity.WorkSchedule, *Pagination) ([]*entity.WorkSchedule, error)
	Count(*entity.WorkSchedule) (int32, error)
	Update(*entity.WorkSchedule) (*entity.WorkSchedule, error)
	FindByPk(en *entity.WorkSchedule) (*entity.WorkSchedule, error)
}
type workSchedule struct {
	db *gorm.DB
}

func NewWorkSchedule(db *gorm.DB) WorkSchedule {
	return &workSchedule{
		db: db,
	}
}
func (o *workSchedule) Create(en *entity.WorkSchedule) (*entity.WorkSchedule, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *workSchedule) Delete(uu *entity.WorkSchedule) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *workSchedule) Select(en *entity.WorkSchedule, pg *Pagination) ([]*entity.WorkSchedule, error) {
	sql := o.db.Model(&entity.WorkSchedule{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	if en.IsNextDay != 0 {
		sql = sql.Where("is_next_day=?", en.IsNextDay)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	var rows []*entity.WorkSchedule
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *workSchedule) Count(en *entity.WorkSchedule) (int32, error) {
	sql := o.db.Model(&entity.WorkSchedule{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	if en.IsNextDay != 0 {
		sql = sql.Where("is_next_day=?", en.IsNextDay)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *workSchedule) Update(en *entity.WorkSchedule) (*entity.WorkSchedule, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *workSchedule) FindByPk(en *entity.WorkSchedule) (*entity.WorkSchedule, error) {
	res := &entity.WorkSchedule{}
	tx := o.db.Model(res).Where("pk = ?", en.Pk).Find(res)
	return res, tx.Error
}
