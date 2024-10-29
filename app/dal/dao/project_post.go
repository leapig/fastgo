package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type ProjectPost interface {
	Create(*entity.ProjectPost) (*entity.ProjectPost, error)
	Delete(*entity.ProjectPost) error
	Select(*entity.ProjectPost, *Pagination) ([]*entity.ProjectPost, error)
	Count(*entity.ProjectPost) (int32, error)
	Update(*entity.ProjectPost) (*entity.ProjectPost, error)
	FindByPk(en *entity.ProjectPost) (*entity.ProjectPost, error)
	SelectModel(en *entity.ProjectPost, pg *Pagination) ([]*model.ProjectPost, error)
}
type projectPost struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewProjectPost(db *gorm.DB, rs *helper.Redis) ProjectPost {
	return &projectPost{
		db: db,
		rs: rs,
	}
}
func (o *projectPost) Create(en *entity.ProjectPost) (*entity.ProjectPost, error) {
	tx := o.db.Create(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("create fatal")
	}
	return en, tx.Error
}

func (o *projectPost) Delete(uu *entity.ProjectPost) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *projectPost) Select(en *entity.ProjectPost, pg *Pagination) ([]*entity.ProjectPost, error) {
	sql := o.db.Model(&entity.ProjectMember{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	if en.Status != 0 {
		sql = sql.Where("status=?", en.Status)
	}
	if en.AddressCode != "" {
		sql = sql.Where("address_code like ?", en.AddressCode+"%")
	}
	if en.PostType != 0 {
		sql = sql.Where("post_type=?", en.PostType)
	}
	if en.IsAlarm != 0 {
		sql = sql.Where("is_alarm=?", en.IsAlarm)
	}
	var rows []*entity.ProjectPost
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *projectPost) Count(en *entity.ProjectPost) (int32, error) {
	sql := o.db.Model(&entity.ProjectMember{})
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Address != "" {
		sql = sql.Where("address like ?", "%"+en.Address+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_pk=?", en.ProjectPk)
	}
	if en.Status != 0 {
		sql = sql.Where("status=?", en.Status)
	}
	if en.AddressCode != "" {
		sql = sql.Where("address_code like ?", en.AddressCode+"%")
	}
	if en.PostType != 0 {
		sql = sql.Where("post_type=?", en.PostType)
	}
	if en.IsAlarm != 0 {
		sql = sql.Where("is_alarm=?", en.IsAlarm)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *projectPost) Update(en *entity.ProjectPost) (*entity.ProjectPost, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *projectPost) FindByPk(en *entity.ProjectPost) (*entity.ProjectPost, error) {
	res := &entity.ProjectPost{}
	tx := o.db.Model(res).Where("pk = ?", en.Pk).Find(res)
	return res, tx.Error
}
func (o *projectPost) SelectModel(en *entity.ProjectPost, pg *Pagination) ([]*model.ProjectPost, error) {
	sql := o.db.Model(&model.ProjectPost{}).
		Joins("LEFT JOIN `enterprise` ON `enterprise`.`pk` = `project_post`.`enterprise_pk`").
		Select("`enterprise`.`name` as enterprise_name,`project_post`.*").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.Name != "" {
		sql = sql.Where("project_post.name like ?", "%"+en.Name+"%")
	}
	if en.Address != "" {
		sql = sql.Where("project_post.address like ?", "%"+en.Address+"%")
	}
	if en.AddressCode != "" {
		sql = sql.Where("project_post.address_code like ?", en.AddressCode+"%")
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("project_post.enterprise_pk=?", en.EnterprisePk)
	}
	if en.ProjectPk != 0 {
		sql = sql.Where("project_post.project_pk=?", en.ProjectPk)
	}
	if en.Status != 0 {
		sql = sql.Where("project_post.status=?", en.Status)
	}
	if en.PostType != 0 {
		sql = sql.Where("project_post.post_type=?", en.PostType)
	}
	if en.IsAlarm != 0 {
		sql = sql.Where("project_post.is_alarm=?", en.IsAlarm)
	}
	var rows []*model.ProjectPost
	tx := sql.Order("project_post.create_at desc").Find(&rows)
	return rows, tx.Error
}
