package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/library/helper"
)

var S Svc

// Svc 定义接口规范
type Svc interface {
	Wechat() Wechat
	UserPermission() UserPermission
	Permission() Permission
	Role() Role
	RoleGroup() RoleGroup
	RoleGroupPermission() RoleGroupPermission
	RolePermission() RolePermission
	PermissionGroup() PermissionGroup
	PermissionGroupPermission() PermissionGroupPermission
	User() User
	UserClient() UserClient
	Enterprise() Enterprise
	Files() Files
	UserProfession() UserProfession
	UserCredentials() UserCredentials
	PageResource() PageResource
	InterfaceResource() InterfaceResource
	PageInterface() PageInterface
	MenuResource() MenuResource
	Sms() Sms
	UserRealNameAuthenticationLog() UserRealNameAuthenticationLog
	EnterpriseAreaPermission() EnterpriseAreaPermission
	Department() Department
	Member() Member
	EnterpriseUser() EnterpriseUser
	RS() *helper.Redis
}

// svc 接口规范实现类
type svc struct {
	dao dao.Dao
}

// NewSvc 实例化接口规范实现类
func NewSvc(dao dao.Dao) Svc {
	return svc{
		dao: dao,
	}
}

func InitSvc(dao dao.Dao) {
	S = NewSvc(dao)
}

func (svc svc) RS() *helper.Redis {
	return svc.dao.Redis()
}

func (svc svc) UserPermission() UserPermission {
	return NewUserPermission(svc.dao)
}
func (svc svc) Permission() Permission {
	return NewPermission(svc.dao)
}

func (svc svc) Role() Role {
	return NewRole(svc.dao)
}
func (svc svc) RoleGroup() RoleGroup {
	return NewRoleGroup(svc.dao)
}
func (svc svc) RoleGroupPermission() RoleGroupPermission {
	return NewRoleGroupPermission(svc.dao)
}
func (svc svc) RolePermission() RolePermission {
	return NewRolePermission(svc.dao)
}
func (svc svc) PermissionGroup() PermissionGroup {
	return NewPermissionGroup(svc.dao)
}
func (svc svc) PermissionGroupPermission() PermissionGroupPermission {
	return NewPermissionGroupPermission(svc.dao)
}
func (svc svc) User() User {
	return NewUser(svc.dao)
}
func (svc svc) UserClient() UserClient {
	return NewUserClient(svc.dao)
}
func (svc svc) Enterprise() Enterprise {
	return NewEnterprise(svc.dao)
}

func (svc svc) Files() Files {
	return NewFiles(svc.dao)
}

func (svc svc) UserProfession() UserProfession {
	return NewUserProfession(svc.dao)
}

func (svc svc) UserCredentials() UserCredentials {
	return NewUserCredentials(svc.dao)
}

func (svc svc) PageResource() PageResource {
	return NewPageResource(svc.dao)
}
func (svc svc) InterfaceResource() InterfaceResource {
	return NewInterfaceResource(svc.dao)
}
func (svc svc) PageInterface() PageInterface {
	return NewPageInterface(svc.dao)
}
func (svc svc) MenuResource() MenuResource {
	return NewMenuResource(svc.dao)
}

func (svc svc) Wechat() Wechat {
	return NewWechat(svc.dao)
}

func (svc svc) Sms() Sms {
	return NewSms(svc.dao)
}

func (svc svc) UserRealNameAuthenticationLog() UserRealNameAuthenticationLog {
	return NewUserRealNameAuthenticationLog(svc.dao)
}

func (svc svc) EnterpriseUser() EnterpriseUser {
	return NewEnterpriseUser(svc.dao)
}

func (svc svc) EnterpriseAreaPermission() EnterpriseAreaPermission {
	return NewEnterpriseAreaPermission(svc.dao)
}

func (svc svc) Position() Position {
	return NewPosition(svc.dao)
}

func (svc svc) Department() Department {
	return NewDepartment(svc.dao)
}

func (svc svc) Member() Member {
	return NewMember(svc.dao)
}
