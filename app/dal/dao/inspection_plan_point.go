package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"gorm.io/gorm"
)

type InspectionPlanPoint interface {
	Create(*entity.InspectionPlanPoint) (*entity.InspectionPlanPoint, error)
	Delete(*entity.InspectionPlanPoint) error
	DeleteByPlanPk(uu *entity.InspectionPlanPoint) error
	Select(*entity.InspectionPlanPoint, *Pagination) ([]*entity.InspectionPlanPoint, error)
	Count(*entity.InspectionPlanPoint) (int32, error)
	Update(*entity.InspectionPlanPoint) (*entity.InspectionPlanPoint, error)
	FindByPk(en *entity.InspectionPlanPoint) (*entity.InspectionPlanPoint, error)
}
type inspectionPlanPoint struct {
	db *gorm.DB
}

func NewInspectionPlanPoint(db *gorm.DB) InspectionPlanPoint {
	return &inspectionPlanPoint{
		db: db,
	}
}
func (o *inspectionPlanPoint) Create(en *entity.InspectionPlanPoint) (*entity.InspectionPlanPoint, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *inspectionPlanPoint) Delete(uu *entity.InspectionPlanPoint) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *inspectionPlanPoint) DeleteByPlanPk(uu *entity.InspectionPlanPoint) error {
	tx := o.db.Where("plan_pk = ? ", uu.PlanPk).Delete(&uu)
	return tx.Error
}
func (o *inspectionPlanPoint) Select(en *entity.InspectionPlanPoint, pg *Pagination) ([]*entity.InspectionPlanPoint, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PlanPk != 0 {
		sql = sql.Where("plan_pk = ?", en.PlanPk)
	}
	if en.PointPk != 0 {
		sql = sql.Where("point_pk = ?", en.PointPk)
	}
	var rows []*entity.InspectionPlanPoint
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *inspectionPlanPoint) Count(en *entity.InspectionPlanPoint) (int32, error) {
	sql := o.db.Model(&entity.InspectionPlanPoint{})
	if en.PlanPk != 0 {
		sql = sql.Where("plan_pk = ?", en.PlanPk)
	}
	if en.PointPk != 0 {
		sql = sql.Where("point_pk = ?", en.PointPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *inspectionPlanPoint) Update(log *entity.InspectionPlanPoint) (*entity.InspectionPlanPoint, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *inspectionPlanPoint) FindByPk(en *entity.InspectionPlanPoint) (*entity.InspectionPlanPoint, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
