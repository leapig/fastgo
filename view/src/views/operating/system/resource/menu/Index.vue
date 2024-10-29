<script lang="ts">
export default { name: 'OperatingSystemResourceMenu' }
</script>
<script lang="ts" setup>
import { getMenuListApi, createMenuApi, updateMenuApi, removeMenuApi } from '@/api/operating/system/resource/menu'
import { getPageListApi } from '@/api/operating/system/resource/page'
import { objectToUnderscore } from '@/utils/common'
import { isHttp } from '@/utils/validate'

const { proxy } = getCurrentInstance() as any

const list = ref<Array<any>>([])
const loading = ref(false)
const tableExpanding = ref(false)
const isExpandAll = ref(false)
const showSearch = ref(true)
const formLoading = ref(false)
const showForm = ref(false)

const resourceKeyValidator = function (rule: any, value: string, callback: Function) {
  if (form.value.menu_type === 1 && !value) {
    callback(new Error('请选择页面'))
  } else if (form.value.menu_type === 2) {
    if (!value) {
      callback(new Error('请输入链接地址'))
    } else if (!isHttp(value)) {
      callback(new Error('链接地址不合法'))
    }
  }
  callback()
}

const pathValidator = function (rule: any, value: string, callback: Function) {
  if (form.value.menu_type === 1 || form.value.menu_type === 3) {
    if (!value) {
      callback(new Error('请输入路由地址'))
    }
  }
  callback()
}

const data = reactive({
  queryParams: {
    menu_name: ''
  },
  form: {
    pk: '',
    parent_pk: '1',
    menu_name: '',
    menu_type: 1,
    resource_key: '',
    path: '',
    icon: '',
    sort: 0
  },
  rules: {
    menu_name: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
    menu_type: [{ required: true, message: '请选择菜单类型', trigger: 'blur' }],
    resource_key: [{ required: true, validator: resourceKeyValidator, trigger: 'blur' }],
    path: [{ required: true, validator: pathValidator, trigger: 'blur' }]
  },
  menuTypeOptions: [
    { label: '页面', value: 1 },
    { label: '链接', value: 2 },
    { label: '目录', value: 3 },
  ],
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
const { queryParams, form, rules, menuTypeOptions, formPage } = toRefs(data)

const treeSelectList = computed<Array<any>>(() => {
  return [{ pk: '1', menuName: "主目录", childMenus: list.value }]
})

/** 获取列表 */
function getList() {
  loading.value = true
  getMenuListApi(queryParams.value).then((response: any) => {
    list.value = response.data.rows
    loading.value = false
  })
}
/** 查询 */
function handleQuery() {
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
    parent_pk: '1',
    menu_name: '',
    menu_type: 1,
    resource_key: '',
    path: '',
    icon: '',
    sort: 0
  }
  const formRef = proxy.$refs['formRef']
  if (formRef) {
    formRef.resetFields()
  }
}
/** 展开/折叠操作 */
function toggleExpandAll() {
  tableExpanding.value = true
  isExpandAll.value = !isExpandAll.value
  nextTick(() => {
    tableExpanding.value = false
  })
}
/** 打开表单 */
function handleForm(row?: any) {
  resetForm()
  if (row) {
    form.value = objectToUnderscore(row)
  }
  showForm.value = true
}
/** 打开添加子菜单 */
function handleAdd(row?: any) {
  resetForm()
  if (row) {
    form.value.parent_pk = row.pk
  }
  showForm.value = true
}
/** 确认表单 */
function confirmForm() {
  proxy.$refs['formRef'].validate((valid: boolean) => {
    if (valid) {
      formLoading.value = true
      if (form.value.pk) {
        updateMenuApi(form.value).then(() => {
          ElMessage.success('修改成功')
          formLoading.value = false
          showForm.value = false
          getList()
        }).catch(() => {
          formLoading.value = false
        })
      } else {
        createMenuApi(form.value).then(() => {
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
/** 表单获取页面列表 */
function formGetPageList(name?: string) {
  formPage.value.queryParams.page_name = name || ''
  formPage.value.loading = true
  getPageListApi(formPage.value.queryParams).then((response: any) => {
    formPage.value.list = response.data.rows
    formPage.value.loading = false
  })
}
/** 删除 */
function handleRemove(row: any) {
  if (row.pk) {
    ElMessageBox.confirm(`确定删除菜单（${row.menuName}）吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removeMenuApi(row.pk).then(() => {
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
  formGetPageList()
})
</script>

<template>
  <div class="app-container">
    <!-- 页头 -->
    <page-header></page-header>
    <!-- 搜索表单 -->
    <el-form :model="queryParams" ref="queryRef" inline v-show="showSearch" label-width="70px" @submit.prevent>
      <el-form-item label="名称" prop="menu_name">
        <el-input v-model="queryParams.menu_name" placeholder="请输入菜单名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
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
      <el-col :span="1.5">
        <el-button type="info" plain @click="toggleExpandAll">
          <el-icon><i-ep-sort /></el-icon>
          <span>展开/折叠</span>
        </el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>
    <!-- 表格栏 -->
    <el-table
      v-if="!tableExpanding"
      v-loading="loading"
      :data="list"
      row-key="pk"
      :default-expand-all="isExpandAll"
      :tree-props="{ children: 'childMenus' }"
    >
      <el-table-column label="名称" prop="menuName" width="280px" show-overflow-tooltip />
      <el-table-column label="类型" align="center" width="90px" show-overflow-tooltip>
        <template #default="scope">
          <el-tag type="primary" v-if="scope.row.menuType === 1">页面</el-tag>
          <el-tag type="info" v-else-if="scope.row.menuType === 2">链接</el-tag>
          <el-tag type="success" v-else-if="scope.row.menuType === 3">目录</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="目标资源" show-overflow-tooltip>
        <template #default="scope">
          <span v-if="scope.row.menuType === 1">{{ scope.row.pageMessage.pageName }}</span>
          <span v-else-if="scope.row.menuType === 2">{{ scope.row.resourceKey }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标" align="center" prop="icon" width="80px">
        <template #default="scope">
          <el-icon size="18">
            <svg-icon :icon-name="scope.row.icon" v-if="scope.row.icon.startsWith('icon-')" />
            <component :is="$icon[scope.row.icon]" v-else />
          </el-icon>
        </template>
      </el-table-column>
      <el-table-column label="排序" align="center" prop="sort" width="60px" />
      <el-table-column label="添加时间" align="center" prop="createTime" width="180px" />
      <el-table-column label="修改时间" align="center" prop="updateTime" width="180px" />
      <el-table-column label="操作" align="center" width="240" class-name="small-padding fixed-width">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-button link type="success" @click="handleAdd(scope.row)" v-if="scope.row.menuType === 3" v-writePermi>
              <el-icon><i-ep-plus /></el-icon>
              <span>新增</span>
            </el-button>
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
    <!-- 新增/修改抽屉 -->
    <el-drawer v-model="showForm" :size="500" append-to-body>
      <template #header>
        <h4>{{ form.pk ? '修改' : '新增' }}菜单</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="110px">
            <el-form-item label="父级" prop="parent_pk">
              <el-tree-select
                v-model="form.parent_pk"
                placeholder="请选择菜单"
                :data="treeSelectList"
                :props="{ value: 'pk', label: 'menuName', children: 'childMenus' }"
                value-key="pk"
                check-strictly
                :render-after-expand="false"
                show-checkbox
                style="width: 300px"
              />
            </el-form-item>
            <el-form-item label="名称" prop="menu_name">
              <el-input v-model="form.menu_name" placeholder="请输入菜单名称" style="width: 300px" />
            </el-form-item>
            <el-form-item label="类型" prop="menu_type">
              <el-select v-model="form.menu_type" placeholder="请选择菜单类型" style="width: 300px">
                <el-option v-for="item in menuTypeOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="页面" prop="resource_key" v-if="form.menu_type === 1">
              <el-select-v2
                v-model="form.resource_key"
                placeholder="请选择页面，输入名称筛选"
                style="width: 300px"
                filterable
                remote
                :remote-method="formGetPageList"
                :options="formPage.list"
                :props="{ value: 'pk', label: 'pageName' }"
                :loading="formPage.loading"
              />
            </el-form-item>
            <el-form-item label="链接地址" prop="resource_key" v-if="form.menu_type === 2">
              <el-input v-model="form.resource_key" placeholder="请输入链接地址" style="width: 300px" />
            </el-form-item>
            <el-form-item label="路由地址" prop="path" v-if="form.menu_type === 1 || form.menu_type === 3">
              <el-input v-model="form.path" placeholder="请输入路由地址" style="width: 300px" />
            </el-form-item>
            <el-form-item label="图标" prop="icon">
              <icon-select v-model="form.icon" width="300px" />
            </el-form-item>
            <el-form-item label="显示排序" prop="sort">
              <el-input-number v-model="form.sort" controls-position="right" :min="0" />
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
