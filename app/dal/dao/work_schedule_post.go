package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"gorm.io/gorm"
)

type WorkSchedulePost interface {
	Create(*entity.WorkSchedulePost) (*entity.WorkSchedulePost, error)
	Delete(*entity.WorkSchedulePost) error
	Select(*entity.WorkSchedulePost, *Pagination) ([]*entity.WorkSchedulePost, error)
	Count(*entity.WorkSchedulePost) (int32, error)
	Update(*entity.WorkSchedulePost) (*entity.WorkSchedulePost, error)
	FindByPk(en *entity.WorkSchedulePost) (*entity.WorkSchedulePost, error)
	SelectModelByPostPkOrSchedulePK(en *entity.WorkSchedulePost, pg *Pagination) ([]*model.WorkSchedulePost, error)
}
type workSchedulePost struct {
	db *gorm.DB
}

func NewWorkSchedulePost(db *gorm.DB) WorkSchedulePost {
	return &workSchedulePost{
		db: db,
	}
}
func (o *workSchedulePost) Create(en *entity.WorkSchedulePost) (*entity.WorkSchedulePost, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *workSchedulePost) Delete(uu *entity.WorkSchedulePost) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *workSchedulePost) Select(en *entity.WorkSchedulePost, pg *Pagination) ([]*entity.WorkSchedulePost, error) {
	sql := o.db.Model(&entity.WorkSchedulePost{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PostPk != 0 {
		sql = sql.Where("post_pk=?", en.PostPk)
	}
	if en.SchedulePk != 0 {
		sql = sql.Where("schedule_pk=?", en.SchedulePk)
	}
	var rows []*entity.WorkSchedulePost
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *workSchedulePost) Count(en *entity.WorkSchedulePost) (int32, error) {
	sql := o.db.Model(&entity.WorkSchedulePost{})
	if en.PostPk != 0 {
		sql = sql.Where("post_pk=?", en.PostPk)
	}
	if en.SchedulePk != 0 {
		sql = sql.Where("schedule_pk=?", en.SchedulePk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *workSchedulePost) Update(en *entity.WorkSchedulePost) (*entity.WorkSchedulePost, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *workSchedulePost) FindByPk(en *entity.WorkSchedulePost) (*entity.WorkSchedulePost, error) {
	res := &entity.WorkSchedulePost{}
	tx := o.db.Model(res).Where("pk = ?", en.Pk).Find(res)
	return res, tx.Error
}
func (o *workSchedulePost) SelectModelByPostPkOrSchedulePK(en *entity.WorkSchedulePost, pg *Pagination) ([]*model.WorkSchedulePost, error) {
	sql := o.db.Model(&model.WorkSchedulePost{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("ProjectPost")
	sql.Preload("WorkSchedule")
	if en.PostPk != 0 {
		sql = sql.Where("work_schedule_post.post_pk=?", en.PostPk)
	}
	if en.SchedulePk != 0 {
		sql = sql.Where("work_schedule_post.schedule_pk=?", en.SchedulePk)
	}
	var rows []*model.WorkSchedulePost
	tx := sql.Order("work_schedule_post.create_at desc").Find(&rows)
	return rows, tx.Error
}
