<script lang="ts" setup>
import VueQrcode from 'vue-qrcode'
const props = defineProps({
  qrCodeValue: {
    type: String,
    default: ''
  },
  qrCodeSize: {
    type: Number,
    default: 150
  },
  color: {
    default: {
      dark: '#000000ff',
      light: '#ffffffff'
    }
  },
  downloadName: {
    type: String,
    default: '二维码'
  }
})

const downloadQrCode = () => {
  let evnet = document.getElementById(`${props.qrCodeValue}qrCodeDom`) as any
  let qrCodeSrc = evnet.src
  const link = document.createElement('a')
  link.href = qrCodeSrc
  link.download = props.downloadName
  link.click()
}
</script>

<template>
  <div :style="{ width: qrCodeSize + 'px' }">
    <VueQrcode :value="qrCodeValue" :width="qrCodeSize" type="image/png" :color="color" :id="qrCodeValue+'qrCodeDom'"></VueQrcode>
    <div style="cursor: pointer;color: rgb(64, 158, 255); display: flex; align-items: center; justify-content: center;" @click="downloadQrCode">
      <el-icon><i-ep-download /></el-icon>点击下载
    </div>
  </div>
</template>

<style lang="scss" scoped></style>
