<script lang="ts" setup>
import type { TabPaneName } from 'element-plus'
import usePermissionStore from '@/stores/permission'
import cache from '@/plugins/cache'

const route = useRoute()
const router = useRouter()
const permissionStore = usePermissionStore()

const props = defineProps({
  /**
   * 标签栏列表
   * 结构：[{label:'标签名',path:'子页面路由，也用作标签key',component:'页面组件路径，校验权限用'},...]
   */
  tabs: {
    type: Array,
    required: true
  },
  /** 父页面路由 */
  pagePath: {
    type: String,
    required: true
  },
  /** 页面状态参数 */
  pageState: {
    type: Object
  },
  /** 参数中表示唯一值的字段 */
  pKey: {
    type: String,
    default: 'pk'
  }
})

const activeTab = ref('')
const pk = ref('')

const tabs = computed<Array<any>>(() => {
  return props.tabs as Array<any>
})

const cacheKey = computed<string>(() => {
  if (pk.value) {
    return `${props.pagePath}-${pk.value}`
  } else {
    return `${props.pagePath}`
  }
})

const toPath = computed<string>(() => {
  if (pk.value) {
    return `${props.pagePath}/${activeTab.value}/${pk.value}`
  } else {
    return `${props.pagePath}/${activeTab.value}`
  }
})

/** 初始化 */
function initTabs() {
  if (props.pageState && props.pageState[props.pKey]) {
    pk.value = props.pageState[props.pKey]
    activeTab.value = cache.local.get(cacheKey.value) as string
    clearActiveTabCache()
  }
  if (route.path === props.pagePath) {
    if (!activeTab.value) {
      for (const ind in props.tabs) {
        const tab = props.tabs[ind] as any
        if (permissionStore.hasPagePermi(tab.component)) {
          activeTab.value = tab.path
          cache.local.set(cacheKey.value, activeTab.value)
          router.replace({ path: toPath.value, state: props.pageState })
          break
        }
      }
    } else {
      router.replace({ path: toPath.value, state: props.pageState })
    }
  }
}
/** Tab切换 */
function handleTabChange(name: TabPaneName) {
  if (pk.value) {
    cache.local.set(cacheKey.value, name)
    router.replace({ path: toPath.value, state: props.pageState })
  }
}
/** 清除当前Tab缓存 */
function clearActiveTabCache() {
  cache.local.remove(cacheKey.value)
}

onMounted(() => {
  initTabs()
})

onUnmounted(() => {
  if (!route.path.startsWith(props.pagePath)) {
    clearActiveTabCache()
  }
})
</script>

<template>
  <!-- 标签栏 -->
  <el-tabs v-model="activeTab" @tab-change="handleTabChange">
    <template v-for="(tab, index) in tabs" :key="index">
      <el-tab-pane :label="tab.label" :name="tab.path" v-if="permissionStore.hasPagePermi(tab.component)"></el-tab-pane>
    </template>
  </el-tabs>
  <!-- 子页面 -->
  <section class="app-main">
    <router-view v-slot="{ Component, route }">
      <component :is="Component" :key="route.path" />
    </router-view>
  </section>
</template>

<style lang="scss" scoped>
.app-main {
  width: 100%;
  margin-top: 10px;
  position: relative;
  overflow: hidden;
}
</style>
