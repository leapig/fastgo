package dao

import (
	"errors"
	"fmt"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type Enterprise interface {
	Create(*entity.Enterprise) (*model.Enterprise, error)
	Delete(*entity.Enterprise) error
	Select(*entity.Enterprise, *Pagination) ([]*model.Enterprise, error)
	Count(*entity.Enterprise) (int32, error)
	Update(*entity.Enterprise) (*model.Enterprise, error)
	FindByPk(en *entity.Enterprise) (*model.Enterprise, error)
	FindEnterpriseByUserPk(en *model.EnterpriseModel) ([]*entity.Enterprise, error)
	SelectEnterpriseForEnterpriseAreaPermission(list []*entity.EnterpriseAreaPermission, en *entity.Enterprise, pg *entity.Pagination) ([]*model.Enterprise, error)
	CountEnterpriseForEnterpriseAreaPermission(list []*entity.EnterpriseAreaPermission, en *entity.Enterprise) (int32, error)
}
type enterprise struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewEnterprise(db *gorm.DB, rs *helper.Redis) Enterprise {
	return &enterprise{
		db: db,
		rs: rs,
	}
}
func (o *enterprise) Create(en *entity.Enterprise) (*model.Enterprise, error) {
	tx := o.db.Create(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("delete fatal")
	}
	return o.FindByPk(en)
}

func (o *enterprise) Delete(uu *entity.Enterprise) error {
	tx := o.db.Where("Pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *enterprise) Select(en *entity.Enterprise, pg *Pagination) ([]*model.Enterprise, error) {
	sql := o.db.Model(&model.Enterprise{}).Joins(
		"LEFT JOIN `user` ON `user`.`pk` = `enterprise`.`corporate_pk`").Select(
		"`user`.`name` as corporate_name,`user`.`phone` as corporate_phone,`enterprise`.*").Limit(
		int(pg.Size)).Offset(
		int((pg.Page - 1) * pg.Size))
	if en.Name != "" {
		sql = sql.Where("enterprise.name like ?", "%"+en.Name+"%")
	}
	if en.Serial != "" {
		sql = sql.Where("enterprise.serial like ?", "%"+en.Serial+"%")
	}
	//if en.Country != "" {
	//	sql = sql.Where("enterprise.country = ?", en.Country)
	//}
	//if en.Province != "" {
	//	sql = sql.Where("enterprise.province = ?", en.Province)
	//}
	//if en.City != "" {
	//	sql = sql.Where("enterprise.city = ?", en.City)
	//}
	//if en.District != "" {
	//	sql = sql.Where("enterprise.district = ?", en.District)
	//}
	//if en.County != "" {
	//	sql = sql.Where("enterprise.county = ?", en.County)
	//}
	if en.AddressCode != "" {
		sql = sql.Where("enterprise.address_code like ?", en.AddressCode+"%")
	}

	if en.Type != 0 {
		sql = sql.Where("enterprise.type = ?", en.Type)
	}
	var rows []*model.Enterprise
	tx := sql.Order("enterprise.create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *enterprise) Count(en *entity.Enterprise) (int32, error) {
	sql := o.db.Model(&entity.Enterprise{})
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Serial != "" {
		sql = sql.Where("enterprise.serial like ?", "%"+en.Serial+"%")
	}
	//if en.Country != "" {
	//	sql = sql.Where("enterprise.country = ?", en.Country)
	//}
	//if en.Province != "" {
	//	sql = sql.Where("enterprise.province = ?", en.Province)
	//}
	//if en.City != "" {
	//	sql = sql.Where("enterprise.city = ?", en.City)
	//}
	//if en.District != "" {
	//	sql = sql.Where("enterprise.district = ?", en.District)
	//}
	//if en.County != "" {
	//	sql = sql.Where("enterprise.county = ?", en.County)
	//}
	if en.AddressCode != "" {
		sql = sql.Where("enterprise.address_code like ?", en.AddressCode+"%")
	}
	if en.Type != 0 {
		sql = sql.Where("type = ?", en.Type)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *enterprise) Update(en *entity.Enterprise) (*model.Enterprise, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return nil, errors.New("update fatal")
	}
	return o.FindByPk(en)
}
func (o *enterprise) FindByPk(en *entity.Enterprise) (*model.Enterprise, error) {
	res := &model.Enterprise{}
	tx := o.db.Model(res).Joins(
		"LEFT JOIN `user` ON `user`.`pk` = `enterprise`.`corporate_pk`").Select(
		"`user`.`name` as corporate_name,`user`.`phone` as corporate_phone,`enterprise`.*").Where("enterprise.pk = ?", en.Pk).Find(res)
	return res, tx.Error
}

func (o *enterprise) FindEnterpriseByUserPk(en *model.EnterpriseModel) ([]*entity.Enterprise, error) {
	rows := make([]*entity.Enterprise, 0)
	sql := o.db.Select("enterprise.*").Joins("LEFT JOIN `enterprise_user` ON `enterprise_user`.`enterprise_pk` = `enterprise`.`pk`").Group("enterprise.pk")
	tx := sql.Where("enterprise_user.`user_pk` = ?", en.EnterpriseUser.UserPk).Find(&rows)
	return rows, tx.Error
}

func (o *enterprise) SelectEnterpriseForEnterpriseAreaPermission(list []*entity.EnterpriseAreaPermission, en *entity.Enterprise, pg *entity.Pagination) ([]*model.Enterprise, error) {
	sql := o.db.Model(&model.Enterprise{})
	sql = sql.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	for _, value := range list {
		str := "1=1"
		//if value.Province != "" {
		//	//sql.Where("province = ?", value.Province)
		//	str += fmt.Sprintf(" and province = '%s' ", value.Province)
		//}
		//if value.City != "" {
		//	//sql.Where("city = ?", value.City)
		//	str += fmt.Sprintf(" and city = '%s' ", value.City)
		//}
		//if value.District != "" {
		//	//sql.Where("district = ?", value.District)
		//	str += fmt.Sprintf(" and district = '%s' ", value.District)
		//}
		//if value.County != "" {
		//	//sql.Where("county = ?", value.County)
		//	str += fmt.Sprintf(" and county = '%s' ", value.County)
		//}
		//if en.Type != 0 {
		//	str += fmt.Sprintf(" and type = %d", en.Type)
		//}
		//if en.Name != "" {
		//	str += fmt.Sprintf(" and name like %" + en.Name + "%")
		//}
		str += fmt.Sprintf(" and address_code  like '%s'", value.AddressCode+"%")
		sql = sql.Or(str)
	}
	var row []*model.Enterprise
	tx := sql.Find(&row)
	return row, tx.Error
}
func (o *enterprise) CountEnterpriseForEnterpriseAreaPermission(list []*entity.EnterpriseAreaPermission, en *entity.Enterprise) (int32, error) {
	sql := o.db.Model(&model.Enterprise{})
	for _, value := range list {
		str := "1=1"
		//if value.Province != "" {
		//	//sql.Where("province = ?", value.Province)
		//	str += fmt.Sprintf(" and province = '%s' ", value.Province)
		//}
		//if value.City != "" {
		//	//sql.Where("city = ?", value.City)
		//	str += fmt.Sprintf(" and city = '%s' ", value.City)
		//}
		//if value.District != "" {
		//	//sql.Where("district = ?", value.District)
		//	str += fmt.Sprintf(" and district = '%s' ", value.District)
		//}
		//if value.County != "" {
		//	//sql.Where("county = ?", value.County)
		//	str += fmt.Sprintf(" and county = '%s' ", value.County)
		//}
		//if en.Type != 0 {
		//	str += fmt.Sprintf(" and type = %d", en.Type)
		//}
		//if en.Name != "" {
		//	str += fmt.Sprintf(" and name like %" + en.Name + "%")
		//}
		str += fmt.Sprintf(" and address_code  like '%s'", value.AddressCode+"%")
		sql = sql.Or(str)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}
