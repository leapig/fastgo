import request from '@/utils/request'

/** 获取监管单位列表 */
export function getSuperviseEnterpriseList(params: any) {
  return request({
    url: '/core/enterprise/supervise',
    method: 'get',
    params
  })
}

/** 新增监管单位 */
export function createSuperviseEnterprise(data: any) {
  return request({
    url: '/core/enterprise/supervise',
    method: 'post',
    data
  })
}

/** 修改监管单位 */
export function updateSuperviseEnterprise(data: any) {
  return request({
    url: '/core/enterprise/supervise',
    method: 'put',
    data
  })
}


/** 获取监管单位辖区权限 */
export function getSuperviseAreaApi(params: any) {
  return request({
    url: '/core/enterprise/area/permission',
    method: 'get',
    params
  })
}

/** 新增监管单位辖区权限 */
export function addSuperviseAreaApi(data: any) {
  return request({
    url: '/core/enterprise/area/permission',
    method: 'post',
    data
  })
}
/** 删除监管单位辖区权限 */
export function deleteSuperviseAreaApi(data: any) {
  return request({
    url: '/core/enterprise/area/permission',
    method: 'delete',
    data
  })
}