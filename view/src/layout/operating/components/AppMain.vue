<script lang="ts" setup>
import usePermissionStore from '@/stores/permission'
import type { Menu } from '@/stores/permission'

const permissionStore = usePermissionStore()

const routes = computed(() => permissionStore.routes)

const cacheViews = ref<Array<string>>([])

onMounted(() => {
  initRoutes()
})

function initRoutes() {
  cacheRoutes(routes.value)
}

function cacheRoutes(routes: Menu[]) {
  routes.forEach(route => {
    if (route.component && route.is_cache && route.name) {
      cacheViews.value.push(route.name)
    }
    if (route.children && route.children.length > 0) {
      cacheRoutes(route.children)
    }
  })
}
</script>

<template>
  <section class="app-main">
    <router-view v-slot="{ Component, route }">
      <transition name="fade-transform" mode="out-in">
        <keep-alive :include="cacheViews">
          <component :is="Component" :key="route.path" />
        </keep-alive>
      </transition>
    </router-view>
  </section>
</template>

<style lang="scss" scoped>
.app-main {
  width: 100%;
  position: relative;
  overflow: hidden;
}
</style>
