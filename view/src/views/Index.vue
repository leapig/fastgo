<script lang="ts" setup>
import useAppStore from '@/stores/app'
import useSettingsStore from '@/stores/settings'
import useUserStore from '@/stores/user'
import LoginBgImg from '@/assets/images/LoginBg.png'
import { getTenantListApi } from '@/api/login'

const router = useRouter()
const appStore = useAppStore()
const settingStore = useSettingsStore()
const userStore = useUserStore()

const userInfo = ref<any>({})
const tenantList = ref<Array<any>>([])
const loading = ref(false)
const copyright = settingStore.copyright

/** 获取租户列表 */
function getTenantList() {
  loading.value = true
  getTenantListApi()
    .then((response: any) => {
      loading.value = false
      tenantList.value = response.data.rows
    })
    .catch((error: any) => {
      loading.value = false
      ElMessage.error(`获取租户列表失败:${error.errorMessage}`)
    })
}
/** 登录到租户 */
function accessToTenant(item: any) {
  const enterprise = {
    pk: item.pk,
    name: item.name,
    type: item.type
  }
  userStore.accessToTenant(enterprise).then(() => {
    if (item.type === 1000) {
      // 运营平台
      router.push({ path: settingStore.operatingIndexPage })
    } else if (item.type === 1 || item.type === 2 || item.type === 3) {
      // 1、监管单位；2、安保单位；3、从业单位；
      router.push({ path: settingStore.manageIndexPage })
    }
  }).catch((error) => {
    ElMessage.error(`登录租户失败:${error.errorMessage}`)
  })
}

onMounted(() => {
  document.title = `${import.meta.env.VITE_APP_TITLE} | 选择租户`
  appStore.clearUsePlatform()
  userStore.clearTenantAccess()
  getTenantList()
})
</script>

<template>
  <div class="page">
    <div class="bg-box">
      <img :src="LoginBgImg" />
    </div>
    <div class="tenant-box">
      <div class="title">
        <span>请选择您要登录的租户</span>
      </div>
      <div class="scrollbar-box" v-loading="loading">
        <el-scrollbar height="calc(90vh - 100px - 20px)">
          <div class="list-box">
            <div class="list-item" :span="3" v-for="item in tenantList" :key="item.pk" @click="accessToTenant(item)">
              <img :src="LoginBgImg" />
              <div class="label">
                <span>{{ item.name }}</span>
              </div>
            </div>
            <div class="list-empty" v-if="!tenantList.length && !loading">
              <div>您好，暂无可登录的租户，请联系负责人添加之后再登录！</div>
            </div>
          </div>
        </el-scrollbar>
      </div>
    </div>
    <div class="copyright-box">
      <span>{{ copyright }}</span>
    </div>
  </div>
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
  .tenant-box {
    position: absolute;
    width: calc(100vw - 10vh);
    height: 90vh;
    top: 5vh;
    left: 5vh;
    border-radius: 4px;
    background-color: #ffffff;
    box-shadow: 0 1px 5px 0 rgba(14, 33, 39, 0.2);
    z-index: 101;
    .title {
      position: relative;
      display: flex;
      height: 100px;
      align-items: center;
      justify-content: center;
      font-size: 30px;
      .avatar {
        position: absolute;
        top: 30px;
        right: 30px;
        cursor: pointer;
      }
    }
    .scrollbar-box {
      display: flex;
      align-items: center;
      justify-content: center;
      height: calc(90vh - 100px - 20px);
      .list-box {
        padding: 0 10vh 20px 10vh;
        .list-item {
          display: inline-block;
          width: 150px;
          height: 210px;
          border: 1px solid #f0f0f0;
          border-radius: 2px;
          margin: 10px;
          cursor: pointer;
          img {
            width: 150px;
            height: 150px;
            border-radius: 2px 2px 0 0;
            object-fit: cover;
          }
          .label {
            position: relative;
            width: 150px;
            height: 60px;
            span {
              position: absolute;
              width: 140px;
              top: 50%;
              left: 50%;
              transform: translate(-50%, -50%);
              font-size: 16px;
              overflow: hidden;
              text-overflow: ellipsis;
              display: -webkit-box;
              -webkit-box-orient: vertical;
              -webkit-line-clamp: 2;
            }
          }
        }
        .list-item:hover {
          box-shadow: 0 0 8px 0 rgba(14, 33, 39, 0.3);
        }
        .list-empty {
          color: var(--el-text-color-secondary);
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
