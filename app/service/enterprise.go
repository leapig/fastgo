package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
)

type Enterprise interface {
	ReadUserEnterprise(userPk int64) ([]*model.Enterprise, error)
	ReadAppUserEnterprise(userPk int64) ([]*model.Enterprise, error)
	Create(*entity.Enterprise) (*model.Enterprise, error)
	Delete(*entity.Enterprise) error
	Select(*entity.Enterprise, *dao.Pagination) ([]*model.Enterprise, *dao.Pagination, error)
	Update(*entity.Enterprise) (*model.Enterprise, error)
	FindByPk(en *entity.Enterprise) (*model.Enterprise, error)
	SelectEnterpriseForEnterpriseAreaPermission(list []*entity.EnterpriseAreaPermission, en *entity.Enterprise, pg *entity.Pagination) ([]*model.Enterprise, int32, error)
}

type enterprise struct {
	dao dao.Dao
}

func NewEnterprise(dao dao.Dao) Enterprise {
	return &enterprise{dao: dao}
}

// ReadUserEnterprise 获取租户列表方法
func (o *enterprise) ReadUserEnterprise(userPk int64) ([]*model.Enterprise, error) {
	models := make([]*model.Enterprise, 0)
	//判断是否有运营平台管理权限
	if bol, bolErr := o.dao.UserPermission().CheckUserEnterprisePlatformPermission(userPk, 1); bolErr == nil && bol {
		models = append(models, &model.Enterprise{Enterprise: entity.Enterprise{Pk: 1, Name: "运营平台", Type: 1000}})
	}
	if res, err := o.dao.Enterprise().FindEnterpriseByUserPk(&model.EnterpriseModel{EnterpriseUser: entity.EnterpriseUser{UserPk: userPk}}); err == nil {
		for _, value := range res {
			if bol, bolErr := o.dao.UserPermission().CheckUserEnterprisePlatformPermission(userPk, value.Pk); bolErr == nil && bol {
				models = append(models, &model.Enterprise{Enterprise: entity.Enterprise{
					Pk:    value.Pk,
					Name:  value.Name,
					Cover: value.Cover,
					Type:  value.Type,
				}})
			}
		}
		return models, err
	} else {
		return nil, err
	}
}

// ReadAppUserEnterprise 获取租户列表方法(小程序）
func (o *enterprise) ReadAppUserEnterprise(userPk int64) ([]*model.Enterprise, error) {
	models := make([]*model.Enterprise, 0)
	if res, err := o.dao.Enterprise().FindEnterpriseByUserPk(&model.EnterpriseModel{EnterpriseUser: entity.EnterpriseUser{UserPk: userPk}}); err == nil {
		for _, value := range res {
			if bol, bolErr := o.dao.UserPermission().CheckUserEnterpriseAppletPermission(userPk, value.Pk); bolErr == nil && bol {
				models = append(models, &model.Enterprise{Enterprise: entity.Enterprise{
					Pk:    value.Pk,
					Name:  value.Name,
					Cover: value.Cover,
					Type:  value.Type,
				}})
			}
		}
		return models, err
	} else {
		return nil, err
	}
}

func (o *enterprise) Create(in *entity.Enterprise) (*model.Enterprise, error) {
	in.Pk = helper.GetRid(helper.Enterprise)
	en, err := o.dao.Enterprise().Create(in)
	if err == nil {
		//创建根部门  创建初始租户权限角色
		o.dao.Department().Create(&entity.Department{
			Pk:           helper.GetRid(helper.Department),
			Name:         en.Name,
			EnterprisePk: en.Pk,
			ParentPk:     en.Pk,
		})
		rl, _ := o.dao.Role().Create(&entity.Role{
			Pk:           helper.GetRid(helper.Role),
			RoleName:     "超级管理员",
			EnterprisePk: en.Pk,
		})
		baseRole, _ := o.dao.Role().Create(&entity.Role{
			Pk:           helper.GetRid(helper.Role),
			RoleName:     "基础用户",
			EnterprisePk: en.Pk,
		})
		var permissionGroupPk int64
		permissionGroupList := make([]*entity.PermissionGroup, 0)
		var permissionGroupErr error
		if in.Type == 1 {
			permissionGroupList, permissionGroupErr = o.dao.PermissionGroup().ListSupervisePermissionGroup()
			if permissionGroupErr != nil {
				return nil, permissionGroupErr
			}
			permissionGroupPk = 3042882110889984
		} else if in.Type == 2 {
			permissionGroupList, permissionGroupErr = o.dao.PermissionGroup().ListTenantPermissionGroup()
			if permissionGroupErr != nil {
				return nil, permissionGroupErr
			}
			permissionGroupPk = 3042881070702592
		} else if in.Type == 3 {
			permissionGroupList, permissionGroupErr = o.dao.PermissionGroup().ListUserPlatformPermissionGroup()
			if permissionGroupErr != nil {
				return nil, permissionGroupErr
			}
			permissionGroupPk = 3042881708236800
		}
		//基础用户角色赋予权限组
		_, _ = o.dao.RolePermission().Create(&entity.RolePermission{
			Pk:             helper.GetRid(helper.RolePermission),
			RolePk:         baseRole.Pk,
			PermissionPk:   permissionGroupPk,
			PermissionType: 2,
		})
		for _, value := range permissionGroupList {
			_, _ = o.dao.RolePermission().Create(&entity.RolePermission{
				Pk:             helper.GetRid(helper.RolePermission),
				RolePk:         rl.Pk,
				PermissionPk:   value.Pk,
				PermissionType: 2,
			})
		}
		return en, err
	} else {
		return nil, err
	}

}
func (o *enterprise) Delete(in *entity.Enterprise) error {
	return o.dao.Enterprise().Delete(in)
}
func (o *enterprise) Select(in *entity.Enterprise, pg *dao.Pagination) ([]*model.Enterprise, *dao.Pagination, error) {
	if rows, err := o.dao.Enterprise().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Enterprise().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *enterprise) Update(in *entity.Enterprise) (*model.Enterprise, error) {
	return o.dao.Enterprise().Update(in)
}
func (o *enterprise) FindByPk(en *entity.Enterprise) (*model.Enterprise, error) {
	return o.dao.Enterprise().FindByPk(en)
}

// SelectEnterpriseForEnterpriseAreaPermission 根据辖区查询单位
func (o *enterprise) SelectEnterpriseForEnterpriseAreaPermission(list []*entity.EnterpriseAreaPermission, en *entity.Enterprise, pg *entity.Pagination) ([]*model.Enterprise, int32, error) {
	if row, err := o.dao.Enterprise().SelectEnterpriseForEnterpriseAreaPermission(list, en, pg); err == nil {
		count, _ := o.dao.Enterprise().CountEnterpriseForEnterpriseAreaPermission(list, en)
		return row, count, err
	} else {
		return nil, 0, err
	}
}
