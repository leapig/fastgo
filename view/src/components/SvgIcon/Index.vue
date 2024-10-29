<script lang="ts" setup>
const props = defineProps({
  iconName: {
    type: String,
    required: true
  }
})

// 图标在 iconfont 中的名字
const iconName = computed(() => {
  return `#${props.iconName}`
})
</script>

<template>
  <svg aria-hidden="true">
    <use :xlink:href="iconName" />
  </svg>
</template>

<style lang="scss" scoped></style>
