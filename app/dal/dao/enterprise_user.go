package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type EnterpriseUser interface {
	Create(*entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	CreateEnterpriseUserAndUser(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	//CreateOrSelectEnterpriseUserAndUser(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	Delete(*entity.EnterpriseUser) error
	Select(*entity.EnterpriseUser, *Pagination) ([]*entity.EnterpriseUser, error)
	Count(*entity.EnterpriseUser) (int32, error)
	Update(*entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	FindByPk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	FindByUserPkAndEnterprisePk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	SelectWithDetail(en *entity.EnterpriseUser, pg *Pagination) ([]*entity.EnterpriseUser, error)
	SelectNoPage(en *model.EnterpriseUser) ([]*model.EnterpriseUser, error)
}
type enterpriseUser struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewEnterpriseUser(db *gorm.DB, rs *helper.Redis) EnterpriseUser {
	return &enterpriseUser{
		db: db,
		rs: rs,
	}
}
func (o *enterpriseUser) Create(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *enterpriseUser) CreateEnterpriseUserAndUser(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		var userMessage *entity.User
		userSql := o.db.Model(&entity.User{}).Where("phone = ?", en.Phone).First(&userMessage)
		if userSql.Error == nil && userMessage != nil && userMessage.Pk != 0 {
			en.UserPk = userMessage.Pk
		} else {
			userPk := helper.GetRid(helper.User)
			en.UserPk = userPk
			if userCreateSql := o.db.Create(&entity.User{
				Pk:       userPk,
				Name:     en.Name,
				Phone:    en.Phone,
				Gender:   en.Gender,
				Birthday: en.Birthday,
				Height:   en.Height,
				Weight:   en.Weight,
			}); userCreateSql.Error != nil {
				return userCreateSql.Error
			}
		}
		//判断人员是否存在，如存在则不再添加
		count, countErr := o.Count(&entity.EnterpriseUser{UserPk: en.UserPk, EnterprisePk: en.EnterprisePk})
		if countErr == nil {
			if count > 0 {
				//return errors.New("人员已存在")
				eu, euErr := o.FindByUserPkAndEnterprisePk(&entity.EnterpriseUser{UserPk: en.UserPk, EnterprisePk: en.EnterprisePk})
				if euErr == nil {
					en = eu
				} else {
					return euErr
				}
			} else {
				if enterpriseUserCreateSql := o.db.Create(en); enterpriseUserCreateSql.Error != nil {
					return enterpriseUserCreateSql.Error
				}
			}
		} else {
			return countErr
		}
		var baseRole *entity.Role
		baseRoleSql := o.db.Model(&entity.Role{}).Where("enterprise_pk = ?", en.EnterprisePk).Where("role_name = ?", "基础用户").Order("create_at desc").First(&baseRole)
		if baseRoleSql.Error == nil && baseRole.Pk != 0 {
			var userPermissionCount int64
			userPermissionCountSql := o.db.Model(&entity.UserPermission{}).Where("user_pk = ? ", en.UserPk).Where("enterprise_pk = ?", en.EnterprisePk).
				Where("permission_type = ?", 2).Where("permission_pk = ? ", baseRole.Pk).Count(&userPermissionCount)
			if userPermissionCountSql.Error == nil && userPermissionCount > 0 {
				return nil
			}
			if userPermissionCreateSql := o.db.Create(&entity.UserPermission{
				Pk:             helper.GetRid(helper.UserPermission),
				UserPk:         en.UserPk,
				EnterprisePk:   en.EnterprisePk,
				PermissionPk:   baseRole.Pk,
				PermissionType: 2,
			}); userPermissionCreateSql.Error != nil {
				return userPermissionCreateSql.Error
			}
		}
		return nil
	})
	return en, err
}

func (o *enterpriseUser) Delete(uu *entity.EnterpriseUser) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *enterpriseUser) Select(en *entity.EnterpriseUser, pg *Pagination) ([]*entity.EnterpriseUser, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Phone != "" {
		sql = sql.Where("phone = ?", en.Phone)
	}
	if en.Gender != 0 {
		sql = sql.Where("gender = ?", en.Gender)
	}
	if en.JobTitle != "" {
		sql = sql.Where("job_title = ?", en.JobTitle)
	}
	if en.JobNumber != "" {
		sql = sql.Where("job_number like ?", "%"+en.JobNumber+"%")
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	var rows []*entity.EnterpriseUser
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *enterpriseUser) Count(en *entity.EnterpriseUser) (int32, error) {
	sql := o.db.Model(&entity.EnterpriseUser{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Phone != "" {
		sql = sql.Where("phone = ?", en.Phone)
	}
	if en.Gender != 0 {
		sql = sql.Where("gender = ?", en.Gender)
	}
	if en.JobTitle != "" {
		sql = sql.Where("job_title = ?", en.JobTitle)
	}
	if en.JobNumber != "" {
		sql = sql.Where("job_number like ?", "%"+en.JobNumber+"%")
	}
	if en.Status != 0 {
		sql = sql.Where("status = ?", en.Status)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *enterpriseUser) Update(log *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	if log.Pk != 0 {
		tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
		if tx.RowsAffected == 0 {
			return log, errors.New("修改失败")
		}
	} else if log.UserPk != 0 {
		tx := o.db.Where("user_pk = ?", log.UserPk).Updates(&log)
		if tx.RowsAffected == 0 {
			return log, errors.New("修改失败")
		}
	}
	return log, nil
}
func (o *enterpriseUser) FindByPk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *enterpriseUser) FindByUserPkAndEnterprisePk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	sql := o.db.Model(&entity.EnterpriseUser{})
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	tx := sql.Find(&en)
	return en, tx.Error
}

func (o *enterpriseUser) SelectWithDetail(en *entity.EnterpriseUser, pg *Pagination) ([]*entity.EnterpriseUser, error) {
	sql := o.db.Model(&entity.EnterpriseUser{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Phone != "" {
		sql = sql.Where("phone = ?", en.Phone)
	}
	if en.Gender != 0 {
		sql = sql.Where("gender = ?", en.Gender)
	}
	if en.JobTitle != "" {
		sql = sql.Where("job_title = ?", en.JobTitle)
	}
	if en.JobNumber != "" {
		sql = sql.Where("job_number like ?", "%"+en.JobNumber+"%")
	}
	var rows []*entity.EnterpriseUser
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *enterpriseUser) SelectNoPage(en *model.EnterpriseUser) ([]*model.EnterpriseUser, error) {
	sql := o.db.Model(&model.EnterpriseUser{}).Preload("Enterprise")
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.Name != "" {
		sql = sql.Where("name like ?", "%"+en.Name+"%")
	}
	if en.Phone != "" {
		sql = sql.Where("phone = ?", en.Phone)
	}
	if en.Gender != 0 {
		sql = sql.Where("gender = ?", en.Gender)
	}
	if en.JobTitle != "" {
		sql = sql.Where("job_title = ?", en.JobTitle)
	}
	if en.JobNumber != "" {
		sql = sql.Where("job_number like ?", "%"+en.JobNumber+"%")
	}
	var rows []*model.EnterpriseUser
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
