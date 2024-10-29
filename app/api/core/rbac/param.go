package rbac

type Pagination struct {
	Page int32 `json:"page" form:"page" query:"page"`
	Size int32 `json:"size" form:"size" query:"size"`
}

type EnterpriseUserPermissionQueryBody struct {
	UserPk string `json:"user_pk"  form:"user_pk" query:"user_pk"`
}
type GetPermissionReq struct {
	OperationType  int32  `json:"operation_type"  form:"operation_type" query:"operation_type"`
	Resource       string `json:"resource"  form:"resource" query:"resource"`
	ResourceType   int32  `json:"resource_type"  form:"resource_type" query:"resource_type"`
	PermissionName string `json:"permission_name"  form:"permission_name" query:"permission_name"`
	Pagination
}
type GetPermissionGroupReq struct {
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	GroupName    string `json:"group_name"  form:"group_name" query:"group_name"`
	Pagination
}
type PostPermissionGroupReq struct {
	Remark    string `json:"remark"  form:"remark" query:"remark"`
	GroupName string `json:"group_name"  form:"group_name" query:"group_name"`
}
type PutPermissionGroupReq struct {
	Pk        string `json:"pk"  form:"pk" query:"pk"`
	GroupName string `json:"group_name"  form:"group_name" query:"group_name"`
	Remark    string `json:"remark"  form:"remark" query:"remark"`
}
type PostPermissionGroupPermissionReq struct {
	PermissionPk      string `json:"permission_pk"  form:"permission_pk" query:"permission_pk"`
	PermissionGroupPk string `json:"permission_group_pk"  form:"permission_group_pk" query:"permission_group_pk"`
}
type PostRoleReq struct {
	Remark   string `json:"remark"  form:"remark" query:"remark"`
	RoleName string `json:"role_name"  form:"role_name" query:"role_name"`
}
type PutRoleReq struct {
	Pk       string `json:"pk"  form:"pk" query:"pk"`
	RoleName string `json:"role_name"  form:"role_name" query:"role_name"`
	Remark   string `json:"remark"  form:"remark" query:"remark"`
}
type GetRoleReq struct {
	RoleName string `json:"role_name"  form:"role_name" query:"role_name"`
	Pagination
}
type PostRoleGroupReq struct {
	Remark        string `json:"remark"  form:"remark" query:"remark"`
	EnterprisePk  string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RoleGroupName string `json:"role_group_name"  form:"role_group_name" query:"role_group_name"`
}
type PostRoleGroupPermissionReq struct {
	RoleGroupPk    string `json:"role_group_pk"  form:"role_group_pk" query:"role_group_pk"`
	PermissionType int32  `json:"permission_type"  form:"permission_type" query:"permission_type"`
	PermissionPk   string `json:"permission_pk"  form:"permission_pk" query:"permission_pk"`
}
type PutRoleGroupReq struct {
	Pk            string `json:"pk"  form:"pk" query:"pk"`
	RoleGroupName string `json:"role_group_name"  form:"role_group_name" query:"role_group_name"`
	Remark        string `json:"remark"  form:"remark" query:"remark"`
}
type GetRoleGroupReq struct {
	RoleGroupName string `json:"role_group_name"  form:"role_group_name" query:"role_group_name"`
	Pagination
}
type PostRolePermissionReq struct {
	RolePk         string `json:"role_pk"  form:"role_pk" query:"role_pk"`
	PermissionType int32  `json:"permission_type"  form:"permission_type" query:"permission_type"`
	PermissionPk   string `json:"permission_pk"  form:"permission_pk" query:"permission_pk"`
}
type GetRoleUserReq struct {
	RolePk string `json:"role_pk"  form:"role_pk" query:"role_pk"`
	Pagination
}
type GetRoleGroupUserReq struct {
	RoleGroupPk string `json:"role_group_pk"  form:"role_group_pk" query:"role_group_pk"`
	Pagination
}
type GetUserPermissionReq struct {
	UserPk       string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
}
type PostUserPermissionReq struct {
	UserPk         string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	PermissionType int32  `json:"permission_type"  form:"permission_type" query:"permission_type"`
	PermissionPk   string `json:"permission_pk"  form:"permission_pk" query:"permission_pk" example:"${权限主键}"`
}

type CreatePermissionReq struct {
	OperationType  int32  `json:"operation_type"  form:"operation_type" query:"operation_type"`
	EnterprisePk   string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	Resource       string `json:"resource"  form:"resource" query:"resource"`
	PermissionName string `json:"permission_name"  form:"permission_name" query:"permission_name"`
	ResourceType   int32  `json:"resource_type"  form:"resource_type" query:"resource_type"`
	Visibility     int32  `json:"visibility"  form:"visibility" query:"visibility"`
}

type PermissionQueryBody struct {
	UserPk       string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
}
type PermissionGroupQueryBody struct {
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	GroupName    string `json:"group_name"  form:"group_name" query:"group_name"`
	GroupType    int32  `json:"group_type"  form:"group_type" query:"group_type"`
	PageQfqz
}
type RoleGroupQueryBody struct {
	EnterprisePk  string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RoleGroupName string `json:"role_group_name"  form:"role_group_name" query:"role_group_name"`
	PageQfqz
}
type PermissionListQueryBody struct {
	EnterprisePk   string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	OperationType  int32  `json:"operation_type"  form:"operation_type" query:"operation_type"`
	Resource       string `json:"resource"  form:"resource" query:"resource"`
	ResourceType   int32  `json:"resource_type"  form:"resource_type" query:"resource_type"`
	PermissionName string `json:"permission_name"  form:"permission_name" query:"permission_name"`
	Visibility     int32  `json:"visibility"  form:"visibility" query:"visibility"`
	PageQfqz
}
type RoleQueryBody struct {
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RoleName     string `json:"role_name"  form:"role_name" query:"role_name"`
	PageQfqz
}
type RoleUserQueryBody struct {
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RolePk       string `json:"role_pk"  form:"role_pk" query:"role_pk"`
	PageQfqz
}
type RoleGroupUserQueryBody struct {
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RoleGroupPk  string `json:"role_group_pk"  form:"role_group_pk" query:"role_group_pk"`
	PageQfqz
}

type CreateUserPermissionReq struct {
	PermissionType int32  `json:"permission_type"  form:"permission_type" query:"permission_type"`
	PermissionPk   string `json:"permission_pk"  form:"permission_pk" query:"permission_pk" example:"${权限主键}"`
	UserPk         string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	EnterprisePk   string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
}
type CreateRoleGroupReq struct {
	Remark        string `json:"remark"  form:"remark" query:"remark"`
	EnterprisePk  string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RoleGroupName string `json:"role_group_name"  form:"role_group_name" query:"role_group_name"`
}
type CreateRoleGroupPermissionReq struct {
	RoleGroupPk    string `json:"role_group_pk"  form:"role_group_pk" query:"role_group_pk"`
	PermissionType int32  `json:"permission_type"  form:"permission_type" query:"permission_type"`
	PermissionPk   string `json:"permission_pk"  form:"permission_pk" query:"permission_pk"`
}
type UpdaterRoleGroupReq struct {
	Pk            string `json:"pk"  form:"pk" query:"pk"`
	RoleGroupName string `json:"role_group_name"  form:"role_group_name" query:"role_group_name"`
	Remark        string `json:"remark"  form:"remark" query:"remark"`
}
type DeleteReq struct {
	Pk string `json:"pk"  form:"pk" query:"pk"`
}
type CreateRoleReq struct {
	Remark       string `json:"remark"  form:"remark" query:"remark"`
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	RoleName     string `json:"role_name"  form:"role_name" query:"role_name"`
}
type UpdaterRoleReq struct {
	Pk       string `json:"pk"  form:"pk" query:"pk"`
	RoleName string `json:"role_name"  form:"role_name" query:"role_name"`
	Remark   string `json:"remark"  form:"remark" query:"remark"`
}

type UpdaterPermissionReq struct {
	Pk             string `json:"pk"  form:"pk" query:"pk"`
	OperationType  int32  `json:"operation_type"  form:"operation_type" query:"operation_type"`
	Resource       string `json:"resource"  form:"resource" query:"resource"`
	PermissionName string `json:"permission_name"  form:"permission_name" query:"permission_name"`
	ResourceType   int32  `json:"resource_type"  form:"resource_type" query:"resource_type"`
	Visibility     int32  `json:"visibility"  form:"visibility" query:"visibility"`
}
type CreatePermissionGroupReq struct {
	Remark       string `json:"remark"  form:"remark" query:"remark"`
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	GroupName    string `json:"group_name"  form:"group_name" query:"group_name"`
	GroupType    int32  `json:"group_type"  form:"group_type" query:"group_type"`
}
type CreatePermissionGroupPermissionReq struct {
	PermissionPk      string `json:"permission_pk"  form:"permission_pk" query:"permission_pk"`
	PermissionGroupPk string `json:"permission_group_pk"  form:"permission_group_pk" query:"permission_group_pk"`
}
type UpdaterPermissionGroupReq struct {
	Pk        string `json:"pk"  form:"pk" query:"pk"`
	GroupName string `json:"group_name"  form:"group_name" query:"group_name"`
	Remark    string `json:"remark"  form:"remark" query:"remark"`
	GroupType int32  `json:"group_type"  form:"group_type" query:"group_type"`
}

type CreatePageResourceReq struct {
	PagePath      string `json:"page_path"  form:"page_path" query:"page_path"`
	Component     string `json:"component"  form:"component" query:"component"`
	ComponentName string `json:"component_name"  form:"component_name" query:"component_name"`
	PageName      string `json:"page_name"  form:"page_name" query:"page_name"`
	IsCache       int32  `json:"is_cache"  form:"is_cache" query:"is_cache"`
	PageType      int32  `json:"page_type"  form:"page_type" query:"page_type"`
}
type UpdatePageResourceReq struct {
	Pk            string `json:"pk"  form:"pk" query:"pk"`
	PagePath      string `json:"page_path"  form:"page_path" query:"page_path"`
	Component     string `json:"component"  form:"component" query:"component"`
	ComponentName string `json:"component_name"  form:"component_name" query:"component_name"`
	PageName      string `json:"page_name"  form:"page_name" query:"page_name"`
	IsCache       int32  `json:"is_cache"  form:"is_cache" query:"is_cache"`
	PageType      int32  `json:"page_type"  form:"page_type" query:"page_type"`
}
type PageResourceQueryBody struct {
	PagePath      string `json:"page_path"  form:"page_path" query:"page_path"`
	Component     string `json:"component"  form:"component" query:"component"`
	ComponentName string `json:"component_name"  form:"component_name" query:"component_name"`
	IsCache       string `json:"is_cache"  form:"is_cache" query:"is_cache"`
	PageName      string `json:"page_name"  form:"page_name" query:"page_name"`
	PageType      int32  `json:"page_type"  form:"page_type" query:"page_type"`
	PageQfqz
}
type CreatePageInterfaceReq struct {
	PagePk        string `json:"page_pk"  form:"page_pk" query:"page_pk"`
	InterfacePk   string `json:"interface_pk"  form:"interface_pk" query:"interface_pk"`
	OperationType int32  `json:"operation_type"  form:"operation_type" query:"operation_type"`
}
type CreateInterfaceResourceReq struct {
	InterfaceKey  string `json:"interface_key"  form:"interface_key" query:"interface_key"`
	InterfaceName string `json:"interface_name"  form:"interface_name" query:"interface_name"`
	InterfaceWay  string `json:"interface_way"  form:"interface_way" query:"interface_way"`
	InterfaceUrl  string `json:"interface_url"  form:"interface_url" query:"interface_url"`
}
type UpdateInterfaceResourceReq struct {
	Pk            string `json:"pk"  form:"pk" query:"pk"`
	InterfaceKey  string `json:"interface_key"  form:"interface_key" query:"interface_key"`
	InterfaceName string `json:"interface_name"  form:"interface_name" query:"interface_name"`
	InterfaceWay  string `json:"interface_way"  form:"interface_way" query:"interface_way"`
	InterfaceUrl  string `json:"interface_url"  form:"interface_url" query:"interface_url"`
}
type InterfaceResourceQueryBody struct {
	InterfaceKey  string `json:"interface_key"  form:"interface_key" query:"interface_key"`
	InterfaceName string `json:"interface_name"  form:"interface_name" query:"interface_name"`
	InterfaceWay  string `json:"interface_way"  form:"interface_way" query:"interface_way"`
	InterfaceUrl  string `json:"interface_url"  form:"interface_url" query:"interface_url"`
	PageQfqz
}
type CreateMenuResourceReq struct {
	MenuType    int32  `json:"menu_type"  form:"menu_type" query:"menu_type"`
	MenuName    string `json:"menu_name"  form:"menu_name" query:"menu_name"`
	ResourceKey string `json:"resource_key"  form:"resource_key" query:"resource_key"`
	ParentPk    string `json:"parent_pk"  form:"parent_pk" query:"parent_pk"`
	Sort        int32  `json:"sort"  form:"sort" query:"sort"`
	Icon        string `json:"icon"  form:"icon" query:"icon"`
	Path        string `json:"path"  form:"path" query:"path"`
}
type UpdateMenuResourceReq struct {
	Pk          string `json:"pk"  form:"pk" query:"pk"`
	MenuType    int32  `json:"menu_type"  form:"menu_type" query:"menu_type"`
	MenuName    string `json:"menu_name"  form:"menu_name" query:"menu_name"`
	ResourceKey string `json:"resource_key"  form:"resource_key" query:"resource_key"`
	ParentPk    string `json:"parent_pk"  form:"parent_pk" query:"parent_pk"`
	Sort        int32  `json:"sort"  form:"sort" query:"sort"`
	Icon        string `json:"icon"  form:"icon" query:"icon"`
	Path        string `json:"path"  form:"path" query:"path"`
}
type MenuResourceQueryBody struct {
	MenuName string `json:"menu_name"  form:"menu_name" query:"menu_name"`
}
type CreateRolePermissionReq struct {
	RolePk         string `json:"role_pk"  form:"role_pk" query:"role_pk"`
	PermissionType int32  `json:"permission_type"  form:"permission_type" query:"permission_type"`
	PermissionPk   string `json:"permission_pk"  form:"permission_pk" query:"permission_pk"`
}

type PageQfqz struct {
	Page int32 `json:"page" form:"page" query:"page"`
	Size int32 `json:"size" form:"size" query:"size"`
}
