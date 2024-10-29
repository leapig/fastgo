import request from '@/utils/request'

/** 获取角色组列表 */
export function getRoleGroupList(params: any) {
  return request({
    url: '/core/rbac/system/role/group',
    method: 'get',
    params
  })
}

/** 新增角色组 */
export function createRoleGroup(data: any) {
  return request({
    url: '/core/rbac/system/role/group',
    method: 'post',
    data
  })
}

/** 修改角色组 */
export function updateRoleGroup(data: any) {
  return request({
    url: '/core/rbac/system/role/group',
    method: 'put',
    data
  })
}

/** 删除角色组 */
export function removeRoleGroup(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/role/group',
    method: 'delete',
    data
  })
}

/** 角色组添加角色/权限 */
export function addRoleGroupPermission(data: any) {
  return request({
    url: '/core/rbac/system/role/group/permission',
    method: 'post',
    data
  })
}

/** 角色组移除角色/权限 */
export function removeRoleGroupPermission(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/role/group/permission',
    method: 'delete',
    data
  })
}

/** 获取角色组人员列表 */
export function getRoleGroupUserList(params: any) {
  return request({
    url: '/core/rbac/system/role/group/user',
    method: 'get',
    params
  })
}
