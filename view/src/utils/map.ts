export type GeoLocation = {
  fullAddress: string // 完整地址
  province: string // 省
  city: string // 市
  district: string // 区/县
  township: string // 镇/街道
  street: string // 道路
  streetNumber: string // 门牌号
  adCode: string // 省市区三级行政区划编码
  townCode: string // 省市区街道四级行政区划编码
}

/**
 * 获取位置
 *
 * @returns {*}
 */
export function getCurrentPosition() {
  let position: any = undefined
  if ('geolocation' in navigator) {
    navigator.geolocation.getCurrentPosition(
      (po) => {
        position = {
          latitude: po.coords.latitude,
          longitude: po.coords.longitude
        }
      },
      (error) => {
        console.warn('获取位置失败：', error.message)
      }
    )
  } else {
    console.warn('浏览器不支持Geolocation')
  }
  return position
}
