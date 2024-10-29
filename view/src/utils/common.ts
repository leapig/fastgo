/**
 * 通用方法封装
 */

/**
 * 日期格式化
 *
 * @param {*} time 时间
 * @param {string} pattern 格式化模板
 * @returns
 */
export function parseTime(time: any, pattern?: string) {
  if (arguments.length === 0 || !time) {
    return null
  }
  const format = pattern || '{y}-{m}-{d} {h}:{i}:{s}'
  let date
  if (typeof time === 'object') {
    date = time
  } else {
    if (typeof time === 'string' && /^[0-9]+$/.test(time)) {
      time = parseInt(time)
    } else if (typeof time === 'string') {
      time = time
        .replace(new RegExp(/-/gm), '/')
        .replace('T', ' ')
        .replace(new RegExp(/\.[\d]{3}/gm), '')
    }
    if (typeof time === 'number' && time.toString().length === 10) {
      time = time * 1000
    }
    date = new Date(time)
  }
  const formatObj: { [key: string]: number } = {
    y: date.getFullYear(),
    m: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    i: date.getMinutes(),
    s: date.getSeconds(),
    a: date.getDay()
  }
  const time_str = format.replace(/{(y|m|d|h|i|s|a)+}/g, (result, ...key: any[]): string => {
    const value = formatObj[key[0]]
    let res = ''
    // Note: getDay() returns 0 on Sunday
    if (key[0] === 'a') {
      return ['日', '一', '二', '三', '四', '五', '六'][value]
    }
    if (result.length > 0 && value < 10) {
      res = '0' + value
    }
    return res || '0'
  })
  return time_str
}

/**
 * 时间字符串转时间对象
 *
 * @param {string} str 时间字符串
 * @returns
 */
export function strToDate(str: string) {
  return new Date(Date.parse(str))
}

/**
 * 时间对象转字符串
 *
 * @param {*} date 时间对象
 * @param {string} format 字符串格式
 * @returns
 */
export function formatDateToString(date: any, format?: string) {
  const formatDate = format ? format : 'YYYY-MM-DD hh:mm:ss'
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')
  return formatDate.replace('YYYY', year).replace('MM', month).replace('DD', day).replace('hh', hours).replace('mm', minutes).replace('ss', seconds)
}

/**
 * 时分秒字符串去掉秒
 * @param time 
 */
export function timeStrRemoveSecond(time: string) {
  return time && time.length === 8 ? time.substring(0, 5) : time
}

/**
 * 计算年龄
 *
 * @param {string|Date} birth 生日
 * @returns
 */
export function getAge(birth: string | Date) {
  let birthday: Date | undefined = undefined
  if (birth) {
    if (typeof birth === 'string') {
      birthday = strToDate(birth)
    } else {
      birthday = birth
    }
  }
  if (birthday) {
    const now = new Date().getTime()
    const bt = birthday.getTime()
    return Math.floor((now - bt) / (12 * 30 * 24 * 60 * 60 * 1000))
  }
  return null
}

/**
 * 构造树型结构数据
 *
 * @param {Array} data 数据源
 * @param {string} id id字段 默认 'id'
 * @param {string} parentId 父节点字段 默认 'parentId'
 * @param {string} children 孩子节点字段 默认 'children'
 */
export function handleTree(data: any[], id?: string, parentId?: string, children?: string) {
  const config: { [key: string]: any } = {
    id: id || 'id',
    parentId: parentId || 'parentId',
    childrenList: children || 'children'
  }
  const childrenListMap: { [key: string]: any } = {}
  const nodeIds: { [key: string]: any } = {}
  const tree: { [key: string]: any }[] = []

  for (const d of data) {
    const parentId = d[config.parentId]
    if (childrenListMap[parentId] == null) {
      childrenListMap[parentId] = []
    }
    nodeIds[d[config.id]] = d
    childrenListMap[parentId].push(d)
  }

  for (const d of data) {
    const parentId = d[config.parentId]
    if (nodeIds[parentId] == null) {
      tree.push(d)
    }
  }

  for (const t of tree) {
    adaptToChildrenList(t)
  }

  function adaptToChildrenList(o: any) {
    if (childrenListMap[o[config.id]] !== null) {
      o[config.childrenList] = childrenListMap[o[config.id]]
    }
    if (o[config.childrenList]) {
      for (const c of o[config.childrenList]) {
        adaptToChildrenList(c)
      }
    }
  }
  return tree
}

/**
 * 参数处理
 *
 * @param {*} params 参数
 */
export function tansParams(params: { [key: string]: any }) {
  let result = ''
  for (const propName of Object.keys(params)) {
    const value = params[propName]
    const part = encodeURIComponent(propName) + '='
    if (value !== null && value !== '' && typeof value !== 'undefined') {
      if (typeof value === 'object') {
        for (const key of Object.keys(value)) {
          if (value[key] !== null && value[key] !== '' && typeof value[key] !== 'undefined') {
            const params = propName + '[' + key + ']'
            const subPart = encodeURIComponent(params) + '='
            result += subPart + encodeURIComponent(value[key]) + '&'
          }
        }
      } else {
        result += part + encodeURIComponent(value) + '&'
      }
    }
  }
  return result
}

/**
 * 下划线转驼峰
 * @param {string} str
 * @returns
 */
export function convertToCamelCase(str: string) {
  return str.toLowerCase().replace(/_(.)/g, function (match, group1) {
    return group1.toUpperCase()
  })
}

/**
 * 对象参数下划线转驼峰
 * @param {any} obj
 * @returns
 */
export function objectToCamelCase(obj: any) {
  if (obj) {
    if (typeof obj === 'object' && !Array.isArray(obj)) {
      const newObj: any = {}
      for (const key in obj) {
        let value = obj[key]
        if (typeof value === 'object') {
          value = objectToCamelCase(value)
        }
        const camelKey = convertToCamelCase(key)
        newObj[camelKey] = value
      }
      return newObj
    } else if (Array.isArray(obj)) {
      const newArr: Array<any> = []
      for (const index in obj) {
        const newObj = objectToCamelCase(obj[index])
        newArr.push(newObj)
      }
      return newArr
    }
  }
  return null
}

/**
 * 驼峰转下划线
 * @param {string} str
 * @returns
 */
export function convertToUnderscore(str: string) {
  return str.replace(/([a-z0-9])([A-Z])/g, '$1_$2').toLowerCase()
}

/**
 * 对象参数驼峰转下划线
 * @param {any} obj
 * @returns
 */
export function objectToUnderscore(obj: any) {
  if (obj) {
    if (typeof obj === 'object' && !Array.isArray(obj)) {
      const newObj: any = {}
      for (const key in obj) {
        let value = obj[key]
        if (typeof value === 'object') {
          value = objectToUnderscore(value)
        }
        const underKey = convertToUnderscore(key)
        newObj[underKey] = value
      }
      return newObj
    } else if (Array.isArray(obj)) {
      const newArr: Array<any> = []
      for (const index in obj) {
        const newObj = objectToUnderscore(obj[index])
        newArr.push(newObj)
      }
      return newArr
    }
  }
  return null
}
