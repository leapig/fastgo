<script lang="ts" setup>
import AMapLoader from '@amap/amap-jsapi-loader'
import { getCurrentPosition } from '@/utils/map'
import type { GeoLocation } from '@/utils/map'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => {
      return {}
    }
  },
  aKey: {
    type: String
  },
  showBottomTips: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'select'])

let amap: any = null
let map: any = null
let marker: any = null
let geocoder: any = null

/** 地图初始化 */
function initMap() {
  const mapParams: any = {
    viewMode: '2D',
    zoom: 15
  }
  const position = getCurrentPosition()
  if (position !== undefined) {
    mapParams.center = [position.longitude, position.latitude]
  }
  AMapLoader.load({
    key: import.meta.env.VITE_APP_AMAP_KEY,
    version: '2.0',
    plugins: ['AMap.Geocoder']
  })
    .then((AMap) => {
      amap = AMap
      map = new AMap.Map(props.aKey ? 'container-' + props.aKey : 'container', mapParams)
      geocoder = new AMap.Geocoder({ city: '全国' })
      map.setDefaultCursor('move')
      /** 定义地图点击事件 */
      map.on('click', (e: any) => {
        gencodeToLocation(e.lnglat.getLng(), e.lnglat.getLat()).then((location) => {
          addMark(e.lnglat.getLng(), e.lnglat.getLat(), location.fullAddress)
          emit('update:modelValue', { longitude: e.lnglat.getLng(), latitude: e.lnglat.getLat() })
          emit('select', location)
        })
      })
      if (props.modelValue && props.modelValue.longitude && props.modelValue.latitude) {
        gencodeToLocation(props.modelValue.longitude, props.modelValue.latitude).then((location) => {
          addMark(props.modelValue.longitude, props.modelValue.latitude, location.fullAddress)
        })
      }
    })
    .catch((e: any) => {
      console.error('Init amap failed: ', e)
    })
}
/** 添加地图标记 */
function addMark(longitude: number, latitude: number, title: string) {
  marker && removeMark()
  marker = new amap.Marker({
    position: new amap.LngLat(longitude, latitude),
    title
  })
  marker.on('dblclick', () => {
    removeMark()
    emit('update:modelValue', {})
    emit('select', {})
  })
  map.add(marker)
  map.setCenter([longitude, latitude])
}
/** 移除地图标记 */
function removeMark() {
  map.remove(marker)
}
/** 坐标转地址 */
function gencodeToLocation(longitude: number, latitude: number) {
  return new Promise<GeoLocation>((resolve, reject) => {
    const lnglat = [longitude, latitude]
    geocoder.getAddress(lnglat, (status: string, result: any) => {
      if (status === 'complete' && result.info === 'OK') {
        const ac = result.regeocode.addressComponent
        const geoLocation: GeoLocation = {
          fullAddress: result.regeocode.formattedAddress,
          province: ac.province,
          city: ac.city,
          district: ac.district,
          township: ac.township,
          street: ac.street,
          streetNumber: ac.streetNumber,
          adCode: ac.adcode,
          townCode: ac.towncode
        }
        resolve(geoLocation)
      } else {
        reject('Gen location failed.')
      }
    })
  })
}
/** 设置地址并标记 */
function setAddressAndMark(address: string) {
  if (map && geocoder) {
    geocoder.getLocation(address, (status: string, result: any) => {
      if (status === 'complete' && result.info === 'OK') {
        if (result.resultNum > 0) {
          const codes = result.geocodes[0]
          const location = codes.location
          const data = {
            longitude: location.lng,
            latitude: location.lat
          }
          addMark(location.lng, location.lat, codes.formattedAddress)
          emit('update:modelValue', data)
          gencodeToLocation(location.lng, location.lat).then((location) => {
            emit('select', location)
          })//通过坐标再次转地址获取完整的省市区街道的地址回传
        } else {
          console.warn('Amap get address codes failed')
        }
      } else {
        console.error('Amap get address codes failed')
      }
    })
  } else {
    console.error('Amap not inited')
  }
}

defineExpose({
  setAddressAndMark
})

onMounted(() => {
  window._AMapSecurityConfig = {
    securityJsCode: import.meta.env.VITE_APP_AMAP_SECRET
  }
  initMap()
})

onUnmounted(() => {
  map?.clearMap() //清除所有覆盖物
  map?.destroy()
})
</script>

<template>
  <div class="box">
    <div :id="aKey ? 'container-' + aKey : 'container'" class="amap" :style="{ height: showBottomTips ? 'calc(100% - 36px)' : '100%' }"></div>
    <div v-if="showBottomTips" class="tips">
      <el-icon size="14" color="#646464"><svg-icon icon-name="icon-question-circle" /></el-icon>
      <span style="margin-left: 5px">点击地图标点，双击删除标点</span>
    </div>  
  </div>
</template>

<style lang="scss" scoped>
.box {
  width: 100%;
  height: 100%;
}
.amap {
  width: 100%;
  border-radius: 2px;
  box-shadow: 0 1px 5px 0 rgba(14, 33, 39, 0.2);
}
.tips {
  display: flex;
  width: 100%;
  padding: 8px;
  color: var(--el-text-color-regular);
  font-size: 14px;
  align-items: center;
}
:deep(.amap-logo),
:deep(.amap-copyright) {
  display: none !important;
}
</style>
