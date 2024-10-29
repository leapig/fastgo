import request from '@/utils/request'

/** 获取角色列表 */
export function getRoleList(params: any) {
  return request({
    url: '/core/rbac/system/role',
    method: 'get',
    params
  })
}

/** 新增角色 */
export function createRole(data: any) {
  return request({
    url: '/core/rbac/system/role',
    method: 'post',
    data
  })
}

/** 修改角色 */
export function updateRole(data: any) {
  return request({
    url: '/core/rbac/system/role',
    method: 'put',
    data
  })
}

/** 删除角色 */
export function removeRole(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/role',
    method: 'delete',
    data
  })
}

/** 角色添加权限 */
export function addRolePermission(data: any) {
  return request({
    url: '/core/rbac/system/role/permission',
    method: 'post',
    data
  })
}

/** 角色移除权限 */
export function removeRolePermission(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/role/permission',
    method: 'delete',
    data
  })
}

/** 获取角色人员列表 */
export function getRoleUserList(params: any) {
  return request({
    url: '/core/rbac/system/role/user',
    method: 'get',
    params
  })
}
