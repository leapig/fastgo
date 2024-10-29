import request from '@/utils/request'

/** 获取页面列表 */
export function getPageListApi(params: any) {
  return request({
    url: '/core/rbac/resource/page',
    method: 'get',
    params
  })
}

/** 新增页面 */
export function createPageApi(data: any) {
  return request({
    url: '/core/rbac/resource/page',
    method: 'post',
    data
  })
}

/** 修改页面 */
export function updatePageApi(data: any) {
  return request({
    url: '/core/rbac/resource/page',
    method: 'put',
    data
  })
}

/** 删除页面 */
export function removePageApi(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/resource/page',
    method: 'delete',
    data
  })
}

/** 页面接口关联 */
export function linkPageInterfaceApi(data: any) {
  return request({
    url: '/core/rbac/resource/page/interface',
    method: 'post',
    data
  })
}

/** 删除接口关联 */
export function removePageInterfaceApi(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/resource/page/interface',
    method: 'delete',
    data
  })
}
