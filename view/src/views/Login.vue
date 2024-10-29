<script lang="ts" setup>
import useAppStore from '@/stores/app'
import useSettingsStore from '@/stores/settings'
import useUserStore from '@/stores/user'
import backgroundImage from '@/assets/images/LoginBg.png'
const appStore = useAppStore()
const settingStore = useSettingsStore()
const userStore = useUserStore()

const loadingQrcode = ref(false)
const wechatLogin = ref({
  qrcode: '',
  key: ''
})
const showProtocol = ref(false)
const showProtocolType = ref<any>(undefined)
const copyright = settingStore.copyright

/** 获取微信登录二维码 */
function getWechatLoginQrcode() {
  loadingQrcode.value = true
  userStore.startWechatLogin().then((response: any) => {
    loadingQrcode.value = false
    wechatLogin.value.qrcode = response.qrcode
    wechatLogin.value.key = response.key
  }).catch((error) => {
    loadingQrcode.value = false
    ElMessage.error(`获取登录二维码失败:${error.errorMessage}`)
  })
}
/** 刷新微信登录二维码 */
function refreshWechatLoginQrcode() {
  loadingQrcode.value = true
  userStore.refreshWechatLogin().then((response: any) => {
    loadingQrcode.value = false
    wechatLogin.value.qrcode = response.qrcode
    wechatLogin.value.key = response.key
  }).catch((error) => {
    loadingQrcode.value = false
    ElMessage.error(`获取二维码失败:${error.errorMessage}`)
  })
}
/** 打开协议对话框 */
function handleProtocol(type: string) {
  showProtocolType.value = type
  showProtocol.value = true
}

onMounted(() => {
  document.title = `${import.meta.env.VITE_APP_TITLE} | 登录`
  appStore.clearUsePlatform()
  userStore.clearUserAccess()
  getWechatLoginQrcode()
})

onUnmounted(() => {
  userStore.stopWechatLogin()
})
</script>

<template>
  <div class="page">
    <div class="bg-box">
      <img :src="backgroundImage"/>
    </div>
    <div class="opt-box">
      <div class="box">
        <!-- 扫码登录 -->
        <div class="form-box">
          <div class="title">
            <span>微信扫码登录</span>
          </div>
          <div class="qrcode-box">
            <div class="qrcode" v-loading="loadingQrcode">
              <div class="cover" v-if="!loadingQrcode" @click="refreshWechatLoginQrcode">
                <el-icon size="30" color="#fff">
                  <i-ep-refresh />
                </el-icon>
              </div>
              <el-image fit="scale-down" style="width: 200px; height: 200px;" :src="wechatLogin.qrcode" alt='安保易'>
                <template #error>
                  <div class="image-slot">
                    <el-icon size="30" color="var(--el-text-color-secondary)">
                      <svg-icon icon-name="icon-picture-error"/>
                    </el-icon>
                  </div>
                </template>
              </el-image>
            </div>
          </div>
          <div class="tips">
            <span>登录视为您已同意</span>
            <el-button class="btn" type="primary" link @click="handleProtocol('bindThirdAccount')">第三方账号绑定协议</el-button>
            <span>、</span>
            <el-button class="btn" type="primary" link @click="handleProtocol('user')">用户协议</el-button>
            <span>、</span>
            <el-button class="btn" type="primary" link @click="handleProtocol('privacyPolicy')">隐私政策</el-button>
            <span>、</span>
            <el-button class="btn" type="primary" link @click="handleProtocol('productService')">产品服务协议</el-button>
          </div>
        </div>
      </div>
    </div>
    <div class="copyright-box">
      <span>{{ copyright }}</span>
    </div>
  </div>
  <!-- 协议对话框 -->
  <protocol v-model="showProtocol" :type="showProtocolType" :closable-after="5"></protocol>
</template>

<style lang="scss" scoped>
.page {
  position: relative;
  width: 100vw;
  height: 100vh;
  .bg-box {
    position: absolute;
    width: 100%;
    z-index: 100;
    img {
      width: 100%;
      height: 100vh;
      object-fit: cover;
    }
  }
  .opt-box {
    position: absolute;
    width: 430px;
    height: 370px;
    top: 50%;
    left: 66%;
    border-radius: 4px;
    background-color: #ffffff;
    transform: translate(-50%, -50%);
    z-index: 101;
    .box {
      display: flex;
      position: relative;
      width: 430px;
      height: 350px;
      padding-top: 20px;
      overflow: hidden;
      .change-btn {
        position: absolute;
        top: 0;
        right: 0;
        width: 100px;
        height: 70px;
        padding-left: 10px;
        border-bottom-left-radius: 100px 70px;
        border-top-right-radius: 4px;
        background-color: skyblue;
        line-height: 55px;
        color: white;
        font-size: 14px;
        text-align: center;
        cursor: pointer;
      }
      .form-box {
        width: 100%;
        .title {
          display: flex;
          height: 50px;
          align-items: center;
          justify-content: center;
          font-size: 18px;
        }
        .qrcode-box {
          display: flex;
          height: 250px;
          align-items: center;
          justify-content: center;
          .qrcode {
            position: relative;
            width: 200px;
            height: 200px;
            .cover {
              display: none;
              position: absolute;
              width: 100%;
              height: 100%;
              z-index: 111;
              align-items: center;
              justify-content: center;
              cursor: pointer;
              background: rgba(0, 0, 0, 0.3);
            }
            &:hover .cover {
              display: flex;
            }
            :deep(.image-slot) {
              display: flex;
              justify-content: center;
              align-items: center;
              width: 100%;
              height: 100%;
              background: var(--el-fill-color-light);
            }
          }
        }
        .form {
          height: 225px;
          padding: 50px 30px 0;
          .flex {
            display: flex;
          }
          .agree-text {
            display: flex;
            flex: 1;
            font-size: 12px;
            margin-left: 5px;
            align-items: center;
            .btn {
              font-size: 12px;
              margin: 0;
              padding: 0;
            }
          }
        }
        .tips {
          display: flex;
          width: 430px;
          height: 50px;
          align-items: center;
          justify-content: center;
          font-size: 12px;
          .btn {
            font-size: 12px;
            margin: 0;
            padding: 0;
          }
        }
      }
    }
  }
  .copyright-box {
    position: absolute;
    left: 50%;
    bottom: 0;
    transform: translate(-50%, -50%);
    color: #fff;
    font-size: 14px;
    z-index: 102;
  }
}
</style>
