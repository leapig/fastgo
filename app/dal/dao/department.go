package dao

import (
	"errors"
	"fmt"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Department interface {
	Create(*entity.Department) (*entity.Department, error)
	Delete(*entity.Department) error
	Select(*entity.Department, *Pagination) ([]*entity.Department, error)
	Count(*entity.Department) (int32, error)
	Update(*entity.Department) (*entity.Department, error)
	FindByPk(en *entity.Department) (*entity.Department, error)
	SelectAllDepartmentList(en *entity.DepartmentList) ([]*entity.DepartmentList, error)
	FindByEnterprisePkAndName(en *entity.Department) (*entity.Department, error)
}
type department struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewDepartment(db *gorm.DB, rs *helper.Redis) Department {
	return &department{
		db: db,
		rs: rs,
	}
}

func (o *department) Create(en *entity.Department) (*entity.Department, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *department) Delete(uu *entity.Department) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *department) Select(en *entity.Department, pg *Pagination) ([]*entity.Department, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	var rows []*entity.Department
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *department) Count(en *entity.Department) (int32, error) {
	sql := o.db.Model(&entity.Department{})
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.ParentPk != 0 {
		sql = sql.Where("parent_pk = ?", en.ParentPk)
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *department) Update(en *entity.Department) (*entity.Department, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *department) FindByPk(en *entity.Department) (*entity.Department, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

func (o *department) SelectAllDepartmentList(en *entity.DepartmentList) ([]*entity.DepartmentList, error) {
	sql := o.db.Model(&entity.DepartmentList{})
	if en.Name != "" {
		sql = sql.Where("name like ?", fmt.Sprint("%", en.Name, "%"))
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ParentPk != 0 {
		sql = sql.Where("parent_pk = ?", en.ParentPk)
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	var rows []*entity.DepartmentList
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *department) FindByEnterprisePkAndName(en *entity.Department) (*entity.Department, error) {
	sql := o.db.Model(&entity.Department{})
	if en.Name != "" {
		sql = sql.Where("name = ?", en.Name)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.ParentPk != 0 {
		sql = sql.Where("parent_pk = ?", en.ParentPk)
	}
	if en.Pk != 0 {
		sql = sql.Where("pk = ?", en.Pk)
	}
	var result *entity.Department
	res := sql.First(&result)
	return result, res.Error
}
