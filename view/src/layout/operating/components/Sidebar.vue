<script lang="ts" setup>
import useAppStore from '@/stores/app'
import usePermissionStore from '@/stores/permission'

const route = useRoute()
const appStore = useAppStore()
const permissionStore = usePermissionStore()

const isCollapse = computed(() => !appStore.operating.sidebar.opened)

const activeMenu = computed<string>(() => route.fullPath)

const basePath = computed<string>(() => {
  return '/operating/' + appStore.operating.useTab
})

const menus = computed<Array<any>>(() => {
  const list = []
  const nowTab: string = appStore.operating.useTab
  if (nowTab) {
    const operatingPermissionMenus: any = permissionStore.operating.menus
    const tabMenus = operatingPermissionMenus[nowTab]
    if (tabMenus && tabMenus.length) {
      list.push(...tabMenus)
    }
  }
  return list
})
</script>

<template>
  <div class="has-logo">
    <operating-components-logo :collapse="isCollapse" />
    <el-scrollbar wrap-class="scrollbar-wrapper" height="calc(100vh - 50px)">
      <el-menu :default-active="activeMenu" :collapse="isCollapse" background-color="#fff" unique-opened :collapse-transition="false" mode="vertical">
        <operating-components-sidebar-item v-for="menu in menus" :key="basePath + '/' + menu.path" :item="menu" :base-path="basePath" />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<style lang="scss" scoped></style>
