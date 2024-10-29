package dao

import (
	"errors"
	"fmt"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"strconv"
)

type UserPermission interface {
	Create(*entity.UserPermission) (*entity.UserPermission, error)
	CreateForCorporate(en *entity.UserPermission) error
	Delete(*entity.UserPermission) error
	DeleteByUserPkAndEnterprisePk(uu *entity.UserPermission) error
	Select(*entity.UserPermission, *Pagination) ([]*entity.UserPermission, error)
	FindList(p *entity.UserPermission) ([]*entity.UserPermission, error)
	Count(p *entity.UserPermission) (int32, error)
	Update(*entity.UserPermission) (*entity.UserPermission, error)
	SelectAllUserPermission(en *entity.UserPermission) ([]*entity.UserPermission, error)
	SelectWithUser(en *entity.UserPermission, pg *Pagination) ([]*model.UserPermissionModel, error)
	SelectWithEnterpriseUser(en *entity.UserPermission, pg *Pagination) ([]*model.UserPermissionForEnterprise, error)
	CheckUserEnterprisePlatformPermission(userPk, enterprisePk int64) (bool, error)
	CheckUserEnterpriseAppletPermission(userPk, enterprisePk int64) (bool, error)
	FindByPk(en *entity.UserPermission) (*entity.UserPermission, error)
	DeleteRedisPermission(enterprisePk, userPk int64) error
}
type userPermission struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUserPermission(db *gorm.DB, rs *helper.Redis) UserPermission {
	return &userPermission{
		db: db,
		rs: rs,
	}
}
func (o *userPermission) Create(en *entity.UserPermission) (*entity.UserPermission, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}
func (o *userPermission) CreateForCorporate(en *entity.UserPermission) error {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		var userMessage *entity.User
		if userSql := o.db.Model(&entity.User{}).Where("pk = ? ", en.UserPk).Order("create_at desc").First(&userMessage); userSql.Error != nil || userMessage.Pk == 0 {
			return userSql.Error
		}
		var departmentMessage *entity.Department
		if departmentSql := o.db.Model(&entity.Department{}).Where("enterprise_pk = ? ", en.EnterprisePk).Where("parent_pk = ?", en.EnterprisePk).Order("create_at desc").First(&departmentMessage); departmentSql.Error != nil {
			return departmentSql.Error
		}
		var enterpriseUserMessage *entity.EnterpriseUser
		var enterpriseUserCount int64
		enterpriseUserSql := o.db.Model(&entity.EnterpriseUser{}).Where("user_pk = ?", en.UserPk).Where("enterprise_pk = ?", en.EnterprisePk).First(&enterpriseUserMessage).Count(&enterpriseUserCount)
		if enterpriseUserSql.Error == nil || enterpriseUserCount > 0 {
			if enterpriseUserMessage.Status == 2 {
				return errors.New("该人员已离职")
			}
			var memberCount int64
			memberCountSql := o.db.Model(&entity.Member{}).Where("user_pk = ?", enterpriseUserMessage.Pk).Where("enterprise_pk = ?", en.EnterprisePk).Where("department_pk = ? ", departmentMessage.Pk).Count(&memberCount)
			if memberCountSql.Error != nil || memberCount <= 0 {
				if memberSql := o.db.Create(&entity.Member{
					Pk:           helper.GetRid(helper.Member),
					UserPk:       enterpriseUserMessage.Pk,
					EnterprisePk: en.EnterprisePk,
					DepartmentPk: departmentMessage.Pk,
				}); memberSql.Error != nil {
					return memberSql.Error
				}
			}
		} else {
			enterpriseUserPk := helper.GetRid(helper.EnterpriseUser)
			if enterpriseUserCreateSql := o.db.Create(&entity.EnterpriseUser{
				Pk:           enterpriseUserPk,
				EnterprisePk: en.EnterprisePk,
				UserPk:       en.UserPk,
				Name:         userMessage.Name,
				Phone:        userMessage.Phone,
				Gender:       userMessage.Gender,
				Birthday:     userMessage.Birthday,
				Height:       userMessage.Height,
				Weight:       userMessage.Weight,
				Status:       1,
			}); enterpriseUserCreateSql.Error != nil {
				return enterpriseUserCreateSql.Error
			}
			if memberSql := o.db.Create(&entity.Member{
				Pk:           helper.GetRid(helper.Member),
				UserPk:       enterpriseUserPk,
				EnterprisePk: en.EnterprisePk,
				DepartmentPk: departmentMessage.Pk,
			}); memberSql.Error != nil {
				return memberSql.Error
			}
		}
		var roleMessage *entity.Role
		if roleSql := o.db.Model(&entity.Role{}).Where("role_name = ? ", "超级管理员").Where("enterprise_pk = ?", en.EnterprisePk).Order("create_at desc").First(&roleMessage); roleSql.Error != nil {
			return roleSql.Error
		}
		var userPermissionCount int64
		userPermissionCountSql := o.db.Model(&entity.UserPermission{}).Where("user_pk = ?", en.UserPk).Where("enterprise_pk = ?", en.EnterprisePk).Where("permission_pk = ? ", roleMessage.Pk).Count(&userPermissionCount)
		if userPermissionCountSql.Error != nil || userPermissionCount <= 0 {
			if userPermissionSql := o.db.Create(&entity.UserPermission{
				Pk:             helper.GetRid(helper.UserPermission),
				UserPk:         en.UserPk,
				EnterprisePk:   en.EnterprisePk,
				PermissionPk:   roleMessage.Pk,
				PermissionType: 2,
			}); userPermissionSql.Error != nil {
				return userPermissionSql.Error
			}
		}
		return nil
	})
	return err
}

func (o *userPermission) Delete(uu *entity.UserPermission) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *userPermission) DeleteByUserPkAndEnterprisePk(uu *entity.UserPermission) error {
	tx := o.db.Where("enterprise_pk= ? and user_pk = ? ", uu.EnterprisePk, uu.UserPk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *userPermission) Select(en *entity.UserPermission, pg *Pagination) ([]*entity.UserPermission, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PermissionType != 0 {
		sql = sql.Where("permission_type  =?", en.PermissionType)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk=?", en.PermissionPk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	var rows []*entity.UserPermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *userPermission) Update(log *entity.UserPermission) (*entity.UserPermission, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *userPermission) SelectAllUserPermission(en *entity.UserPermission) ([]*entity.UserPermission, error) {
	sql := o.db.Model(&entity.UserPermission{})
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.PermissionType != 0 {
		sql = sql.Where("permission_type  =?", en.PermissionType)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk=?", en.PermissionPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	var rows []*entity.UserPermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *userPermission) FindList(p *entity.UserPermission) ([]*entity.UserPermission, error) {
	sql := o.db.Model(&entity.UserPermission{})
	if p.UserPk != 0 {
		sql.Where("user_pk = ?", p.UserPk)
	}
	var rows []*entity.UserPermission
	tx := sql.Find(&rows)
	return rows, tx.Error
}
func (o *userPermission) FindByPk(en *entity.UserPermission) (*entity.UserPermission, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}

// CheckUserEnterprisePlatformPermission 通过userPK和EnterprisePk判断是否有后台登录权限
func (o *userPermission) CheckUserEnterprisePlatformPermission(userPk, enterprisePk int64) (bool, error) {
	permissionMessage := &entity.Permission{}
	if enterprisePk == 1 {
		tx := o.db.Model(permissionMessage).Where("pk = ?", 2999834460028928).First(&permissionMessage)
		if tx.Error != nil || permissionMessage.Pk == 0 {
			return false, errors.New("系统错误！！")
		}
	} else {
		enterpriseMessage := &entity.Enterprise{}
		enterpriseSql := o.db.Model(&entity.Enterprise{}).Where("pk = ?", enterprisePk).First(&enterpriseMessage)
		if enterpriseSql.Error != nil || enterpriseMessage.Type == 0 {
			return false, errors.New("系统错误！！")
		}
		if enterpriseMessage.Type == 2 {
			tx := o.db.Model(permissionMessage).Where("pk = ?", 2999835802206208).First(&permissionMessage)
			if tx.Error != nil || permissionMessage.Pk == 0 {
				return false, errors.New("系统错误！！")
			}
		} else if enterpriseMessage.Type == 3 {
			tx := o.db.Model(permissionMessage).Where("pk = ?", 3014166634823680).First(&permissionMessage)
			if tx.Error != nil || permissionMessage.Pk == 0 {
				return false, errors.New("系统错误！！")
			}
		} else if enterpriseMessage.Type == 1 {
			tx := o.db.Model(permissionMessage).Where("pk = ?", 3014167775674368).First(&permissionMessage)
			if tx.Error != nil || permissionMessage.Pk == 0 {
				return false, errors.New("系统错误！！")
			}
		}
	}
	sql := o.db.Model(&entity.UserPermission{}).
		Where("enterprise_pk = ? ", enterprisePk).
		Where("user_pk = ?", userPk)
	var rows []*entity.UserPermission
	tx := sql.Find(&rows)
	if tx.Error != nil || len(rows) <= 0 {
		return false, tx.Error
	}
	for _, z := range rows {
		if z.PermissionType == 1 {
			// 人员绑定的角色组
			sql1 := o.db.Model(&entity.RoleGroupPermission{}).
				Where("role_group_pk = ? ", z.PermissionPk)
			var roleGroupPermissionRows []*entity.RoleGroupPermission
			tx1 := sql1.Find(&roleGroupPermissionRows)
			if tx1.Error == nil && len(roleGroupPermissionRows) > 0 {
				for _, p := range roleGroupPermissionRows {
					if p.PermissionType == 1 {
						//角色
						sql3 := o.db.Model(&entity.RolePermission{}).
							Where("role_pk = ? ", p.PermissionPk)
						var rolePermissionRows []*entity.RolePermission
						tx3 := sql3.Find(&rolePermissionRows)
						if tx3.Error == nil && len(rolePermissionRows) > 0 {
							for _, r := range rolePermissionRows {
								if r.PermissionType == 1 {
									//权限
									if r.PermissionPk == permissionMessage.Pk {
										return true, nil
									}
								} else if r.PermissionType == 2 {
									//权限组
									sql2 := o.db.Model(&entity.PermissionGroupPermission{}).
										Where("permission_group_pk = ? ", r.PermissionPk).
										Where("permission_pk = ? ", permissionMessage.Pk)
									var count int64
									tx2 := sql2.Count(&count)
									if tx2.Error == nil && int32(count) == 1 {
										return true, nil
									}
								}
							}
						}
					} else if p.PermissionType == 2 {
						//权限
						if p.PermissionPk == permissionMessage.Pk {
							return true, nil
						}
					} else if p.PermissionType == 3 {
						//权限组
						sql2 := o.db.Model(&entity.PermissionGroupPermission{}).
							Where("permission_group_pk = ? ", p.PermissionPk).
							Where("permission_pk = ? ", permissionMessage.Pk)
						var count int64
						tx2 := sql2.Count(&count)
						if tx2.Error == nil && int32(count) == 1 {
							return true, nil
						}
					}
				}
			}
		} else if z.PermissionType == 2 {
			// 人员绑定的角色
			sql1 := o.db.Model(&entity.RolePermission{}).
				Where("role_pk = ? ", z.PermissionPk)
			var rolePermissionRows []*entity.RolePermission
			tx1 := sql1.Find(&rolePermissionRows)
			if tx1.Error == nil && len(rolePermissionRows) > 0 {
				for _, p := range rolePermissionRows {
					if p.PermissionType == 1 {
						//权限
						if p.PermissionPk == permissionMessage.Pk {
							return true, nil
						}
					} else if p.PermissionType == 2 {
						//权限组
						sql2 := o.db.Model(&entity.PermissionGroupPermission{}).
							Where("permission_group_pk = ? ", p.PermissionPk).
							Where("permission_pk = ? ", permissionMessage.Pk)
						var count int64
						tx2 := sql2.Count(&count)
						if tx2.Error == nil && int32(count) == 1 {
							return true, nil
						}
					}
				}
			}
		}
	}
	return false, tx.Error
}

// CheckUserEnterpriseAppletPermission 通过userPK和EnterprisePk判断是否有小程序登录权限
func (o *userPermission) CheckUserEnterpriseAppletPermission(userPk, enterprisePk int64) (bool, error) {
	permissionMessage := &entity.Permission{}
	if enterprisePk == 1 {
		tx := o.db.Model(permissionMessage).Where("pk = ?", 3014172406185984).First(&permissionMessage)
		if tx.Error != nil || permissionMessage.Pk == 0 {
			return false, errors.New("系统错误！！")
		}
	} else {
		enterpriseMessage := &entity.Enterprise{}
		enterpriseSql := o.db.Model(&entity.Enterprise{}).Where("pk = ?", enterprisePk).First(&enterpriseMessage)
		if enterpriseSql.Error != nil || enterpriseMessage.Type == 0 {
			return false, errors.New("系统错误！！")
		}
		if enterpriseMessage.Type == 2 {
			tx := o.db.Model(permissionMessage).Where("pk = ?", 2999842278211584).First(&permissionMessage)
			if tx.Error != nil || permissionMessage.Pk == 0 {
				return false, errors.New("系统错误！！")
			}
		} else if enterpriseMessage.Type == 3 {
			tx := o.db.Model(permissionMessage).Where("pk = ?", 3014171600879616).First(&permissionMessage)
			if tx.Error != nil || permissionMessage.Pk == 0 {
				return false, errors.New("系统错误！！")
			}
		} else if enterpriseMessage.Type == 1 {
			tx := o.db.Model(permissionMessage).Where("pk = ?", 3014172875948032).First(&permissionMessage)
			if tx.Error != nil || permissionMessage.Pk == 0 {
				return false, errors.New("系统错误！！")
			}
		}
	}
	sql := o.db.Model(&entity.UserPermission{}).
		Where("enterprise_pk = ? ", enterprisePk).
		Where("user_pk = ?", userPk)
	var rows []*entity.UserPermission
	tx := sql.Find(&rows)
	if tx.Error != nil || len(rows) <= 0 {
		return false, tx.Error
	}
	for _, z := range rows {
		if z.PermissionType == 1 {
			// 人员绑定的角色组
			sql1 := o.db.Model(&entity.RoleGroupPermission{}).
				Where("role_group_pk = ? ", z.PermissionPk)
			var roleGroupPermissionRows []*entity.RoleGroupPermission
			tx1 := sql1.Find(&roleGroupPermissionRows)
			if tx1.Error == nil && len(roleGroupPermissionRows) > 0 {
				for _, p := range roleGroupPermissionRows {
					if p.PermissionType == 1 {
						//角色
						sql3 := o.db.Model(&entity.RolePermission{}).
							Where("role_pk = ? ", p.PermissionPk)
						var rolePermissionRows []*entity.RolePermission
						tx3 := sql3.Find(&rolePermissionRows)
						if tx3.Error == nil && len(rolePermissionRows) > 0 {
							for _, r := range rolePermissionRows {
								if r.PermissionType == 1 {
									//权限
									if r.PermissionPk == permissionMessage.Pk {
										return true, nil
									}
								} else if r.PermissionType == 2 {
									//权限组
									sql2 := o.db.Model(&entity.PermissionGroupPermission{}).
										Where("permission_group_pk = ? ", r.PermissionPk).
										Where("permission_pk = ? ", permissionMessage.Pk)
									var count int64
									tx2 := sql2.Count(&count)
									if tx2.Error == nil && int32(count) == 1 {
										return true, nil
									}
								}
							}
						}
					} else if p.PermissionType == 2 {
						//权限
						if p.PermissionPk == permissionMessage.Pk {
							return true, nil
						}
					} else if p.PermissionType == 3 {
						//权限组
						sql2 := o.db.Model(&entity.PermissionGroupPermission{}).
							Where("permission_group_pk = ? ", p.PermissionPk).
							Where("permission_pk = ? ", permissionMessage.Pk)
						var count int64
						tx2 := sql2.Count(&count)
						if tx2.Error == nil && int32(count) == 1 {
							return true, nil
						}
					}
				}
			}
		} else if z.PermissionType == 2 {
			// 人员绑定的角色
			sql1 := o.db.Model(&entity.RolePermission{}).
				Where("role_pk = ? ", z.PermissionPk)
			var rolePermissionRows []*entity.RolePermission
			tx1 := sql1.Find(&rolePermissionRows)
			if tx1.Error == nil && len(rolePermissionRows) > 0 {
				for _, p := range rolePermissionRows {
					if p.PermissionType == 1 {
						//权限
						if p.PermissionPk == permissionMessage.Pk {
							return true, nil
						}
					} else if p.PermissionType == 2 {
						//权限组
						sql2 := o.db.Model(&entity.PermissionGroupPermission{}).
							Where("permission_group_pk = ? ", p.PermissionPk).
							Where("permission_pk = ? ", permissionMessage.Pk)
						var count int64
						tx2 := sql2.Count(&count)
						if tx2.Error == nil && int32(count) == 1 {
							return true, nil
						}
					}
				}
			}
		}
	}
	return false, tx.Error
}
func (o *userPermission) Count(p *entity.UserPermission) (int32, error) {
	sql := o.db.Model(&entity.UserPermission{})
	if p.Pk != 0 {
		sql.Where("pk = ?", p.Pk)
	}
	if p.UserPk != 0 {
		sql = sql.Where("user_pk = ?", p.UserPk)
	}
	if p.PermissionType != 0 {
		sql = sql.Where("permission_type  =?", p.PermissionType)
	}
	if p.PermissionPk != 0 {
		sql = sql.Where("permission_pk=?", p.PermissionPk)
	}
	if p.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", p.EnterprisePk)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

type ErrorInfo struct {
	UserID uint
	RoleID uint
	Error  error
}

func (o *userPermission) SelectWithUser(en *entity.UserPermission, pg *Pagination) ([]*model.UserPermissionModel, error) {
	sql := o.db.Model(&model.UserPermission{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	sql.Preload("User", func(sql *gorm.DB) *gorm.DB {
		err := sql.Error
		if err != nil {
			fmt.Println("Error occurred during preload:", err)
		}
		return sql
	})
	if en.PermissionType != 0 {
		sql = sql.Where("user_permission.permission_type  =?", en.PermissionType)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("user_permission.permission_pk = ?", en.PermissionPk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_permission.user_pk = ?", en.UserPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("user_permission.enterprise_pk = ?", en.EnterprisePk)
	}
	var rows []*model.UserPermissionModel
	tx := sql.Order("user_permission.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *userPermission) SelectWithEnterpriseUser(en *entity.UserPermission, pg *Pagination) ([]*model.UserPermissionForEnterprise, error) {
	sql := o.db.Model(&model.UserPermissionForEnterprise{}).Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.PermissionType != 0 {
		sql = sql.Where("permission_type  =?", en.PermissionType)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk = ?", en.PermissionPk)
	}
	if en.UserPk != 0 {
		sql = sql.Where("user_pk = ?", en.UserPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("enterprise_pk = ?", en.EnterprisePk)
	}
	var rows []*model.UserPermissionForEnterprise
	tx := sql.Order("create_at desc").Find(&rows)
	if tx.Error == nil && rows != nil && len(rows) > 0 {
		for _, v := range rows {
			if v.UserPk != 0 && v.EnterprisePk != 0 {
				var enterpriseUserMessage *entity.EnterpriseUser
				userSql := o.db.Model(&entity.EnterpriseUser{}).Where("user_pk = ?", v.UserPk).Where("enterprise_pk = ?", v.EnterprisePk).First(&enterpriseUserMessage)
				if userSql.Error == nil && enterpriseUserMessage != nil && enterpriseUserMessage.Pk != 0 {
					v.EnterpriseUser = enterpriseUserMessage
				}
			}
		}
	}
	return rows, tx.Error
}
func (o *userPermission) DeleteRedisPermission(enterprisePk, userPk int64) error {
	redisName := "permissionKey" + "_" + strconv.FormatInt(enterprisePk, 10)
	if enterprisePk != 0 && userPk != 0 {
		redisKey := strconv.FormatInt(userPk, 10)
		_, err := o.rs.HDel(redisName, redisKey)
		return err
	} else if enterprisePk != 0 && userPk == 0 {
		_, err := o.rs.Del(redisName)
		return err
	} else {
		return errors.New("参数不合法！！")
	}
}
