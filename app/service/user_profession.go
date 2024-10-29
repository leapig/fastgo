package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
)

type UserProfession interface {
	GetProfessionUsers(user *model.QfqzProfessionUserModel, e *entity.Pagination) ([]*model.QfqzProfessionUserModel, int32, error)
	SelectUserProfession(in *model.UserProfessionQuery, pg *dao.Pagination) ([]*model.UserProfession, *dao.Pagination, error)
	SelectUserProfessionForPlatform(in *model.UserProfessionForPlatformQuery, pg *dao.Pagination) ([]*model.UserProfessionForPlatform, *dao.Pagination, error)
	SelectProjectForUserProfession(in *model.ProjectQueryForUserProfession) ([]*model.Project, error)
	CheckUserProfessionIsRealName(userName string, userPk int64) int32
	SelectUserCredentials(in *entity.UserCredentials) ([]*entity.UserCredentials, error)
	SelectUserEnterpriseByUserPk(userPk int64) ([]*model.UserProfessionWithEnterpriseNameAndStatus, error)
	FindEnterpriseAddressCodeList(areaPermission *entity.EnterpriseAreaPermission) ([]*entity.EnterpriseAreaPermission, error)
	SelectUserProfessionForSupervisePlatform(in *model.UserProfessionForSupervisePlatformQuery, pg *dao.Pagination) ([]*model.UserProfessionForPlatform, *dao.Pagination, error)
	SelectProjectNoPage(in *entity.Project) ([]*entity.Project, error)
}

// userProfession 接口规范实现类
type userProfession struct {
	dao dao.Dao
}

// NewUserProfession 实例化接口规范实现类
func NewUserProfession(dao dao.Dao) UserProfession {
	return &userProfession{dao: dao}
}

func (o *userProfession) GetProfessionUsers(user *model.QfqzProfessionUserModel, pg *entity.Pagination) ([]*model.QfqzProfessionUserModel, int32, error) {
	return o.dao.UserProfession().GetProfessionUsers(user, pg)
}

func (o *userProfession) SelectUserProfession(in *model.UserProfessionQuery, pg *dao.Pagination) ([]*model.UserProfession, *dao.Pagination, error) {
	if rows, count, err := o.dao.UserProfession().SelectUserProfession(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total = count
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *userProfession) SelectUserProfessionForPlatform(in *model.UserProfessionForPlatformQuery, pg *dao.Pagination) ([]*model.UserProfessionForPlatform, *dao.Pagination, error) {
	if rows, count, err := o.dao.UserProfession().SelectUserProfessionForPlatform(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total = count
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *userProfession) SelectProjectForUserProfession(in *model.ProjectQueryForUserProfession) ([]*model.Project, error) {
	return o.dao.Project().SelectProjectForUserProfession(in)
}
func (o *userProfession) CheckUserProfessionIsRealName(userName string, userPk int64) int32 {
	count, _ := o.dao.UserRealNameAuthenticationLog().Count(&entity.UserRealNameAuthenticationLog{
		UserPk: userPk,
		Name:   userName,
	})
	if count > 0 {
		return 1
	}
	return 0
}
func (o *userProfession) SelectUserCredentials(in *entity.UserCredentials) ([]*entity.UserCredentials, error) {
	return o.dao.UserCredentials().FindList(in)
}
func (o *userProfession) SelectUserEnterpriseByUserPk(userPk int64) ([]*model.UserProfessionWithEnterpriseNameAndStatus, error) {
	return o.dao.UserProfession().SelectByUserPk(&entity.UserProfession{
		UserPk: userPk,
	})
}

// FindEnterpriseAddressCodeList 查询监管单位管理辖区
func (o *userProfession) FindEnterpriseAddressCodeList(areaPermission *entity.EnterpriseAreaPermission) ([]*entity.EnterpriseAreaPermission, error) {
	return o.dao.EnterpriseAreaPermission().FindList(areaPermission)
}

func (o *userProfession) SelectUserProfessionForSupervisePlatform(in *model.UserProfessionForSupervisePlatformQuery, pg *dao.Pagination) ([]*model.UserProfessionForPlatform, *dao.Pagination, error) {
	if rows, count, err := o.dao.UserProfession().SelectUserProfessionForSupervisePlatform(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total = count
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *userProfession) SelectProjectNoPage(in *entity.Project) ([]*entity.Project, error) {
	return o.dao.Project().SelectNoPage(in)
}
