<script lang="ts" setup>

const { proxy } = getCurrentInstance() as any

const props = defineProps({
  /* 配置项 */
  objMap: {
    type: Object,
    default: () => {
      return {
        value: 'id', // ID字段名
        label: 'label', // 显示名称
        children: 'children' // 子级字段名
      }
    }
  },
  /* 自动收起 */
  accordion: {
    type: Boolean,
    default: () => {
      return false
    }
  },
  /** 当前双向数据绑定的值 */
  modelValue: {
    type: [String, Number],
    default: ''
  },
  /** 当前的数据 */
  options: {
    type: Array,
    default: () => []
  },
  /** 输入框内部的文字 */
  placeholder: {
    type: String,
    default: ''
  },
  /** 是否允许清空 */
  clearable: {
    type: Boolean,
    default: () => {
      return false
    }
  },
  /** 是否禁用 */
  disabled: {
    type: Boolean,
    default: () => {
      return false
    }
  }
})

const emit = defineEmits(['update:modelValue'])

const valueId = computed<string|number>({
  get: () => props.modelValue,
  set: (val: string|number) => {
    emit('update:modelValue', val)
  }
});
const valueTitle = ref('')
const defaultExpandedKey = ref<(string|number)[]>([])

function initHandle() {
  nextTick(() => {
    const selectedValue = valueId.value;
    if(selectedValue !== null && typeof (selectedValue) !== 'undefined') {
      const node = proxy.$refs.selectTree.getNode(selectedValue)
      if (node) {
        valueTitle.value = node.data[props.objMap.label]
        proxy.$refs.selectTree.setCurrentKey(selectedValue) // 设置默认选中
        defaultExpandedKey.value = [selectedValue] // 设置默认展开
      }
    } else {
      clearHandle()
    }
  })
}
function handleNodeClick(node: any) {
  valueTitle.value = node[props.objMap.label]
  valueId.value = node[props.objMap.value]
  defaultExpandedKey.value = []
  proxy.$refs.treeSelect.blur()
  selectFilterData('')
}
function selectFilterData(val: string) {
  proxy.$refs.selectTree.filter(val)
}
function filterNode(value: string, data: any) {
  if (!value) return true
  return data[props.objMap['label']].indexOf(value) !== -1
}
function clearHandle() {
  valueTitle.value = ''
  valueId.value = ''
  defaultExpandedKey.value = []
  clearSelected()
}
function clearSelected() {
  const allNode = document.querySelectorAll('#tree-option .el-tree-node')
  allNode.forEach((element) => element.classList.remove('is-current'))
}

onMounted(() => {
  initHandle()
})

watch(valueId, () => {
  initHandle()
})
</script>

<template>
  <div class="el-tree-select">
    <el-select
      style="width: 100%"
      v-model="valueId"
      ref="treeSelect"
      :filterable="true"
      :clearable="clearable"
      @clear="clearHandle"
      :filter-method="selectFilterData"
      :placeholder="placeholder"
      :disabled="disabled"
    >
      <el-option :value="valueId" :label="valueTitle">
        <el-tree
          id="tree-option"
          ref="selectTree"
          :accordion="accordion"
          :data="options"
          :props="objMap"
          :node-key="objMap.value"
          :expand-on-click-node="false"
          :default-expanded-keys="defaultExpandedKey"
          :filter-node-method="filterNode"
          @node-click="handleNodeClick"
        ></el-tree>
      </el-option>
    </el-select>
  </div>
</template>

<style lang='scss' scoped>
.el-tree-select {
  width: 100%;
}

.el-scrollbar .el-scrollbar__view .el-select-dropdown__item {
  padding: 0;
  background-color: #fff;
  height: auto;
}

.el-select-dropdown__item.selected {
  font-weight: normal;
}

ul li .el-tree .el-tree-node__content {
  height: auto;
  padding: 0 20px;
  box-sizing: border-box;
}

:deep(.el-tree-node__content:hover),
:deep(.el-tree-node__content:active),
:deep(.is-current > div:first-child),
:deep(.el-tree-node__content:focus) {
  background-color: mix(#fff, #409eff, 90%);
  color: #409eff;
}
</style>