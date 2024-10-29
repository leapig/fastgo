package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type EnterpriseAreaPermission interface {
	Create(areaPermission *entity.EnterpriseAreaPermission) (*entity.EnterpriseAreaPermission, error)
	CheckRepetition(areaPermission *entity.EnterpriseAreaPermission) (bool, error)
	Delete(uu *entity.EnterpriseAreaPermission) error
	Select(uu *entity.EnterpriseAreaPermission, page *entity.Pagination) ([]*entity.EnterpriseAreaPermission, error)
	Count(uu *entity.EnterpriseAreaPermission) (int32, error)
	FindList(areaPermission *entity.EnterpriseAreaPermission) ([]*entity.EnterpriseAreaPermission, error)
}

type enterpriseAreaPermission struct {
	db *gorm.DB
}

func NewEnterpriseAreaPermission(db *gorm.DB) EnterpriseAreaPermission {
	return &enterpriseAreaPermission{
		db: db,
	}
}

func (o *enterpriseAreaPermission) Create(en *entity.EnterpriseAreaPermission) (*entity.EnterpriseAreaPermission, error) {
	en.Pk = helper.GetRid(helper.EnterpriseAreaPermission)
	tx := o.db.Create(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("delete fatal")
	}
	return en, tx.Error
}

func (o *enterpriseAreaPermission) Delete(uu *entity.EnterpriseAreaPermission) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *enterpriseAreaPermission) Select(uu *entity.EnterpriseAreaPermission, pg *entity.Pagination) ([]*entity.EnterpriseAreaPermission, error) {
	sql := o.db.Model(&entity.EnterpriseAreaPermission{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if uu.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", uu.EnterprisePk)
	}
	var rows []*entity.EnterpriseAreaPermission
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *enterpriseAreaPermission) Count(uu *entity.EnterpriseAreaPermission) (int32, error) {
	sql := o.db.Model(&entity.EnterpriseAreaPermission{})
	if uu.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", uu.EnterprisePk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

// CheckRepetition 查询重复 true 无重复  false 重复或错误
func (o *enterpriseAreaPermission) CheckRepetition(areaPermission *entity.EnterpriseAreaPermission) (bool, error) {
	sql := o.db.Model(&entity.EnterpriseAreaPermission{})
	if areaPermission.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", areaPermission.EnterprisePk)
	} else {
		return false, errors.New("请输入参数")
	}
	if areaPermission.Province != "" {
		sql.Where("province = ?", areaPermission.Province)
	} else {
		sql.Where("province is null")
	}
	if areaPermission.City != "" {
		sql.Where("city = ?", areaPermission.City)
	} else {
		sql.Where("city is null")
	}
	if areaPermission.District != "" {
		sql.Where("district = ?", areaPermission.District)
	} else {
		sql.Where("district is null")
	}
	if areaPermission.County != "" {
		sql.Where("county = ?", areaPermission.County)
	} else {
		sql.Where("county is null")
	}
	var count int64
	tx := sql.Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	} else {
		if count > 0 {
			return false, tx.Error
		} else {
			return true, tx.Error
		}
	}
}

func (o *enterpriseAreaPermission) FindList(uu *entity.EnterpriseAreaPermission) ([]*entity.EnterpriseAreaPermission, error) {
	sql := o.db.Model(&entity.EnterpriseAreaPermission{})
	if uu.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", uu.EnterprisePk)
	}
	var rows []*entity.EnterpriseAreaPermission
	tx := sql.Find(&rows)
	return rows, tx.Error
}
