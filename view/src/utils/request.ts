import axios from 'axios'
import errorCode from '@/utils/errorCode'
import { tansParams } from '@/utils/common'
import cache from '@/plugins/cache'
import { saveAs } from 'file-saver'
import useUserStore from '@/stores/user'
import useSettingsStore from '@/stores/settings'

import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style

NProgress.configure({
  showSpinner: false
})

axios.defaults.headers['Content-Type'] = 'application/json;charset=utf-8'
// 创建axios实例
const service = axios.create({
  // axios中请求配置有baseURL选项，表示请求URL公共部分
  baseURL: import.meta.env.VITE_APP_BASE_API,
  // 超时
  timeout: 30000
})
// request拦截器
service.interceptors.request.use(
  async (config) => {
    if (!useSettingsStore().isSilentRequestUrl(config.url)) {
      NProgress.start()
    }
    // 是否需要防止数据重复提交
    const isRepeatSubmit = (config.headers || {}).repeatSubmit === false
    // 是否使用自定义token
    const useCustomToken = (config.headers || {}).useCustomToken === 'true'
    if (!useCustomToken) {
      // 授权token
      const token = await useUserStore().getUserToken()
      if (token) {
        config.headers['Authorization'] = `Bearer ${token}`
      }
    }
    // get请求映射params参数
    if (config.method === 'get' && config.params) {
      let url = config.url + '?' + tansParams(config.params)
      url = url.slice(0, -1)
      config.params = {}
      config.url = url
    }
    if (!isRepeatSubmit && (config.method === 'post' || config.method === 'put')) {
      const requestObj = {
        url: config.url,
        data: typeof config.data === 'object' ? JSON.stringify(config.data) : config.data,
        time: new Date().getTime()
      }
      const requestSize = Object.keys(JSON.stringify(requestObj)).length // 请求数据大小
      const limitSize = 5 * 1024 * 1024 // 限制存放数据5M
      if (requestSize >= limitSize) {
        console.warn(`[${config.url}]: ` + '请求数据大小超出允许的5M限制，无法进行防重复提交验证。')
        return config
      }
      const sessionObj = cache.session.getJSON('sessionObj')
      if (sessionObj === undefined || sessionObj === null || sessionObj === '') {
        cache.session.setJSON('sessionObj', requestObj)
      } else {
        const s_url = sessionObj.url // 请求地址
        const s_data = sessionObj.data // 请求数据
        const s_time = sessionObj.time // 请求时间
        const interval = 1000 // 间隔时间(ms)，小于此时间视为重复提交
        if (s_data === requestObj.data && requestObj.time - s_time < interval && s_url === requestObj.url && !config.headers.isRefreshToken) {
          const message = '数据正在处理，请勿重复提交'
          console.warn(`[${s_url}]: ` + message)
          return Promise.reject(new Error(message))
        } else {
          cache.session.setJSON('sessionObj', requestObj)
        }
      }
    }
    return config
  },
  (error) => {
    console.error(error)
    Promise.reject(error)
  }
)
// 响应拦截器
service.interceptors.response.use(
  (res) => {
    if (!useSettingsStore().isSilentRequestUrl(res.config.url)) {
      NProgress.done()
    }
    // 请求成功状态
    const success = res.data.success || false
    // 未设置状态码则默认成功状态
    const code = res.data.errorCode + '' || '0'
    // 获取错误信息
    const msg = res.data.errorMessage || errorCode[code] || errorCode['default']
    // 二进制数据则直接返回
    if (res.request.responseType === 'blob' || res.request.responseType === 'arraybuffer') {
      return res.data
    }
    if (code !== '0' || !success) {
      if (!useSettingsStore().isSilentRequestUrl(res.config.url)) {
        ElMessage({ message: msg, type: 'error' })
      }
      return Promise.reject('error')
    } else {
      return Promise.resolve(res.data)
    }
  },
  (error) => {
    if (!useSettingsStore().isSilentRequestUrl(error.config.url)) {
      NProgress.done()
      const { message } = error
      if (error.response.status === 400) {
        const msg = error.response.data.errorMessage || '接口请求失败，请稍后再试'
        ElMessage({ message: msg, type: 'error', duration: 5 * 1000 })
        return Promise.reject(error.response.data)
      } else if (error.response.status === 403) {
        ElMessageBox.alert('登录状态已过期，请重新登录', '系统提示', {
          confirmButtonText: '确定',
          showClose: false,
          type: 'warning'
        }).then(() => {
          location.href = '/login'
        })
        return Promise.reject('无效的会话，或者会话已过期，请重新登录。')
      } else if (message === 'Network Error') {
        ElMessage({ message: '后端接口连接异常', type: 'error', duration: 5 * 1000 })
      } else if (message.includes('timeout')) {
        ElMessage({ message: '系统接口请求超时', type: 'error', duration: 5 * 1000 })
      } else if (message.includes('Request failed with status code')) {
        ElMessage({ message: `系统接口${message.substr(message.length - 3)}异常`, type: 'error', duration: 5 * 1000 })
      }
    } else {
      if (error.response.status === 400) {
        return Promise.reject(error.response.data)
      } else if (error.response.status === 403) {
        ElMessageBox.alert('登录状态已过期，请重新登录', '系统提示', {
          confirmButtonText: '确定',
          showClose: false,
          type: 'warning'
        }).then(() => {
          location.href = '/login'
        })
        return Promise.reject('无效的会话，或者会话已过期，请重新登录。')
      }
    }
    return Promise.reject(error)
  }
)

export default service

// 通用下载方法
export function download(url: string, params: object, filename: string, config: object) {
  return service
    .post(url, params, {
      transformRequest: [
        (params) => {
          return tansParams(params)
        }
      ],
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      responseType: 'blob',
      timeout: 180000, // 30分钟超时
      onDownloadProgress: (progressEvent) => {
        if (progressEvent.total) {
          const percent = Math.round((progressEvent.loaded * 100) / progressEvent.total)
          NProgress.set(percent / 100)
        }
      },
      ...config
    })
    .then(async (res) => {
      const isBlob = res.data.type !== 'application/json'
      if (isBlob) {
        const blob = new Blob([res.data])
        saveAs(blob, filename)
      } else {
        const resText = await res.data.text()
        const rspObj = JSON.parse(resText)
        const errMsg = errorCode[rspObj.code] || rspObj.msg || errorCode['default']
        ElMessage.error(errMsg)
      }
    })
    .catch((r) => {
      console.error(r)
      ElMessage.error('下载文件出现错误，请联系管理员！')
    })
}

