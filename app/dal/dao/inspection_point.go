package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"gorm.io/gorm"
)

type InspectionPoint interface {
	Create(*entity.InspectionPoint) (*entity.InspectionPoint, error)
	Delete(*entity.InspectionPoint) error
	Select(*entity.InspectionPoint, *Pagination) ([]*entity.InspectionPoint, error)
	Count(*entity.InspectionPoint) (int32, error)
	Update(*entity.InspectionPoint) (*entity.InspectionPoint, error)
	FindByPk(en *entity.InspectionPoint) (*entity.InspectionPoint, error)
}
type inspectionPoint struct {
	db *gorm.DB
}

func NewInspectionPoint(db *gorm.DB) InspectionPoint {
	return &inspectionPoint{
		db: db,
	}
}
func (o *inspectionPoint) Create(en *entity.InspectionPoint) (*entity.InspectionPoint, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *inspectionPoint) Delete(uu *entity.InspectionPoint) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *inspectionPoint) Select(en *entity.InspectionPoint, pg *Pagination) ([]*entity.InspectionPoint, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.Level != 0 {
		sql = sql.Where("level = ?", en.Level)
	}
	var rows []*entity.InspectionPoint
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *inspectionPoint) Count(en *entity.InspectionPoint) (int32, error) {
	sql := o.db.Model(&entity.InspectionPoint{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.Level != 0 {
		sql = sql.Where("level = ?", en.Level)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *inspectionPoint) Update(log *entity.InspectionPoint) (*entity.InspectionPoint, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *inspectionPoint) FindByPk(en *entity.InspectionPoint) (*entity.InspectionPoint, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
