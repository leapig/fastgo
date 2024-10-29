<script setup lang="ts">
type Select = 'staffSize' | 'deviceType' | 'projectType' | 'postType' | 'oAType' | 'sealType'
type Show = 'select' | 'show'

const props = defineProps({
  intoValue: {},
  size: {
    default: '170px',
    type: String
  },
  type: {
    default: 'select', //show只是展示 select只是选择
    type: String as PropType<Show>
  },
  selectType: {
    type: String as PropType<Select>
  },
  placeholder: {
    default: '请选择人员规模',
    type: String
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

const data = {
  /* 公司规模 */
  staffSize: [
    { label: '50人以下', value: '1', type: 'primary' },
    { label: '50-99人', value: '2', type: 'primary' },
    { label: '100-199人', value: '3', type: 'primary' },
    { label: '200-499人', value: '4', type: 'primary' },
    { label: '500-999人', value: '5', type: 'primary' },
    { label: '1000-1999人', value: '6', type: 'primary' },
    { label: '2000人以上', value: '7', type: 'primary' }
  ],
  /* 设备类型 */
  deviceType: [
    { value: 1, label: '大喇叭 (TSJ-S1)', type: 'primary' },
    { value: 2, label: '对讲机 (TSJ-S5)', type: 'primary' },
    { value: 3, label: '执法记录仪 (TSJ-S7)', type: 'primary' },
    { value: 4, label: '工牌 (TSJ-S3-Q)', type: 'primary' },
    { value: 5, label: '肩灯 (TSJ-S2-Q)', type: 'primary' },
    { value: 6, label: '肩灯 (TSJ-S2-T)', type: 'primary' },
    { value: 7, label: '工牌 (TSJ-S3-T)', type: 'primary' },
    { value: 8, label: '手表 (TSJ-S4-T)', type: 'primary' },
    { value: 9, label: '执法记录仪 (TSJ-S5-PRO)', type: 'primary' },
    { value: 10, label: '对讲机 (TSJ-S5-MAX)', type: 'primary' },
    { value: 11, label: '报警器', type: 'primary' }
  ],
  /* 项目状态 */
  projectType: [
    { value: 1, label: '未开始', type: 'info' },
    { value: 2, label: '进行中', type: 'primary' },
    { value: 3, label: '已完结', type: 'success' },
    { value: 4, label: '已续签', type: 'success' },
    { value: 5, label: '已过期', type: 'warning' }
  ],
  /* 岗位类型 */
  postType: [
    { value: 1, label: '固定岗', type: 'primary' },
    { value: 2, label: '巡逻岗', type: 'success' },
    { value: 3, label: '机动岗', type: 'warning' }
  ],
  /* 设备类型 */
  oAType: [
    { value: 1, label: '加班申请', type: 'primary' },
    { value: 2, label: '补卡申请', type: 'primary' },
    { value: 3, label: '请假申请', type: 'primary' },
    { value: 4, label: '离职申请', type: 'primary' },
    { value: 5, label: '报销申请', type: 'primary' },
    { value: 6, label: '物资申请', type: 'primary' },
    { value: 7, label: '用章申请', type: 'primary' },
    { value: 8, label: 'OA申请', type: 'primary' }
  ],
  /* 用章 */
  sealType: [
    {
      label: '公章',
      value: 1,
      type: 'primary'
    },
    {
      label: '合同专用章',
      value: 2,
      type: 'primary'
    },
    {
      label: '财务专用章',
      value: 3,
      type: 'primary'
    },
    {
      label: '法人章',
      value: 4,
      type: 'primary'
    },
    {
      label: '发票专用章',
      value: 5,
      type: 'primary'
    }
  ]
}
const emit = defineEmits(['update:intoValue', 'change'])

const selectValue: any = computed({
  set(newV) {
    emit('update:intoValue', newV)
  },
  get() {
    return props.intoValue
  }
})

function change(e: any) {
  emit('change', e)
}
</script>

<template>
  <el-select v-model="selectValue" :placeholder="placeholder" :style="{ width: size }" v-if="type === 'select'" @change="change" :disabled="disabled">
    <el-option v-for="item in data[selectType as Select]" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <div v-if="type === 'show'">
    <div v-for="item in data[selectType as Select]" :key="item.value">
      <el-tag :type="item.type as any" v-if="item.value === intoValue">{{ item.label }}</el-tag>
    </div>
  </div>
</template>

<style lang="scss" scoped></style>
