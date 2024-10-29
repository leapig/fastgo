package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type PatrolPoint interface {
	Create(route *entity.PatrolPoint) (*entity.PatrolPoint, error)
	Delete(uu *entity.PatrolPoint) error
	Update(uu *entity.PatrolPoint) (*entity.PatrolPoint, error)
	Select(en *entity.PatrolPoint, pg *entity.Pagination) ([]*entity.PatrolPoint, error)
	Count(en *entity.PatrolPoint) (int32, error)
	List(en *entity.PatrolPoint) ([]*entity.PatrolPoint, error)
}
type patrolPoint struct {
	db *gorm.DB
}

func NewPatrolPoint(db *gorm.DB) PatrolPoint {
	return &patrolPoint{
		db: db,
	}
}

func (o *patrolPoint) Create(en *entity.PatrolPoint) (*entity.PatrolPoint, error) {
	en.Pk = helper.GetRid(helper.PatrolPoint)
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *patrolPoint) Delete(uu *entity.PatrolPoint) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *patrolPoint) Update(en *entity.PatrolPoint) (*entity.PatrolPoint, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}

func (o *patrolPoint) Select(en *entity.PatrolPoint, pg *entity.Pagination) ([]*entity.PatrolPoint, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	var rows []*entity.PatrolPoint
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql.Where("name like ?", "%"+en.Name+"%")
	}
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *patrolPoint) List(en *entity.PatrolPoint) ([]*entity.PatrolPoint, error) {
	sql := o.db.Model(&entity.PatrolPoint{})
	if en.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql.Where("name like ?", "%"+en.Name+"%")
	}
	var rows []*entity.PatrolPoint
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *patrolPoint) Count(en *entity.PatrolPoint) (int32, error) {
	sql := o.db.Model(&entity.PatrolPoint{})
	if en.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql.Where("name like ?", "%"+en.Name+"%")
	}
	var res int64
	tx := sql.Count(&res)
	return int32(res), tx.Error
}
