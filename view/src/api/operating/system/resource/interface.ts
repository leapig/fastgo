import request from '@/utils/request'

/** 获取接口列表 */
export function getInterfaceListApi(params: any) {
  return request({
    url: '/core/rbac/resource/interface',
    method: 'get',
    params
  })
}

/** 新增接口 */
export function createInterfaceApi(data: any) {
  return request({
    url: '/core/rbac/resource/interface',
    method: 'post',
    data
  })
}

/** 修改接口 */
export function updateInterfaceApi(data: any) {
  return request({
    url: '/core/rbac/resource/interface',
    method: 'put',
    data
  })
}

/** 删除接口 */
export function removeInterfaceApi(pk: string) {
  const data = { pk }
  return request({
    url: '/core/rbac/resource/interface',
    method: 'delete',
    data
  })
}
