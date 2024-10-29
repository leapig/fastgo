package dao

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
	"strconv"
)

type Member interface {
	FindLeader(entity.Department) (*model.Member, error)
	Create(*entity.Member) (*entity.Member, error)
	Delete(*entity.Member) error
	Update(*entity.Member) error
	List(*entity.Member) ([]*entity.Member, error)
	CreateMemberUser(member *entity.Member, user *entity.User, userCredentials []*entity.UserCredentials) (*entity.Member, error)
	SelectMemberDetail(m *model.Member, pg *entity.Pagination) ([]*model.Member, error)
	Count(en *model.Member) (int32, error)
	DeleteByEnterprisePkAndUserPk(m *entity.Member) error
	UpdateMemberUser(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error)
	UpdateUserDepartmentRelation(en *entity.Member, departmentPks []int64) error
}

type member struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewMember(db *gorm.DB, rs *helper.Redis) Member {
	return &member{
		db: db,
		rs: rs,
	}
}
func (o *member) CreateMemberUser(member *entity.Member, user *entity.User, userCredentials []*entity.UserCredentials) (*entity.Member, error) {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		var userCount int64
		userCountSql := o.db.Model(&entity.User{}).Where("phone = ?", user.Phone).Count(&userCount)
		if userCountSql.Error == nil && userCount >= 1 {
			return errors.New("手机号已存在！")
		}
		userPk := helper.GetRid(helper.User)
		if userSql := o.db.Create(&entity.User{
			Pk:    userPk,
			Phone: user.Phone,
			Name:  user.Name,
		}); userSql.Error != nil {
			return userSql.Error
		}
		member.Pk = helper.GetRid(helper.Member)
		if memberSql := o.db.Create(&entity.Member{
			Pk:           member.Pk,
			EnterprisePk: member.EnterprisePk,
			DepartmentPk: member.DepartmentPk,
			UserPk:       userPk,
			IsLeader:     member.IsLeader,
			IsMain:       member.IsMain,
			JobTitle:     member.JobTitle,
			JobNumber:    member.JobNumber,
		}); memberSql.Error != nil {
			return memberSql.Error
		}
		if len(userCredentials) > 0 {
			for _, v := range userCredentials {
				v.Pk = helper.GetRid(helper.UserCredentials)
				v.UserPk = userPk
				if rs := tx.Create(&v); rs.Error != nil {
					return rs.Error
				}
			}
		}
		return nil
	})
	return member, err
}

func (o *member) Create(m *entity.Member) (*entity.Member, error) {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		var memberCount int64
		memberCountSql := o.db.Model(&entity.Member{}).
			Where("department_pk = ?", m.DepartmentPk).
			Where("enterprise_pk = ?", m.EnterprisePk).
			Where("user_pk = ?", m.UserPk).Count(&memberCount)
		if memberCountSql.Error == nil && memberCount >= 1 {
			return errors.New("人员已存在部门！")
		}
		m.Pk = helper.GetRid(helper.Member)
		if rs := tx.Create(&m); rs.Error != nil {
			return rs.Error
		}
		return nil
	})
	return m, err
}

func (o *member) Update(m *entity.Member) error {
	tx := o.db.Where("enterprise_pk=? and department_pk=? and user_pk=?", m.EnterprisePk, m.DepartmentPk, m.UserPk).Updates(&m)
	return tx.Error
}

func (o *member) List(m *entity.Member) ([]*entity.Member, error) {
	sql := o.db.Model(&entity.Member{})
	if m.UserPk != 0 {
		sql = sql.Where("user_pk = ?", m.UserPk)
	}
	rows := make([]*entity.Member, 0)
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *member) Delete(m *entity.Member) error {
	tx := o.db.Where("pk=?", m.Pk).Delete(&m)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}
func (o *member) DeleteByEnterprisePkAndUserPk(m *entity.Member) error {
	tx := o.db.Where("user_pk =?", m.UserPk).Where("enterprise_pk =?", m.EnterprisePk).Delete(&m)
	if tx.RowsAffected == 0 {
		return errors.New("delete fatal")
	}
	return tx.Error
}

func (o *member) SelectMemberDetail(m *model.Member, pg *entity.Pagination) ([]*model.Member, error) {
	sql := o.db.Limit(int(pg.Size)).Offset(int((pg.Page - 1) * pg.Size)).Preload("EnterpriseUser").Joins("left join enterprise_user on enterprise_user.pk=member.user_pk")
	if m.EnterprisePk != 0 {
		sql = sql.Where("member.enterprise_pk = ?", m.EnterprisePk)
	}
	if m.DepartmentPk != 0 {
		sql = sql.Where("member.department_pk = ?", m.DepartmentPk)
	}
	if m.IsLeader != 0 {
		sql = sql.Where("member.is_leader = ?", m.IsLeader)
	}
	if m.JobTitle != "" {
		sql = sql.Where("member.job_title = ?", m.JobTitle)
	}
	if m.JobNumber != 0 {
		sql = sql.Where("member.job_number = ?", m.JobNumber)
	}
	if m.EnterpriseUser.Name != "" {
		sql = sql.Where("enterprise_user.name like ?", "%"+m.EnterpriseUser.Name+"%")
	}
	// 大于4位的纯数字
	if _, err := strconv.Atoi(m.EnterpriseUser.Phone); err == nil && len(m.EnterpriseUser.Phone) > 4 {
		sql.Where("enterprise_user.phone like ?", "%"+m.EnterpriseUser.Phone+"%")
	}
	rows := make([]*model.Member, 0)
	tx := sql.Find(&rows)
	return rows, tx.Error
}
func (o *member) Count(en *model.Member) (int32, error) {
	sql := o.db.Model(&model.Member{}).Preload("EnterpriseUser").Joins("left join enterprise_user on enterprise_user.pk=member.user_pk")
	if en.UserPk != 0 {
		sql = sql.Where("member.user_pk = ?", en.UserPk)
	}
	if en.EnterprisePk != 0 {
		sql = sql.Where("member.enterprise_pk = ?", en.EnterprisePk)
	}
	if en.DepartmentPk != 0 {
		sql = sql.Where("member.department_pk = ?", en.DepartmentPk)
	}
	if en.IsLeader != 0 {
		sql = sql.Where("member.is_leader = ?", en.IsLeader)
	}
	if en.JobTitle != "" {
		sql = sql.Where("member.job_title = ?", en.JobTitle)
	}
	if en.JobNumber != 0 {
		sql = sql.Where("member.job_number = ?", en.JobNumber)
	}
	if en.EnterpriseUser != nil {
		if en.EnterpriseUser.Name != "" {
			sql = sql.Where("enterprise_user.name like ?", "%"+en.EnterpriseUser.Name+"%")
		}
		// 大于4位的纯数字
		if _, err := strconv.Atoi(en.EnterpriseUser.Phone); err == nil && len(en.EnterpriseUser.Phone) > 4 {
			sql.Where("enterprise_user.phone like ?", "%"+en.EnterpriseUser.Name+"%")
		}
	}
	var count int64
	tx := sql.Count(&count)
	return int32(count), tx.Error
}
func (o *member) UpdateMemberUser(en *entity.EnterpriseUser) (*entity.EnterpriseUser, error) {
	tx := o.db.Where("pk = ?", en.Pk).Updates(&en)
	if tx.RowsAffected == 0 {
		return en, errors.New("update fatal")
	}
	return en, tx.Error
}
func (o *member) UpdateUserDepartmentRelation(en *entity.Member, departmentPks []int64) error {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		members := make([]*entity.Member, 0)
		memberSql := o.db.Model(&entity.Member{}).
			Where("enterprise_pk = ?", en.EnterprisePk).
			Where("user_pk = ?", en.UserPk).
			Find(&members)
		if memberSql.Error == nil && len(members) > 0 {
			for _, v := range members {
				if !contains(departmentPks, v.DepartmentPk) {
					var m *entity.Member
					m.Pk = v.Pk
					deleteSql := o.db.Where("pk=?", m.Pk).Delete(&m)
					if deleteSql.RowsAffected == 0 {
						return errors.New("delete fatal")
					}
				}
			}
		}
		for _, departmentPk := range departmentPks {
			var memberCount int64
			memberCountSql := o.db.Model(&entity.Member{}).
				Where("enterprise_pk = ?", en.EnterprisePk).
				Where("user_pk = ?", en.UserPk).
				Where("department_pk = ?", departmentPk).
				Count(&memberCount)
			if memberCountSql.Error == nil && memberCount >= 1 {
				continue
			}
			if memberSql := o.db.Create(&entity.Member{
				Pk:           helper.GetRid(helper.Member),
				EnterprisePk: en.EnterprisePk,
				DepartmentPk: departmentPk,
				UserPk:       en.UserPk,
			}); memberSql.Error != nil {
				return memberSql.Error
			}
		}
		return nil
	})
	return err
}
func contains(slice []int64, value int64) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// FindLeader 查找部门领导
func (o *member) FindLeader(d entity.Department) (*model.Member, error) {
	var mem *model.Member
	sql := o.db.Model(&model.Member{}).Preload("EnterpriseUser")
	tx := sql.Where("enterprise_pk=? and department_pk=? and is_leader=?", d.EnterprisePk, d.Pk, 1).First(&mem)
	return mem, tx.Error
}
