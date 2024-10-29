import request from '@/utils/request'

/** 获取保安公司列表 */
export function getSecurityEnterpriseList(params: any) {
  return request({
    url: '/core/enterprise/security',
    method: 'get',
    params
  })
}

/** 新增保安公司 */
export function createSecurityEnterprise(data: any) {
  return request({
    url: '/core/enterprise/security',
    method: 'post',
    data
  })
}

/** 修改保安公司 */
export function updateSecurityEnterprise(data: any) {
  return request({
    url: '/core/enterprise/security',
    method: 'put',
    data
  })
}
