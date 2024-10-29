import request from '@/utils/request'

/** 修改单位负责人 */
export function updateEnterpriseCorporateApi(data: any) {
  return request({
    url: '/core/enterprise/corporate',
    method: 'put',
    data
  })
}
