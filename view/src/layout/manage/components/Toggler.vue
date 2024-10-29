<script lang="ts" setup>
import useAppStore from '@/stores/app'

const appStore = useAppStore()

const sidebar = appStore.manage.sidebar

/** 展开/收起侧边栏 */
function toggleSideBar() {
  appStore.toggleSideBar(false)
}
</script>

<template>
  <div class="sidebar-toggler-container">
    <div class="toggle" :class="{ active: !sidebar.opened }" v-if="!sidebar.hide" @click="toggleSideBar">
      <el-icon :size="20">
        <i-ep-fold />
      </el-icon>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.sidebar-toggler-container {
  position: relative;
  display: flex;
  width: 100%;
  border-top: 1px solid #f0f0f0;
  overflow: hidden;
  .toggle {
    display: flex;
    width: 100%;
    height: 39px;
    padding: 0 15px;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    &:hover {
      color: var(--el-menu-active-color);
      background: rgba(0, 0, 0, 0.025);
    }
    &.active {
      transform: rotate(180deg);
    }
  }
}
</style>
