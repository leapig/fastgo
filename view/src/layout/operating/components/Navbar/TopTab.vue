<script lang="ts" setup>
import useAppStore from '@/stores/app'
import usePermissionStore from '@/stores/permission'

const appStore = useAppStore()
const permissionStore = usePermissionStore()

const tabsRef = ref<HTMLDivElement>()
const tabRef = ref<Array<any>>([])
const tabs = ref<Array<any>>([])
const maxWidth = ref(0)
const scrollIndex = ref(0)
const showCtl = ref(false)

/**
 * 初始加载
 */
function initLoad() {
  tabs.value = permissionStore.operating.tabs
  nextTick(() => {
    loadScroll()
  })
}

/**
 * 加载滚动条
 */
function loadScroll() {
  if (tabsRef.value) {
    maxWidth.value = tabsRef.value.scrollWidth - tabsRef.value.clientWidth
    showCtl.value = tabsRef.value.scrollWidth > tabsRef.value.clientWidth
  }
}

function setTabRef(element: any, index: number) {
  tabRef.value[index] = element
}

/**
 * 选择Tab
 */
function selectTab(key: string) {
  appStore.setOperatingUseTab(key)
}

/**
 * 翻页
 * @param arg 翻页数
 */
function tabMove(arg: number) {
  if (maxWidth.value !== tabsRef.value!.scrollLeft || arg < 0) {
    const index = scrollIndex.value + arg
    if (index >= 0) {
      const tab = tabRef.value[index]
      if (tab) {
        scrollIndex.value = index
        tab.scrollIntoView({ block: 'start', behavior: 'smooth', inline: 'start' })
      }
    }
  }
}

watch(appStore.operating, (value, oldValue) => {
  if (value.sidebar.opened !== oldValue.sidebar.opened) {
    setTimeout(() => {
      loadScroll()
    }, 330)
  }
})

watch(permissionStore.operating, () => {
  initLoad()
})

onMounted(() => {
  initLoad()
})
</script>

<template>
  <template v-if="tabs.length > 0">
    <div class="box">
      <div class="tabs" ref="tabsRef">
        <div
          class="tab"
          :class="{ active: appStore.operating.useTab === item.key }"
          v-for="(item, index) in tabs"
          :key="'tab-' + index"
          :ref="(element) => setTabRef(element, index)"
          @click="selectTab(item.key)"
        >
          <span>{{ item.label }}</span>
        </div>
      </div>
      <div class="ctls" v-if="showCtl">
        <el-icon class="ctl" size="20" color="#A8ABB2" @click="tabMove(-1)">
          <i-ep-caret-left />
        </el-icon>
        <el-icon class="ctl" size="20" color="#A8ABB2" @click="tabMove(1)">
          <i-ep-caret-right />
        </el-icon>
      </div>
    </div>
  </template>
</template>

<style lang="scss" scoped>
.box {
  display: flex;
  .tabs {
    position: relative;
    flex: 1;
    width: 0;
    display: flex;
    overflow-x: hidden;
    .tab {
      display: flex;
      height: 50px;
      padding: 0 20px;
      align-items: center;
      justify-content: center;
      font-size: 14px;
      cursor: pointer;
      white-space: nowrap;
      &:hover {
        background: var(--el-color-primary-light-9);
      }
      &.active {
        height: 48px;
        border-bottom: 2px solid var(--el-color-primary);
        background: var(--el-color-primary-light-8);
        color: var(--el-color-primary);
      }
    }
  }
  .ctls {
    float: right;
    display: flex;
    height: 50px;
    padding: 0 10px;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    .ctl {
      cursor: pointer;
      &:hover {
        color: #606266;
      }
    }
  }
}
</style>
