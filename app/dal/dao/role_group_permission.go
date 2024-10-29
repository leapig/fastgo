package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type RoleGroupPermission interface {
	Create(*entity.RoleGroupPermission) (*entity.RoleGroupPermission, error)
	Delete(*entity.RoleGroupPermission) error
	DeleteByGroupPk(uu *entity.RoleGroupPermission) error
	Select(*entity.RoleGroupPermission, *Pagination) ([]*entity.RoleGroupPermission, error)
	SelectAllRoleGroupPermission(en *entity.RoleGroupPermission) ([]*entity.RoleGroupPermission, error)
	Count(*entity.RoleGroupPermission) (int32, error)
	Update(*entity.RoleGroupPermission) (*entity.RoleGroupPermission, error)
	FindByPk(*entity.RoleGroupPermission) (*entity.RoleGroupPermission, error)
	SelectRoleGroupPermissionByRoleGroupPk(en *entity.RoleGroupPermission) ([]*model.RoleGroupPermissionModel, error)
	SelectPermissionByRoleGroupPK(roleGroupPk int64) ([]*entity.Permission, error)
}
type roleGroupPermission struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewRoleGroupPermission(db *gorm.DB, rs *helper.Redis) RoleGroupPermission {
	return &roleGroupPermission{
		db: db,
		rs: rs,
	}
}
func (o *roleGroupPermission) Create(en *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *roleGroupPermission) Delete(uu *entity.RoleGroupPermission) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *roleGroupPermission) DeleteByGroupPk(uu *entity.RoleGroupPermission) error {
	tx := o.db.Where("role_group_pk=?", uu.RoleGroupPk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *roleGroupPermission) Select(en *entity.RoleGroupPermission, pg *Pagination) ([]*entity.RoleGroupPermission, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))

	var rows []*entity.RoleGroupPermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *roleGroupPermission) Count(en *entity.RoleGroupPermission) (int32, error) {
	sql := o.db.Model(&entity.RoleGroupPermission{})
	if en.RoleGroupPk != 0 {
		sql.Where("role_group_pk=?", en.RoleGroupPk)
	}
	if en.PermissionPk != 0 {
		sql.Where("permission_pk=?", en.PermissionPk)
	}
	if en.PermissionType != 0 {
		sql.Where("permission_type=?", en.PermissionType)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *roleGroupPermission) Update(log *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}
func (o *roleGroupPermission) FindByPk(en *entity.RoleGroupPermission) (*entity.RoleGroupPermission, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
func (o *roleGroupPermission) SelectRoleGroupPermissionByRoleGroupPk(en *entity.RoleGroupPermission) ([]*model.RoleGroupPermissionModel, error) {
	sql := o.db.Model(&model.RoleGroupPermissionModel{})
	sql.Preload("Role").Preload("Permission").Preload("PermissionGroup")
	if en.RoleGroupPk != 0 {
		sql = sql.Where("role_group_permission.role_group_pk = ?", en.RoleGroupPk)
	}
	var rows []*model.RoleGroupPermissionModel
	tx := sql.Order("role_group_permission.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *roleGroupPermission) SelectAllRoleGroupPermission(en *entity.RoleGroupPermission) ([]*entity.RoleGroupPermission, error) {
	sql := o.db.Model(&entity.RoleGroupPermission{})
	if en.RoleGroupPk != 0 {
		sql = sql.Where("role_group_pk = ?", en.RoleGroupPk)
	}
	var rows []*entity.RoleGroupPermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

// SelectPermissionByRoleGroupPK 通过角色组PK获取单一权限项列表
func (o *roleGroupPermission) SelectPermissionByRoleGroupPK(roleGroupPk int64) ([]*entity.Permission, error) {
	var roleGroupPermissions []*model.RoleGroupPermission
	tx := o.db.Model(&model.RoleGroupPermission{}).Where("role_group_pk = ?", roleGroupPk).Order("create_at desc").Find(&roleGroupPermissions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	associationPermissionGroupPermissions := func(permissionGroupPermissionList []*model.PermissionGroupPermissionModel) error {
		for _, p := range permissionGroupPermissionList {
			if p.Permission == nil {
				err := o.db.Model(p).Association("Permission").Find(&p.Permission)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	associationRolePermissions := func(rolePermissionList []*model.RolePermission) error {
		for _, p := range rolePermissionList {
			if p.PermissionType == 1 && p.Permission == nil {
				err := o.db.Model(p).Association("Permission").Find(&p.Permission)
				if err != nil {
					return err
				}
			} else if p.PermissionType == 2 {
				err := associationPermissionGroupPermissions(p.PermissionGroupPermission)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	for _, rp := range roleGroupPermissions {
		var rolePermissions []*model.RolePermission
		if rp.PermissionType == 1 {
			err := o.db.Model(rp).Association("RolePermission").Find(&rolePermissions)
			if err != nil {
				return nil, err
			}
			err = associationRolePermissions(rolePermissions)
			if err != nil {
				return nil, err
			}
			rp.RolePermission = rolePermissions
		} else if rp.PermissionType == 2 {
			err := o.db.Model(rp).Association("Permission").Find(&rp.Permission)
			if err != nil {
				return nil, err
			}
		} else if rp.PermissionType == 3 {
			var permissionGroupPermissions []*model.PermissionGroupPermissionModel
			err := o.db.Model(rp).Association("PermissionGroupPermission").Find(&permissionGroupPermissions)
			if err != nil {
				return nil, err
			}
			err = associationPermissionGroupPermissions(permissionGroupPermissions)
			if err != nil {
				return nil, err
			}
			rp.PermissionGroupPermission = permissionGroupPermissions
		}
	}
	result := make([]*entity.Permission, 0)
	if roleGroupPermissions != nil && len(roleGroupPermissions) > 0 {
		for _, v := range roleGroupPermissions {
			switch v.PermissionType {
			case 1:
				if v.RolePermission != nil && len(v.RolePermission) > 0 {
					for _, z := range v.RolePermission {
						switch z.PermissionType {
						case 1:
							if z.Permission != nil {
								result = append(result, z.Permission)
							}
						case 2:
							if z.PermissionGroupPermission != nil && len(z.PermissionGroupPermission) > 0 {
								for _, p := range z.PermissionGroupPermission {
									if p.Permission != nil {
										result = append(result, p.Permission)
									}
								}
							}

						default:
							continue
						}
					}
				}
			case 2:
				if v.Permission != nil {
					result = append(result, v.Permission)
				}
			case 3:
				if v.PermissionGroupPermission != nil && len(v.PermissionGroupPermission) > 0 {
					for _, z := range v.PermissionGroupPermission {
						if z.Permission != nil {
							result = append(result, z.Permission)
						}
					}
				}
			default:
				continue
			}
		}
	}
	return result, nil
}
