import request from '@/utils/request'

/** 获取用户列表 */
export function getUserList(params: any) {
  return request({
    url: '/core/user/list',
    method: 'get',
    params
  })
}