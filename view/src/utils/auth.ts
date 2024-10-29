import Cookies from 'js-cookie'

const userAccessTokenKey = 'UserAccessToken' // 用户短token
const userAccessTokenExpiresKey = 'UserAccessTokenExpires' // 用户短token有效期
const userRefreshTokenKey = 'UserRefreshToken' // 用户刷新token
const userRefreshTokenExpiresKey = 'UserRefreshTokenExpires' // 用户刷新token有效期

const tenantAccessTokenKey= 'TenantAccessToken' // 租户短token
const tenantAccessTokenExpiresKey = 'TenantAccessTokenExpires' // 租户短token有效期
const tenantRefreshTokenKey = 'TenantRefreshToken' // 租户刷新token
const tenantRefreshTokenExpiresKey = 'TenantRefreshTokenExpires' // 租户刷新token有效期

/**
 * 获取用户短token
 * @returns 
 */
export function getUserAccessToken() {
  const t = Cookies.get(userAccessTokenExpiresKey)
  if (!t || parseInt(t) > new Date().getTime()) {
    return Cookies.get(userAccessTokenKey)
  } else {
    return ''
  }
}
/**
 * 设置用户短token
 * @param userAccessToken 用户短token
 * @param expireIn 有效时间,秒s
 * @returns 
 */
export function setUserAccessToken(userAccessToken: string, expireIn?: number) {
  if (expireIn && expireIn > 30) {
    const t = new Date().getTime() + (expireIn - 30) * 1000
    Cookies.set(userAccessTokenExpiresKey, t + '')
    const expires = new Date(t)
    return Cookies.set(userAccessTokenKey, userAccessToken, { expires })
  } else {
    return Cookies.set(userAccessTokenKey, userAccessToken)
  }
}
/**
 * 移除用户短token
 * @returns 
 */
export function removeUserAccessToken() {
  return Cookies.remove(userAccessTokenKey)
}

/**
 * 获取用户刷新token
 * @returns 
 */
export function getUserRefreshToken() {
  const t = Cookies.get(userRefreshTokenExpiresKey)
  if (!t || parseInt(t) > new Date().getTime()) {
    return Cookies.get(userRefreshTokenKey)
  } else {
    return ''
  }
}
/**
 * 设置用户刷新token
 * @param userRefreshToken 用户刷新token
 * @param expireIn 有效时间,秒s
 * @returns 
 */
export function setUserRefreshToken(userRefreshToken: string, expireIn?: number) {
  if (expireIn && expireIn > 30) {
    const t = new Date().getTime() + (expireIn - 30) * 1000
    Cookies.set(userRefreshTokenExpiresKey, t + '')
    const expires = new Date(t)
    return Cookies.set(userRefreshTokenKey, userRefreshToken, { expires })
  } else {
    return Cookies.set(userRefreshTokenKey, userRefreshToken)
  }
}
/**
 * 移除用户刷新token
 * @returns 
 */
export function removeUserRefreshToken() {
  return Cookies.remove(userRefreshTokenKey)
}

/**
 * 获取租户短token
 * @returns 
 */
export function getTenantAccessToken() {
  const t = Cookies.get(tenantAccessTokenExpiresKey)
  if (!t || parseInt(t) > new Date().getTime()) {
    return Cookies.get(tenantAccessTokenKey)
  } else {
    return ''
  }
}
/**
 * 设置租户短token
 * @param tenantAccessToken 租户短token
 * @param expireIn 有效时间,秒s
 * @returns 
 */
export function setTenantAccessToken(tenantAccessToken: string, expireIn?: number) {
  if (expireIn && expireIn > 30) {
    const t = new Date().getTime() + (expireIn - 30) * 1000
    Cookies.set(tenantAccessTokenExpiresKey, t + '')
    const expires = new Date(t)
    return Cookies.set(tenantAccessTokenKey, tenantAccessToken, { expires })
  } else {
    return Cookies.set(tenantAccessTokenKey, tenantAccessToken)
  }
}
/**
 * 移除租户短token
 * @returns 
 */
export function removeTenantAccessToken() {
  return Cookies.remove(tenantAccessTokenKey)
}

/**
 * 获取租户刷新token
 * @returns 
 */
export function getTenantRefreshToken() {
  const t = Cookies.get(tenantRefreshTokenExpiresKey)
  if (!t || parseInt(t) > new Date().getTime()) {
    return Cookies.get(tenantRefreshTokenKey)
  } else {
    return ''
  }
}
/**
 * 设置租户刷新token
 * @param tenantRefreshToken 租户刷新token
 * @param expireIn 有效时间,秒s
 * @returns 
 */
export function setTenantRefreshToken(tenantRefreshToken: string, expireIn?: number) {
  if (expireIn && expireIn > 30) {
    const t = new Date().getTime() + (expireIn - 30) * 1000
    Cookies.set(tenantRefreshTokenExpiresKey, t + '')
    const expires = new Date(t)
    return Cookies.set(tenantRefreshTokenKey, tenantRefreshToken, { expires })
  } else {
    return Cookies.set(tenantRefreshTokenKey, tenantRefreshToken)
  }
}
/**
 * 移除租户刷新token
 * @returns 
 */
export function removeTenantRefreshToken() {
  return Cookies.remove(tenantRefreshTokenKey)
}
