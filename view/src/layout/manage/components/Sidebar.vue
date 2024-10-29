<script lang="ts" setup>
import useAppStore from '@/stores/app'
import usePermissionStore from '@/stores/permission'

const route = useRoute()
const appStore = useAppStore()
const permissionStore = usePermissionStore()

const hover = ref<boolean>(false)

const isCollapse = computed(() => !appStore.manage.sidebar.opened)

const isHover = computed<boolean>({
  get: () => {
    return hover.value
  },
  set: (val: boolean) => {
    hover.value = val
  }
})

const activeMenu = computed<string>(() => route.fullPath)

const menus = computed<any>(() => {
  const list = []
  const managePermission: any = permissionStore.manage
  const menus = managePermission.menus
  list.push(...menus)
  return list
})

function toggleSideBar() {
  appStore.toggleSideBar(false)
}
</script>

<template>
  <div class="has-logo sidebar-container" @mouseover="isHover = true" @mouseout="isHover = false">
    <manage-components-user />
    <el-scrollbar wrap-class="scrollbar-wrapper" :height="isCollapse ? 'calc(100vh - 100px)' : 'calc(100vh - 150px)'">
      <el-menu :default-active="activeMenu" :collapse="isCollapse" background-color="#fff" unique-opened :collapse-transition="false" mode="vertical">
        <manage-components-sidebar-item v-for="menu in menus" :key="'/manage' + menu.path" :item="menu" base-path="/manage" />
      </el-menu>
    </el-scrollbar>
    <manage-components-toggler />
  </div>
  <div class="scrollbar-box" v-if="isCollapse" @mouseover="isHover = true" @mouseout="isHover = false">
    <div class="collapse-btn" :class="{ 'hover': isHover }" @click="toggleSideBar">
      <el-icon size="16">
        <svg-icon icon-name="icon-toggle-right"/>
      </el-icon>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to { /* 动画结束时 */
    opacity: 1;
  }
}
.scrollbar-box {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  width: 74px;
  z-index: 1000;
  .collapse-btn {
    position: absolute;
    display: flex;
    top: 0;
    bottom: 0;
    right: 0;
    width: 20px;
    height: 100%;
    background: radial-gradient(ellipse at left, rgba(0, 21, 41, 0.1) 10% 5%, rgba(0, 0, 0, 0) 80%);
    color: rgba(0, 21, 41, 0.4);
    align-items: center;
    justify-content: center;
    opacity: 0;
    cursor: pointer;
    &.hover {
      opacity: 0;
      animation: fadeIn 0.6s ease-in-out forwards;
    }
    &:hover {
      color: var(--el-menu-active-color);
    }
  }
}
</style>
