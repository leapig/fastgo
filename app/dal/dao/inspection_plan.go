package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"gorm.io/gorm"
)

type InspectionPlan interface {
	Create(*entity.InspectionPlan) (*entity.InspectionPlan, error)
	Delete(*entity.InspectionPlan) error
	Select(*entity.InspectionPlan, *Pagination) ([]*entity.InspectionPlan, error)
	Count(*entity.InspectionPlan) (int32, error)
	Update(*entity.InspectionPlan) (*entity.InspectionPlan, error)
	FindByPk(en *entity.InspectionPlan) (*entity.InspectionPlan, error)
	SelectModel(en *entity.InspectionPlan, pg *Pagination) ([]*model.InspectionPlan, error)
}
type inspectionPlan struct {
	db *gorm.DB
}

func NewInspectionPlan(db *gorm.DB) InspectionPlan {
	return &inspectionPlan{
		db: db,
	}
}
func (o *inspectionPlan) Create(en *entity.InspectionPlan) (*entity.InspectionPlan, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *inspectionPlan) Delete(uu *entity.InspectionPlan) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *inspectionPlan) Select(en *entity.InspectionPlan, pg *Pagination) ([]*entity.InspectionPlan, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	var rows []*entity.InspectionPlan
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *inspectionPlan) Count(en *entity.InspectionPlan) (int32, error) {
	sql := o.db.Model(&entity.InspectionPlan{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *inspectionPlan) Update(log *entity.InspectionPlan) (*entity.InspectionPlan, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *inspectionPlan) FindByPk(en *entity.InspectionPlan) (*entity.InspectionPlan, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *inspectionPlan) SelectModel(en *entity.InspectionPlan, pg *Pagination) ([]*model.InspectionPlan, error) {
	sql := o.db.Model(&model.InspectionPlan{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("InspectionPlanPoint")
	sql.Preload("InspectionPlanPoint.InspectionPoint")
	if en.EnterprisePk != 0 {
		sql = sql.Where("inspection_plan.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("inspection_plan.project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql = sql.Where("inspection_plan.name like ?", "%"+en.Name+"%")
	}
	var rows []*model.InspectionPlan
	tx := sql.Order("inspection_plan.create_at desc").Find(&rows)
	return rows, tx.Error
}
