<script lang="ts" setup>
import useAppStore from '@/stores/app'

const appStore = useAppStore()

const sidebar = appStore.manage.sidebar
</script>

<template>
  <div :class="{ hideSidebar: !sidebar.opened }" class="app-wrapper">
    <manage-components-sidebar v-if="!sidebar.hide" />
    <div :class="{ sidebarHide: sidebar.hide }" class="main-container">
      <el-scrollbar wrap-class="page-container" noresize>
        <manage-components-app-main />
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
