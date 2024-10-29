package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Project interface {
	Create(*entity.Project) (*entity.Project, error)
	Delete(*entity.Project) error
	Select(*entity.Project, *Pagination) ([]*entity.Project, error)
	Count(*entity.Project) (int32, error)
	Update(*entity.Project) (*entity.Project, error)
	FindByPk(en *entity.Project) (*entity.Project, error)
	SelectForEnterprise(en *model.ProjectQuery, pg *Pagination) ([]*model.Project, int32, error)
	SelectProjectForUserProfession(en *model.ProjectQueryForUserProfession) ([]*model.Project, error)
	SelectNoPage(en *entity.Project) ([]*entity.Project, error)
}
type project struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewProject(db *gorm.DB, rs *helper.Redis) Project {
	return &project{
		db: db,
		rs: rs,
	}
}
func (o *project) Create(en *entity.Project) (*entity.Project, error) {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if createProjectSql := o.db.Create(&en); createProjectSql.Error != nil {
			return createProjectSql.Error
		} else if createProjectSql.RowsAffected == 0 {
			return errors.New("create fatal")
		}
		//新增项目企业关联关系
		projectEnterpriseRelationPk := helper.GetRid(helper.ProjectEnterpriseRelation)
		if projectEnterpriseRelationCreateSql := o.db.Create(&entity.ProjectEnterpriseRelation{
			Pk:           projectEnterpriseRelationPk,
			ProjectPk:    en.Pk,
			EnterprisePk: en.CreateEnterprisePk,
		}); projectEnterpriseRelationCreateSql.Error != nil {
			return projectEnterpriseRelationCreateSql.Error
		}
		return nil
	})
	return en, err
}

func (o *project) Delete(uu *entity.Project) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *project) Select(en *entity.Project, pg *Pagination) ([]*entity.Project, error) {
	sql := o.db.Model(&entity.Project{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.CreateEnterprisePk != 0 {
		sql = sql.Where("create_enterprise_pk=?", en.CreateEnterprisePk)
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.AddressCode != "" {
		sql = sql.Where("address_code like ?", en.AddressCode+"%")
	}
	if en.Province != "" {
		sql = sql.Where("province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("district = ?", en.District)
	}
	if en.County != "" {
		sql = sql.Where("county = ?", en.County)
	}
	var rows []*entity.Project
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *project) SelectForEnterprise(en *model.ProjectQuery, pg *Pagination) ([]*model.Project, int32, error) {
	sql := o.db.Model(&model.Project{}).
		Joins("left join enterprise_user d on d.pk = project.corporate_pk").
		Joins("left join enterprise e on e.pk = project.create_enterprise_pk").
		Select("project.*", "d.name  as corporate_name ", "d.phone as corporate_phone", "e.name as create_enterprise_name").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
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
		sql = sql.Where("project.address like ?", en.ProjectAddress+"%")
	}
	if en.AddressCode != "" {
		sql = sql.Where("project.address_code like ?", en.AddressCode+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("project.pk in (?)", o.db.Table("project_enterprise_relation").Where("enterprise_pk=?", en.EnterprisePk).Select("project_pk"))
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

func (o *project) Count(en *entity.Project) (int32, error) {
	sql := o.db.Model(&entity.Project{})
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.CreateEnterprisePk != 0 {
		sql = sql.Where("create_enterprise_pk=?", en.CreateEnterprisePk)
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.Province != "" {
		sql = sql.Where("province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("district = ?", en.District)
	}
	if en.County != "" {
		sql = sql.Where("county = ?", en.County)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *project) Update(en *entity.Project) (*entity.Project, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *project) FindByPk(en *entity.Project) (*entity.Project, error) {
	res := &entity.Project{}
	tx := o.db.Model(res).Where("pk = ?", en.Pk).Find(res)
	return res, tx.Error
}
func (o *project) SelectProjectForUserProfession(en *model.ProjectQueryForUserProfession) ([]*model.Project, error) {
	sql := o.db.Model(&model.Project{}).
		Joins("left join project_member m on m.project_pk = project.pk and m.delete_at is null").
		Select("project.*")
	if en.EnterpriseUserPk != 0 {
		sql = sql.Where("m.enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	var rows []*model.Project
	tx := sql.Order("project.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *project) SelectNoPage(en *entity.Project) ([]*entity.Project, error) {
	sql := o.db.Model(&entity.Project{})
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.CreateEnterprisePk != 0 {
		sql = sql.Where("create_enterprise_pk=?", en.CreateEnterprisePk)
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.AddressCode != "" {
		sql = sql.Where("address_code like ?", en.AddressCode+"%")
	}
	if en.Province != "" {
		sql = sql.Where("province = ?", en.Province)
	}
	if en.City != "" {
		sql = sql.Where("city = ?", en.City)
	}
	if en.District != "" {
		sql = sql.Where("district = ?", en.District)
	}
	if en.County != "" {
		sql = sql.Where("county = ?", en.County)
	}
	var rows []*entity.Project
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
