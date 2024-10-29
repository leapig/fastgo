package dao

import (
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

// Dao 定义接口规范
type Dao interface {
	Permission() Permission
	PermissionGroup() PermissionGroup
	UserPermission() UserPermission
	PermissionGroupPermission() PermissionGroupPermission
	Role() Role
	RoleGroup() RoleGroup
	RoleGroupPermission() RoleGroupPermission
	RolePermission() RolePermission
	User() User
	UserClient() UserClient
	Enterprise() Enterprise
	UserRealNameAuthenticationLog() UserRealNameAuthenticationLog
	Files() Files
	UserProfession() UserProfession
	UserCredentials() UserCredentials
	InterfaceResource() InterfaceResource
	PageResource() PageResource
	PageInterface() PageInterface
	MenuResource() MenuResource
	Redis() *helper.Redis
	Sms() Sms
	Department() Department
	Project() Project
	ProjectEnterpriseRelation() ProjectEnterpriseRelation
	ProjectMember() ProjectMember
	ProjectPost() ProjectPost
	Position() Position
	ApplicationRecord() ApplicationRecord
	Member() Member
	UserEntry() UserEntry
	EnterpriseUser() EnterpriseUser
	EnterpriseAreaPermission() EnterpriseAreaPermission
	EnterpriseUserAttachment() EnterpriseUserAttachment
	TurnoverRecord() TurnoverRecord
	EnterpriseRelation() EnterpriseRelation
	LogisticsPurchase() LogisticsPurchase
	TemporaryWorker() TemporaryWorker
	LogisticsWarehousing() LogisticsWarehousing
	LogisticsInventory() LogisticsInventory
	LogisticsDelivery() LogisticsDelivery
	WorkSchedule() WorkSchedule
	WorkSchedulePost() WorkSchedulePost
	PatrolRoute() PatrolRoute
	PatrolPoint() PatrolPoint
	PatrolRoutePoint() PatrolRoutePoint
	PatrolRouteRecord() PatrolRouteRecord
	PatrolRouteRecordPoint() PatrolRouteRecordPoint
	InspectionPlan() InspectionPlan
	InspectionPlanPoint() InspectionPlanPoint
	InspectionPlanUser() InspectionPlanUser
	InspectionPoint() InspectionPoint
	Car() Car
	CarApply() CarApply
}

// dao 接口规范实现类
type dao struct {
	Db *gorm.DB
	Rs *helper.Redis
}

// NewDao 实例化接口规范实现类
func NewDao(db *gorm.DB, rs *helper.Redis) Dao {
	return &dao{
		Db: db,
		Rs: rs,
	}
}

type Pagination struct {
	Page    int32  `json:"page"`
	Size    int32  `json:"size"`
	Cursor  int32  `json:"cursor"`
	Total   int32  `json:"total"`
	Keyword string `json:"keyword"`
}

func (d *dao) Redis() *helper.Redis {
	return d.Rs
}
func (d *dao) UserPermission() UserPermission {
	return NewUserPermission(d.Db, d.Rs)
}
func (d *dao) Permission() Permission {
	return NewPermission(d.Db, d.Rs)
}
func (d *dao) PermissionGroup() PermissionGroup {
	return NewPermissionGroup(d.Db, d.Rs)
}
func (d *dao) PermissionGroupPermission() PermissionGroupPermission {
	return NewPermissionGroupPermission(d.Db, d.Rs)
}
func (d *dao) Role() Role {
	return NewRole(d.Db, d.Rs)
}

func (d *dao) RoleGroup() RoleGroup {
	return NewRoleGroup(d.Db, d.Rs)
}
func (d *dao) RoleGroupPermission() RoleGroupPermission {
	return NewRoleGroupPermission(d.Db, d.Rs)
}
func (d *dao) RolePermission() RolePermission {
	return NewRolePermission(d.Db, d.Rs)
}
func (d *dao) User() User {
	return NewUser(d.Db, d.Rs)
}
func (d *dao) UserClient() UserClient {
	return NewUserClient(d.Db, d.Rs)
}
func (d *dao) Enterprise() Enterprise {
	return NewEnterprise(d.Db, d.Rs)
}
func (d *dao) UserRealNameAuthenticationLog() UserRealNameAuthenticationLog {
	return NewUserRealNameAuthenticationLog(d.Db, d.Rs)
}

func (d *dao) Files() Files {
	return NewFilesDao(d.Db, d.Rs)
}

func (d *dao) UserProfession() UserProfession {
	return NewUserProfession(d.Db, d.Rs)
}

func (d *dao) UserCredentials() UserCredentials {
	return NewUserCredentials(d.Db, d.Rs)
}

func (d *dao) InterfaceResource() InterfaceResource {
	return NewInterfaceResource(d.Db, d.Rs)
}
func (d *dao) PageResource() PageResource {
	return NewPageResource(d.Db, d.Rs)
}
func (d *dao) PageInterface() PageInterface {
	return NewPageInterface(d.Db, d.Rs)
}
func (d *dao) MenuResource() MenuResource {
	return NewMenuResource(d.Db, d.Rs)
}
func (d *dao) Sms() Sms {
	return NewSms(d.Db, d.Rs)
}
func (d *dao) Department() Department {
	return NewDepartment(d.Db, d.Rs)
}

func (d *dao) Project() Project {
	return NewProject(d.Db, d.Rs)
}
func (d *dao) ProjectEnterpriseRelation() ProjectEnterpriseRelation {
	return NewProjectEnterpriseRelation(d.Db, d.Rs)
}
func (d *dao) ProjectMember() ProjectMember {
	return NewProjectMember(d.Db, d.Rs)
}
func (d *dao) ProjectPost() ProjectPost {
	return NewProjectPost(d.Db, d.Rs)
}

func (d *dao) Position() Position {
	return NewPosition(d.Db, d.Rs)
}

func (d *dao) ApplicationRecord() ApplicationRecord {
	return NewApplicationRecord(d.Db, d.Rs)
}

func (d *dao) Member() Member {
	return NewMember(d.Db, d.Rs)
}

func (d *dao) UserEntry() UserEntry {
	return NewUserEntry(d.Db, d.Rs)
}

func (d *dao) EnterpriseUser() EnterpriseUser {
	return NewEnterpriseUser(d.Db, d.Rs)
}

func (d *dao) EnterpriseAreaPermission() EnterpriseAreaPermission {
	return NewEnterpriseAreaPermission(d.Db)
}

func (d *dao) EnterpriseUserAttachment() EnterpriseUserAttachment {
	return NewEnterpriseUserAttachment(d.Db)
}
func (d *dao) EnterpriseRelation() EnterpriseRelation {
	return NewEnterpriseRelation(d.Db)
}

func (d *dao) TurnoverRecord() TurnoverRecord {
	return NewTurnoverRecord(d.Db)
}

func (d *dao) LogisticsPurchase() LogisticsPurchase {
	return NewLogisticsPurchase(d.Db)
}

func (d *dao) TemporaryWorker() TemporaryWorker {
	return NewTemporaryWorker(d.Db)
}

func (d *dao) LogisticsWarehousing() LogisticsWarehousing {
	return NewLogisticsWarehousing(d.Db)
}

func (d *dao) LogisticsInventory() LogisticsInventory {
	return NewLogisticsInventory(d.Db)
}

func (d *dao) LogisticsDelivery() LogisticsDelivery {
	return NewLogisticsDelivery(d.Db)
}
func (d *dao) WorkSchedule() WorkSchedule {
	return NewWorkSchedule(d.Db)
}
func (d *dao) WorkSchedulePost() WorkSchedulePost {
	return NewWorkSchedulePost(d.Db)
}

func (d *dao) PatrolRoute() PatrolRoute {
	return NewPatrolRoute(d.Db)
}

func (d *dao) PatrolPoint() PatrolPoint {
	return NewPatrolPoint(d.Db)
}

func (d *dao) PatrolRoutePoint() PatrolRoutePoint {
	return NewPatrolRoutePoint(d.Db)
}
func (d *dao) InspectionPoint() InspectionPoint {
	return NewInspectionPoint(d.Db)
}
func (d *dao) InspectionPlanUser() InspectionPlanUser {
	return NewInspectionPlanUser(d.Db)
}
func (d *dao) InspectionPlanPoint() InspectionPlanPoint {
	return NewInspectionPlanPoint(d.Db)
}
func (d *dao) InspectionPlan() InspectionPlan {
	return NewInspectionPlan(d.Db)
}

func (d *dao) PatrolRouteRecord() PatrolRouteRecord {
	return NewPatrolRouteRecord(d.Db)
}
func (d *dao) PatrolRouteRecordPoint() PatrolRouteRecordPoint {
	return NewPatrolRouteRecordPoint(d.Db)
}

func (d *dao) Car() Car {
	return NewCar(d.Db)
}

func (d *dao) CarApply() CarApply {
	return NewCarApply(d.Db)
}
