import tenantRequest from '@/utils/tenantRequest'

/** 获取登录权限菜单 */
export function getMenusApi() {
  return tenantRequest({
    url: '/core/jwt/tenant/menu',
    method: 'get'
  })
}

/** 获取登录权限页面 */
export function getPagesApi() {
  return tenantRequest({
    url: '/core/jwt/tenant/page',
    method: 'get'
  })
}