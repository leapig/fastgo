package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type ProjectEnterpriseRelation interface {
	Create(*entity.ProjectEnterpriseRelation) (*entity.ProjectEnterpriseRelation, error)
	Delete(*entity.ProjectEnterpriseRelation) error
	Select(*entity.ProjectEnterpriseRelation, *Pagination) ([]*entity.ProjectEnterpriseRelation, error)
	SelectWithProject(en *entity.ProjectEnterpriseRelation, pg *Pagination) ([]*model.ProjectEnterpriseRelation, error)
	SelectByProjectPk(en *entity.ProjectEnterpriseRelation) ([]*model.ProjectEnterpriseRelationForRelationEnterprise, error)
	Count(*entity.ProjectEnterpriseRelation) (int32, error)
	Update(*entity.ProjectEnterpriseRelation) (*entity.ProjectEnterpriseRelation, error)
	FindByPk(en *entity.ProjectEnterpriseRelation) (*entity.ProjectEnterpriseRelation, error)
	SelectAllEnterpriseRelation(*entity.ProjectEnterpriseRelation) ([]*entity.ProjectEnterpriseRelation, error)

	SelectProjectRelationEnterprise(en *model.ProjectEnterpriseRelationQuery, pg *Pagination) ([]*model.ProjectEnterpriseRelationForProject, int32, error)
}
type projectEnterpriseRelation struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewProjectEnterpriseRelation(db *gorm.DB, rs *helper.Redis) ProjectEnterpriseRelation {
	return &projectEnterpriseRelation{
		db: db,
		rs: rs,
	}
}
func (o *projectEnterpriseRelation) Create(en *entity.ProjectEnterpriseRelation) (*entity.ProjectEnterpriseRelation, error) {
	tx := o.db.Create(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("create fatal")
	}
	return en, tx.Error
}

func (o *projectEnterpriseRelation) Delete(uu *entity.ProjectEnterpriseRelation) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *projectEnterpriseRelation) Select(en *entity.ProjectEnterpriseRelation, pg *Pagination) ([]*entity.ProjectEnterpriseRelation, error) {
	sql := o.db.Model(&entity.ProjectEnterpriseRelation{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	var rows []*entity.ProjectEnterpriseRelation
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *projectEnterpriseRelation) SelectWithProject(en *entity.ProjectEnterpriseRelation, pg *Pagination) ([]*model.ProjectEnterpriseRelation, error) {
	sql := o.db.Model(&model.ProjectEnterpriseRelation{}).
		Preload("Project").
		Limit(int(pg.Size)).
		Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("project_enterprise_relation.enterprise_pk=?", en.EnterprisePk)
	}
	var rows []*model.ProjectEnterpriseRelation
	tx := sql.Order("project_enterprise_relation.create_at desc").Find(&rows)
	if tx.Error == nil && rows != nil && len(rows) > 0 {
		for _, v := range rows {
			if v.Project != nil && v.Project.CreateEnterprisePk != 0 && v.Project.CorporatePk != 0 {
				var userMessage *entity.EnterpriseUser
				userSql := o.db.Model(&entity.EnterpriseUser{}).
					Where("enterprise_pk = ?", v.Project.CreateEnterprisePk).
					Where("user_pk = ?", v.Project.CorporatePk).
					First(&userMessage)
				if userSql.Error == nil && userMessage != nil && userMessage.Name != "" && userMessage.Phone != "" {
					v.CorporateName = userMessage.Name
					v.CorporatePhone = userMessage.Phone
				}
			}
			if v.Project != nil && v.Project.CreateEnterprisePk != 0 {
				var enterpriseMessage *entity.Enterprise
				enterpriseSql := o.db.Model(&entity.Enterprise{}).
					Where("pk = ?", v.Project.CreateEnterprisePk).
					First(&enterpriseMessage)
				if enterpriseSql.Error == nil && enterpriseMessage != nil {
					v.CreateEnterpriseName = enterpriseMessage.Name
				}
			}
		}
	}
	return rows, tx.Error
}
func (o *projectEnterpriseRelation) SelectByProjectPk(en *entity.ProjectEnterpriseRelation) ([]*model.ProjectEnterpriseRelationForRelationEnterprise, error) {
	sql := o.db.Model(&model.ProjectEnterpriseRelationForRelationEnterprise{}).
		Joins("LEFT JOIN `enterprise` ON `enterprise`.`pk` = `project_enterprise_relation`.`enterprise_pk`").
		Select("`enterprise`.`name` as enterprise_name,`project_enterprise_relation`.*")
	if en.ProjectPk != 0 {
		sql = sql.Where("project_enterprise_relation.project_pk=?", en.ProjectPk)
	}
	var rows []*model.ProjectEnterpriseRelationForRelationEnterprise
	tx := sql.Order("project_enterprise_relation.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *projectEnterpriseRelation) Count(en *entity.ProjectEnterpriseRelation) (int32, error) {
	sql := o.db.Model(&entity.ProjectEnterpriseRelation{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *projectEnterpriseRelation) Update(en *entity.ProjectEnterpriseRelation) (*entity.ProjectEnterpriseRelation, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *projectEnterpriseRelation) FindByPk(en *entity.ProjectEnterpriseRelation) (*entity.ProjectEnterpriseRelation, error) {
	res := &entity.ProjectEnterpriseRelation{}
	tx := o.db.Model(res).Where("pk = ?", en.Pk).Find(res)
	return res, tx.Error
}
func (o *projectEnterpriseRelation) SelectAllEnterpriseRelation(en *entity.ProjectEnterpriseRelation) ([]*entity.ProjectEnterpriseRelation, error) {
	sql := o.db.Model(&entity.ProjectEnterpriseRelation{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	var rows []*entity.ProjectEnterpriseRelation
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *projectEnterpriseRelation) SelectProjectRelationEnterprise(en *model.ProjectEnterpriseRelationQuery, pg *Pagination) ([]*model.ProjectEnterpriseRelationForProject, int32, error) {
	sql := o.db.Model(&model.ProjectEnterpriseRelation{}).
		Joins("left join enterprise e on e.pk = project_enterprise_relation.enterprise_pk").
		Joins("left join enterprise_user d on d.user_pk = e.corporate_pk and d.enterprise_pk = e.pk").
		Select("project_enterprise_relation.*", "d.name  as corporate_name ", "d.phone as corporate_phone",
			"e.name as enterprise_name", "e.cover as cover", "e.serial as serial", "e.staff_size as staff_size", "e.license as license",
			"e.country as country", "e.province as province", "e.city as city", "e.district as district", "e.county as county", "e.site as site",
			"e.longitude as longitude", "e.latitude as latitude", "e.type as type", "e.address_code as address_code").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.ProjectPk != 0 {
		sql = sql.Where("project_enterprise_relation.project_pk  = ? ", en.ProjectPk)
	}
	if en.EnterpriseName != "" {
		sql = sql.Where("e.name like ?", "%"+en.EnterpriseName+"%")
	}
	if en.CorporateName != "" {
		sql = sql.Where("d.name like ?", "%"+en.CorporateName+"%")
	}
	if en.CorporatePhone != "" {
		sql = sql.Where("d.phone like ?", "%"+en.CorporatePhone+"%")
	}
	var rows []*model.ProjectEnterpriseRelationForProject
	var count int64
	tx := sql.Order("project_enterprise_relation.create_at desc").Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}
