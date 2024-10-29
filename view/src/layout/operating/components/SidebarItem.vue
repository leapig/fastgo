<script lang="ts" setup>
import { isExternal } from '@/utils/validate'

const props = defineProps({
  item: {
    type: Object,
    required: true
  },
  isNest: {
    type: Boolean,
    default: false
  },
  basePath: {
    type: String,
    default: ''
  }
})

function resolvePath(routePath: string) {
  if (isExternal(routePath)) {
    return routePath
  }
  if (isExternal(props.basePath)) {
    return props.basePath
  }
  routePath = routePath ? '/' + routePath : ''
  return props.basePath + routePath
}
</script>

<template>
  <div v-if="!item.hidden">
    <template v-if="!item.children || item.children.length === 0">
      <router-link :to="resolvePath(item.path)" class="no-underline">
        <el-menu-item :index="resolvePath(item.path)" :class="{ 'submenu-title-noDropdown': !isNest }">
          <el-icon size="20">
            <svg-icon :icon-name="item.icon" v-if="item.icon && item.icon.startsWith('icon-')" />
            <component :is="item.icon ? $icon[item.icon] : $icon['Menu']" v-else />
          </el-icon>
          <template #title>
            <span class="menu-title">{{ item.title }}</span>
          </template>
        </el-menu-item>
      </router-link>
    </template>
    <el-sub-menu v-else :index="resolvePath(item.path)">
      <template #title>
        <el-icon size="20">
          <svg-icon :icon-name="item.icon" v-if="item.icon && item.icon.startsWith('icon-')" />
          <component :is="item.icon ? $icon[item.icon] : $icon['Menu']" v-else />
        </el-icon>
        <span class="menu-title">{{ item.title }}</span>
      </template>
      <sidebar-item class="nest-menu" v-for="child in item.children" :key="resolvePath(item.path) + '/' + child.path" is-nest :item="child" :base-path="resolvePath(item.path)" />
    </el-sub-menu>
  </div>
</template>

<style lang="scss" scoped></style>
