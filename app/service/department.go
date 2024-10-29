package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
)

type Department interface {
	FindAllDepartByEnterprisePk(en *entity.Department) ([]*pb.DepartmentList, error)
	Create(*entity.Department) (*entity.Department, error)
	Update(*entity.Department, string) (*entity.Department, error)
	Delete(*entity.Department) error
	FindByPk(en *entity.Department) (*entity.Department, error)
	Select(*entity.Department, *dao.Pagination) ([]*entity.Department, *dao.Pagination, error)
	FindByEnterprisePkAndName(en *entity.Department) (*entity.Department, error)
}

type department struct {
	dao dao.Dao
}

func NewDepartment(dao dao.Dao) Department {
	return &department{dao: dao}
}
func (o *department) Create(in *entity.Department) (*entity.Department, error) {
	in.Pk = helper.GetRid(helper.Department)
	return o.dao.Department().Create(in)
}
func (o *department) Delete(in *entity.Department) error {
	if count, err := o.dao.Department().Count(&entity.Department{
		ParentPk: in.Pk,
	}); err == nil && count > 0 {
		return errors.New("存在子部门无法删除！")
	}
	if topDepartmentCount, err := o.dao.Department().Count(&entity.Department{
		ParentPk: in.EnterprisePk,
		Pk:       in.Pk,
	}); err == nil && topDepartmentCount > 0 {
		return errors.New("顶层部门无法删除！")
	}
	if memberCount, err := o.dao.Member().Count(&model.Member{
		DepartmentPk: in.Pk,
	}); err == nil && memberCount > 0 {
		return errors.New("部门存在人员无法删除！")
	}
	return o.dao.Department().Delete(&entity.Department{Pk: in.Pk})
}
func (o *department) Select(in *entity.Department, pg *dao.Pagination) ([]*entity.Department, *dao.Pagination, error) {
	if rows, err := o.dao.Department().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.Department().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *department) Update(in *entity.Department, leader string) (*entity.Department, error) {
	// 修改主管
	if mem, _ := o.dao.Member().FindLeader(*in); mem.UserPk != 0 && leader != "" && mem.UserPk != utils.StringToInt64(leader) {
		_ = o.dao.Member().Update(&entity.Member{
			EnterprisePk: mem.EnterprisePk,
			DepartmentPk: mem.DepartmentPk,
			UserPk:       mem.UserPk,
			IsLeader:     2})
		_ = o.dao.Member().Update(&entity.Member{
			EnterprisePk: mem.EnterprisePk,
			DepartmentPk: mem.DepartmentPk,
			UserPk:       utils.StringToInt64(leader),
			IsLeader:     1})
	}
	return o.dao.Department().Update(in)
}
func (o *department) FindAllDepartByEnterprisePk(en *entity.Department) ([]*pb.DepartmentList, error) {
	allTopDepartment, err := o.dao.Department().SelectAllDepartmentList(&entity.DepartmentList{
		EnterprisePk: en.EnterprisePk,
	})
	if err != nil || allTopDepartment == nil || len(allTopDepartment) <= 0 {
		return nil, err
	}
	result := make([]*pb.DepartmentList, 0)
	for _, v := range allTopDepartment {
		dep := &pb.DepartmentList{
			Pk:           utils.Int64ToString(v.Pk),
			EnterprisePk: utils.Int64ToString(v.EnterprisePk),
			Name:         v.Name,
			ParentPk:     utils.Int64ToString(v.ParentPk),
			CreateAt:     utils.TimeToString(v.CreatedAt),
			UpdateAt:     utils.TimeToString(v.UpdatedAt),
		}
		if mem, err := o.dao.Member().FindLeader(entity.Department{EnterprisePk: v.EnterprisePk, Pk: v.Pk}); err == nil && mem.EnterpriseUser != nil && mem.EnterpriseUser.Pk != 0 {
			dep.Leader = &pb.Member{
				Pk:       utils.Int64ToString(mem.Pk),
				Name:     mem.EnterpriseUser.Name,
				IsLeader: mem.IsLeader,
				IsMain:   mem.IsMain,
				Phone:    mem.EnterpriseUser.Phone,
				UserPk:   utils.Int64ToString(mem.EnterpriseUser.Pk),
			}
		}
		result = append(result, dep)
	}
	tree := buildTree(result, utils.Int64ToString(en.EnterprisePk))
	return tree, nil
}

func buildTree(data []*pb.DepartmentList, parentPk string) []*pb.DepartmentList {
	var tree []*pb.DepartmentList
	for _, item := range data {
		if item.ParentPk == parentPk {
			item.Rows = buildTree(data, item.Pk)
			tree = append(tree, item)
		}
	}
	return tree
}
func (o *department) FindByEnterprisePkAndName(en *entity.Department) (*entity.Department, error) {
	departmentMessage, err := o.dao.Department().FindByEnterprisePkAndName(&entity.Department{
		EnterprisePk: en.EnterprisePk,
		Name:         en.Name,
		ParentPk:     en.ParentPk,
	})
	return departmentMessage, err
}

func (o *department) FindByPk(en *entity.Department) (*entity.Department, error) {
	return o.dao.Department().FindByPk(en)
}
