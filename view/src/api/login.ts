import request from '@/utils/request'
import tenantRequest from '@/utils/tenantRequest'

/** 获取微信登录二维码 */
export function getWechatLoginQrcodeApi() {
  return request({
    url: '/core/jwt/wechat/qrcode',
    method: 'get'
  })
}

/** 获取登录状态 */
export function getLoginStatusApi(params: any) {
  return request({
    url: '/core/jwt/token/user_token',
    method: 'get',
    params
  })
}

/** 获取可登录租户列表 */
export function getTenantListApi() {
  return request({
    url: '/core/jwt/tenant/list',
    method: 'get'
  })
}

/** 登录租户 */
export function accessTenantApi(data: any) {
  return request({
    url: '/core/jwt/token/tenant_token',
    method: 'post',
    data
  })
}

/** 刷新用户短token */
export function refreshUserTokenApi(refreshToken: string) {
  return request({
    url: '/core/jwt/token/user_token',
    method: 'put',
    headers: {
      Authorization: `Bearer ${refreshToken}`,
      useCustomToken: 'true'
    }
  })
}

/** 刷新租户token */
export function refreshTenantTokenApi(refreshToken: string) {
  return tenantRequest({
    url: '/core/jwt/token/tenant_token',
    method: 'put',
    headers: {
      Authorization: `Bearer ${refreshToken}`,
      useCustomToken: 'true'
    }
  })
}

/** 获取登录用户信息 */
export function getUserInfoApi() {
  return tenantRequest({
    url: '/core/organization/user/info',
    method: 'get'
  })
}