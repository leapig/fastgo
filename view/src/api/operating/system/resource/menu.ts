import request from '@/utils/request'

/** 获取菜单列表 */
export function getMenuListApi(params: any) {
  return request({
    url: '/core/rbac/resource/menu',
    method: 'get',
    params
  })
}

/** 新增菜单 */
export function createMenuApi(data: any) {
  return request({
    url: '/core/rbac/resource/menu',
    method: 'post',
    data
  })
}

/** 修改菜单 */
export function updateMenuApi(data: any) {
  return request({
    url: '/core/rbac/resource/menu',
    method: 'put',
    data
  })
}

/** 删除菜单 */
export function removeMenuApi(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/resource/menu',
    method: 'delete',
    data
  })
}