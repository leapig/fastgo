<script lang="ts" setup>
defineProps({
  title: {
    type: String
  },
  icon: {
    type: String
  }
})

const emit = defineEmits(['icon-click'])

const route = useRoute()

const routeTitle = ref('')
const routeIcon = ref('')

function iconClick () {
  emit('icon-click')
}

onMounted(() => {
  routeTitle.value = route.meta.title as string || ''
  routeIcon.value = route.meta.icon as string || ''
})
</script>

<template>
  <div class="page-header" v-if="title || routeTitle">
    <template v-if="icon || routeIcon">
      <el-icon size="22" class="icon" v-if="icon" @click="iconClick">
        <svg-icon :icon-name="icon" v-if="icon.startsWith('icon-')" />
        <component :is="$icon[icon]" v-else />
      </el-icon>
      <el-icon size="22" class="icon" v-else-if="routeIcon" @click="iconClick">
        <svg-icon :icon-name="routeIcon" v-if="routeIcon.startsWith('icon-')" />
        <component :is="$icon[routeIcon]" v-else />
      </el-icon>
    </template>
    <div class="title">{{ title || routeTitle }}</div>
  </div>
</template>

<style lang="scss" scoped>
.page-header {
  display: flex;
  padding: 0 10px 15px 10px;
  margin-bottom: 20px;
  border-bottom: 1px solid var(--el-border-color-light);
  align-items: center;
  .icon {
    margin-right: 15px;
    cursor: pointer;
  }
  .title {
    font-size: 20px;
  }
}
</style>
