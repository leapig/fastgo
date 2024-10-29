package dao

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type UserProfession interface {
	Create(*entity.UserProfession) (*entity.UserProfession, error)
	Count(*entity.UserProfession) (int32, error)
	SelectUserProfession(en *model.UserProfessionQuery, pg *Pagination) ([]*model.UserProfession, int32, error)
	SelectUserProfessionForPlatform(en *model.UserProfessionForPlatformQuery, pg *Pagination) ([]*model.UserProfessionForPlatform, int32, error)
	GetProfessionUsers(model *model.QfqzProfessionUserModel, pg *entity.Pagination) ([]*model.QfqzProfessionUserModel, int32, error)
	SelectByUserPk(en *entity.UserProfession) ([]*model.UserProfessionWithEnterpriseNameAndStatus, error)
	SelectUserProfessionForSupervisePlatform(en *model.UserProfessionForSupervisePlatformQuery, pg *Pagination) ([]*model.UserProfessionForPlatform, int32, error)
}

type userProfession struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUserProfession(db *gorm.DB, rs *helper.Redis) UserProfession {
	return &userProfession{
		db: db,
		rs: rs,
	}
}

func (o *userProfession) Create(en *entity.UserProfession) (*entity.UserProfession, error) {
	en.Pk = helper.GetRid(helper.UserProfession)
	tx := o.db.Create(&en)
	return en, tx.Error
}
func (o *userProfession) Count(en *entity.UserProfession) (int32, error) {
	sql := o.db.Model(&entity.UserProfession{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.EnterpriseUserPk != 0 {
		sql = sql.Where("enterprise_user_pk = ?", en.EnterpriseUserPk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.Type != 0 {
		sql = sql.Where("type = ?", en.Type)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}
func (o *userProfession) SelectUserProfession(en *model.UserProfessionQuery, pg *Pagination) ([]*model.UserProfession, int32, error) {
	sql := o.db.Model(&model.UserProfession{}).
		Joins("left join enterprise_user d on d.pk = user_profession.enterprise_user_pk").
		Select("user_profession.*", "d.name  as user_name ", "d.phone as user_phone", "d.status as user_status",
			"d.gender as user_gender", "d.birthday as user_birthday", "d.height as user_height", "d.weight as user_weight").
		Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.UserProfessionType != 0 {
		sql = sql.Where("user_profession.type = ?", en.UserProfessionType)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("user_profession.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.UserStatus != 0 {
		sql = sql.Where("d.status = ?", en.UserStatus)
	}
	if en.UserGender != 0 {
		sql = sql.Where("d.gender = ?", en.UserGender)
	}
	if en.UserName != "" {
		sql = sql.Where("d.name like ?", "%"+en.UserName+"%")
	}
	if en.UserPhone != "" {
		sql = sql.Where("d.phone like ?", "%"+en.UserPhone+"%")
	}
	var rows []*model.UserProfession
	var count int64
	tx := sql.Order("user_profession.create_at desc").Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}

func (o *userProfession) SelectUserProfessionForPlatform(en *model.UserProfessionForPlatformQuery, pg *Pagination) ([]*model.UserProfessionForPlatform, int32, error) {
	sql := o.db.Model(&model.UserProfessionForPlatform{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql = sql.Where("pk in (?)", o.db.Table("user_profession").Order("create_at desc").Distinct("user_pk"))
	var rows []*model.UserProfessionForPlatform
	if en.UserGender != 0 {
		sql = sql.Where("gender = ?", en.UserGender)
	}
	if en.UserName != "" {
		sql = sql.Where("name like ?", "%"+en.UserName+"%")
	}
	if en.UserPhone != "" {
		sql = sql.Where("phone like ?", "%"+en.UserPhone+"%")
	}
	var count int64
	tx := sql.Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}
func (o *userProfession) GetProfessionUsers(profession *model.QfqzProfessionUserModel, pg *entity.Pagination) ([]*model.QfqzProfessionUserModel, int32, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("User")
	sql.Preload("UserCredentials")
	if profession.User.Phone != "" {
		sql.Where("user.phone like ?", "%"+profession.User.Phone+"%")
	}
	if profession.User.Name != "" {
		sql.Where("user.name like ?", "%"+profession.User.Name+"%")
	}
	rows := make([]*model.QfqzProfessionUserModel, 0)
	tx := sql.Find(&rows)
	count, _ := o.ProfessionCount(profession)
	return rows, count, tx.Error
}

func (o *userProfession) ProfessionCount(profession *model.QfqzProfessionUserModel) (int32, error) {
	sql := o.db.Model(&model.QfqzProfessionUserModel{})
	if profession.User.Phone != "" {
		sql.Where("phone like ?", "%"+profession.User.Phone+"%")
	}
	if profession.User.Name != "" {
		sql.Where("name like ?", "%"+profession.User.Name+"%")
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error

}
func (o *userProfession) SelectByUserPk(en *entity.UserProfession) ([]*model.UserProfessionWithEnterpriseNameAndStatus, error) {
	sql := o.db.Model(&model.UserProfessionWithEnterpriseNameAndStatus{}).
		Joins("left join enterprise_user d on d.pk = user_profession.enterprise_user_pk").
		Joins("left join enterprise e on e.pk = user_profession.enterprise_pk").
		Select("user_profession.*", "d.status  as status ", "e.name as enterprise_name")
	if en.UserPk != 0 {
		sql = sql.Where("user_profession.user_pk = ?", en.UserPk)
	}
	var rows []*model.UserProfessionWithEnterpriseNameAndStatus
	tx := sql.Order("user_profession.create_at desc").Find(&rows)
	return rows, tx.Error
}

// SelectUserProfessionForSupervisePlatform 监管单位查询从业人员
func (o *userProfession) SelectUserProfessionForSupervisePlatform(en *model.UserProfessionForSupervisePlatformQuery, pg *Pagination) ([]*model.UserProfessionForPlatform, int32, error) {
	sql := o.db.Model(&model.UserProfessionForPlatform{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql = sql.Where("pk in (?)", o.db.Table("project_member").Where("project_pk in (?)", en.ProjectPks).Order("create_at desc").Distinct("user_pk"))
	var rows []*model.UserProfessionForPlatform
	if en.UserGender != 0 {
		sql = sql.Where("gender = ?", en.UserGender)
	}
	if en.UserName != "" {
		sql = sql.Where("name like ?", "%"+en.UserName+"%")
	}
	if en.UserPhone != "" {
		sql = sql.Where("phone like ?", "%"+en.UserPhone+"%")
	}
	var count int64
	tx := sql.Find(&rows).Count(&count)
	return rows, int32(count), tx.Error
}
