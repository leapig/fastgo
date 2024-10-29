import request from '@/utils/request'

/** 获取用户账号信息列表 */
export function getUserAccountList(params: any) {
  return request({
    url: '/core/user/account',
    method: 'get',
    params
  })
}

/** 新增用户账号 */
export function createUser(data: any) {
  return request({
    url: '/core/user/account',
    method: 'post',
    data
  })
}

/** 用户手机号更换 */
export function changeUserPhone(data: any) {
  return request({
    url: '/core/user/phone',
    method: 'put',
    data
  })
}

/** 用户第三方信息解绑 */
export function unbindUserClient(data: any) {
  return request({
    url: '/core/user/client',
    method: 'put',
    data
  })
}

/** 获取用户角色/角色组信息列表 */
export function getUserRoleGroupList(params: any) {
  return request({
    url: '/core/rbac/system/user/role/group',
    method: 'get',
    params
  })
}

/** 用户添加角色/角色组 */
export function addUserPermission(data: any) {
  return request({
    url: '/core/rbac/system/user/permission',
    method: 'post',
    data
  })
}

/** 用户移除角色/角色组 */
export function removeUserPermission(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/user/permission',
    method: 'delete',
    data
  })
}

/** 获取用户权限信息列表 */
export function getUserPermissionList(params: any) {
  return request({
    url: '/core/rbac/system/user/permission',
    method: 'get',
    params
  })
}

