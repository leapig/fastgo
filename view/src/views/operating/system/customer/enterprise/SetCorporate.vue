<script lang="ts" setup>
import { getUserList } from '@/api/operating/system/customer/user'
import { updateEnterpriseCorporateApi } from '@/api/operating/system/customer/enterprise'

const props = defineProps({
  /** 对话框显示状态 */
  modelValue: {
    type: Boolean,
    default: false
  },
  /** 需要设置责任人的目标单位信息 */
  enterprise: {
    required: true,
    type: Object
  }
})

const emit = defineEmits(['update:modelValue', 'update:enterprise', 'close'])

const { proxy } = getCurrentInstance() as any

const queryParams = ref({
  name: '',
  phone: '',
  page: 1,
  size: 10
})
const tableList = ref<Array<any>>([])
const tableTotal = ref(0)
const tableLoading = ref(false)
const setLoading = ref(false)

const show = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const enterpriseComp = computed({
  get: () => props.enterprise,
  set: (val) => emit('update:enterprise', val)
})

const isEnterpriseCorporate = computed(() => {
  return function(pk: string) {
    return pk === enterpriseComp.value.corporate_pk
  }
})

/** 获取列表 */
function tableGetList() {
  tableLoading.value = true
  getUserList(queryParams.value).then((response: any) => {
    tableList.value = response.data.rows
    tableTotal.value = response.data.total
    tableLoading.value = false
  })
}
/** 查询 */
function handleQuery() {
  queryParams.value.page = 1
  tableGetList()
}
/** 重置查询 */
function resetQuery() {
  const queryRef = proxy.$refs['queryRef']
  if (queryRef) {
    queryRef.resetFields()
  }
  handleQuery()
}
/** 修改单位负责人 */
function updateEnterpriseCorporate(row: any) {
  setLoading.value = true
  updateEnterpriseCorporateApi({ corporatePk: row.pk, pk: enterpriseComp.value.pk })
    .then(() => {
      const newEnterprise = { ...props.enterprise }
      newEnterprise.corporateName = row.name
      enterpriseComp.value = newEnterprise
      ElMessage.success('设置负责人成功')
      setLoading.value = false
      close()
    })
    .catch(() => {
      setLoading.value = false
    })
}
/** 关闭对话框 */
function close() {
  show.value = false
  emit('close')
}

watch(show, (value) => {
  if (value) {
    tableList.value = []
    tableTotal.value = 0
    tableGetList()
  }
})
</script>

<template>
  <!-- 设置责任人对话框 -->
  <el-dialog :title="`设置责任人 - ${enterpriseComp.name}`" v-model="show" width="45%" top="10vh" append-to-body @close="close">
    <template #default>
      <div v-loading="setLoading">
        <div class="desc-box">
          <el-descriptions :column="1">
            <el-descriptions-item label="当前责任人：">{{ enterpriseComp.corporateName }}</el-descriptions-item>
          </el-descriptions>
        </div>
        <div style="min-height: 500px">
          <!-- 搜索表单 -->
          <el-form :model="queryParams" ref="queryRef" inline label-width="60px" @submit.prevent>
            <el-form-item label="姓名">
              <el-input v-model="queryParams.name" placeholder="请输入姓名" clearable style="width: 160px" />
            </el-form-item>
            <el-form-item label="手机号">
              <el-input v-model="queryParams.phone" placeholder="请输入手机号" clearable style="width: 160px" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleQuery">
                <el-icon><i-ep-search /></el-icon>
                <span>搜索</span>
              </el-button>
              <el-button @click="resetQuery">
                <el-icon><i-ep-refresh /></el-icon>
                <span>重置</span>
              </el-button>
            </el-form-item>
          </el-form>
          <!-- 表格栏 -->
          <el-table v-loading="tableLoading" :data="tableList">
            <el-table-column label="头像" align="center" prop="name" width="100">
              <template #default="scope">
                <div class="table-flex-cell">
                  <avatar v-model="scope.row.avatar" shape="square"></avatar>
                </div>
              </template>
            </el-table-column>
            <el-table-column label="姓名" align="center" prop="name" show-overflow-tooltip />
            <el-table-column label="手机号" align="center" prop="phone" width="210" />
            <el-table-column label="操作" align="center" width="140" class-name="small-padding fixed-width">
              <template #default="scope">
                <div class="table-flex-cell">
                  <el-popconfirm
                    v-if="!isEnterpriseCorporate(scope.row.pk)"
                    title="确定将此账号设置为责任人吗?"
                    width="260"
                    @confirm="updateEnterpriseCorporate(scope.row)"
                    v-writePermi
                  >
                    <template #reference>
                      <el-button link type="primary">
                        <el-icon><i-ep-document /></el-icon>
                        <span>设置责任人</span>
                      </el-button>
                    </template>
                  </el-popconfirm>
                </div>
              </template>
            </el-table-column>
          </el-table>
          <!-- 分页 -->
          <pagination
            v-show="tableTotal > 0"
            :total="tableTotal"
            v-model:page="queryParams.page"
            v-model:limit="queryParams.size"
            layout="prev, pager, next"
            @pagination="tableGetList"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="close">关 闭</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
.desc-box {
  margin: 15px 10px;
  padding-bottom: 10px;
  border-bottom: 1px var(--el-border-color) var(--el-border-style);
}
</style>
