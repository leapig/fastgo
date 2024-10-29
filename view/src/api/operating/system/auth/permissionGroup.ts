import request from '@/utils/request'

/** 获取权限组列表 */
export function getPermissionGroupList(params: any) {
  return request({
    url: '/core/rbac/system/permission/group',
    method: 'get',
    params
  })
}

/** 新增权限组 */
export function createPermissionGroup(data: any) {
  return request({
    url: '/core/rbac/system/permission/group',
    method: 'post',
    data
  })
}

/** 修改权限组 */
export function updatePermissionGroup(data: any) {
  return request({
    url: '/core/rbac/system/permission/group',
    method: 'put',
    data
  })
}

/** 删除权限组 */
export function removePermissionGroup(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/permission/group',
    method: 'delete',
    data
  })
}

/** 权限组添加权限 */
export function addPermissionGroupPermission(data: any) {
  return request({
    url: '/core/rbac/system/permission/group/permission',
    method: 'post',
    data
  })
}

/** 权限组移除权限 */
export function removePermissionGroupPermission(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/permission/group/permission',
    method: 'delete',
    data
  })
}
