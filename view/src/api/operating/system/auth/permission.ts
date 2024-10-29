import request from '@/utils/request'

/** 获取权限列表 */
export function getPermissionList(params: any) {
  return request({
    url: '/core/rbac/system/permission',
    method: 'get',
    params
  })
}

/** 新增权限 */
export function createPermission(data: any) {
  return request({
    url: '/core/rbac/system/permission',
    method: 'post',
    data
  })
}

/** 修改权限 */
export function updatePermission(data: any) {
  return request({
    url: '/core/rbac/system/permission',
    method: 'put',
    data
  })
}

/** 删除权限 */
export function removePermission(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/system/permission',
    method: 'delete',
    data
  })
}