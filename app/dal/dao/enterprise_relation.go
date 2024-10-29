package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"gorm.io/gorm"
)

type EnterpriseRelation interface {
	Create(*entity.EnterpriseRelation) (*entity.EnterpriseRelation, error)
	Delete(*entity.EnterpriseRelation) error
	Select(*entity.EnterpriseRelation, *Pagination) ([]*entity.EnterpriseRelation, error)
	Count(*entity.EnterpriseRelation) (int32, error)
	Update(*entity.EnterpriseRelation) (*entity.EnterpriseRelation, error)
	FindByPk(en *entity.EnterpriseRelation) (*entity.EnterpriseRelation, error)
	DeleteByRelationEnterprisePkAndProjectPk(uu *entity.EnterpriseRelation) error
	SelectCustomerEnterprise(en *model.EnterpriseRelationQuery, pg *Pagination) ([]*model.Enterprise, int32, error)
	SelectCustomerRelationProject(en *model.EnterpriseRelationProjectQuery, pg *Pagination) ([]*model.Project, int32, error)
}
type enterpriseRelation struct {
	db *gorm.DB
}

func NewEnterpriseRelation(db *gorm.DB) EnterpriseRelation {
	return &enterpriseRelation{
		db: db,
	}
}
func (o *enterpriseRelation) Create(en *entity.EnterpriseRelation) (*entity.EnterpriseRelation, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *enterpriseRelation) Delete(uu *entity.EnterpriseRelation) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *enterpriseRelation) Select(en *entity.EnterpriseRelation, pg *Pagination) ([]*entity.EnterpriseRelation, error) {
	sql := o.db.Model(&entity.EnterpriseRelation{})
	sql = o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RelationEnterprisePk != 0 {
		sql = sql.Where("relation_enterprise_pk = ?", en.RelationEnterprisePk)
	}
	var rows []*entity.EnterpriseRelation
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *enterpriseRelation) Count(en *entity.EnterpriseRelation) (int32, error) {
	sql := o.db.Model(&entity.EnterpriseRelation{})
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk = ?", en.ProjectPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.RelationEnterprisePk != 0 {
		sql = sql.Where("relation_enterprise_pk = ?", en.RelationEnterprisePk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *enterpriseRelation) Update(en *entity.EnterpriseRelation) (*entity.EnterpriseRelation, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *enterpriseRelation) FindByPk(en *entity.EnterpriseRelation) (*entity.EnterpriseRelation, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *enterpriseRelation) DeleteByRelationEnterprisePkAndProjectPk(uu *entity.EnterpriseRelation) error {
	tx := o.db.Where("relation_enterprise_pk = ? and  project_pk = ?", uu.RelationEnterprisePk, uu.ProjectPk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

// SelectCustomerEnterprise 查找客户
func (o *enterpriseRelation) SelectCustomerEnterprise(en *model.EnterpriseRelationQuery, pg *Pagination) ([]*model.Enterprise, int32, error) {
	sql := o.db.Model(&model.Enterprise{}).
		Joins("left join enterprise_user d on d.user_pk = enterprise.corporate_pk and d.enterprise_pk = enterprise.pk").
		Select("enterprise.*", "d.name as corporate_name", "d.phone as corporate_phone").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise.pk in (?)", o.db.Table("enterprise_relation").Where("enterprise_pk = ? ", en.EnterprisePk).Distinct("relation_enterprise_pk"))
	}
	if en.CustomerName != "" {
		sql = sql.Where("enterprise.name  like  ?", "%"+en.CustomerName+"%")
	}
	if en.CorporatePhone != "" {
		sql = sql.Where("d.phone  like  ?", "%"+en.CorporatePhone+"%")
	}
	if en.CorporateName != "" {
		sql = sql.Where("d.name  like  ?", "%"+en.CorporateName+"%")
	}
	if en.AddressCode != "" {
		sql = sql.Where("enterprise.address_code  like  ?", "%"+en.AddressCode+"%")
	}
	if en.Country != "" {
		sql = sql.Where("enterprise.country = ?", en.Country)
	}
	if en.Province != "" {
		sql = sql.Where("enterprise.province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("enterprise.city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("enterprise.district = ?", en.District)
	}
	if en.County != "" {
		sql = sql.Where("enterprise.county = ?", en.County)
	}
	var rows []*model.Enterprise
	var count int64
	tx := sql.Order("enterprise.create_at desc").Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}
func (o *enterpriseRelation) SelectCustomerRelationProject(en *model.EnterpriseRelationProjectQuery, pg *Pagination) ([]*model.Project, int32, error) {
	sql := o.db.Model(&model.Project{}).
		Joins("left join enterprise_user d on d.pk = project.corporate_pk").
		Joins("left join enterprise e on e.pk = project.create_enterprise_pk").
		Select("project.*", "d.name  as corporate_name ", "d.phone as corporate_phone", "e.name as create_enterprise_name").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 && en.RelationEnterprisePk != 0 {
		sql = sql.Where("project.pk in (?)", o.db.Table("enterprise_relation").Where("enterprise_pk=? and relation_enterprise_pk=? ", en.EnterprisePk, en.RelationEnterprisePk).
			Or("relation_enterprise_pk=? and enterprise_pk=?", en.EnterprisePk, en.RelationEnterprisePk).Select("project_pk"))
	}
	if en.ProjectName != "" {
		sql = sql.Where("project.name like ?", "%"+en.ProjectName+"%")
	}
	if en.ProjectStatus != 0 {
		sql = sql.Where("project.project_status = ?", en.ProjectStatus)
	}
	if en.Province != "" {
		sql = sql.Where("project.province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("project.city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("project.district = ?", en.District)
	}
	if en.County != "" {
		sql = sql.Where("project.county = ?", en.County)
	}
	if en.ProjectAddress != "" {
		sql = sql.Where("project.address like ?", "%"+en.ProjectAddress+"%")
	}
	if en.AddressCode != "" {
		sql = sql.Where("project.address_code like ?", en.AddressCode+"%")
	}
	if en.CorporateName != "" {
		sql = sql.Where("d.name like ?", "%"+en.CorporateName+"%")
	}
	if en.CorporatePhone != "" {
		sql = sql.Where("d.phone like ?", "%"+en.CorporatePhone+"%")
	}
	if en.CreateEnterpriseName != "" {
		sql = sql.Where("e.name like ?", "%"+en.CreateEnterpriseName+"%")
	}
	var rows []*model.Project
	var count int64
	tx := sql.Order("project.create_at desc").Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}
