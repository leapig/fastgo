package service

import (
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
)

type EnterpriseUser interface {
	Create(*entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	Delete(*entity.EnterpriseUser) error
	Select(*entity.EnterpriseUser, *dao.Pagination) ([]*entity.EnterpriseUser, *dao.Pagination, error)
	Update(*entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	FindByPk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	FindByUserPkAndEnterprisePk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)

	Count(user *entity.EnterpriseUser) (int32, error)
	DimissionEnterpriseUser(en *entity.EnterpriseUser) error
	CreateEnterpriseUser(*entity.EnterpriseUser) (*entity.EnterpriseUser, error)
}

type enterpriseUser struct {
	dao dao.Dao
}

func NewEnterpriseUser(dao dao.Dao) EnterpriseUser {
	return &enterpriseUser{dao: dao}
}
func (o *enterpriseUser) Create(in *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	in.Pk = helper.GetRid(helper.EnterpriseUser)
	return o.dao.EnterpriseUser().CreateEnterpriseUserAndUser(in)
}
func (o *enterpriseUser) Delete(in *entity.EnterpriseUser) error {
	return o.dao.EnterpriseUser().Delete(in)
}
func (o *enterpriseUser) Select(in *entity.EnterpriseUser, pg *dao.Pagination) ([]*entity.EnterpriseUser, *dao.Pagination, error) {
	if rows, err := o.dao.EnterpriseUser().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.EnterpriseUser().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *enterpriseUser) Update(in *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	return o.dao.EnterpriseUser().Update(in)
}
func (o *enterpriseUser) FindByPk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	return o.dao.EnterpriseUser().FindByPk(en)
}

func (o *enterpriseUser) FindByUserPkAndEnterprisePk(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	return o.dao.EnterpriseUser().FindByUserPkAndEnterprisePk(en)
}

func (o *enterpriseUser) DimissionEnterpriseUser(en *entity.EnterpriseUser) error {
	if res, err := o.dao.EnterpriseUser().FindByPk(&entity.EnterpriseUser{Pk: en.Pk}); err == nil {
		list, _ := o.dao.Member().List(&entity.Member{
			UserPk: res.Pk,
		})
		for _, value := range list {
			o.dao.Member().Delete(value)
		}
		_, updateErr := o.dao.EnterpriseUser().Update(&entity.EnterpriseUser{
			Pk:     res.Pk,
			Status: 2,
		})
		if updateErr == nil {
			//删除临时人员记录
			o.dao.TemporaryWorker().DeleteByEnterpriseUserPk(&entity.TemporaryWorker{EnterpriseUserPk: res.Pk})
			delErr := o.dao.UserPermission().DeleteByUserPkAndEnterprisePk(&entity.UserPermission{UserPk: res.UserPk, EnterprisePk: res.EnterprisePk})
			if delErr != nil {
				logger.Error(delErr)
			}
		}
		return updateErr
	} else {
		return err
	}
}

func (o *enterpriseUser) CreateEnterpriseUser(in *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	usr, usrErr := o.dao.User().Find(&entity.User{Pk: in.UserPk})
	if usrErr != nil {
		return nil, usrErr
	}
	res, resErr := o.dao.EnterpriseUser().FindByUserPkAndEnterprisePk(&entity.EnterpriseUser{UserPk: in.UserPk, EnterprisePk: in.EnterprisePk})
	if resErr == nil && res.Pk != 0 {
		return o.dao.EnterpriseUser().Update(&entity.EnterpriseUser{
			Pk:     res.Pk,
			Status: 1,
		})
	} else {
		eu2, eu2Err := o.dao.EnterpriseUser().CreateEnterpriseUserAndUser(&entity.EnterpriseUser{
			Pk:           helper.GetRid(helper.EnterpriseUser),
			UserPk:       in.UserPk,
			EnterprisePk: in.EnterprisePk,
			Name:         usr.Name,
			Phone:        usr.Phone,
			Gender:       usr.Gender,
			Birthday:     usr.Birthday,
			Height:       usr.Height,
			Weight:       usr.Weight,
			Status:       1,
		})
		if eu2Err != nil {
			return nil, errors.New("创建人员错误！")
		} else {
			return eu2, eu2Err
		}
	}
}

func (o *enterpriseUser) Count(user *entity.EnterpriseUser) (int32, error) {
	return o.dao.EnterpriseUser().Count(user)
}
