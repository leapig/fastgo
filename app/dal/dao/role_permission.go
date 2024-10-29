package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type RolePermission interface {
	Create(*entity.RolePermission) (*entity.RolePermission, error)
	Delete(*entity.RolePermission) error
	DeleteByRolePk(uu *entity.RolePermission) error
	Select(*entity.RolePermission, *Pagination) ([]*entity.RolePermission, error)
	Count(*entity.RolePermission) (int32, error)
	Update(*entity.RolePermission) (*entity.RolePermission, error)
	SelectAllRolePermission(en *entity.RolePermission) ([]*entity.RolePermission, error)
	SelectRolePermissionByRolePk(en *entity.RolePermission) ([]*model.RolePermissionModel, error)
	SelectRolePermissionByRolePkTest(en *entity.RolePermission) ([]*model.RolePermissionModel, error)
	SelectPermissionByRolePK(rolePk int64) ([]*entity.Permission, error)
	FindByPk(en *entity.RolePermission) (*entity.RolePermission, error)
}
type rolePermission struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewRolePermission(db *gorm.DB, rs *helper.Redis) RolePermission {
	return &rolePermission{
		db: db,
		rs: rs,
	}
}
func (o *rolePermission) Create(en *entity.RolePermission) (*entity.RolePermission, error) {
	tx := o.db.Create(&en)
	return en, tx.Error
}

func (o *rolePermission) Delete(uu *entity.RolePermission) error {
	tx := o.db.Where("pk=?", uu.Pk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *rolePermission) DeleteByRolePk(uu *entity.RolePermission) error {
	tx := o.db.Where("role_pk=?", uu.RolePk).Delete(&uu)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *rolePermission) Select(en *entity.RolePermission, pg *Pagination) ([]*entity.RolePermission, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size))
	if en.RolePk != 0 {
		sql = sql.Where("role_pk = ?", en.RolePk)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk = ?", en.PermissionPk)
	}
	if en.PermissionType != 0 {
		sql = sql.Where("permission_type = ?", en.PermissionType)
	}
	var rows []*entity.RolePermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *rolePermission) Count(en *entity.RolePermission) (int32, error) {
	sql := o.db.Model(&entity.RolePermission{})
	if en.RolePk != 0 {
		sql = sql.Where("role_pk = ?", en.RolePk)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk = ?", en.PermissionPk)
	}
	if en.PermissionType != 0 {
		sql = sql.Where("permission_type = ?", en.PermissionType)
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}

func (o *rolePermission) Update(log *entity.RolePermission) (*entity.RolePermission, error) {
	tx := o.db.Where("pk = ?", log.Pk).Updates(&log)
	if tx.RowsAffected == 0 {
		return log, errors.New("update fatal")
	}
	return log, tx.Error
}

func (o *rolePermission) SelectAllRolePermission(en *entity.RolePermission) ([]*entity.RolePermission, error) {
	sql := o.db.Model(&entity.RolePermission{})
	if en.RolePk != 0 {
		sql = sql.Where("role_pk = ?", en.RolePk)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("permission_pk = ?", en.PermissionPk)
	}
	if en.PermissionType != 0 {
		sql = sql.Where("permission_type = ?", en.PermissionType)
	}
	var rows []*entity.RolePermission
	tx := sql.Order("create_at desc").Find(&rows)
	return rows, tx.Error
}

func (o *rolePermission) SelectRolePermissionByRolePk(en *entity.RolePermission) ([]*model.RolePermissionModel, error) {
	sql := o.db.Model(&model.RolePermissionModel{})
	if en.RolePk != 0 {
		sql = sql.Where("role_permission.role_pk = ?", en.RolePk)
	}
	var rows []*model.RolePermissionModel
	tx := sql.Order("role_permission.create_at desc").Find(&rows)
	return rows, tx.Error
}
func (o *rolePermission) SelectRolePermissionByRolePkTest(en *entity.RolePermission) ([]*model.RolePermissionModel, error) {
	sql := o.db.Model(&model.RolePermissionModel{}).Preload("Permission").Preload("PermissionGroup")
	if en.RolePk != 0 {
		sql = sql.Where("role_permission.role_pk = ?", en.RolePk)
	}
	if en.PermissionPk != 0 {
		sql = sql.Where("role_permission.permission_pk = ?", en.PermissionPk)
	}
	if en.PermissionType != 0 {
		sql = sql.Where("role_permission.permission_type = ?", en.PermissionType)
	}
	var rows []*model.RolePermissionModel
	tx := sql.Order("role_permission.create_at desc").Find(&rows)
	return rows, tx.Error
}

// SelectPermissionByRolePK 通过角色PK获取单一权限项列表
func (o *rolePermission) SelectPermissionByRolePK(rolePk int64) ([]*entity.Permission, error) {
	var rolePermissions []*model.RolePermission
	tx := o.db.Model(&model.RolePermission{}).Where("role_pk = ?", rolePk).Order("create_at desc").Find(&rolePermissions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	associationPermissions := func(permissionGroupPermissionList []*model.PermissionGroupPermissionModel) error {
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
	for _, rp := range rolePermissions {
		if rp.PermissionType == 1 {
			err := o.db.Model(rp).Association("Permission").Find(&rp.Permission)
			if err != nil {
				return nil, err
			}
		} else if rp.PermissionType == 2 {
			var permissionGroupPermissions []*model.PermissionGroupPermissionModel
			err := o.db.Model(rp).Association("PermissionGroupPermission").Find(&permissionGroupPermissions)
			if err != nil {
				return nil, err
			}
			err = associationPermissions(permissionGroupPermissions)
			if err != nil {
				return nil, err
			}
			rp.PermissionGroupPermission = permissionGroupPermissions
		}
	}
	result := make([]*entity.Permission, 0)
	if rolePermissions != nil && len(rolePermissions) > 0 {
		for _, v := range rolePermissions {
			switch v.PermissionType {
			case 1:
				if v.Permission != nil {
					result = append(result, v.Permission)
				}
			case 2:
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
func (o *rolePermission) FindByPk(en *entity.RolePermission) (*entity.RolePermission, error) {
	tx := o.db.Where("pk = ?", en.Pk).Find(&en)
	return en, tx.Error
}
