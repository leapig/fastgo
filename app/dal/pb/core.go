package pb

type Enterprise struct {
	Pk                        string                      `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	Cover                     string                      `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover"`
	Name                      string                      `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Serial                    string                      `protobuf:"bytes,4,opt,name=serial,proto3" json:"serial"`
	License                   string                      `protobuf:"bytes,5,opt,name=license,proto3" json:"license"`
	Country                   string                      `protobuf:"bytes,6,opt,name=country,proto3" json:"country"`
	Province                  string                      `protobuf:"bytes,7,opt,name=province,proto3" json:"province"`
	City                      string                      `protobuf:"bytes,8,opt,name=city,proto3" json:"city"`
	District                  string                      `protobuf:"bytes,9,opt,name=district,proto3" json:"district"`
	County                    string                      `protobuf:"bytes,10,opt,name=county,proto3" json:"county"`
	Site                      string                      `protobuf:"bytes,11,opt,name=site,proto3" json:"site"`
	Longitude                 string                      `protobuf:"bytes,12,opt,name=longitude,proto3" json:"longitude"`
	Latitude                  string                      `protobuf:"bytes,13,opt,name=latitude,proto3" json:"latitude"`
	CorporatePk               string                      `protobuf:"bytes,14,opt,name=corporatePk,proto3" json:"corporatePk"`
	CorporateName             string                      `protobuf:"bytes,15,opt,name=corporateName,proto3" json:"corporateName"`
	CorporatePhone            string                      `protobuf:"bytes,16,opt,name=corporatePhone,proto3" json:"corporatePhone"`
	EnterpriseAreaPermissions []*EnterpriseAreaPermission `protobuf:"bytes,22,rep,name=enterpriseAreaPermissions,proto3" json:"enterpriseAreaPermissions"`
	Type                      int32                       `protobuf:"varint,17,opt,name=type,proto3" json:"type"`
	StaffSize                 string                      `protobuf:"bytes,18,opt,name=staff_size,json=staffSize,proto3" json:"staff_size"`
	CreateAt                  string                      `protobuf:"bytes,19,opt,name=createAt,proto3" json:"createAt"`
	Ctx                       *Ctx                        `protobuf:"bytes,20,opt,name=ctx,proto3" json:"ctx"`
	AddressCode               string                      `protobuf:"bytes,21,opt,name=addressCode,proto3" json:"addressCode"`
}

type EnterpriseAreaPermission struct {
	Pk           string `protobuf:"bytes,1,opt,name=Pk,proto3" json:"Pk"`
	EnterprisePk string `protobuf:"bytes,2,opt,name=EnterprisePk,proto3" json:"EnterprisePk"`
	Province     string `protobuf:"bytes,3,opt,name=province,proto3" json:"province"`
	City         string `protobuf:"bytes,4,opt,name=city,proto3" json:"city"`
	District     string `protobuf:"bytes,5,opt,name=district,proto3" json:"district"`
	County       string `protobuf:"bytes,6,opt,name=county,proto3" json:"county"`
	AddressCode  string `protobuf:"bytes,7,opt,name=address_code,json=addressCode,proto3" json:"address_code"`
}

type Ctx struct {
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size"`
}

type Resp struct {
	Page   int32       `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Size   int32       `protobuf:"varint,2,opt,name=size,proto3" json:"size"`
	Cursor int32       `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor"`
	Total  int32       `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
	Rows   interface{} `protobuf:"bytes,5,rep,name=rows,proto3" json:"rows"`
}

type User struct {
	Pk         string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Phone      string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone"`
	IdCard     string `protobuf:"bytes,5,opt,name=idCard,proto3" json:"idCard"`
	Birthday   string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday"`
	IsRealName int32  `protobuf:"varint,7,opt,name=isRealName,proto3" json:"isRealName"`
	LiveCity   string `protobuf:"bytes,8,opt,name=liveCity,proto3" json:"liveCity"`
	CreateAt   string `protobuf:"bytes,9,opt,name=createAt,proto3" json:"createAt"`
	LastLiveAt string `protobuf:"bytes,10,opt,name=lastLiveAt,proto3" json:"lastLiveAt"`
	Gender     int32  `protobuf:"varint,11,opt,name=gender,proto3" json:"gender"`
	Height     string `protobuf:"bytes,12,opt,name=height,proto3" json:"height"`
	Weight     string `protobuf:"bytes,13,opt,name=weight,proto3" json:"weight"`
}

type UserAccount struct {
	Pk                       string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	Name                     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Phone                    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone"`
	Role                     string `protobuf:"bytes,5,opt,name=role,proto3" json:"role"`
	WxOfficialAccountsOpenId string `protobuf:"bytes,6,opt,name=wxOfficialAccountsOpenId,proto3" json:"wxOfficialAccountsOpenId"`
	WxMiniProgramOpenId      string `protobuf:"bytes,7,opt,name=wxMiniProgramOpenId,proto3" json:"wxMiniProgramOpenId"`
	CreateAt                 string `protobuf:"bytes,8,opt,name=createAt,proto3" json:"createAt"`
	UserPk                   string `protobuf:"bytes,9,opt,name=user_pk,json=userPk,proto3" json:"user_pk"`
}

type UserCredentials struct {
	Pk            string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Serial        string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial"`
	Type          int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type"`
	UserPk        string `protobuf:"bytes,5,opt,name=user_pk,json=userPk,proto3" json:"user_pk"`
	FrontFileName string `protobuf:"bytes,6,opt,name=front_file_name,json=frontFileName,proto3" json:"front_file_name"`
	BackFileName  string `protobuf:"bytes,7,opt,name=back_file_name,json=backFileName,proto3" json:"back_file_name"`
}

type InterfaceResource struct {
	Pk            string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	InterfaceKey  string `protobuf:"bytes,2,opt,name=interfaceKey,proto3" json:"interfaceKey"`
	InterfaceName string `protobuf:"bytes,3,opt,name=interfaceName,proto3" json:"interfaceName"`
	InterfaceWay  string `protobuf:"bytes,4,opt,name=interfaceWay,proto3" json:"interfaceWay"`
	InterfaceUrl  string `protobuf:"bytes,5,opt,name=interfaceUrl,proto3" json:"interfaceUrl"`
}

type PageResourceWithInterfaceMessage struct {
	Pk            string                  `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	PagePath      string                  `protobuf:"bytes,2,opt,name=pagePath,proto3" json:"pagePath"`
	Component     string                  `protobuf:"bytes,3,opt,name=component,proto3" json:"component"`
	ComponentName string                  `protobuf:"bytes,4,opt,name=componentName,proto3" json:"componentName"`
	IsCache       int32                   `protobuf:"varint,5,opt,name=isCache,proto3" json:"isCache"`
	PageName      string                  `protobuf:"bytes,6,opt,name=pageName,proto3" json:"pageName"`
	PageType      int32                   `protobuf:"varint,7,opt,name=pageType,proto3" json:"pageType"`
	Rows          []*PageInterfaceMessage `protobuf:"bytes,8,rep,name=rows,proto3" json:"rows"`
}

type PageInterfaceMessage struct {
	Pk            string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	OperationType int32  `protobuf:"varint,2,opt,name=operationType,proto3" json:"operationType"`
	InterfaceKey  string `protobuf:"bytes,3,opt,name=interfaceKey,proto3" json:"interfaceKey"`
	InterfaceName string `protobuf:"bytes,4,opt,name=interfaceName,proto3" json:"interfaceName"`
	InterfaceWay  string `protobuf:"bytes,5,opt,name=interfaceWay,proto3" json:"interfaceWay"`
	InterfaceUrl  string `protobuf:"bytes,6,opt,name=interfaceUrl,proto3" json:"interfaceUrl"`
	InterfacePk   string `protobuf:"bytes,7,opt,name=interfacePk,proto3" json:"interfacePk"`
}

type Permission struct {
	Pk             string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	OperationType  int32  `protobuf:"varint,2,opt,name=operationType,proto3" json:"operationType"`
	Resource       string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource"`
	ResourceType   int32  `protobuf:"varint,4,opt,name=resourceType,proto3" json:"resourceType"`
	EnterprisePk   string `protobuf:"bytes,5,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	PermissionName string `protobuf:"bytes,6,opt,name=permissionName,proto3" json:"permissionName"`
	Visibility     int32  `protobuf:"varint,7,opt,name=visibility,proto3" json:"visibility"`
}

type PermissionWithDetail struct {
	Pk             string        `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	OperationType  int32         `protobuf:"varint,2,opt,name=operationType,proto3" json:"operationType"`
	Resource       string        `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource"`
	ResourceType   int32         `protobuf:"varint,4,opt,name=resourceType,proto3" json:"resourceType"`
	EnterprisePk   string        `protobuf:"bytes,5,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	PermissionName string        `protobuf:"bytes,6,opt,name=permissionName,proto3" json:"permissionName"`
	Visibility     int32         `protobuf:"varint,7,opt,name=visibility,proto3" json:"visibility"`
	PageResource   *PageResource `protobuf:"bytes,8,opt,name=pageResource,proto3" json:"pageResource"`
	MenuResource   *MenuResource `protobuf:"bytes,9,opt,name=menuResource,proto3" json:"menuResource"`
}

type PageResource struct {
	Pk            string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	PagePath      string `protobuf:"bytes,2,opt,name=pagePath,proto3" json:"pagePath"`
	Component     string `protobuf:"bytes,3,opt,name=component,proto3" json:"component"`
	ComponentName string `protobuf:"bytes,4,opt,name=componentName,proto3" json:"componentName"`
	IsCache       int32  `protobuf:"varint,5,opt,name=isCache,proto3" json:"isCache"`
	PageName      string `protobuf:"bytes,6,opt,name=pageName,proto3" json:"pageName"`
	PageType      int32  `protobuf:"varint,7,opt,name=pageType,proto3" json:"pageType"`
}

type MenuResource struct {
	Pk          string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	MenuType    int32  `protobuf:"varint,2,opt,name=menuType,proto3" json:"menuType"`
	MenuName    string `protobuf:"bytes,3,opt,name=menuName,proto3" json:"menuName"`
	ResourceKey string `protobuf:"bytes,4,opt,name=resourceKey,proto3" json:"resourceKey"`
	ParentPk    string `protobuf:"bytes,5,opt,name=parentPk,proto3" json:"parentPk"`
	Icon        string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon"`
	Sort        int32  `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`
	Path        string `protobuf:"bytes,8,opt,name=path,proto3" json:"path"`
}

type PermissionGroupWithPermission struct {
	Pk           string                          `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	GroupName    string                          `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName"`
	Remark       string                          `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk string                          `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	GroupType    int32                           `protobuf:"varint,5,opt,name=group_type,json=groupType,proto3" json:"group_type"`
	Rows         []*PermissionForPermissionGroup `protobuf:"bytes,6,rep,name=rows,proto3" json:"rows"`
}

type PermissionForPermissionGroup struct {
	Pk             string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	OperationType  int32  `protobuf:"varint,2,opt,name=operationType,proto3" json:"operationType"`
	Resource       string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource"`
	ResourceType   int32  `protobuf:"varint,4,opt,name=resourceType,proto3" json:"resourceType"`
	EnterprisePk   string `protobuf:"bytes,5,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	PermissionName string `protobuf:"bytes,6,opt,name=permissionName,proto3" json:"permissionName"`
	RelationPk     string `protobuf:"bytes,7,opt,name=relationPk,proto3" json:"relationPk"`
}

type RoleForUserPermission struct {
	Pk           string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	RoleName     string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName"`
	Remark       string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk string `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	RelationPk   string `protobuf:"bytes,5,opt,name=relationPk,proto3" json:"relationPk"`
}

type RoleGroupForUserPermission struct {
	Pk            string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	RoleGroupName string `protobuf:"bytes,2,opt,name=roleGroupName,proto3" json:"roleGroupName"`
	Remark        string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk  string `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	RelationPk    string `protobuf:"bytes,5,opt,name=relationPk,proto3" json:"relationPk"`
}

type UserPermissionRoleAndRoleGroupResp struct {
	RoleRows      []*RoleForUserPermission      `protobuf:"bytes,1,rep,name=roleRows,proto3" json:"roleRows"`
	RoleGroupRows []*RoleGroupForUserPermission `protobuf:"bytes,2,rep,name=roleGroupRows,proto3" json:"roleGroupRows"`
}

type MenuResourceResp struct {
	Pk               string              `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	MenuType         int32               `protobuf:"varint,2,opt,name=menuType,proto3" json:"menuType"`
	MenuName         string              `protobuf:"bytes,3,opt,name=menuName,proto3" json:"menuName"`
	ResourceKey      string              `protobuf:"bytes,4,opt,name=resourceKey,proto3" json:"resourceKey"`
	ParentPk         string              `protobuf:"bytes,5,opt,name=parentPk,proto3" json:"parentPk"`
	Icon             string              `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon"`
	Sort             int32               `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`
	MenuResourceRows []*MenuResourceResp `protobuf:"bytes,8,rep,name=menuResourceRows,proto3" json:"menuResourceRows"`
}

type FullUserPermissionResp struct {
	PageResourceRows []*PageResource     `protobuf:"bytes,1,rep,name=pageResourceRows,proto3" json:"pageResourceRows"`
	MenuResourceRows []*MenuResourceResp `protobuf:"bytes,2,rep,name=menuResourceRows,proto3" json:"menuResourceRows"`
	InterfaceKey     []string            `protobuf:"bytes,3,rep,name=interfaceKey,proto3" json:"interfaceKey"`
	PathWay          []string            `protobuf:"bytes,4,rep,name=pathWay,proto3" json:"pathWay"`
}

type MenuListForRedis struct {
	Pk            string              `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	MenuType      int32               `protobuf:"varint,2,opt,name=menuType,proto3" json:"menuType"`
	MenuName      string              `protobuf:"bytes,3,opt,name=menuName,proto3" json:"menuName"`
	ResourceKey   string              `protobuf:"bytes,4,opt,name=resourceKey,proto3" json:"resourceKey"`
	ParentPk      string              `protobuf:"bytes,5,opt,name=parentPk,proto3" json:"parentPk"`
	Icon          string              `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon"`
	Sort          int32               `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`
	Path          string              `protobuf:"bytes,8,opt,name=path,proto3" json:"path"`
	Component     string              `protobuf:"bytes,9,opt,name=component,proto3" json:"component"`
	ComponentName string              `protobuf:"bytes,10,opt,name=component_name,json=componentName,proto3" json:"component_name"`
	IsCache       int32               `protobuf:"varint,11,opt,name=is_cache,json=isCache,proto3" json:"is_cache"`
	ChildMenus    []*MenuListForRedis `protobuf:"bytes,12,rep,name=childMenus,proto3" json:"childMenus"`
}

type RoleWithPermission struct {
	Pk                  string                     `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	RoleName            string                     `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName"`
	Remark              string                     `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk        string                     `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	PermissionRows      []*PermissionForRole       `protobuf:"bytes,5,rep,name=permissionRows,proto3" json:"permissionRows"`
	PermissionGroupRows []*PermissionGroupForRole  `protobuf:"bytes,6,rep,name=permissionGroupRows,proto3" json:"permissionGroupRows"`
	UserRows            []*UserForRoleAndRoleGroup `protobuf:"bytes,7,rep,name=userRows,proto3" json:"userRows"`
}

type PermissionForRole struct {
	Pk             string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	OperationType  int32  `protobuf:"varint,2,opt,name=operationType,proto3" json:"operationType"`
	Resource       string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource"`
	ResourceType   int32  `protobuf:"varint,4,opt,name=resourceType,proto3" json:"resourceType"`
	EnterprisePk   string `protobuf:"bytes,5,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	PermissionName string `protobuf:"bytes,6,opt,name=permissionName,proto3" json:"permissionName"`
	RelationPk     string `protobuf:"bytes,7,opt,name=relationPk,proto3" json:"relationPk"`
}

type PermissionGroupForRole struct {
	Pk           string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	GroupName    string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName"`
	Remark       string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk string `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	RelationPk   string `protobuf:"bytes,5,opt,name=relationPk,proto3" json:"relationPk"`
}

type UserForRoleAndRoleGroup struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
}

type RoleGroupWithPermission struct {
	Pk                  string                         `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	RoleGroupName       string                         `protobuf:"bytes,2,opt,name=roleGroupName,proto3" json:"roleGroupName"`
	Remark              string                         `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk        string                         `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	RoleRows            []*RoleForRoleGroup            `protobuf:"bytes,5,rep,name=roleRows,proto3" json:"roleRows"`
	PermissionRows      []*PermissionForRoleGroup      `protobuf:"bytes,6,rep,name=permissionRows,proto3" json:"permissionRows"`
	PermissionGroupRows []*PermissionGroupForRoleGroup `protobuf:"bytes,7,rep,name=permissionGroupRows,proto3" json:"permissionGroupRows"`
	UserRows            []*UserForRoleAndRoleGroup     `protobuf:"bytes,8,rep,name=userRows,proto3" json:"userRows"`
}

type RoleForRoleGroup struct {
	Pk           string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	RoleName     string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName"`
	Remark       string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk string `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	RelationPk   string `protobuf:"bytes,5,opt,name=relationPk,proto3" json:"relationPk"`
}

type PermissionForRoleGroup struct {
	Pk             string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	OperationType  int32  `protobuf:"varint,2,opt,name=operationType,proto3" json:"operationType"`
	Resource       string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource"`
	ResourceType   int32  `protobuf:"varint,4,opt,name=resourceType,proto3" json:"resourceType"`
	EnterprisePk   string `protobuf:"bytes,5,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	PermissionName string `protobuf:"bytes,6,opt,name=permissionName,proto3" json:"permissionName"`
	RelationPk     string `protobuf:"bytes,7,opt,name=relationPk,proto3" json:"relationPk"`
}

type PermissionGroupForRoleGroup struct {
	Pk           string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	GroupName    string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName"`
	Remark       string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	EnterprisePk string `protobuf:"bytes,4,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	RelationPk   string `protobuf:"bytes,5,opt,name=relationPk,proto3" json:"relationPk"`
}

type UserMessageForRoleAndRoleGroup struct {
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Phone      string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	RelationPk string `protobuf:"bytes,4,opt,name=relationPk,proto3" json:"relationPk"`
}

type LoginQrCodeResp struct {
	QrCode string `protobuf:"bytes,1,opt,name=qrCode,proto3" json:"qrCode"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key"`
}

type MenuList struct {
	Pk          string                            `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	MenuType    int32                             `protobuf:"varint,2,opt,name=menuType,proto3" json:"menuType"`
	MenuName    string                            `protobuf:"bytes,3,opt,name=menuName,proto3" json:"menuName"`
	ResourceKey string                            `protobuf:"bytes,4,opt,name=resourceKey,proto3" json:"resourceKey"`
	ParentPk    string                            `protobuf:"bytes,5,opt,name=parentPk,proto3" json:"parentPk"`
	Icon        string                            `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon"`
	Sort        int32                             `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`
	Path        string                            `protobuf:"bytes,8,opt,name=path,proto3" json:"path"`
	PageMessage *PageResourceWithInterfaceMessage `protobuf:"bytes,9,opt,name=pageMessage,proto3" json:"pageMessage"`
	ChildMenus  []*MenuList                       `protobuf:"bytes,10,rep,name=childMenus,proto3" json:"childMenus"`
	CreateTime  string                            `protobuf:"bytes,11,opt,name=createTime,proto3" json:"createTime"`
	UpdateTime  string                            `protobuf:"bytes,12,opt,name=updateTime,proto3" json:"updateTime"`
}

type DepartmentList struct {
	Pk           string            `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	EnterprisePk string            `protobuf:"bytes,2,opt,name=enterprise_pk,json=enterprisePk,proto3" json:"enterprise_pk"`
	Name         string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	ParentPk     string            `protobuf:"bytes,4,opt,name=parent_pk,json=parentPk,proto3" json:"parent_pk"`
	Leader       *Member           `protobuf:"bytes,5,opt,name=leader,proto3" json:"leader"`
	CreateAt     string            `protobuf:"bytes,6,opt,name=create_at,json=createAt,proto3" json:"create_at"`
	UpdateAt     string            `protobuf:"bytes,7,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	Rows         []*DepartmentList `protobuf:"bytes,8,rep,name=rows,proto3" json:"rows"`
}

type Member struct {
	Pk        string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	UserPk    string `protobuf:"bytes,2,opt,name=user_pk,json=userPk,proto3" json:"user_pk"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone"`
	IsLeader  int32  `protobuf:"varint,6,opt,name=is_leader,json=isLeader,proto3" json:"is_leader"`
	IsMain    int32  `protobuf:"varint,7,opt,name=is_main,json=isMain,proto3" json:"is_main"`
	JobNumber int32  `protobuf:"varint,8,opt,name=job_number,json=jobNumber,proto3" json:"job_number"`
	JobTitle  string `protobuf:"bytes,9,opt,name=job_title,json=jobTitle,proto3" json:"job_title"`
}

type DepartMember struct {
	Pk        string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	UserPk    string `protobuf:"bytes,2,opt,name=user_pk,json=userPk,proto3" json:"user_pk"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone"`
	IsLeader  int32  `protobuf:"varint,6,opt,name=is_leader,json=isLeader,proto3" json:"is_leader"`
	IsMain    int32  `protobuf:"varint,7,opt,name=is_main,json=isMain,proto3" json:"is_main"`
	JobNumber string `protobuf:"bytes,8,opt,name=job_number,json=jobNumber,proto3" json:"job_number"`
	JobTitle  string `protobuf:"bytes,9,opt,name=job_title,json=jobTitle,proto3" json:"job_title"`
}

type EnterpriseUser struct {
	Pk                    string                      `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	EnterprisePk          string                      `protobuf:"bytes,2,opt,name=enterprise_pk,json=enterprisePk,proto3" json:"enterprise_pk"`
	UserPk                string                      `protobuf:"bytes,3,opt,name=user_pk,json=userPk,proto3" json:"user_pk"`
	Name                  string                      `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Phone                 string                      `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone"`
	Height                string                      `protobuf:"bytes,6,opt,name=height,proto3" json:"height"`
	Weight                string                      `protobuf:"bytes,7,opt,name=weight,proto3" json:"weight"`
	Birthday              string                      `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday"`
	Gender                int32                       `protobuf:"varint,10,opt,name=gender,proto3" json:"gender"`
	JobTitle              string                      `protobuf:"bytes,11,opt,name=job_title,json=jobTitle,proto3" json:"job_title"`
	JobNumber             string                      `protobuf:"bytes,12,opt,name=job_number,json=jobNumber,proto3" json:"job_number"`
	DepartmentPk          []string                    `protobuf:"bytes,13,rep,name=department_pk,json=departmentPk,proto3" json:"department_pk"`
	Nation                string                      `protobuf:"bytes,14,opt,name=nation,proto3" json:"nation"`
	Education             string                      `protobuf:"bytes,15,opt,name=education,proto3" json:"education"`
	EmergencyContact      string                      `protobuf:"bytes,16,opt,name=emergency_contact,json=emergencyContact,proto3" json:"emergency_contact"`
	EmergencyContactPhone string                      `protobuf:"bytes,17,opt,name=emergency_contact_phone,json=emergencyContactPhone,proto3" json:"emergency_contact_phone"`
	EntryTime             string                      `protobuf:"bytes,18,opt,name=entry_time,json=entryTime,proto3" json:"entry_time"`
	TermOfContract        string                      `protobuf:"bytes,19,opt,name=term_of_contract,json=termOfContract,proto3" json:"term_of_contract"`
	Address               string                      `protobuf:"bytes,20,opt,name=address,proto3" json:"address"`
	ContractFileName      []*EnterpriseUserAttachment `protobuf:"bytes,21,rep,name=contract_file_name,json=contractFileName,proto3" json:"contract_file_name"`
	BodyReportFileNames   []*EnterpriseUserAttachment `protobuf:"bytes,22,rep,name=body_report_file_names,json=bodyReportFileNames,proto3" json:"body_report_file_names"`
	ReviewFileNames       []*EnterpriseUserAttachment `protobuf:"bytes,23,rep,name=review_file_names,json=reviewFileNames,proto3" json:"review_file_names"`
	Status                int32                       `protobuf:"varint,24,opt,name=status,proto3" json:"status"`
	DepartmentRow         []*EnterpriseUserDepartment `protobuf:"bytes,25,rep,name=department_row,json=departmentRow,proto3" json:"department_row"`
}

type EnterpriseUserAttachment struct {
	Pk           string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	EnterprisePk string `protobuf:"bytes,2,opt,name=enterprisePk,proto3" json:"enterprisePk"`
	UserPk       string `protobuf:"bytes,3,opt,name=userPk,proto3" json:"userPk"`
	FileName     string `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName"`
	FileType     string `protobuf:"bytes,5,opt,name=fileType,proto3" json:"fileType"`
	Url          string `protobuf:"bytes,6,opt,name=url,proto3" json:"url"`
}

type EnterpriseUserDepartment struct {
	Pk           string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	DepartmentPk string `protobuf:"bytes,3,opt,name=department_pk,json=departmentPk,proto3" json:"department_pk"`
}
