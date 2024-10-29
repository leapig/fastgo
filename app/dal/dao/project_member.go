package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type ProjectMember interface {
	Create(*entity.ProjectMember) (*entity.ProjectMember, error)
	Delete(*entity.ProjectMember) error
	Select(*entity.ProjectMember, *Pagination) ([]*entity.ProjectMember, error)
	SelectGetUserPks(en *entity.ProjectMember) ([]int64, error)
	Count(*entity.ProjectMember) (int32, error)
	Update(*entity.ProjectMember) (*entity.ProjectMember, error)
	FindByPk(en *entity.ProjectMember) (*entity.ProjectMember, error)
	FindByUserPkAndProjectPk(en *entity.ProjectMember) (*entity.ProjectMember, error)
	SelectProjectEnterpriseUser(en *model.ProjectMemberQueryBody, pg *Pagination) ([]*model.ProjectMember, int32, error)
}
type projectMember struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewProjectMember(db *gorm.DB, rs *helper.Redis) ProjectMember {
	return &projectMember{
		db: db,
		rs: rs,
	}
}
func (o *projectMember) Create(en *entity.ProjectMember) (*entity.ProjectMember, error) {
	tx := o.db.Create(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("create fatal")
	}
	return en, tx.Error
}

func (o *projectMember) Delete(uu *entity.ProjectMember) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *projectMember) Select(en *entity.ProjectMember, pg *Pagination) ([]*entity.ProjectMember, error) {
	sql := o.db.Model(&entity.ProjectMember{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	if en.EnterpriseUserPk != 0 {
		sql = sql.Where("enterprise_user_pk=?", en.EnterpriseUserPk)
	}
	var rows []*entity.ProjectMember
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *projectMember) SelectGetUserPks(en *entity.ProjectMember) ([]int64, error) {
	sql := o.db.Model(&entity.ProjectMember{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	var result []int64
	tx := sql.Pluck("enterprise_user_pk", &result)
	return result, tx.Error
}
func (o *projectMember) Count(en *entity.ProjectMember) (int32, error) {
	sql := o.db.Model(&entity.ProjectMember{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	if en.EnterpriseUserPk != 0 {
		sql = sql.Where("enterprise_user_pk=?", en.EnterpriseUserPk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *projectMember) Update(en *entity.ProjectMember) (*entity.ProjectMember, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *projectMember) FindByPk(en *entity.ProjectMember) (*entity.ProjectMember, error) {
	res := &entity.ProjectMember{}
	tx := o.db.Model(res).Where("pk = ?", en.Pk).Find(res)
	return res, tx.Error
}
func (o *projectMember) FindByUserPkAndProjectPk(en *entity.ProjectMember) (*entity.ProjectMember, error) {
	res := &entity.ProjectMember{}
	tx := o.db.Model(res).Where("enterprise_user_pk = ?", en.EnterpriseUserPk).
		Where("project_pk = ?", en.ProjectPk).
		First(res)
	return res, tx.Error
}
func (o *projectMember) SelectProjectEnterpriseUser(en *model.ProjectMemberQueryBody, pg *Pagination) ([]*model.ProjectMember, int32, error) {
	sql := o.db.Model(&model.ProjectMember{}).
		Joins("left join enterprise_user d on d.pk = project_member.enterprise_user_pk").Joins("left join enterprise e on e.pk = project_member.enterprise_pk").
		Select("project_member.*", "d.name  as name ", "d.phone as phone", "d.gender as gender", "d.birthday as birthday", "d.education as education", "d.height as height", "d.weight as weight", "e.name as enterprise_name").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("project_member.enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_member.project_pk=?", en.ProjectPk)
	}
	if en.Name != "" {
		sql = sql.Where("d.name like ?", "%"+en.Name+"%")
	}
	if en.Phone != "" {
		sql = sql.Where("d.phone like ?", "%"+en.Phone+"%")
	}
	if en.Gender != 0 {
		sql = sql.Where("d.gender=?", en.Gender)
	}
	var rows []*model.ProjectMember
	var count int64
	tx := sql.Order("project_member.create_at desc").Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}
