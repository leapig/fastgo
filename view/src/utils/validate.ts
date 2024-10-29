import { validateIdCard } from './idcard'

/**
 * 判断url是否是http或https
 *
 * @param {string} url
 * @returns {boolean}
 */
export function isHttp(url: string) {
  return url.indexOf('http://') !== -1 || url.indexOf('https://') !== -1
}

/**
 * 判断url是否为外链
 *
 * @param {string} url
 * @returns {boolean}
 */
export function isExternal(url: string) {
  const e = /^(https?:|mailto:|tel:)/
  return e.test(url)
}

/**
 * 判断字符串是否为Base64图片
 * @param {string} str 
 * @returns {boolean}
 */
export function isBase64Image(str: string) {
  const e = /^data:image\/([a-z]+);base64,/
  return e.test(str)
}

/**
 * 判断页面路径是否存在
 * @param {string} viewPath 页面路径
 * @returns {boolean}
 */
export function existView(viewPath: string) {
  const modules = import.meta.glob('./../views/**/*.vue')
  let exist = false
  for (const path in modules) {
    const dir = path.split('views/')[1].split('.vue')[0]
    if (dir === viewPath) {
      exist = true
      break
    }
  }
  return exist
}

/**
 * 判断页面路径是否存在(element plus form rules)
 * 
 * @param rule 规则
 * @param value 原始值
 * @param callback 回调
 */
export function existViewValidator(rule: any, value: string, callback: Function) {
  if (value && !existView(value)) {
    callback(new Error('页面组件不存在'))
  }
  callback()
}

/**
 * 判断是否电话号码
 *
 * @param {string} phone 电话号码（固话、手机号码）
 * @returns {boolean}
 */
export function isPhone(phone: string) {
  const e = /^((0\d{2,3}-\d{7,8})|(\d{7,8})|(1[3456789]\d{9}))$/
  return e.test(phone)
}

/**
 * 判断是否电话号码(element plus form rules)
 * 
 * @param rule 规则
 * @param value 原始值
 * @param callback 回调
 */
export function phoneValidator(rule: any, value: string, callback: Function) {
  if (value && !isPhone(value)) {
    callback(new Error('请输入正确的电话号码'))
  }
  callback()
}

/**
 * 判断是否手机号码
 *
 * @param {string} mobilePhone 手机号码
 * @returns {boolean}
 */
export function isMobilePhone(mobilePhone: string) {
  const e = /^(1[3456789]\d{9})$/
  return e.test(mobilePhone)
}

/**
 * 判断是否手机号码(element plus form rules)
 * 
 * @param rule 规则
 * @param value 原始值
 * @param callback 回调
 */
export function mobilePhoneValidator(rule: any, value: string, callback: Function) {
  if (value && !isMobilePhone(value)) {
    callback(new Error('请输入正确的手机号码'))
  }
  callback()
}

/**
 * 判断是否邮箱
 *
 * @param {string} email 邮箱
 * @returns {boolean}
 */
export function isEmail(email: string) {
  const e = /^\w+([-\.]\w+)*@\w+([\.-]\w+)*\.\w{2,4}$/ // eslint-disable-line
  return e.test(email)
}

/**
 * 判断是否车牌号
 *
 * @param {string} carNumber 车牌号
 * @returns {boolean}
 */
export function isCarNumber(carNumber: string) {
  const e = /^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}(([A-Z0-9]{4}[A-Z0-9挂学警港澳]{1}$)|([0-9]{5}[DF]$)|([DF][A-HJ-NP-Z0-9][0-9]{4}$))/
  return e.test(carNumber)
}

/**
 * 身份证校验
 *
 * @param {string} idcard 身份证号
 * @returns {boolean}
 */
export function isIdcard(idcard: string) {
  return validateIdCard(idcard)
}

/**
 * 判断字符串是否是金额数值
 * @param {string} price 金额字符
 * @returns {boolean}
 */
export function isPrice(price: string) {
  return !isNaN(price as any) && !isNaN(parseFloat(price))
}

/**
 * 判断是否金额数值(element plus form rules)
 * 
 * @param rule 规则
 * @param value 原始值
 * @param callback 回调
 */
export function priceValidator (rule: any, value: string, callback: Function) {
  if (value && !isPrice(value)) {
    callback(new Error('金额格式错误'))
  }
  callback()
}
