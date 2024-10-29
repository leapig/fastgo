package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type User interface {
	Find(*entity.User) (*entity.User, error)
	Save(*entity.User) (*entity.User, error)
	UpdatePhone(*entity.User) error
	SelectAccountPage(*entity.User, *entity.Pagination) ([]*model.QfqzAccountModel, int64, error)
	SelectEnterpriseAccountPage(user *model.QfqzAccountModelForEnterprise, page *entity.Pagination) ([]*model.QfqzAccountModelForEnterprise, int64, error)
	SelectUser(*entity.User, *entity.Pagination) ([]*model.QfqzUserModel, int64, error)
	FindUser(*entity.User) (*model.QfqzUserModel, error)
	FindOaOpenid(unionId string) (openid string)
	UpdateBaseInfo(*entity.User) error
}

// user 接口规范实现类
type user struct {
	dao dao.Dao
}

// NewUser 实例化接口规范实现类
func NewUser(dao dao.Dao) User {
	return &user{dao: dao}
}

func (u *user) Find(user *entity.User) (*entity.User, error) {
	return u.dao.User().Find(user)
}

func (u *user) Save(user *entity.User) (*entity.User, error) {
	if row, err := u.Find(&entity.User{Phone: user.Phone}); err == nil && row.Pk != 0 {
		return nil, errors.New("该手机号已被注册！")
	}
	user.Pk = helper.GetRid(helper.User)
	return u.dao.User().Create(user)
}

func (u *user) UpdatePhone(user *entity.User) error {
	if row, err := u.Find(&entity.User{Phone: user.Phone}); err == nil && row.Pk != user.Pk && row.Pk != 0 {
		return errors.New("该手机号已被占用")
	}
	return u.dao.User().Update(user)
}

func (u *user) SelectAccountPage(user *entity.User, page *entity.Pagination) ([]*model.QfqzAccountModel, int64, error) {
	if res, err := u.dao.User().SelectAccount(user, page); err == nil {
		count, _ := u.dao.User().Count(user)
		return res, count, err
	} else {
		return nil, 0, err
	}
}

func (u *user) FindOaOpenid(unionId string) (openid string) {
	return u.dao.User().FindOaOpenid(unionId)
}

func (u *user) SelectEnterpriseAccountPage(user *model.QfqzAccountModelForEnterprise, page *entity.Pagination) ([]*model.QfqzAccountModelForEnterprise, int64, error) {
	if res, err := u.dao.User().SelectEnterpriseAccount(user, page); err == nil {
		count, _ := u.dao.User().CountEnterprise(user)
		return res, count, err
	} else {
		return nil, 0, err
	}
}
func (u *user) SelectUser(user *entity.User, page *entity.Pagination) ([]*model.QfqzUserModel, int64, error) {
	if res, err := u.dao.User().SelectUser(user, page); err == nil {
		count, _ := u.dao.User().Count(user)
		return res, count, err
	} else {
		return nil, 0, err
	}
}

func (u *user) FindUser(user *entity.User) (*model.QfqzUserModel, error) {
	return u.dao.User().FindUser(user)
}

func (u *user) UpdateBaseInfo(en *entity.User) error {
	return u.dao.User().Update(en)
}
