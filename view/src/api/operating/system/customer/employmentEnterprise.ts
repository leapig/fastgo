import request from '@/utils/request'

/** 获取业主单位列表 */
export function getEmploymentEnterpriseList(params: any) {
  return request({
    url: '/core/enterprise/employment',
    method: 'get',
    params
  })
}

/** 新增业主单位 */
export function createEmploymentEnterprise(data: any) {
  return request({
    url: '/core/enterprise/employment',
    method: 'post',
    data
  })
}

/** 修改业主单位 */
export function updateEmploymentEnterprise(data: any) {
  return request({
    url: '/core/enterprise/employment',
    method: 'put',
    data
  })
}
