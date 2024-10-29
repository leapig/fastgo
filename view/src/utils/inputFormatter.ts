import { isPrice } from './validate'

/**
 * 输入框输入值格式化工具
 */

/**
 * 金额格式化
 * @param {string} value 输入框内容
 * @returns {string}
 */
export function priceFormatter(value: string) {
  if (value && isPrice(value)) {
    const vals = value.split('.')
    if (vals.length > 1) {
        return (parseInt(vals[0]) + '').replace(/\B(?=(\d{3})+(?!\d))/g, ',') + '.' + vals[1]
    } else {
        return (parseInt(vals[0]) + '').replace(/\B(?=(\d{3})+(?!\d))/g, ',')
    }
  } else {
    return ''
  }
}

/**
 * 金额格式化值中提取值
 * @param {string} value 输入框内容
 * @returns {string}
 */
export function priceParser(value: string) {
  return value.replace(/\$\s?|(,*)/g, '')
}
