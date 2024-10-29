package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"sort"
)

type MenuResource interface {
	Create(*entity.MenuResource) (*entity.MenuResource, error)
	Delete(*entity.MenuResource) error
	Select(*entity.MenuResource, *dao.Pagination) ([]*model.MenuResourceModel, *dao.Pagination, error)
	Update(*entity.MenuResource) (*entity.MenuResource, error)
	FindByPk(en *entity.MenuResource) (*entity.MenuResource, error)
	SelectAllMenuWithDetail(en *entity.MenuResource) ([]*pb.MenuList, error)
}

type menuResource struct {
	dao dao.Dao
}

func NewMenuResource(dao dao.Dao) MenuResource {
	return &menuResource{dao: dao}
}
func (o *menuResource) Create(in *entity.MenuResource) (*entity.MenuResource, error) {
	in.Pk = helper.GetRid(helper.MenuResource)
	return o.dao.MenuResource().Create(in)
}
func (o *menuResource) Delete(in *entity.MenuResource) error {
	if count, err := o.dao.Permission().Count(&entity.Permission{Resource: in.Pk, ResourceType: 1}); err != nil || count > 0 {
		return errors.New("该项已被使用，请先删除使用者")
	}
	if count, err := o.dao.MenuResource().Count(&entity.MenuResource{ParentPk: in.Pk}); err != nil || count > 0 {
		return errors.New("请先删除子菜单")
	}
	return o.dao.MenuResource().Delete(in)
}
func (o *menuResource) Select(in *entity.MenuResource, pg *dao.Pagination) ([]*model.MenuResourceModel, *dao.Pagination, error) {
	if rows, err := o.dao.MenuResource().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.MenuResource().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *menuResource) Update(in *entity.MenuResource) (*entity.MenuResource, error) {
	return o.dao.MenuResource().Update(in)
}
func (o *menuResource) FindByPk(en *entity.MenuResource) (*entity.MenuResource, error) {
	return o.dao.MenuResource().FindByPk(en)
}
func (o *menuResource) SelectAllMenuWithDetail(en *entity.MenuResource) ([]*pb.MenuList, error) {
	menu, err := o.dao.MenuResource().SelectAllMenu(en)
	if err != nil {
		return nil, err
	}
	rows := make([]*pb.MenuList, 0)
	if menu != nil && len(menu) > 0 {
		for _, v := range menu {
			if v.MenuType == 1 && v.ResourceKey != "" {
				if pageResourceMessage, err3 := o.dao.PageResource().FindByPk(&entity.PageResource{
					Pk: utils.StringToInt64(v.ResourceKey),
				}); err3 == nil && pageResourceMessage != nil {
					rows = append(rows, &pb.MenuList{
						Pk:          utils.Int64ToString(v.Pk),
						MenuType:    v.MenuType,
						MenuName:    v.MenuName,
						ResourceKey: v.ResourceKey,
						ParentPk:    utils.Int64ToString(v.ParentPk),
						Icon:        v.Icon,
						Sort:        v.Sort,
						CreateTime:  utils.TimeToString(v.CreatedAt),
						UpdateTime:  utils.TimeToString(v.UpdatedAt),
						Path:        v.Path,
						PageMessage: &pb.PageResourceWithInterfaceMessage{
							Pk:            utils.Int64ToString(pageResourceMessage.Pk),
							PageName:      pageResourceMessage.PageName,
							PagePath:      pageResourceMessage.PagePath,
							Component:     pageResourceMessage.Component,
							ComponentName: pageResourceMessage.ComponentName,
							IsCache:       pageResourceMessage.IsCache,
						},
						ChildMenus: nil,
					})
				} else {
					rows = append(rows, &pb.MenuList{
						Pk:          utils.Int64ToString(v.Pk),
						MenuType:    v.MenuType,
						MenuName:    v.MenuName,
						ResourceKey: v.ResourceKey,
						ParentPk:    utils.Int64ToString(v.ParentPk),
						Icon:        v.Icon,
						Sort:        v.Sort,
						Path:        v.Path,
						CreateTime:  utils.TimeToString(v.CreatedAt),
						UpdateTime:  utils.TimeToString(v.UpdatedAt),
						PageMessage: nil,
						ChildMenus:  nil,
					})
				}
			} else {
				rows = append(rows, &pb.MenuList{
					Pk:          utils.Int64ToString(v.Pk),
					MenuType:    v.MenuType,
					MenuName:    v.MenuName,
					ResourceKey: v.ResourceKey,
					ParentPk:    utils.Int64ToString(v.ParentPk),
					Icon:        v.Icon,
					Sort:        v.Sort,
					Path:        v.Path,
					CreateTime:  utils.TimeToString(v.CreatedAt),
					UpdateTime:  utils.TimeToString(v.UpdatedAt),
					PageMessage: nil,
					ChildMenus:  nil,
				})
			}
		}
	} else {
		return nil, nil
	}
	lists := sortMenu(rows, "1")
	recursiveSort(lists)
	return lists, nil
}
func recursiveSort(menus []*pb.MenuList) {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	for i := range menus {
		if len(menus[i].ChildMenus) > 0 {
			recursiveSort(menus[i].ChildMenus)
		}
	}
}
func sortMenu(data []*pb.MenuList, parentPk string) []*pb.MenuList {
	var tree []*pb.MenuList
	for _, item := range data {
		if item.ParentPk == parentPk {
			item.ChildMenus = sortMenu(data, item.Pk)
			tree = append(tree, item)
		}
	}
	return tree
}
