<script lang="ts" setup>
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import * as content from './content'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  type: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'user' | 'privacyPolicy' | 'productService' | 'bindThirdAccount', unknown>>,
    default: 'user'
  },
  /** 打开后超过多久可关闭，单位秒 */
  closableAfter: {
    type: Number,
    default: 0
  }
})

const emit = defineEmits(['update:modelValue'])

const closableTimeout = ref(0)

const show = computed({
  get: () => props.modelValue,
  set: (val: boolean) => {
    emit('update:modelValue', val)
  }
})

const title = computed(() => {
  return content[props.type].title || ''
})

const lines = computed(() => {
  return content[props.type].lines || []
})

watch(show, (value) => {
  if (value && props.closableAfter > 0) {
    closableTimeout.value = props.closableAfter
    const closableTimer = setInterval(() => {
      if (closableTimeout.value > 0) {
        closableTimeout.value--
      } else {
        clearInterval(closableTimer)
      }
    }, 1000)
  }
})
</script>

<template>
  <el-dialog
    v-model="show"
    :title="title"
    width="600"
    top="5vh"
    center
    destroy-on-close
    append-el-diato-body
    :show-close="false"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
  >
    <el-scrollbar height="77vh">
      <div class="line" v-for="(line, index) in lines" :key="'line-' + index">
        <span>{{ line }}</span>
      </div>
    </el-scrollbar>
    <template #footer>
      <div style="flex: auto">
        <el-button type="primary" :disabled="closableTimeout > 0" @click="show = false">
          <span>确 定{{ closableTimeout > 0 ? ` ${closableTimeout}s` : '' }}</span>
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
.line {
  font-size: 14px;
  text-indent: 2em;
}
</style>
