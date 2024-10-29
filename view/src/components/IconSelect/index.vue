<script lang="ts" setup>
import { ClickOutside as vClickOutside } from 'element-plus'
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import { defaultType, types, icons } from './requireIcons'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: '点击选择图标'
  },
  width: {
    type: String,
    default: '100%'
  }
})

const emit = defineEmits(['update:modelValue'])

const { proxy } = getCurrentInstance() as any

const popoverPlacement = ref<EpPropMergeType<StringConstructor, 'bottom-start' | 'top-start', unknown>>('bottom-start')
const showChooseIcon = ref(false)
const searchName = ref('')
const selectType = ref(defaultType)

const selectedIcon = computed({
  get: () => props.modelValue,
  set: (val) => {
    emit('update:modelValue', val)
  }
})

const iconList = computed<Array<string>>(() => {
  let list = (icons as any)[selectType.value]
  if (searchName.value) {
    list = list.filter((item: string) => item.indexOf(searchName.value) > -1)
  }
  return list
})

/** 显示图标选择框 */
function beforeShowSelectIcon() {
  searchName.value = ''
  const iconInputRef = proxy.$refs['iconInputRef']
  if (iconInputRef) {
    const iconInputRect = iconInputRef.ref.getBoundingClientRect()
    if (window.innerHeight - 300 < iconInputRect.top) {
      popoverPlacement.value = 'top-start'
    } else {
      popoverPlacement.value = 'bottom-start'
    }
  }
}
/** 图标外层点击隐藏图标选择框 */
function hideSelectIcon(event: any) {
  if (showChooseIcon.value) {
    const elem = event.relatedTarget || event.srcElement || event.target || event.currentTarget
    const className = elem.className
    const parentClassName = elem.parentElement.parentElement.className
    if (className === 'el-input__inner' && parentClassName.indexOf('icon-search') > -1) {
      showChooseIcon.value = true 
    } else if (className === 'el-radio-button__inner' && parentClassName.indexOf('icon-type') > -1) {
      showChooseIcon.value = true
    } else {
      showChooseIcon.value = false
    }
  }
}
/** 选择图标 */
function selectIcon(icon: string) {
  selectedIcon.value = icon
}
</script>

<template>
  <el-popover
    :placement="popoverPlacement"
    :width="540"
    v-model:visible="showChooseIcon"
    trigger="click"
    @before-enter="beforeShowSelectIcon"
  >
    <template #reference>
      <el-input
        ref="iconInputRef"
        v-model="selectedIcon"
        :placeholder="placeholder"
        v-click-outside="hideSelectIcon"
        readonly
        :style="{ width }"
      >
        <template #prefix>
          <template v-if="selectedIcon">
            <el-icon size="18" color="var(--el-input-text-color)">
              <svg-icon :icon-name="selectedIcon" v-if="selectedIcon.startsWith('icon-')" />
              <component :is="$icon[selectedIcon]" v-else />
            </el-icon>
          </template>
          <template v-else>
            <el-icon><i-ep-search /></el-icon>
          </template>
        </template>
      </el-input>
    </template>
    <div class="icon-body">
      <el-input
        ref="iconSearchRef"
        v-model="searchName"
        class="icon-search"
        clearable
        placeholder="输入图标名称搜索"
      >
        <template #suffix>
          <el-icon><i-ep-search /></el-icon>
        </template>
      </el-input>
      <div class="icon-type">
        <el-radio-group v-model="selectType" class="icon-type">
          <el-radio-button :label="item.title" :value="item.key" v-for="item in types" :key="item.key" />
        </el-radio-group>
      </div>
      <el-scrollbar height="200px">
        <div class="icon-list">
          <div class="list-container">
            <div v-for="(item, index) in iconList" class="icon-item-wrapper" :class="{ active: selectedIcon === item }" :key="index" @click="selectIcon(item)">
              <div class="icon-item">
                <el-icon size="25">
                  <svg-icon :icon-name="item" v-if="item.startsWith('icon-')" />
                  <component :is="$icon[item]" v-else />
                </el-icon>
                <span>{{ item }}</span>
              </div>
            </div>
          </div>
        </div>
      </el-scrollbar>
    </div>
  </el-popover>
</template>

<style lang="scss" scoped>
.icon-body {
  padding: 7px;
  .icon-search {
    position: relative;
    margin-bottom: 10px;
  }
  .icon-type {
    margin-bottom: 10px;
  }
  .icon-list {
    height: 200px;
    padding-right: 5px;
    .list-container {
      display: flex;
      flex-wrap: wrap;
      .icon-item-wrapper {
        width: calc(100% / 3 - 2px);
        height: 35px;
        line-height: 25px;
        cursor: pointer;
        display: flex;
        align-items: center;
        .icon-item {
          display: flex;
          max-width: 100%;
          height: 25px;
          padding: 0 5px;
          span {
            display: inline-block;
            vertical-align: -0.15em;
            fill: currentColor;
            padding-left: 5px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
          }
        }
        &:hover {
          background: #ececec;
          border-radius: 3px;
        }
        &.active {
          background: #ececec;
          border-radius: 3px;
        }
      }
    }
  }
}
</style>