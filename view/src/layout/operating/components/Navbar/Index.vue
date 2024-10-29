<script lang="ts" setup>
import useAppStore from '@/stores/app'
import useUserStore from '@/stores/user'
import useSettingsStore from '@/stores/settings'

const appStore = useAppStore()
const userStore = useUserStore()
const settingsStore = useSettingsStore()

const sidebar = appStore.operating.sidebar

const userInfo = ref<any>({})

/** 展开/收起侧边栏 */
function toggleSideBar() {
  appStore.toggleSideBar(false)
}
/** 下拉菜单选择 */
function handleCommand(command: any) {
  switch (command) {
    case 'tenant':
    userStore.goSelectTenant()
      break
    case 'logout':
    userStore.logout()
      break
    default:
      break
  }
}

onMounted(() => {
  appStore.setOperatingUseTab('system')
  userStore.setUser({})
  userStore.getUserInfo().then((user) => {
    userInfo.value = user
  })
})
</script>

<template>
  <div class="navbar" :style="{ background: settingsStore.navbarColor || '#fff' }">
    <div class="toggle" :class="{ active: !sidebar.opened }" v-if="!sidebar.hide" @click="toggleSideBar">
      <el-icon :size="20"><i-ep-fold /></el-icon>
    </div>
    <div class="tabs">
      <operating-components-navbar-top-tab />
    </div>
    <div class="right-menu">
      <div class="user-info">
        <div class="text">{{ userInfo.name || '' }}</div>
        <div class="text">{{ userInfo.role || '' }}</div>
      </div>
      <div class="avatar-container">
        <el-dropdown @command="handleCommand" class="right-menu-item hover-effect" trigger="click" placement="bottom-end">
          <div class="avatar-wrapper">
            <!-- <avatar v-model="userInfo.avatar"></avatar> -->
            <el-icon><i-ep-caret-bottom /></el-icon>
          </div>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="tenant">
                <el-icon :size="18">
                  <i-ep-menu />
                </el-icon>
                <span>选择租户</span>
              </el-dropdown-item>
              <el-dropdown-item command="logout">
                <el-icon :size="18">
                  <svg-icon icon-name="icon-logout" />
                </el-icon>
                <span>退出登录</span>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.navbar {
  position: relative;
  display: flex;
  height: 50px;
  margin-bottom: 2px;
  overflow: hidden;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  .toggle {
    display: flex;
    height: 50px;
    padding: 0 15px;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    &:hover {
      background: rgba(0, 0, 0, 0.025);
    }
    &.active {
      transform: rotate(180deg);
    }
  }
  .tabs {
    flex: 1;
  }
  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;
    display: flex;
    &:focus {
      outline: none;
    }
    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;
      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;
      }
    }
    .user-info {
      display: inline-block;
      height: 40px;
      width: 80px;
      line-height: 20px;
      padding-left: 10px;
      margin: 5px 0;
      // border-left: 1px solid #f0f0f0;
      cursor: default;
      .text {
        width: 75px;
        color: var(--el-text-color-primary);
        font-size: 12px;
        text-align: right;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }
    }
    .avatar-container {
      margin-right: 20px;
      .avatar-wrapper {
        margin-top: 5px;
        position: relative;
        i {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}
</style>
