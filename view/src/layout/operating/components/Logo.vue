<script lang="ts" setup>
import useSettingsStore from '@/stores/settings'
import logo from '@/assets/images/Logo.jpg'

defineProps({
  collapse: {
    type: Boolean,
    required: true
  }
})

const settingsStore = useSettingsStore()

const title = import.meta.env.VITE_APP_TITLE
</script>

<template>
  <div class="sidebar-logo-container" :class="{ collapse: collapse }" :style="{ background: settingsStore.navbarColor || '#fff' }">
    <transition name="sidebarLogoFade">
      <div v-if="collapse" key="collapse" class="sidebar-logo-link">
        <div v-if="logo" class="sidebar-logo">
          <!-- <img :src="logo" /> -->
        </div>
        <!-- <h1 v-else class="sidebar-title">{{ title }}</h1> -->
      </div>
      <div v-else key="expand" class="sidebar-logo-link">
        <div v-if="logo" class="sidebar-logo">
          <!-- <img :src="logo" /> -->
        </div>
        <!-- <h1 class="sidebar-title">{{ title }}</h1> -->
      </div>
    </transition>
  </div>
</template>

<style lang="scss" scoped>
.sidebarLogoFade-enter-active {
  transition: opacity 1.5s;
}
.sidebarLogoFade-enter, .sidebarLogoFade-leave-to {
  opacity: 0;
}
.sidebar-logo-container {
  position: relative;
  width: 100%;
  height: 50px;
  line-height: 50px;
  text-align: center;
  overflow: hidden;
  & .sidebar-logo-link {
    display: flex;
    height: 100%;
    width: 100%;
    align-items: center;
    justify-content: center;
    & .sidebar-logo {
      display: inline-block;
      width: 32px;
      height: 32px;
      vertical-align: middle;
      margin-right: 12px;
      img {
        width: 100%;
        height: 100%;
        object-fit: contain;
      }
    }
    & .sidebar-title {
      margin: 0;
      font-weight: 600;
      line-height: 50px;
      font-size: 16px;
      font-family:
        Avenir,
        Helvetica Neue,
        Arial,
        Helvetica,
        sans-serif;
      vertical-align: middle;
      color: var(--el-text-color-primary);
      white-space: nowrap;
      overflow: hidden;
    }
  }
  &.collapse {
    .sidebar-logo {
      margin-right: 0px;
    }
  }
}
</style>
