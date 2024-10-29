<script lang="ts" setup>
import useAppStore from '@/stores/app'
import useUserStore from '@/stores/user'

const appStore = useAppStore()
const userStore = useUserStore()

const userInfo = ref<any>({})

const isCollapse = computed(() => !appStore.manage.sidebar.opened)

const tenantInfo = computed<any>(() => userStore.enterprise)

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
  userStore.setUser({})
  userStore.getUserInfo().then((user) => {
    userInfo.value = user
  })
})
</script>

<template>
  <div class="sidebar-user-container" :class="{ collapse: isCollapse }" style="background: #fff">
    <transition name="sidebarUserFade">
      <div v-if="isCollapse" key="collapse" class="sidebar-user-collapse">
        <el-dropdown @command="handleCommand" trigger="click">
          <avatar v-model="userInfo.avatar"></avatar>
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
      <div v-else key="expand" class="sidebar-user-expand">
        <div class="avatar">
          <avatar :size="60"></avatar>
        </div>
        <div class="info">
          <div class="text">{{ tenantInfo.name || '' }}·{{ userInfo.name || '' }}</div>
        </div>
      </div>
    </transition>
    <div class="sidebar-user-btns" :class="{ collapse: isCollapse }">
      <el-dropdown @command="handleCommand" class="btn" trigger="click" v-if="!isCollapse">
        <el-icon :size="20" color="var(--el-color-primary)">
          <i-ep-setting />
        </el-icon>
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
</template>

<style lang="scss" scoped>
.sidebarUserFade-enter-active {
  transition: opacity 1.5s;
}
.sidebarUserFade-enter, .sidebarUserFade-leave-to {
  opacity: 0;
}
.sidebar-user-container {
  position: relative;
  width: 100%;
  overflow: hidden;
  & .sidebar-user-collapse {
    display: flex;
    width: 100%;
    height: 100%;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }
  & .sidebar-user-expand {
    padding: 10px;
    .avatar {
      display: flex;
      align-items: center;
      height: 60px;
      padding: 0 10px 5px 10px;
    }
    .info {
      line-height: 20px;
      padding: 5px 10px 0 10px;
      .text {
        font-size: 16px;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
      }
    }
  }
  &.collapse {
    height: 60px;
  }
  .sidebar-user-btns {
    display: none;
    position: absolute;
    top: 0;
    right: 0;
    height: 40px;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    .btn {
      display: flex;
      width: 40px;
      height: 40px;
      align-items: center;
      justify-content: center;
    }
  }
  &:hover .sidebar-user-btns {
    display: flex;
  }
  &:not(.collapse) .sidebar-user-btns {
    display: flex;
  }
}
</style>
