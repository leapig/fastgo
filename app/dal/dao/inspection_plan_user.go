package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"gorm.io/gorm"
)

type InspectionPlanUser interface {
	Create(*entity.InspectionPlanUser) (*entity.InspectionPlanUser, error)
	Delete(*entity.InspectionPlanUser) error
	Select(*entity.InspectionPlanUser, *Pagination) ([]*entity.InspectionPlanUser, error)
	Count(*entity.InspectionPlanUser) (int32, error)
	Update(*entity.InspectionPlanUser) (*entity.InspectionPlanUser, error)
	FindByPk(en *entity.InspectionPlanUser) (*entity.InspectionPlanUser, error)
}
type inspectionPlanUser struct {
	db *gorm.DB
}

func NewInspectionPlanUser(db *gorm.DB) InspectionPlanUser {
	return &inspectionPlanUser{
		db: db,
	}
}
func (o *inspectionPlanUser) Create(en *entity.InspectionPlanUser) (*entity.InspectionPlanUser, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *inspectionPlanUser) Delete(uu *entity.InspectionPlanUser) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *inspectionPlanUser) Select(en *entity.InspectionPlanUser, pg *Pagination) ([]*entity.InspectionPlanUser, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PlanPk != 0 {
		sql = sql.Where("plan_pk = ?", en.PlanPk)
	}
	if en.EnterpriseUserPk != 0 {
		sql = sql.Where("enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	var rows []*entity.InspectionPlanUser
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *inspectionPlanUser) Count(en *entity.InspectionPlanUser) (int32, error) {
	sql := o.db.Model(&entity.InspectionPlanUser{})
	if en.PlanPk != 0 {
		sql = sql.Where("plan_pk = ?", en.PlanPk)
	}
	if en.EnterpriseUserPk != 0 {
		sql = sql.Where("enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *inspectionPlanUser) Update(log *entity.InspectionPlanUser) (*entity.InspectionPlanUser, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *inspectionPlanUser) FindByPk(en *entity.InspectionPlanUser) (*entity.InspectionPlanUser, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
