package dao

import (
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type User interface {
	Find(*entity.User) (*entity.User, error)
	Create(*entity.User) (*entity.User, error)
	Update(*entity.User) error
	Count(*entity.User) (int64, error)
	SelectAccount(*entity.User, *entity.Pagination) ([]*model.QfqzAccountModel, error)
	SelectEnterpriseAccount(*model.QfqzAccountModelForEnterprise, *entity.Pagination) ([]*model.QfqzAccountModelForEnterprise, error)
	CountEnterprise(user *model.QfqzAccountModelForEnterprise) (int64, error)
	SelectUser(*entity.User, *entity.Pagination) ([]*model.QfqzUserModel, error)
	FindUser(*entity.User) (*model.QfqzUserModel, error)
	FindOaOpenid(string) string
}

type user struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUser(db *gorm.DB, rs *helper.Redis) User {
	return &user{
		db: db,
		rs: rs,
	}
}

func (o *user) Find(ur *entity.User) (*entity.User, error) {
	sql := o.db.Model(&entity.User{})
	if ur.Pk != 0 {
		sql.Where("pk = ?", ur.Pk)
	}
	if ur.Phone != "" {
		sql.Where("phone = ?", ur.Phone)
	}
	tx := sql.First(&ur)
	return ur, tx.Error
}

// FindOaOpenid 查找公众号opened
func (o *user) FindOaOpenid(unionId string) (openid string) {
	o.db.Model(&entity.UserClient{}).Where("client_type=? and wx_unionid=?", 2, unionId).Select("open_id").Find(&openid)
	return
}

func (o *user) Create(user *entity.User) (*entity.User, error) {
	tx := o.db.Create(&user)
	return user, tx.Error
}

func (o *user) Update(user *entity.User) error {
	return o.db.Transaction(func(tx *gorm.DB) error {
		if res := tx.Where("pk = ?", user.Pk).Updates(user); res.Error != nil {
			return res.Error
		}
		return nil
	})
}

func (o *user) Count(user *entity.User) (int64, error) {
	sql := o.db.Model(&entity.User{})
	if user.Phone != "" {
		sql.Where("phone like ?", "%"+user.Phone+"%")
	}
	if user.Name != "" {
		sql.Where("name like ?", "%"+user.Name+"%")
	}
	var count int64
	tx := sql.Count(&count)
	return count, tx.Error
}

func (o *user) SelectAccount(user *entity.User, pg *entity.Pagination) ([]*model.QfqzAccountModel, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("UserPermissionList")
	sql.Preload("UserClientList")
	if user.Phone != "" {
		sql.Where("user.phone like ?", "%"+user.Phone+"%")
	}
	if user.Name != "" {
		sql.Where("user.name like ?", "%"+user.Name+"%")
	}
	rows := make([]*model.QfqzAccountModel, 0)
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *user) SelectEnterpriseAccount(user *model.QfqzAccountModelForEnterprise, pg *entity.Pagination) ([]*model.QfqzAccountModelForEnterprise, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("UserPermissionList")
	sql.Preload("UserClientList")
	if user.Phone != "" {
		sql.Where("enterprise_user.phone like ?", "%"+user.Phone+"%")
	}
	if user.Name != "" {
		sql.Where("enterprise_user.name like ?", "%"+user.Name+"%")
	}
	if user.EnterprisePk != 0 {
		sql.Where("enterprise_user.enterprise_pk = ?", user.EnterprisePk)
	}
	rows := make([]*model.QfqzAccountModelForEnterprise, 0)
	tx := sql.Find(&rows)
	return rows, tx.Error
}
func (o *user) CountEnterprise(user *model.QfqzAccountModelForEnterprise) (int64, error) {
	sql := o.db.Model(&model.QfqzAccountModelForEnterprise{})
	if user.Phone != "" {
		sql.Where("phone like ?", "%"+user.Phone+"%")
	}
	if user.Name != "" {
		sql.Where("name like ?", "%"+user.Name+"%")
	}
	if user.EnterprisePk != 0 {
		sql.Where("enterprise_pk = ?", user.EnterprisePk)
	}
	var count int64
	tx := sql.Count(&count)
	return count, tx.Error
}
func (o *user) SelectUser(user *entity.User, pg *entity.Pagination) ([]*model.QfqzUserModel, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("UserRealNameAuthenticationLog")
	sql.Preload("UserLiveness")
	sql.Preload("UserCredentials")
	if user.Phone != "" {
		sql.Where("user.phone like ?", "%"+user.Phone+"%")
	}
	if user.Name != "" {
		sql.Where("user.name like ?", "%"+user.Name+"%")
	}
	rows := make([]*model.QfqzUserModel, 0)
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *user) FindUser(user *entity.User) (*model.QfqzUserModel, error) {
	sql := o.db.Model(&model.QfqzUserModel{})
	sql.Preload("UserRealNameAuthenticationLog")
	sql.Preload("UserLiveness")
	sql.Preload("UserCredentials")
	if user.Pk != 0 {
		sql.Where("user.pk = ?", user.Pk)
	}
	if user.Phone != "" {
		sql.Where("user.phone like ?", "%"+user.Phone+"%")
	}
	if user.Name != "" {
		sql.Where("user.name like ?", "%"+user.Name+"%")
	}
	row := &model.QfqzUserModel{}
	tx := sql.First(&row)
	return row, tx.Error
}
