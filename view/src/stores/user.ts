import router from '../router'
import {
  setUserAccessToken,
  setUserRefreshToken,
  setTenantAccessToken,
  setTenantRefreshToken,
  getUserAccessToken,
  getUserRefreshToken,
  getTenantAccessToken,
  getTenantRefreshToken,
  removeUserAccessToken,
  removeUserRefreshToken,
  removeTenantAccessToken,
  removeTenantRefreshToken
} from '@/utils/auth'
import {
  getWechatLoginQrcodeApi,
  getLoginStatusApi,
  accessTenantApi,
  getUserInfoApi,
  refreshUserTokenApi,
  refreshTenantTokenApi
} from '@/api/login'
import cache from '@/plugins/cache'

export type LoginEnterprise = {
  pk?: string
  name?: string
  type?: number
}

const tenantUserCacheKey = 'tenantUserCacheKey' //租户下用户信息
const userCacheKey = 'UserCacheKey' //用户基本信息
const enterpriseCacheKey = 'EnterpriseCacheKey'
const wechatLoginInfo = { qrcode: '', key: '' }
let wechatLoginTimer: number | undefined = undefined

const useUserStore = defineStore('user', {
  state: () => ({
    userTokenRefreshing: false,
    tenantTokenRefreshing: false,
    user: cache.local.getJSON(userCacheKey) || {}, // 当前用户
    enterprise: cache.local.getJSON(enterpriseCacheKey) || {}, // 当前单位(租户)
    tenantUserInfo:cache.local.getJSON(tenantUserCacheKey),//当前租户下的用户信息
  }),
  actions: {
    /**
     * 获取微信登录二维码
     * @returns
     */
    startWechatLogin() {
      return new Promise((resolve: Function, reject: Function) => {
        if (!wechatLoginInfo.qrcode) {
          getWechatLoginQrcodeApi()
            .then((response: any) => {
              const qrcode = 'data:image/jpeg;base64,' + response.data.qrCode
              const res = {
                qrcode,
                key: response.data.key
              }
              wechatLoginInfo.qrcode = qrcode
              wechatLoginInfo.key = response.data.key
              startGetLoginStatus()
              resolve(res)
            })
            .catch((error) => {
              reject(error)
            })
        } else {
          const res = {
            qrcode: wechatLoginInfo.qrcode,
            key: wechatLoginInfo.key
          }
          resolve(res)
        }
      })
    },
    /**
     * 刷新微信登录二维码
     * @returns
     */
    refreshWechatLogin() {
      this.stopWechatLogin()
      return this.startWechatLogin()
    },
    /**
     * 停止微信登录二维码
     */
    stopWechatLogin() {
      stopGetLoginStatus()
    },
    /**
     * 获取登录用户信息 （有无租户共有）
     */
    getUserInfo() {
      return new Promise((resolve: Function, reject: Function) => {
        if (this.user.pk) {
          resolve(this.user)
        } else {
          getUserInfo()
            .then((response: any) => {
              resolve(response)
            })
            .catch((error: any) => {
              reject(error)
            })
        }
      })
    },
    /**
     * 登录到租户
     * @param {LoginEnterprise} enterprise 租户信息
     * @returns
     */
    accessToTenant(enterprise: LoginEnterprise) {
      return new Promise((resolve: Function, reject: Function) => {
        accessTenantApi({ enterprise_pk: enterprise.pk })
          .then((response) => {
            setTenantAccessToken(response.data.TenantAccessToken, response.data.TenantAccessTokenExpireIn)
            setTenantRefreshToken(response.data.TenantRefreshToken)
            this.setEnterprise(enterprise)
            resolve()
          })
          .catch((error: any) => {
            reject(error)
          })
      })
    },
    /**
     * 清除用户登录信息
     */
    clearUserAccess() {
      removeUserAccessToken()
      removeUserRefreshToken()
      this.setUser({})
      this.clearTenantAccess()
    },
    /**
     * 清除租户登录信息
     */
    clearTenantAccess() {
      removeTenantAccessToken()
      removeTenantRefreshToken()
      this.setEnterprise({})
    },
    /**
     * 退出到选择租户页面
     */
    goSelectTenant() {
      ElMessageBox.confirm('确定要返回选择租户界面吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.clearTenantAccess()
          location.replace('/index')
        })
        .catch(() => {})
    },
    /**
     * 退出登录
     */
    logout() {
      ElMessageBox.confirm('确定注销并退出系统吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.clearUserAccess()
          location.replace('/login')
        })
        .catch(() => {})
    },
    /**
     * 获取用户短token(自动刷新)
     */
    getUserToken() {
      return new Promise((resolve: Function) => {
        const userToken = getUserAccessToken()
        const refreshToken = getUserRefreshToken()
        if (userToken) {
          resolve(userToken)
        } else if (this.userTokenRefreshing) {
          let times = 100
          const timer = setInterval(() => {
            const newUserToken = getUserAccessToken()
            if (!this.userTokenRefreshing && newUserToken) {
              resolve(newUserToken)
              clearInterval(timer)
            } else if (times > 0) {
              times--
            } else {
              resolve('')
              clearInterval(timer)
            }
          }, 300)
        } else if (refreshToken) {
          this.userTokenRefreshing = true
          refreshUserTokenApi(refreshToken)
            .then((response) => {
              setUserAccessToken(response.data.UserAccessToken, response.data.ExpireIn)
              setUserRefreshToken(response.data.UserRefreshToken)
              this.userTokenRefreshing = false
              resolve(response.data.UserAccessToken)
            })
            .catch((error: any) => {
              console.error(error)
              this.userTokenRefreshing = false
              resolve('')
            })
        } else {
          resolve('')
        }
      })
    },
    /**
     * 获取租户短token(自动刷新)
     */
    getTenantToken() {
      return new Promise((resolve: Function) => {
        const tenantToken = getTenantAccessToken()
        const refreshToken = getTenantRefreshToken()
        if (tenantToken) {
          resolve(tenantToken)
        } else if (this.tenantTokenRefreshing) {
          let times = 100
          const timer = setInterval(() => {
            const newTenantToken = getTenantAccessToken()
            if (!this.tenantTokenRefreshing && newTenantToken) {
              resolve(newTenantToken)
              clearInterval(timer)
            } else if (times > 0) {
              times--
            } else {
              resolve('')
              clearInterval(timer)
            }
          }, 300)
        } else if (refreshToken) {
          this.tenantTokenRefreshing = true
          refreshTenantTokenApi(refreshToken)
            .then((response) => {
              setTenantAccessToken(response.data.TenantAccessToken, response.data.ExpireIn)
              setTenantRefreshToken(response.data.TenantRefreshToken)
              this.tenantTokenRefreshing = false
              resolve(response.data.TenantAccessToken)
            })
            .catch((error: any) => {
              console.error(error)
              this.tenantTokenRefreshing = false
              resolve('')
            })
        } else {
          resolve('')
        }
      })
    },
    /**
     * 设置当前用户
     */
    setUser(user: any) {
      if (user.pk) {
        this.user = user
        cache.local.setJSON(userCacheKey, user)
      } else {
        this.user = {}
        cache.local.remove(userCacheKey)
      }
    },
    /**
     * 设置当前单位(租户)
     */
    setEnterprise(enterprise: LoginEnterprise) {
      if (enterprise.pk) {
        this.enterprise = enterprise
        cache.local.setJSON(enterpriseCacheKey, enterprise)
      } else {
        this.enterprise = {}
        cache.local.remove(enterpriseCacheKey)
      }
    },
    /**设置当前租户下用户信息 */
    setTenantUser(tenant:any){
      if (tenant.pk) {
        this.tenantUserInfo = tenant
        cache.local.setJSON(tenantUserCacheKey, tenant)
      } else {
        this.tenantUserInfo = {}
        cache.local.remove(tenantUserCacheKey)
      }
    }
  }
})

/* 获取租户下用户专属信息 */
function getUserInfo(){
  return new Promise((resolve: Function, reject: Function) => {
    getUserInfoApi()
      .then((response: any) => {
        resolve(response.data)
      })
      .catch((error: any) => {
        reject(error)
      })
  })
}

// 循环获取登录状态
function startGetLoginStatus() {
  wechatLoginTimer = setInterval(() => {
    getLoginStatusApi({
      key: wechatLoginInfo.key
    })
      .then((response: any) => {
        if (response.data.UserAccessToken) {
          setUserAccessToken(response.data.UserAccessToken, response.data.ExpireIn)
          setUserRefreshToken(response.data.UserRefreshToken)
          ElMessage.success('登录成功')
          stopGetLoginStatus()
          router.replace('/index')
        }
      })
      .catch((error: any) => {
        if (error.errorMessage !== '未登录') {
          console.log('获取扫码登录结果服务异常', error)
        }
      })
  }, 5000)
}

// 停止循环获取登录状态
function stopGetLoginStatus() {
  wechatLoginInfo.qrcode = ''
  wechatLoginInfo.key = ''
  if (wechatLoginTimer) {
    clearInterval(wechatLoginTimer)
    wechatLoginTimer = undefined
  }
}

export default useUserStore
