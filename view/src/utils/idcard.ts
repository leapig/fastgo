/**
 * 身份证校验
 *
 * @param {string} idcard 身份证号
 * @returns {boolean}
 */
export function validateIdCard(idcard: string) {
  idcard = idcard.toUpperCase()
  // 身份证号码为15位或18位，15位时全为数字，18位前17位为数字，最后一位是校验位，可能为数字或字符X
  if (!/(^\d{15}$)|(^\d{17}([0-9]|X)$)/.test(idcard)) {
    return false
  }
  // 分别分析出生日期和校验位
  const len = idcard.length
  if (len === 15) {
    const re = new RegExp(/^(\d{6})(\d{2})(\d{2})(\d{2})(\d{3})$/)
    const arrSplit = idcard.match(re)
    if (arrSplit) {
      const dtmBirth = new Date('19' + arrSplit[2] + '/' + arrSplit[3] + '/' + arrSplit[4])
      const bGoodDay = Number(dtmBirth.getFullYear().toString().substring(2)) === Number(arrSplit[2]) && dtmBirth.getMonth() + 1 === Number(arrSplit[3]) && dtmBirth.getDate() === Number(arrSplit[4])
      if (!bGoodDay) {
        return false
      } else {
        // 将15位身份证转成18位
        idcard = setIdcardToNew(idcard)
        return true
      }
    } else {
      return false
    }
  }
  if (len === 18) {
    const re2 = new RegExp(/^(\d{6})(\d{4})(\d{2})(\d{2})(\d{3})([0-9]|X)$/)
    const arrSplit2 = idcard.match(re2)
    if (arrSplit2) {
      const dtmBirth2 = new Date(arrSplit2[2] + '/' + arrSplit2[3] + '/' + arrSplit2[4])
      const bGoodDay2 = dtmBirth2.getFullYear() === Number(arrSplit2[2]) && dtmBirth2.getMonth() + 1 === Number(arrSplit2[3]) && dtmBirth2.getDate() === Number(arrSplit2[4])
      if (!bGoodDay2) {
        return false
      } else {
        const arrInt2 = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2]
        const arrCh2 = ['1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2']
        let nTemp2 = 0
        for (let j = 0; j < 17; j++) {
          nTemp2 += Number(idcard.substring(j, j + 1)) * arrInt2[j]
        }
        const valnum2 = arrCh2[nTemp2 % 11]
        if (valnum2 !== idcard.substring(17)) {
          return false
        }
        return true
      }
    }
  }
  return false
}

/**
 * 将15位身份证号码转换为18位身份证号码
 *
 * @param {string} idcard 身份证号
 * @returns {string}
 */
export function setIdcardToNew(idcard: string) {
  if (idcard.length === 15) {
    const arrInt = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2]
    const arrCh = ['1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2']
    let nTemp = 0
    idcard = idcard.substring(0, 6) + '19' + idcard.substring(6, idcard.length - 1)
    for (let i = 0; i < 17; i++) {
      nTemp += Number(idcard.substring(i, i + 1)) * arrInt[i]
    }
    idcard += arrCh[nTemp % 11]
    return idcard
  } else {
    return ''
  }
}

/**
 * 从18位身份证号码中获取性别
 *
 * @param {string} idcard 身份证号
 * @returns {number} 1:男,2:女
 */
export function getSexFromIdcard(idcard: string) {
  if (validateIdCard(idcard)) {
    const sex = Number(idcard.substring(16, 17))
    return sex % 2 === 0 ? 2 : 1
  } else {
    return 0
  }
}

/**
 * 从18位身份证号码中获取生日
 *
 * @param {string} idcard 身份证号
 * @returns {string} 生日字符串，格式yyyy-MM-dd
 */
export function getBirthdayFromIdcard(idcard: string) {
  const re = new RegExp(/^(\d{6})(\d{4})(\d{2})(\d{2})(\d{3})([0-9]|X)$/)
  const arrSplit = idcard.match(re)
  if (arrSplit) {
    return arrSplit[2] + '-' + arrSplit[3] + '-' + arrSplit[4]
  } else {
    return ''
  }
}
