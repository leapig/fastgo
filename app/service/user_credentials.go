package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
)

type UserCredentials interface {
	FindCardByUserPk(credentials *entity.UserCredentials) (*entity.UserCredentials, error)
	FindListCardByUserPk(credentials *entity.UserCredentials) ([]*entity.UserCredentials, error)
	Find(credentials *entity.UserCredentials) (*entity.UserCredentials, error)
	Create(credentials *entity.UserCredentials) (*entity.UserCredentials, error)
	Delete(credentials *entity.UserCredentials) error
	Update(credentials *entity.UserCredentials) error
}

// UserCredentials 接口规范实现类
type userCredentials struct {
	dao dao.Dao
}

// NewQfqzUser 实例化接口规范实现类
func NewUserCredentials(dao dao.Dao) UserCredentials {
	return &userCredentials{dao: dao}
}

func (u *userCredentials) FindCardByUserPk(credentials *entity.UserCredentials) (*entity.UserCredentials, error) {
	return u.dao.UserCredentials().FindCardByUserPk(credentials)
}

func (u *userCredentials) FindListCardByUserPk(credentials *entity.UserCredentials) ([]*entity.UserCredentials, error) {
	return u.dao.UserCredentials().FindList(credentials)
}

func (u *userCredentials) Find(credentials *entity.UserCredentials) (*entity.UserCredentials, error) {
	return u.dao.UserCredentials().Find(credentials)
}

func (u *userCredentials) Create(credentials *entity.UserCredentials) (*entity.UserCredentials, error) {
	if res, err := u.dao.UserCredentials().Find(credentials); err == nil && res.Pk != 0 {
		return nil, errors.New("数据重复！")
	}
	return u.dao.UserCredentials().Create(credentials)
}

func (u *userCredentials) Delete(credentials *entity.UserCredentials) error {
	return u.dao.UserCredentials().Delete(credentials)
}

func (u *userCredentials) Update(credentials *entity.UserCredentials) error {
	return u.dao.UserCredentials().Update(credentials)
}
