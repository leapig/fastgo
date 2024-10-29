package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type Member interface {
	Create(m *entity.Member) (*entity.Member, error)
	CreateMemberUser(member *entity.Member, user *entity.User, userCredentials []*entity.UserCredentials) (*entity.Member, error)
	Delete(m *entity.Member) error
	DeleteByEnterprisePkAndUserPk(m *entity.Member) error
	Select(in *model.Member, pg *entity.Pagination) ([]*model.Member, *entity.Pagination, error)
	Count(m *entity.Member) (int32, error)
	UpdateMemberUser(in *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	UpdateUserDepartmentRelation(in *entity.Member, departmentPks []int64) error
	CreateMemberAndEnterpriseUser(eu *entity.EnterpriseUser, in *entity.Member) (*entity.Member, error)
	List(m *entity.Member) ([]*entity.Member, error)
}

// member 接口规范实现类
type member struct {
	dao dao.Dao
}

func NewMember(dao dao.Dao) Member {
	return &member{dao: dao}
}

func (o *member) Create(m *entity.Member) (*entity.Member, error) {
	return o.dao.Member().Create(m)
}
func (o *member) CreateMemberUser(member *entity.Member, user *entity.User, userCredentials []*entity.UserCredentials) (*entity.Member, error) {
	return o.dao.Member().CreateMemberUser(member, user, userCredentials)
}
func (o *member) Delete(m *entity.Member) error {
	return o.dao.Member().Delete(m)
}
func (o *member) DeleteByEnterprisePkAndUserPk(m *entity.Member) error {
	return o.dao.Member().DeleteByEnterprisePkAndUserPk(m)
}
func (o *member) Select(in *model.Member, pg *entity.Pagination) ([]*model.Member, *entity.Pagination, error) {
	if rows, err := o.dao.Member().SelectMemberDetail(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Member().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *member) Count(m *entity.Member) (int32, error) {
	return o.dao.Member().Count(&model.Member{
		UserPk:       m.UserPk,
		EnterprisePk: m.EnterprisePk,
		DepartmentPk: m.DepartmentPk,
	})
}
func (o *member) UpdateMemberUser(in *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	return o.dao.Member().UpdateMemberUser(in)
}
func (o *member) UpdateUserDepartmentRelation(in *entity.Member, departmentPks []int64) error {
	return o.dao.Member().UpdateUserDepartmentRelation(in, departmentPks)
}

func (o *member) CreateMemberAndEnterpriseUser(eu *entity.EnterpriseUser, in *entity.Member) (*entity.Member, error) {
	if eu.UserPk != 0 {
		usr, usrErr := o.dao.User().Find(&entity.User{Pk: eu.UserPk})
		if usrErr != nil {
			return nil, usrErr
		}
		res, resErr := o.dao.EnterpriseUser().FindByUserPkAndEnterprisePk(&entity.EnterpriseUser{UserPk: eu.UserPk, EnterprisePk: eu.EnterprisePk})
		if resErr == nil && res.Pk != 0 {
			o.dao.EnterpriseUser().Update(&entity.EnterpriseUser{
				Pk:     res.Pk,
				Status: 1,
			})
			in.UserPk = res.Pk
			count, _ := o.dao.Member().Count(&model.Member{
				UserPk: res.Pk,
			})
			in.IsLeader = 1
			if count > 0 {
				in.IsMain = 1
			}
			return o.dao.Member().Create(in)
		} else {
			eu2, eu2Err := o.dao.EnterpriseUser().CreateEnterpriseUserAndUser(&entity.EnterpriseUser{
				Pk:           helper.GetRid(helper.EnterpriseUser),
				UserPk:       eu.UserPk,
				EnterprisePk: eu.EnterprisePk,
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
				in.UserPk = eu2.Pk
				count, _ := o.dao.Member().Count(&model.Member{
					UserPk: eu2.Pk,
				})
				in.IsLeader = 1
				if count > 0 {
					in.IsMain = 1
				}
				return o.dao.Member().Create(in)
			}
		}
	} else {
		return nil, errors.New("参数错误")
	}
}

func (o *member) List(m *entity.Member) ([]*entity.Member, error) {
	return o.dao.Member().List(m)
}
