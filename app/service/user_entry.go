package service

import (
	"github.com/go-errors/errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"time"
)

type UserEntry interface {
	Create(en *entity.UserEntry) (*entity.UserEntry, error)
	Select(en *model.UserEntryModel, pg *entity.Pagination, st, et *time.Time) ([]*model.UserEntryModel, int32, error)
	Update(en *entity.UserEntry) (*entity.UserEntry, error)
	FindByPk(en *entity.UserEntry) (*entity.UserEntry, error)
	CreateUserEntry(us *entity.User, e *entity.UserEntry) (*entity.UserEntry, error)
	Count(en *entity.UserEntry, st, et *time.Time) (int32, error)
}

type userEntry struct {
	dao dao.Dao
}

func NewUserEntry(dao dao.Dao) UserEntry {
	return &userEntry{dao: dao}
}

func (o *userEntry) Create(en *entity.UserEntry) (*entity.UserEntry, error) {
	return o.dao.UserEntry().Create(en)
}

func (o *userEntry) Select(en *model.UserEntryModel, pg *entity.Pagination, st, et *time.Time) ([]*model.UserEntryModel, int32, error) {
	res, err := o.dao.UserEntry().Select(en, pg, st, et)
	if err != nil {
		return nil, 0, err
	}
	count, countErr := o.dao.UserEntry().Count(en, st, et)
	if countErr != nil {
		return res, 0, countErr
	}
	return res, count, err
}

func (o *userEntry) Update(en *entity.UserEntry) (*entity.UserEntry, error) {
	return o.dao.UserEntry().Update(en)
}
func (o *userEntry) FindByPk(en *entity.UserEntry) (*entity.UserEntry, error) {
	return o.dao.UserEntry().FindByPk(en)
}

func (o *userEntry) CreateUserEntry(us *entity.User, e *entity.UserEntry) (*entity.UserEntry, error) {
	user, err := o.dao.User().Find(us)
	if err != nil {
		return nil, errors.New("人员未找到！")
	}
	e.UserPk = user.Pk
	//判断是否已入职
	eu, err := o.dao.EnterpriseUser().FindByUserPkAndEnterprisePk(&entity.EnterpriseUser{
		EnterprisePk: e.EnterprisePk,
		UserPk:       user.Pk,
	})
	if eu == nil && eu.Status == 1 {
		return nil, errors.New("该人员已入职")
	}
	if ue, ueErr := o.dao.UserEntry().Find(&entity.UserEntry{
		UserPk:       user.Pk,
		EnterprisePk: e.EnterprisePk,
	}); ueErr == nil {
		if ue.Status == "0" {
			return ue, nil
		} else if ue.Status == "1" {
			return o.dao.UserEntry().Create(e)
		} else if ue.Status == "2" {
			return o.dao.UserEntry().Create(e)
		} else {
			return nil, errors.New("数据错误，请联系管理人员")
		}
	} else {
		if ueErr.Error() == "find fatal" {
			return o.dao.UserEntry().Create(e)
		} else {
			return nil, ueErr
		}
	}

}

func (o *userEntry) Count(en *entity.UserEntry, st, et *time.Time) (int32, error) {
	return o.dao.UserEntry().Count(&model.UserEntryModel{
		Status: en.Status,
	}, st, et)
}
