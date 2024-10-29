<script lang="ts">
export default { name: 'OperatingSystemAuthPermission' }
</script>
<script lang="ts" setup>
import { getPermissionList, createPermission, updatePermission, removePermission } from '@/api/operating/system/auth/permission'
import { getMenuListApi } from '@/api/operating/system/resource/menu'
import { getPageListApi } from '@/api/operating/system/resource/page'
import useUserStore from '@/stores/user'
import { objectToUnderscore } from '@/utils/common'

const { proxy } = getCurrentInstance() as any

const userStore = useUserStore()

const enterprise = ref<{ [key: string]: string }>({})
const list = ref([])
const total = ref(0)
const loading = ref(false)
const showSearch = ref(true)
const formLoading = ref(false)
const showForm = ref(false)

const resourceValidator = function (rule: any, value: string, callback: Function) {
  if (form.value.resource_type === 1 && !value) {
    callback(new Error('请选择菜单'))
  } else if (form.value.resource_type === 2 && !value) {
    callback(new Error('请选择页面'))
  }
  callback()
}

const data = reactive({
  queryParams: {
    permission_name: '',
    visibility: undefined,
    resource_type: undefined,
    enterprise_pk: '',
    page: 1,
    size: 10
  },
  form: {
    pk: '',
    permission_name: '',
    resource_type: 1,
    resource: '',
    operation_type: 2,
    enterprise_pk: '',
    visibility: 1
  },
  rules: {
    permission_name: [{ required: true, message: '请输入权限名称', trigger: 'blur' }],
    resource: [{ required: true, validator: resourceValidator, trigger: 'blur' }]
  },
  visibilityOptions: [
    { label: '运营平台', value: 1 },
    { label: '管理平台', value: 2 },
    { label: '用户平台', value: 3 },
    { label: '监管平台', value: 4 }
  ],
  resourceTypeOptions: [
    { label: '菜单', value: 1 },
    { label: '页面', value: 2 }
  ],
  operationTypeOptions: [
    { label: '只读', value: 1 },
    { label: '读写', value: 2 }
  ],
  formMenu: {
    list: [] as any
  },
  formPage: {
    queryParams: {
      page_name: '',
      page: 1,
      size: 1000
    },
    loading: false,
    list: [] as any
  }
})
const {
  queryParams,
  form,
  rules,
  visibilityOptions,
  resourceTypeOptions,
  operationTypeOptions,
  formMenu,
  formPage
} = toRefs(data)

/** 获取单位信息 */
function getEnterprise() {
  enterprise.value = userStore.enterprise
  getList()
  formGetMenuTreeSelectList()
  formGetPageList()
}
/** 获取列表 */
function getList() {
  loading.value = true
  queryParams.value.enterprise_pk = enterprise.value.pk || '1'
  getPermissionList(queryParams.value).then((response: any) => {
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
    permission_name: '',
    resource_type: 1,
    resource: '',
    operation_type: 2,
    enterprise_pk: enterprise.value.pk || '1',
    visibility: 1
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
        updatePermission(form.value).then(() => {
          ElMessage.success('修改成功')
          formLoading.value = false
          showForm.value = false
          getList()
        }).catch(() => {
          formLoading.value = false
        })
      } else {
        createPermission(form.value).then(() => {
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
/** 表单获取菜单下拉树 */
function formGetMenuTreeSelectList() {
  getMenuListApi({}).then((response: any) => {
    formMenu.value.list = response.data.rows
  })
}
/** 表单获取页面列表 */
function formGetPageList(name?: string) {
  formPage.value.queryParams.page_name = name || ''
  formPage.value.loading = true
  getPageListApi(formPage.value.queryParams).then((response: any) => {
    formPage.value.list = response.data.rows
    formPage.value.loading = false
  })
}
/** 表单资源类型选中 */
function formResourceTypeChange() {
  form.value.resource = ''
}
/** 删除 */
function handleRemove(row: any) {
  if (row.pk) {
    ElMessageBox.confirm(`确定删除权限（${row.permissionName}）吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removePermission(row.pk).then(() => {
        ElMessage.success('删除成功')
        getList()
      }).catch(() => {
        loading.value = false
      })
    }).catch(() => {})
  }
}

onMounted(() => {
  getEnterprise()
})
</script>

<template>
  <div class="app-container">
    <!-- 页头 -->
    <page-header></page-header>
    <!-- 搜索表单 -->
    <el-form :model="queryParams" ref="queryRef" inline v-show="showSearch" label-width="70px" @submit.prevent>
      <el-form-item label="名称" prop="permission_name">
        <el-input v-model="queryParams.permission_name" placeholder="请输入权限名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="所属平台" prop="visibility">
        <el-select v-model="queryParams.visibility" placeholder="请选择所属平台" clearable style="width: 220px" @change="handleQuery">
          <el-option v-for="item in visibilityOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="资源类型" prop="resource_type">
        <el-select v-model="queryParams.resource_type" placeholder="请选择资源类型" clearable style="width: 220px" @change="handleQuery">
          <el-option v-for="item in resourceTypeOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
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
      <el-table-column label="名称" align="center" prop="permissionName" width="420" show-overflow-tooltip />
      <el-table-column label="所属平台" align="center" width="120" show-overflow-tooltip>
        <template #default="scope">
          <el-tag type="success" v-if="scope.row.visibility === 1">运营平台</el-tag>
          <el-tag type="primary" v-else-if="scope.row.visibility === 2">管理平台</el-tag>
          <el-tag type="primary" v-else-if="scope.row.visibility === 3">用户平台</el-tag>
          <el-tag type="primary" v-else-if="scope.row.visibility === 4">监管平台</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作类型" align="center" width="90" show-overflow-tooltip>
        <template #default="scope">
          <el-tag type="primary" v-if="scope.row.operationType === 1">只读</el-tag>
          <el-tag type="success" v-else-if="scope.row.operationType === 2">读写</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="资源类型" align="center" width="90" show-overflow-tooltip>
        <template #default="scope">
          <el-tag type="success" v-if="scope.row.resourceType === 1">菜单</el-tag>
          <el-tag type="primary" v-else-if="scope.row.resourceType === 2">页面</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="目标资源" show-overflow-tooltip>
        <template #default="scope">
          <div v-if="scope.row.resourceType === 1" style="display: flex; align-items: center;">
            <el-icon size="18" v-if="scope.row.menuResource" style="margin-right: 10px;">
              <svg-icon :icon-name="scope.row.menuResource.icon" v-if="scope.row.menuResource.icon.startsWith('icon-')" />
              <component :is="$icon[scope.row.menuResource.icon]" v-else />
            </el-icon>
            <span>{{ scope.row.menuResource?.menuName || '' }}</span>
          </div>
          <span v-else-if="scope.row.resourceType === 2">{{ scope.row.pageResource?.pageName || '' }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="140" class-name="small-padding fixed-width">
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
        <h4>{{ form.pk ? '修改' : '新增' }}权限</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="110px">
            <el-form-item label="名称" prop="permission_name">
              <el-input v-model="form.permission_name" placeholder="请输入权限名称" style="width: 300px" />
            </el-form-item>
            <el-form-item label="所属平台" prop="visibility">
              <el-select v-model="form.visibility" placeholder="请选择所属平台" style="width: 300px">
                <el-option v-for="item in visibilityOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="资源类型" prop="resource_type">
              <el-select v-model="form.resource_type" placeholder="请选择资源类型" clearable style="width: 300px" @change="formResourceTypeChange">
                <el-option v-for="item in resourceTypeOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="菜单" prop="resource" v-if="form.resource_type === 1">
              <el-tree-select
                v-model="form.resource"
                placeholder="请选择菜单"
                :data="formMenu.list"
                :props="{ value: 'pk', label: 'menuName', children: 'childMenus' }"
                value-key="pk"
                check-strictly
                :render-after-expand="false"
                show-checkbox
                style="width: 300px"
              />
            </el-form-item>
            <el-form-item label="页面" prop="resource" v-if="form.resource_type === 2">
              <el-select-v2
                v-model="form.resource"
                placeholder="请选择页面，输入名称筛选"
                filterable
                remote
                :remote-method="formGetPageList"
                :options="formPage.list"
                :props="{ value: 'pk', label: 'pageName' }"
                :loading="formPage.loading"
                style="width: 300px"
              />
            </el-form-item>
            <el-form-item label="操作类型" prop="operation_type">
              <el-radio-group v-model="form.operation_type">
                <el-radio v-for="item in operationTypeOptions" :key="item.value" :value="item.value">{{ item.label }}</el-radio>
              </el-radio-group>
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

