package service

import (
	"encoding/json"
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/dal/pb"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
	"sync"
)

const (
	//permissionKeyName = "permissionKey"
	permissionRedisName = "scope_"
	menuRedisKey        = "_menu"
	pageRedisKey        = "_page"
	interfaceRedisKey   = "_interface"
)

type UserPermission interface {
	Create(*entity.UserPermission) (*entity.UserPermission, error)
	CreateForCorporate(in *entity.UserPermission) error
	Delete(*entity.UserPermission) error
	Select(*entity.UserPermission, *dao.Pagination) ([]*entity.UserPermission, *dao.Pagination, error)
	SelectWithUserMessage(in *entity.UserPermission, pg *dao.Pagination) ([]*model.UserPermissionModel, *dao.Pagination, error)
	SelectWithEnterpriseUserMessage(in *entity.UserPermission, pg *dao.Pagination) ([]*model.UserPermissionForEnterprise, *dao.Pagination, error)
	SelectPermissionByUserPkAndEnterprisePk(in *entity.UserPermission) ([]*entity.Permission, error)
	Update(*entity.UserPermission) (*entity.UserPermission, error)
	SelectAllUserPermission(in *entity.UserPermission) ([]*entity.UserPermission, error)
	SelectUserInterfaceKeyByUserPkAndEnterprisePk(in *entity.UserPermission) ([]string, error)
	CompareUserInterfaceKey(in *entity.UserPermission, interfaceKey string) (bool, error)
	FullRefreshUserPermissionByUserPkAndEnterprisePk(userPk, enterprisePk int64) error
	CheckFullRefreshUserPermissionByUserPkAndEnterprisePk(userPk, enterprisePk int64) bool
	GetUserPermissionByUserPkAndEnterprisePkForRedis(userPk, enterprisePk int64) (*pb.FullUserPermissionResp, error)
}

// openPlatform 接口规范实现类
type userPermission struct {
	dao dao.Dao
}

// NewUserPermission 实例化接口规范实现类
func NewUserPermission(dao dao.Dao) UserPermission {
	return &userPermission{dao: dao}
}
func (o *userPermission) Create(in *entity.UserPermission) (*entity.UserPermission, error) {
	//if err := o.dao.UserPermission().DeleteRedisPermission(in.EnterprisePk, in.UserPk); err != nil {
	//	logger.Error(err)
	//}
	//权限变动标识
	enterprisePkString := strconv.FormatInt(in.EnterprisePk, 10)
	userPkString := strconv.FormatInt(in.UserPk, 10)
	if err := o.dao.Redis().SetEx(enterprisePkString+"_"+userPkString, 60*30, "1"); err != nil {
		logger.Error(err)
	}
	in.Pk = helper.GetRid(helper.UserPermission)
	return o.dao.UserPermission().Create(in)
}
func (o *userPermission) CreateForCorporate(in *entity.UserPermission) error {
	//if err := o.dao.UserPermission().DeleteRedisPermission(in.EnterprisePk, in.UserPk); err != nil {
	//	logger.Error(err)
	//}
	//权限变动标识
	enterprisePkString := strconv.FormatInt(in.EnterprisePk, 10)
	userPkString := strconv.FormatInt(in.UserPk, 10)
	if err := o.dao.Redis().SetEx(enterprisePkString+"_"+userPkString, 60*30, "1"); err != nil {
		logger.Error(err)
	}
	return o.dao.UserPermission().CreateForCorporate(in)
}
func (o *userPermission) Delete(in *entity.UserPermission) error {
	result, err := o.dao.UserPermission().FindByPk(&entity.UserPermission{Pk: in.Pk})
	if err != nil {
		return err
	}
	//if err := o.dao.UserPermission().DeleteRedisPermission(result.EnterprisePk, result.UserPk); err != nil {
	//	logger.Error(err)
	//}
	//权限变动标识
	enterprisePkString := strconv.FormatInt(result.EnterprisePk, 10)
	userPkString := strconv.FormatInt(result.UserPk, 10)
	if err := o.dao.Redis().SetEx(enterprisePkString+"_"+userPkString, 60*30, "1"); err != nil {
		logger.Error(err)
	}
	return o.dao.UserPermission().Delete(in)
}
func (o *userPermission) Select(in *entity.UserPermission, pg *dao.Pagination) ([]*entity.UserPermission, *dao.Pagination, error) {
	if rows, err := o.dao.UserPermission().Select(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.UserPermission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *userPermission) SelectWithUserMessage(in *entity.UserPermission, pg *dao.Pagination) ([]*model.UserPermissionModel, *dao.Pagination, error) {
	if rows, err := o.dao.UserPermission().SelectWithUser(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.UserPermission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *userPermission) SelectWithEnterpriseUserMessage(in *entity.UserPermission, pg *dao.Pagination) ([]*model.UserPermissionForEnterprise, *dao.Pagination, error) {
	if rows, err := o.dao.UserPermission().SelectWithEnterpriseUser(in, pg); err != nil {
		return nil, pg, err
	} else {
		pg.Total, err = o.dao.UserPermission().Count(in)
		pg.Cursor = ((pg.Page - 1) * pg.Size) + int32(len(rows))
		if err != nil {
			return nil, nil, err
		}
		return rows, pg, err
	}
}
func (o *userPermission) Update(in *entity.UserPermission) (*entity.UserPermission, error) {
	return o.dao.UserPermission().Update(in)
}

func (o *userPermission) SelectAllUserPermission(in *entity.UserPermission) ([]*entity.UserPermission, error) {
	return o.dao.UserPermission().SelectAllUserPermission(in)
}

// SelectPermissionByUserPkAndEnterprisePk 通过人员、企业pk获取该人员的权限
func (o *userPermission) SelectPermissionByUserPkAndEnterprisePk(in *entity.UserPermission) ([]*entity.Permission, error) {
	//redisName := permissionKeyName + "_" + strconv.FormatInt(in.EnterprisePk, 10)
	//redisKey := strconv.FormatInt(in.UserPk, 10)
	//if permissionForRedis, err := o.dao.Redis().HGet(redisName, redisKey); err == nil && permissionForRedis != nil {
	//	permissions := make([]*entity.Permission, 0)
	//	if err = json.Unmarshal(permissionForRedis.([]byte), &permissions); err == nil {
	//		return permissions, nil
	//	}
	//}
	if allUserPermission, err := o.dao.UserPermission().SelectAllUserPermission(in); err == nil && allUserPermission != nil && len(allUserPermission) > 0 {
		roles := make([]*entity.UserPermission, 0)
		roleGroups := make([]*entity.UserPermission, 0)
		rolePermissions := make([]*entity.Permission, 0)
		roleGroupPermissions := make([]*entity.Permission, 0)
		result := make([]*entity.Permission, 0)
		for _, item := range allUserPermission {
			if item.PermissionType == 1 {
				roleGroups = append(roleGroups, item)
			} else if item.PermissionType == 2 {
				roles = append(roles, item)
			}
		}
		var wg sync.WaitGroup
		if len(roles) > 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, v := range roles {
					if permissions, err2 := o.dao.RolePermission().SelectPermissionByRolePK(v.PermissionPk); err2 == nil && permissions != nil && len(permissions) > 0 {
						rolePermissions = append(rolePermissions, permissions...)
					}
				}
			}()
		}
		if len(roleGroups) > 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, v := range roleGroups {
					if permissions, err2 := o.dao.RoleGroupPermission().SelectPermissionByRoleGroupPK(v.PermissionPk); err2 == nil && permissions != nil && len(permissions) > 0 {
						roleGroupPermissions = append(roleGroupPermissions, permissions...)
					}
				}
			}()
		}
		wg.Wait()
		if len(rolePermissions) > 0 {
			result = append(result, rolePermissions...)
		}
		if len(roleGroupPermissions) > 0 {
			result = append(result, roleGroupPermissions...)
		}
		result = distance(result)
		//if jsonData, err := json.Marshal(result); err == nil {
		//	err := o.dao.Redis().HSet(redisName, redisKey, string(jsonData))
		//	logger.Error(err)
		//}
		return result, nil
	} else {
		return nil, err
	}
}
func distance(slice []*entity.Permission) []*entity.Permission {
	result := make([]*entity.Permission, 0)
	distanceMap := make(map[int64]bool)
	for _, item := range slice {
		if !distanceMap[item.Pk] {
			distanceMap[item.Pk] = true
			result = append(result, item)
		}
	}
	return result
}

// SelectUserInterfaceKeyByUserPkAndEnterprisePk 获取用户接口权限关键字
func (o *userPermission) SelectUserInterfaceKeyByUserPkAndEnterprisePk(in *entity.UserPermission) ([]string, error) {
	if in.EnterprisePk == 0 || in.UserPk == 0 {
		return nil, errors.New("缺少请求参数！！！")
	}
	allUserPermission, err := o.dao.UserPermission().SelectAllUserPermission(&entity.UserPermission{
		EnterprisePk: in.EnterprisePk,
		UserPk:       in.UserPk,
	})
	if err != nil {
		return nil, err
	}
	if allUserPermission == nil || len(allUserPermission) <= 0 {
		return nil, nil
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	result := make([]string, 0)
	for _, userPermission := range allUserPermission {
		if userPermission.PermissionType == 1 {
			//角色组
			wg.Add(1)
			go func() {
				//角色组获取接口key方法
				defer wg.Done()
				if roleGroupPermissionList, err2 := o.dao.RoleGroupPermission().SelectRoleGroupPermissionByRoleGroupPk(&entity.RoleGroupPermission{
					RoleGroupPk: userPermission.PermissionPk,
				}); err2 == nil && roleGroupPermissionList != nil && len(roleGroupPermissionList) > 0 {
					for _, v := range roleGroupPermissionList {
						switch v.PermissionType {
						case 1:
							//角色
							if v.Role != nil {
								if rolePermissionList, err2 := o.dao.RolePermission().SelectAllRolePermission(&entity.RolePermission{RolePk: v.Role.Pk}); err2 == nil && rolePermissionList != nil && len(rolePermissionList) > 0 {
									for _, rolePermissionMessage := range rolePermissionList {
										switch rolePermissionMessage.PermissionType {
										case 1:
											if permissionMessage, err3 := o.dao.Permission().FindByPk(&entity.Permission{Pk: rolePermissionMessage.PermissionPk}); err3 == nil && permissionMessage != nil {
												key := o.SelectPermissionInterfaceKey(permissionMessage)
												setData(key, &result, &mutex)
											}
										case 2:
											if permissionList, err3 := o.dao.PermissionGroupPermission().SelectAllPermissionGroupPermission(&model.PermissionGroupPermissionModel{
												PermissionGroupPk: rolePermissionMessage.PermissionPk,
											}); err3 == nil && permissionList != nil && len(permissionList) > 0 {
												for _, permissionMessage := range permissionList {
													if permissionMessage.Permission != nil {
														key := o.SelectPermissionInterfaceKey(permissionMessage.Permission)
														setData(key, &result, &mutex)
													}
												}
											}
										default:
											continue
										}
									}
								}
							}
						case 2:
							//单一权限项
							if v.Permission != nil {
								key := o.SelectPermissionInterfaceKey(v.Permission)
								setData(key, &result, &mutex)
							}
						case 3:
							//权限组
							if v.PermissionGroup != nil {
								if permissionList, err2 := o.dao.PermissionGroupPermission().SelectAllPermissionGroupPermission(&model.PermissionGroupPermissionModel{
									PermissionGroupPk: v.PermissionGroup.Pk,
								}); err2 == nil && permissionList != nil && len(permissionList) > 0 {
									for _, permissionMessage := range permissionList {
										if permissionMessage.Permission != nil {
											key := o.SelectPermissionInterfaceKey(permissionMessage.Permission)
											setData(key, &result, &mutex)
										}
									}
								}
							}
						}
					}
				}
			}()
		} else if userPermission.PermissionType == 2 {
			//角色
			wg.Add(1)
			go func() {
				defer wg.Done()
				//角色获取接口key方法
				if rolePermissionList, err2 := o.dao.RolePermission().SelectAllRolePermission(&entity.RolePermission{RolePk: userPermission.PermissionPk}); err2 == nil && rolePermissionList != nil && len(rolePermissionList) > 0 {
					for _, rolePermissionMessage := range rolePermissionList {
						switch rolePermissionMessage.PermissionType {
						case 1:
							if permissionMessage, err3 := o.dao.Permission().FindByPk(&entity.Permission{Pk: rolePermissionMessage.PermissionPk}); err3 == nil && permissionMessage != nil {
								key := o.SelectPermissionInterfaceKey(permissionMessage)
								setData(key, &result, &mutex)
							}
						case 2:
							if permissionList, err3 := o.dao.PermissionGroupPermission().SelectAllPermissionGroupPermission(&model.PermissionGroupPermissionModel{
								PermissionGroupPk: rolePermissionMessage.PermissionPk,
							}); err3 == nil && permissionList != nil && len(permissionList) > 0 {
								for _, permissionMessage := range permissionList {
									if permissionMessage.Permission != nil {
										key := o.SelectPermissionInterfaceKey(permissionMessage.Permission)
										setData(key, &result, &mutex)
									}
								}
							}
						default:
							continue
						}
					}
				}
			}()
		} else {
			continue
		}
	}
	wg.Wait()
	result = utils.RemoveDuplicates(result)
	return result, nil
}

func setData(data []string, resultSlice *[]string, mutex *sync.Mutex) {
	if data == nil || len(data) <= 0 {
		return
	}
	mutex.Lock()
	for _, a := range data {
		*resultSlice = append(*resultSlice, a)
	}
	mutex.Unlock()
}

// SelectPermissionInterfaceKey 获取单一权限项关联到的接口Key
func (o *userPermission) SelectPermissionInterfaceKey(in *entity.Permission) []string {
	if in.Resource == 0 || in.OperationType == 0 {
		return nil
	}
	var operationType int32
	operationType = 0
	if in.OperationType == 1 {
		operationType = 1
	}
	result := make([]string, 0)
	switch in.ResourceType {
	case 1:
		//菜单
		if menu, err := o.dao.MenuResource().FindByPk(&entity.MenuResource{
			Pk: in.Resource,
		}); err == nil && menu != nil && menu.MenuType == 1 && menu.ResourceKey != "" {
			//菜单中挂载了页面资源
			if pageInterfaceMessage, err1 := o.dao.PageInterface().FindByPagePkAndOperationType(&entity.PageInterface{
				PagePk:        utils.StringToInt64(menu.ResourceKey),
				OperationType: operationType,
			}); err1 == nil && pageInterfaceMessage != nil && len(pageInterfaceMessage) > 0 {
				for _, v := range pageInterfaceMessage {
					if v.InterfaceResource != nil && v.InterfaceResource.InterfaceKey != "" {
						result = append(result, v.InterfaceResource.InterfaceKey)
					}
				}
			}
		}
	case 2:
		//页面
		if pageInterfaceMessage, err1 := o.dao.PageInterface().FindByPagePkAndOperationType(&entity.PageInterface{
			PagePk:        in.Resource,
			OperationType: operationType,
		}); err1 == nil && pageInterfaceMessage != nil && len(pageInterfaceMessage) > 0 {
			for _, v := range pageInterfaceMessage {
				if v.InterfaceResource != nil && v.InterfaceResource.InterfaceKey != "" {
					result = append(result, v.InterfaceResource.InterfaceKey)
				}
			}
		}
	default:
	}
	return result
}

// CompareUserInterfaceKey 比对用户是否有接口权限
func (o *userPermission) CompareUserInterfaceKey(in *entity.UserPermission, interfaceKey string) (bool, error) {
	if keys, err := o.SelectUserInterfaceKeyByUserPkAndEnterprisePk(in); err == nil && slices.Contains(keys, interfaceKey) {
		return true, nil
	} else {
		return false, err
	}
}
func (o *userPermission) CheckFullRefreshUserPermissionByUserPkAndEnterprisePk(userPk, enterprisePk int64) bool {
	enterprisePkString := strconv.FormatInt(enterprisePk, 10)
	userPkString := strconv.FormatInt(userPk, 10)
	if status, err := o.dao.Redis().Get(enterprisePkString + "_" + userPkString); err == nil && status == "1" {
		return true
	}
	return false
}

func (o *userPermission) GetUserPermissionByUserPkAndEnterprisePkForRedis(userPk, enterprisePk int64) (*pb.FullUserPermissionResp, error) {
	enterprisePkString := strconv.FormatInt(enterprisePk, 10)
	hashRedisName := permissionRedisName + enterprisePkString
	userPkString := strconv.FormatInt(userPk, 10)
	menuResourceRows := make([]*pb.MenuResourceResp, 0)
	pageResourceRows := make([]*pb.PageResource, 0)
	interfaceKeys := make([]string, 0)
	userPathWays := make([]string, 0)
	if userMenuResource, err := o.dao.Redis().HGet(hashRedisName, userPkString+menuRedisKey); err == nil && userMenuResource != nil {
		menuResources := make([]*entity.MenuResource, 0)
		if err = json.Unmarshal(userMenuResource.([]byte), &menuResources); err == nil {
			if len(menuResources) > 0 {
				for _, v := range menuResources {
					menuResourceRows = append(menuResourceRows, &pb.MenuResourceResp{
						Pk:          utils.Int64ToString(v.Pk),
						MenuType:    v.MenuType,
						MenuName:    v.MenuName,
						ResourceKey: v.ResourceKey,
						ParentPk:    utils.Int64ToString(v.ParentPk),
						Icon:        v.Icon,
						Sort:        v.Sort,
					})
				}
				if len(menuResourceRows) > 1 {
					menuResourceRows = sortMenuResp(menuResourceRows, "1")
				}
				sort.Slice(menuResourceRows, func(i, j int) bool {
					if menuResourceRows[i].Sort == menuResourceRows[j].Sort {
						return i < j
					}
					return menuResourceRows[i].Sort < menuResourceRows[j].Sort
				})
				for _, menu := range menuResourceRows {
					sort.Slice(menu.MenuResourceRows, func(i, j int) bool {
						if menu.MenuResourceRows[i].Sort == menu.MenuResourceRows[j].Sort {
							return i < j
						}
						return menu.MenuResourceRows[i].Sort < menu.MenuResourceRows[j].Sort
					})
				}
			}
		}
	}
	if userPageResource, err := o.dao.Redis().HGet(hashRedisName, userPkString+pageRedisKey); err == nil && userPageResource != nil {
		pageResources := make([]*entity.PageResource, 0)
		if err = json.Unmarshal(userPageResource.([]byte), &pageResources); err == nil {
			if len(pageResources) > 0 {
				for _, v := range pageResources {
					pageResourceRows = append(pageResourceRows, &pb.PageResource{
						Pk:            utils.Int64ToString(v.Pk),
						PagePath:      v.PagePath,
						Component:     v.Component,
						ComponentName: v.ComponentName,
						IsCache:       v.IsCache,
						PageName:      v.PageName,
					})
				}
			}
		}
	}
	if userInterfaceKey, err := o.dao.Redis().HGet(hashRedisName, userPkString+interfaceRedisKey); err == nil && userInterfaceKey != nil {
		if err = json.Unmarshal(userInterfaceKey.([]byte), &interfaceKeys); err != nil {
			logger.Error(err)
		}
	}
	if userPathWay, err := o.dao.Redis().SMembers(hashRedisName + "_" + userPkString); err != nil {
		if err = json.Unmarshal(userPathWay.([]byte), &userPathWays); err != nil {
			logger.Error(err)
		}
	}
	return &pb.FullUserPermissionResp{
		InterfaceKey:     interfaceKeys,
		PathWay:          userPathWays,
		PageResourceRows: pageResourceRows,
		MenuResourceRows: menuResourceRows,
	}, nil
}
func sortMenuResp(data []*pb.MenuResourceResp, parentPk string) []*pb.MenuResourceResp {
	var tree []*pb.MenuResourceResp
	for _, item := range data {
		if item.ParentPk == parentPk {
			item.MenuResourceRows = sortMenuResp(data, item.Pk)
			tree = append(tree, item)
		}
	}
	return tree
}

// FullRefreshUserPermissionByUserPkAndEnterprisePk 全量同步人员权限方法
func (o *userPermission) FullRefreshUserPermissionByUserPkAndEnterprisePk(userPk, enterprisePk int64) error {
	userPermissions, err := o.SelectPermissionByUserPkAndEnterprisePk(&entity.UserPermission{
		UserPk:       userPk,
		EnterprisePk: enterprisePk,
	})
	if err != nil || len(userPermissions) <= 0 {
		return err
	}
	menuPermissions := make([]*entity.Permission, 0)
	pagePermissions := make([]*entity.Permission, 0)
	for _, item := range userPermissions {
		if item.ResourceType == 1 {
			menuPermissions = append(menuPermissions, item)
		} else if item.ResourceType == 2 {
			pagePermissions = append(pagePermissions, item)
		}
	}
	userMenuResource := make([]*model.MenuResourceForRedis, 0) //人员菜单资源
	userPageResourceForMenu := make([]*model.PageResource, 0)  //菜单资源中的页面
	userPageResourceForPage := make([]*model.PageResource, 0)  //页面资源中的页面
	userInterfaceKeyForMenu := make([]string, 0)               //菜单资源中的接口KEY
	userInterfaceKeyForPage := make([]string, 0)               //页面资源中的接口KEY
	userPathWayForMenu := make([]string, 0)                    //菜单资源中的path_method
	userPathWayForPage := make([]string, 0)                    //页面资源中的path_method
	var wg sync.WaitGroup
	if len(menuPermissions) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//挂载菜单的权限项
			for _, v := range menuPermissions {
				var component, componentName string
				var isCache int32
				menu, err := o.dao.MenuResource().FindByPk(&entity.MenuResource{
					Pk: v.Resource,
				})
				if err != nil {
					continue
				}
				if menu.MenuType == 1 {
					//菜单中挂载了页面资源
					page, err := o.dao.PageResource().FindByPk(&entity.PageResource{
						Pk: utils.StringToInt64(menu.ResourceKey),
					})
					if err == nil {
						isCache = page.IsCache
						component = page.Component
						componentName = page.ComponentName
						userPageResourceForMenu = append(userPageResourceForMenu, &model.PageResource{
							Pk:            page.Pk,
							PagePath:      page.PagePath,
							Component:     page.Component,
							ComponentName: page.ComponentName,
							//PageName:      page.PageName,
							//IsCache:       page.IsCache,
							OperationType: v.OperationType,
						})
					}
					if pageInterfaceMessage, err1 := o.dao.PageInterface().FindByPagePkAndOperationType(&entity.PageInterface{
						PagePk:        utils.StringToInt64(menu.ResourceKey),
						OperationType: v.OperationType,
					}); err1 == nil && pageInterfaceMessage != nil && len(pageInterfaceMessage) > 0 {
						for _, z := range pageInterfaceMessage {
							if z.InterfaceResource != nil && z.InterfaceResource.InterfaceKey != "" {
								userInterfaceKeyForMenu = append(userInterfaceKeyForMenu, z.InterfaceResource.InterfaceKey)
							}
							if z.InterfaceResource != nil && z.InterfaceResource.InterfaceWay != "" && z.InterfaceResource.InterfaceUrl != "" {
								userPathWayForMenu = append(userPathWayForMenu, z.InterfaceResource.InterfaceUrl+"_"+z.InterfaceResource.InterfaceWay)
							}
						}
					}
				}
				userMenuResource = append(userMenuResource, &model.MenuResourceForRedis{
					Pk:            menu.Pk,
					MenuType:      menu.MenuType,
					MenuName:      menu.MenuName,
					ResourceKey:   menu.ResourceKey,
					ParentPk:      menu.ParentPk,
					Icon:          menu.Icon,
					Sort:          menu.Sort,
					Path:          menu.Path,
					IsCache:       isCache,
					Component:     component,
					ComponentName: componentName,
				})

			}
		}()
	}
	if len(pagePermissions) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//挂载页面的权限项
			for _, v := range pagePermissions {
				page, err := o.dao.PageResource().FindByPk(&entity.PageResource{
					Pk: v.Resource,
				})
				if err == nil {
					userPageResourceForPage = append(userPageResourceForPage, &model.PageResource{
						Pk:            page.Pk,
						PagePath:      page.PagePath,
						Component:     page.Component,
						ComponentName: page.ComponentName,
						//PageName:      page.PageName,
						//IsCache:       page.IsCache,
						OperationType: v.OperationType,
					})
				}
				if pageInterfaceMessage, err1 := o.dao.PageInterface().FindByPagePkAndOperationType(&entity.PageInterface{
					PagePk:        v.Resource,
					OperationType: v.OperationType,
				}); err1 == nil && pageInterfaceMessage != nil && len(pageInterfaceMessage) > 0 {
					for _, z := range pageInterfaceMessage {
						if z.InterfaceResource != nil && z.InterfaceResource.InterfaceKey != "" {
							userInterfaceKeyForPage = append(userInterfaceKeyForPage, z.InterfaceResource.InterfaceKey)
						}
						if z.InterfaceResource != nil && z.InterfaceResource.InterfaceWay != "" && z.InterfaceResource.InterfaceUrl != "" {
							userPathWayForPage = append(userPathWayForPage, z.InterfaceResource.InterfaceUrl+"_"+z.InterfaceResource.InterfaceWay)
						}
					}
				}
			}
		}()
	}
	wg.Wait()
	userMenuResource = distanceMenuResource(userMenuResource) //人员菜单资源
	userPageResource := make([]*model.PageResource, 0)        //人员页面资源
	userPageResource = append(userPageResource, userPageResourceForPage...)
	userPageResource = append(userPageResource, userPageResourceForMenu...)
	userPageResource = distancePageResource(userPageResource)
	userInterfaceKey := make([]string, 0) //人员接口Key
	userInterfaceKey = append(userInterfaceKey, userInterfaceKeyForPage...)
	userInterfaceKey = append(userInterfaceKey, userInterfaceKeyForMenu...)
	userInterfaceKey = utils.RemoveDuplicates(userInterfaceKey)
	userPathWay := make([]string, 0) //人员Path_Way
	userPathWay = append(userPathWay, userPathWayForPage...)
	userPathWay = append(userPathWay, userPathWayForMenu...)
	userPathWay = utils.RemoveDuplicates(userPathWay)
	enterprisePkString := strconv.FormatInt(enterprisePk, 10)
	hashRedisName := permissionRedisName + enterprisePkString
	userPkString := strconv.FormatInt(userPk, 10)
	menuRedis := make([]*pb.MenuListForRedis, 0)
	if userMenuResource != nil && len(userMenuResource) > 0 {
		menus := make([]*pb.MenuListForRedis, 0)
		for _, v := range userMenuResource {
			menus = append(menus, &pb.MenuListForRedis{
				Pk:            utils.Int64ToString(v.Pk),
				MenuType:      v.MenuType,
				MenuName:      v.MenuName,
				ResourceKey:   v.ResourceKey,
				ParentPk:      utils.Int64ToString(v.ParentPk),
				Icon:          v.Icon,
				Sort:          v.Sort,
				Path:          v.Path,
				IsCache:       v.IsCache,
				ComponentName: v.ComponentName,
				Component:     v.Component,
				ChildMenus:    nil,
			})
		}
		menuRedis = sortMenuForRedis(menus, "1")
		recursiveSortForRedis(menuRedis)
	}
	if userMenuResourceJsonData, err := json.Marshal(menuRedis); err == nil {
		_ = o.dao.Redis().HSet(hashRedisName, userPkString+menuRedisKey, string(userMenuResourceJsonData))
	}
	if userPageResourceJsonData, err := json.Marshal(userPageResource); err == nil {
		_ = o.dao.Redis().HSet(hashRedisName, userPkString+pageRedisKey, string(userPageResourceJsonData))
	}
	if userInterfaceKeyJsonData, err := json.Marshal(userInterfaceKey); err == nil {
		_ = o.dao.Redis().HSet(hashRedisName, userPkString+interfaceRedisKey, string(userInterfaceKeyJsonData))
	}
	setRedisKey := hashRedisName + "_" + userPkString
	_, _ = o.dao.Redis().Del(setRedisKey)
	for _, v := range userPathWay {
		if _, err := o.dao.Redis().SAdd(setRedisKey, v); err != nil {
			logger.Error(err)
		}
	}
	_, _ = o.dao.Redis().Del(enterprisePkString + "_" + userPkString)
	return nil
}

// func (o *userPermission) GetUserPermissionByUserPkAndEnterprisePk(userPk, enterprisePk int64) ([]*entity.Permission, error) {
//
// }
func recursiveSortForRedis(menus []*pb.MenuListForRedis) {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	for i := range menus {
		if len(menus[i].ChildMenus) > 0 {
			recursiveSortForRedis(menus[i].ChildMenus)
		}
	}
}
func sortMenuForRedis(data []*pb.MenuListForRedis, parentPk string) []*pb.MenuListForRedis {
	var tree []*pb.MenuListForRedis
	for _, item := range data {
		if item.ParentPk == parentPk {
			item.ChildMenus = sortMenuForRedis(data, item.Pk)
			tree = append(tree, item)
		}
	}
	return tree
}

//	func distancePageResource(slice []*model.PageResource) []*model.PageResource {
//		result := make([]*model.PageResource, 0)
//		distanceMap := make(map[int64]bool)
//		for _, item := range slice {
//			if !distanceMap[item.Pk] {
//				distanceMap[item.Pk] = true
//				result = append(result, item)
//			}
//		}
//		return result
//	}
func distancePageResource(slice []*model.PageResource) []*model.PageResource {
	result := make([]*model.PageResource, 0)
	distanceMap := make(map[int64]*model.PageResource)
	for _, item := range slice {
		if _, exists := distanceMap[item.Pk]; exists {
			if item.OperationType == 2 {
				distanceMap[item.Pk] = item
			}
		} else {
			distanceMap[item.Pk] = item
		}
	}
	for _, item := range distanceMap {
		result = append(result, item)
	}
	return result
}
func distanceMenuResource(slice []*model.MenuResourceForRedis) []*model.MenuResourceForRedis {
	result := make([]*model.MenuResourceForRedis, 0)
	distanceMap := make(map[int64]bool)
	for _, item := range slice {
		if !distanceMap[item.Pk] {
			distanceMap[item.Pk] = true
			result = append(result, item)
		}
	}
	return result
}
