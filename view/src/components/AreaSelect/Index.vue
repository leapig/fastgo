<script setup lang="ts">
import type { CascaderNodeValue, CascaderNodePathValue, CascaderValue } from 'element-plus'
import areaData from '@/utils/areaData'

// 选择地区组件
const props = defineProps({
  /** 当前选择区域编码绑定值，单选为编码字符串，多选为编码字符串数组 */
  modelValue: {
    type: [String, Array<String>]
  },
  /** 当前选择区域名称绑定值，单选为名称字符串，多选为名称字符串数组 */
  name: {
    type: [String, Array<String>]
  },
  /** 是否多选 */
  multiple: {
    type: Boolean,
    default: false
  },
  /** 是否禁用 */
  disabled: {
    type: Boolean,
    default: false
  },
  /** 是否可清空选项 */
  clearable: {
    type: Boolean,
    default: false
  },
  /** 占位文本 */
  placeholder: {
    type: String,
    default: '请选择地区'
  },
  /** 宽度 */
  width: {
    type: String,
    default: '100%'
  }
})

const emit = defineEmits(['update:modelValue', 'update:name', 'change'])

const selected = computed<CascaderValue>({
  get: () => {
    if (props.multiple) {
      const value = [] as CascaderNodePathValue[]
      if (props.modelValue && props.modelValue.length) {
        for (const ind in props.modelValue as Array<string>) {
          const mv = props.modelValue[ind] as string
          const mva = changeCodeArray(mv) as CascaderNodePathValue
          value.push([...mva])
        }
      }
      return value
    } else {
      const mva = changeCodeArray(props.modelValue as string) as CascaderNodePathValue
      return mva
    }
  },
  set: (val: CascaderValue) => {
    if (props.multiple) {
      const value = [] as Array<string>
      const cv = val as Array<CascaderNodePathValue>
      if (cv && cv.length) {
        for (const ind in cv) {
          const v = cv[ind] as Array<CascaderNodeValue>
          value.push(v[v.length - 1] as string)
        }
      }
      emit('update:modelValue', value)
      updateName(value)
    } else {
      let value = ''
      const cv = val as Array<CascaderNodeValue>
      if (cv && cv.length) {
        value = cv[cv.length - 1] as string
      }
      emit('update:modelValue', value)
      updateName(value)
    }
  }
})
/** 将区域编码转成层级编码数组 */
function changeCodeArray(code: string) {
  const codes = [] as Array<string>
  if (code) {
    if (code.length === 2) {
      codes.push(code)
    } else if (code.length === 4) {
      codes.push(code.slice(0, 2), code)
    } else if (code.length === 6) {
      codes.push(code.slice(0, 2), code.slice(0, 4), code)
    } else if (code.length === 9) {
      codes.push(code.slice(0, 2), code.slice(0, 4), code.slice(0, 6), code)
    }
  }
  return codes
}
/** 更新props.name为选中区域的名称 */
function updateName(code: string | Array<string>) {
  if (props.name !== undefined) {
    if (props.multiple) {
      const names = [] as Array<string>
      const codes = code as Array<string>
      for (const ind in codes) {
        const nCode = codes[ind]
        const codeArray = changeCodeArray(nCode)
        const name = getNameByCodeArray(codeArray)
        names.push(name)
      }
      emit('update:name', names)
    } else {
      const nCode = code as string
      const codeArray = changeCodeArray(nCode)
      const name = getNameByCodeArray(codeArray)
      emit('update:name', name)
    }
  }
}
/** 从层级编码数组中获取最终区域名称 */
function getNameByCodeArray(codeArray: Array<string>) {
  let name = ''
  let nowAreaData = areaData as Array<any>
  for (let i = 0; i < codeArray.length; i++) {
    const codeA = codeArray[i]
    const codeAData = nowAreaData.filter((area: any) => area.code === codeA)
    if (codeAData && codeAData.length) {
      if (i < codeArray.length - 1 && codeAData[0].children) {
        nowAreaData = codeAData[0].children
      } else {
        name = codeAData[0].name
      }
    }
  }
  return name
}
/** 选择区域 */
function handleChange() {
  emit('change', selected.value)
}
</script>

<template>
  <el-cascader
    v-model="selected"
    ref="selectAreaRef"
    :options="areaData"
    :props="{ value: 'code', label: 'name', children: 'children', multiple: multiple, checkStrictly: true }"
    :disabled="disabled"
    :clearable="clearable"
    :style="{ width: width }"
    :placeholder="placeholder"
    :collapse-tags="multiple"
    @change="handleChange"
  />
</template>

<style lang="scss" scoped></style>
