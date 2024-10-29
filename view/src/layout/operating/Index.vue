<script lang="ts" setup>
import useAppStore from '@/stores/app'

const appStore = useAppStore()

const sidebar = appStore.operating.sidebar
</script>

<template>
  <div :class="{ hideSidebar: !sidebar.opened }" class="app-wrapper">
    <operating-components-sidebar v-if="!sidebar.hide" class="sidebar-container" />
    <div :class="{ sidebarHide: sidebar.hide }" class="hasNavbar main-container">
      <operating-components-navbar />
      <el-scrollbar wrap-class="page-container" noresize>
        <operating-components-app-main />
      </el-scrollbar>
      <el-backtop target=".page-container" :right="70" :bottom="70" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.app-wrapper {
  position: relative;
  background-color: #f7f7f7;
  width: 100%;
  height: 100%;
}
</style>
