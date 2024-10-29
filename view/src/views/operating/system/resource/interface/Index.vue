<script lang="ts">
export default { name: 'OperatingSystemResourceInterface' }
</script>
<script lang="ts" setup>
import { getInterfaceListApi, createInterfaceApi, updateInterfaceApi, removeInterfaceApi } from '@/api/operating/system/resource/interface'
import { objectToUnderscore } from '@/utils/common'

const { proxy } = getCurrentInstance() as any

const list = ref([])
const total = ref(0)
const loading = ref(false)
const showSearch = ref(true)
const formLoading = ref(false)
const showForm = ref(false)

const data = reactive({
  queryParams: {
    interface_name: '',
    interface_key: '',
    interface_way: '',
    page: 1,
    size: 10
  },
  form: {
    pk: '',
    interface_name: '',
    interface_key: '',
    interface_way: '',
    interface_url: ''
  },
  rules: {
    interface_name: [{ required: true, message: '请输入接口名称', trigger: 'blur' }],
    interface_way: [{ required: true, message: '请选择接口方法', trigger: 'blur' }],
    interface_url: [{ required: true, message: '请输入接口路径', trigger: 'blur' }]
  },
  interfaceWayOptions: [
    { label: 'GET', value: 'GET' },
    { label: 'POST', value: 'POST' },
    { label: 'PUT', value: 'PUT' },
    { label: 'DELETE', value: 'DELETE' }
  ]
})
const { queryParams, form, rules, interfaceWayOptions } = toRefs(data)

const interfaceWayTagType = computed(() => {
  return function (way: string) {
    if (way.toUpperCase() === 'GET') {
      return 'primary'
    } else if (way.toUpperCase() === 'POST') {
      return 'success'
    } else if (way.toUpperCase() === 'PUT') {
      return 'warning'
    } else if (way.toUpperCase() === 'DELETE') {
      return 'danger'
    } else {
      return 'info'
    }
  }
})

/** 获取列表 */
function getList() {
  loading.value = true
  getInterfaceListApi(queryParams.value).then((response: any) => {
    list.value = response.data.rows
    total.value = response.data.total
    loading.value = false
  })
}
/** 查询 */
function handleQuery() {
  queryParams.value.page = 1
  getList()
}
/** 重置查询 */
function resetQuery() {
  const queryRef = proxy.$refs['queryRef']
  if (queryRef) {
    queryRef.resetFields()
  }
  handleQuery()
}
/** 重置表单 */
function resetForm() {
  form.value = {
    pk: '',
    interface_name: '',
    interface_key: '',
    interface_way: '',
    interface_url: ''
  }
  const formRef = proxy.$refs['formRef']
  if (formRef) {
    formRef.resetFields()
  }
}
/** 打开表单 */
function handleForm(row?: any) {
  resetForm()
  if (row) {
    form.value = objectToUnderscore(row)
  }
  showForm.value = true
}
/** 确认表单 */
function confirmForm() {
  proxy.$refs['formRef'].validate((valid: boolean) => {
    if (valid) {
      formLoading.value = true
      if (form.value.pk) {
        updateInterfaceApi(form.value).then(() => {
          ElMessage.success('修改成功')
          formLoading.value = false
          showForm.value = false
          getList()
        }).catch(() => {
          formLoading.value = false
        })
      } else {
        createInterfaceApi(form.value).then(() => {
          ElMessage.success('新增成功')
          formLoading.value = false
          showForm.value = false
          getList()
        }).catch(() => {
          formLoading.value = false
        })
      }
    }
  })
}
/** 删除 */
function handleRemove(row: any) {
  if (row.pk) {
    ElMessageBox.confirm(`确定删除接口（${row.interfaceName}）吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removeInterfaceApi(row.pk).then(() => {
        ElMessage.success('删除成功')
        loading.value = true
        getList()
      }).catch(() => {
        loading.value = false
      })
    }).catch(() => {})
  }
}

onMounted(() => {
  getList()
})
</script>

<template>
  <div class="app-container">
    <!-- 页头 -->
    <page-header></page-header>
    <!-- 搜索表单 -->
    <el-form :model="queryParams" ref="queryRef" inline v-show="showSearch" label-width="70px" @submit.prevent>
      <el-form-item label="名称" prop="interface_name">
        <el-input v-model="queryParams.interface_name" placeholder="请输入接口名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="方法" prop="interface_way">
        <el-select v-model="queryParams.interface_way" placeholder="请选择接口方法" clearable style="width: 220px" @change="handleQuery">
          <el-option v-for="item in interfaceWayOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="标识符" prop="interface_key">
        <el-input v-model="queryParams.interface_key" placeholder="请输入接口标识符" clearable style="width: 220px" @keyup.enter="handleQuery" />
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
    <!-- 操作栏 -->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5" v-writePermi>
        <el-button type="success" plain @click="handleForm()">
          <el-icon><i-ep-plus /></el-icon>
          <span>新增</span>
        </el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>
    <!-- 表格栏 -->
    <el-table v-loading="loading" :data="list">
      <el-table-column label="序号" align="center" width="90">
        <template #default="scope">
          <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column label="名称" align="center" prop="interfaceName" width="220" show-overflow-tooltip />
      <el-table-column label="方法" align="center" prop="interfaceWay" width="120">
        <template #default="scope">
          <el-tag :type="interfaceWayTagType(scope.row.interfaceWay)">{{ scope.row.interfaceWay }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="路径" align="center" prop="interfaceUrl" show-overflow-tooltip />
      <el-table-column label="标识符" align="center" prop="interfaceKey" show-overflow-tooltip />
      <el-table-column label="操作" align="center" width="160" class-name="small-padding fixed-width">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-button link type="primary" @click="handleForm(scope.row)" v-writePermi>
              <el-icon><i-ep-edit /></el-icon>
              <span>修改</span>
            </el-button>
            <el-button link type="danger" @click="handleRemove(scope.row)" v-writePermi>
              <el-icon><i-ep-delete /></el-icon>
              <span>删除</span>
            </el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.page" v-model:limit="queryParams.size" @pagination="getList" />
    <!-- 新增/修改抽屉 -->
    <el-drawer v-model="showForm" :size="500" append-to-body>
      <template #header>
        <h4>{{ form.pk ? '修改' : '新增' }}接口</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="110px">
            <el-form-item label="名称" prop="interface_name">
              <el-input v-model="form.interface_name" placeholder="请输入接口名称" style="width: 300px" />
            </el-form-item>
            <el-form-item label="方法" prop="interface_way">
              <el-select v-model="form.interface_way" placeholder="请选择接口方法" clearable style="width: 300px">
                <el-option v-for="item in interfaceWayOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="路径" prop="interface_url">
              <el-input v-model="form.interface_url" placeholder="请输入接口路径" style="width: 300px" />
            </el-form-item>
            <el-form-item label="标识符" prop="interface_key">
              <el-input v-model="form.interface_key" placeholder="请输入接口标识符" style="width: 300px" />
            </el-form-item>
          </el-form>
        </div>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="showForm = false">取 消</el-button>
          <el-button type="primary" :loading="formLoading" style="width: 140px" @click="confirmForm" v-writePermi>
            <el-icon v-if="!formLoading"><i-ep-edit /></el-icon>
            <span>确 定</span>
          </el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<style lang="scss" scoped></style>

