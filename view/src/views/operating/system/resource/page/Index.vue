<script lang="ts">
export default { name: 'OperatingSystemResourcePage' }
</script>
<script lang="ts" setup>
import {
  getPageListApi,
  createPageApi,
  updatePageApi,
  removePageApi,
  linkPageInterfaceApi,
  removePageInterfaceApi
} from '@/api/operating/system/resource/page'
import { getInterfaceListApi } from '@/api/operating/system/resource/interface'
import { objectToUnderscore } from '@/utils/common'
import { existViewValidator } from '@/utils/validate'

const { proxy } = getCurrentInstance() as any

const list = ref([])
const total = ref(0)
const loading = ref(false)
const showSearch = ref(true)
const formLoading = ref(false)
const showForm = ref(false)
const linkInterfaceLoading = ref(false)
const showLinkInterface = ref(false)
const interfaceBoxTableHeight = ref(0)

const componentValidator = function (rule: any, value: string, callback: Function) {
  if (form.value.page_type === 1) {
    existViewValidator(rule, value, callback)
  } else {
    callback()
  }
}

const data = reactive({
  queryParams: {
    page_name: '',
    page_type: '',
    component_name: '',
    page: 1,
    size: 10
  },
  form: {
    pk: '',
    page_name: '',
    page_type: 1,
    component: '',
    component_name: '',
    is_cache: 1
  },
  rules: {
    page_name: [{ required: true, message: '请输入页面名称', trigger: 'blur' }],
    component: [
      { required: true, message: '请输入组件路径', trigger: 'blur' },
      { validator: componentValidator, trigger: 'blur' }
    ]
  },
  linkInterface: {
    page: {
      pk: '',
      page_name: '',
      rows: [] as any
    },
    table: {
      queryParams: {
        interface_name: '',
        interface_url: '',
        page: 1,
        size: 10
      },
      list: [],
      total: 0,
      loading: false
    }
  },
  pageTypeOptions: [
    { label: '平台', value: 1 },
    { label: '小程序', value: 2 },
  ]
})
const { queryParams, form, rules, linkInterface, pageTypeOptions } = toRefs(data)

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

const isPageInterfaceExist = computed(() => {
  return function (interface_pk: string) {
    let exist = false
    for (const index in linkInterface.value.page.rows) {
      if (linkInterface.value.page.rows[index].interface_pk === interface_pk) {
        exist = true
        break
      }
    }
    return exist
  }
})

/** 获取列表 */
function getList() {
  loading.value = true
  getPageListApi(queryParams.value).then((response: any) => {
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
    page_name: '',
    page_type: 1,
    component: '',
    component_name: '',
    is_cache: 1
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
      if (form.value.page_type === 2) {
        form.value.is_cache = 0
      }
      formLoading.value = true
      if (form.value.pk) {
        updatePageApi(form.value).then(() => {
          ElMessage.success('修改成功')
          formLoading.value = false
          showForm.value = false
          getList()
        }).catch(() => {
          formLoading.value = false
        })
      } else {
        createPageApi(form.value).then(() => {
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
    ElMessageBox.confirm(`确定删除页面（${row.pageName}）吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removePageApi(row.pk).then(() => {
        ElMessage.success('删除成功')
        loading.value = true
        getList()
      }).catch(() => {
        loading.value = false
      })
    }).catch(() => {})
  }
}
/** 打开关联接口 */
function handleLinkInterface(row: any) {
  linkInterface.value.page = objectToUnderscore(row)
  showLinkInterface.value = true
  linkInterface.value.table.queryParams.size = 10
  resetLinkInterfaceQuery()
}
/** 关联接口 - 接口查询 */
function handleLinkInterfaceQuery() {
  linkInterface.value.table.queryParams.page = 1
  linkInterfaceGetInterfaceList()
}
/** 关联接口 - 重置接口查询 */
function resetLinkInterfaceQuery() {
  const queryRef = proxy.$refs['linkInterfaceQueryRef']
  if (queryRef) {
    queryRef.resetFields()
  }
  handleLinkInterfaceQuery()
}
/** 关联接口 - 获取接口列表 */
function linkInterfaceGetInterfaceList() {
  linkInterface.value.table.loading = true
  getInterfaceListApi(linkInterface.value.table.queryParams).then((response: any) => {
    linkInterface.value.table.list = response.data.rows
    linkInterface.value.table.total = response.data.total
    linkInterface.value.table.loading = false
    interfaceBoxTableHeight.value = proxy.$refs['linkInterfaceBoxTableRef'].offsetHeight
  })
}
/** 关联接口 - 添加关联接口 */
function handleAddPageInterface(row: any) {
  const item = objectToUnderscore(row)
  item.interface_pk = item.pk
  delete item.pk
  const data = {
    interface_pk: item.interface_pk,
    operation_type: item.interface_way.toUpperCase() === 'GET' ? 1 : 2,
    page_pk: linkInterface.value.page.pk
  }
  linkInterfaceLoading.value = true
  linkPageInterfaceApi(data).then((response) => {
    linkInterfaceLoading.value = false
    if (response.data.pk) {
      ElMessage.success('添加成功')
      item.pk = response.data.pk
      linkInterface.value.page.rows.push(item)
    }
  }).catch(() => {
    linkInterfaceLoading.value = false
  })
}
/** 关联接口 - 移除关联接口 */
function handleRemovePageInterface(index: number) {
  const item = linkInterface.value.page.rows[index]
  if (item) {
    ElMessageBox.confirm(`确定移除页面接口（${item.interface_name}）关联吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      linkInterfaceLoading.value = true
      removePageInterfaceApi(item.pk).then(() => {
        ElMessage.success('移除成功')
        linkInterfaceLoading.value = false
        linkInterface.value.page.rows.splice(index, 1)
      }).catch(() => {
        linkInterfaceLoading.value = false
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
      <el-form-item label="名称" prop="page_name">
        <el-input v-model="queryParams.page_name" placeholder="请输入页面名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="类型" prop="page_type">
        <el-select v-model="queryParams.page_type" placeholder="请选页面类型" clearable style="width: 220px" @change="handleQuery">
          <el-option v-for="item in pageTypeOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="组件名称" prop="component_name">
        <el-input v-model="queryParams.component_name" placeholder="请输入组件名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
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
      <el-table-column label="名称" align="center" prop="pageName" width="320" show-overflow-tooltip />
      <el-table-column label="类型" align="center" width="90">
        <template #default="scope">
          <el-tag type="success" v-if="scope.row.pageType === 1">平台</el-tag>
          <el-tag type="primary" v-else-if="scope.row.pageType === 2">小程序</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="组件名称" align="center" prop="componentName" show-overflow-tooltip />
      <el-table-column label="组件路径" align="center" prop="component" show-overflow-tooltip />
      <el-table-column label="是否缓存" align="center" width="120">
        <template #default="scope">
          <el-tag type="success" v-if="scope.row.isCache">是</el-tag>
          <el-tag type="info" v-else>否</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="240" class-name="small-padding fixed-width">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-button link type="primary" @click="handleForm(scope.row)" v-writePermi>
              <el-icon><i-ep-edit /></el-icon>
              <span>修改</span>
            </el-button>
            <el-button link type="primary" @click="handleLinkInterface(scope.row)" v-writePermi>
              <el-icon><i-ep-link /></el-icon>
              <span>关联接口</span>
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
        <h4>{{ form.pk ? '修改' : '新增' }}页面</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="110px">
            <el-form-item label="名称" prop="page_name">
              <el-input v-model="form.page_name" placeholder="请输入页面名称" style="width: 300px" />
            </el-form-item>
            <el-form-item label="类型" prop="page_type">
              <el-select v-model="form.page_type" placeholder="请选页面类型" style="width: 300px">
                <el-option v-for="item in pageTypeOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="组件名称" prop="component_name">
              <el-input v-model="form.component_name" placeholder="请输入组件名称" style="width: 300px" />
            </el-form-item>
            <el-form-item label="组件路径" prop="component">
              <el-input v-model="form.component" placeholder="请输入组件路径" style="width: 300px" />
            </el-form-item>
            <el-form-item label="是否缓存" prop="is_cache" v-if="form.page_type === 1">
              <el-switch v-model="form.is_cache" :active-value="1" :inactive-value="0" />
              <el-tooltip effect="dark" content="要使页面缓存生效，必须输入页面组件名称，且与代码中指定的组件名称一致。" placement="bottom">
                <el-icon size="18" color="var(--el-text-color-regular)" style="margin: 0 10px"><i-ep-info-filled /></el-icon>
              </el-tooltip>
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
    <!-- 关联接口对话框 -->
    <el-dialog :title="`关联接口 - ${linkInterface.page.page_name}`" v-model="showLinkInterface" width="80%" top="10vh" append-to-body @close="getList">
      <template #default>
        <div class="link-interface-box">
          <div class="table" ref="linkInterfaceBoxTableRef">
            <!-- 搜索表单 -->
            <el-form :model="linkInterface.table.queryParams" ref="linkInterfaceQueryRef" inline label-width="50px" @submit.prevent>
              <el-form-item label="名称" prop="interface_name">
                <el-input v-model="linkInterface.table.queryParams.interface_name" placeholder="请输入接口名称" clearable style="width: 180px" @keyup.enter="handleLinkInterfaceQuery" />
              </el-form-item>
              <el-form-item label="路径" prop="interface_url">
                <el-input v-model="linkInterface.table.queryParams.interface_url" placeholder="请输入接口路径" clearable style="width: 200px" @keyup.enter="handleLinkInterfaceQuery" />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleLinkInterfaceQuery">
                  <el-icon><i-ep-search /></el-icon>
                  <span>搜索</span>
                </el-button>
                <el-button @click="resetLinkInterfaceQuery">
                  <el-icon><i-ep-refresh /></el-icon>
                  <span>重置</span>
                </el-button>
              </el-form-item>
            </el-form>
            <!-- 表格栏 -->
            <el-table v-loading="linkInterface.table.loading" :data="linkInterface.table.list">
              <el-table-column label="序号" align="center" width="80">
                <template #default="scope">
                  <span>{{ (linkInterface.table.queryParams.page - 1) * linkInterface.table.queryParams.size + scope.$index + 1 }}</span>
                </template>
              </el-table-column>
              <el-table-column label="名称" align="center" prop="interfaceName" width="280" show-overflow-tooltip />
              <el-table-column label="方法" align="center" prop="interfaceWay" width="90">
                <template #default="scope">
                  <div class="table-flex-cell">
                    <el-tag :type="interfaceWayTagType(scope.row.interfaceWay)">{{ scope.row.interfaceWay }}</el-tag>
                  </div>
                </template>
              </el-table-column>
              <el-table-column label="路径" align="center" prop="interfaceUrl" show-overflow-tooltip />
              <el-table-column label="标识符" align="center" prop="interfaceKey" width="180" show-overflow-tooltip />
              <el-table-column label="操作" align="center" width="90" class-name="small-padding fixed-width">
                <template #default="scope">
                  <div class="table-flex-cell">
                    <el-button link type="primary" @click="handleAddPageInterface(scope.row)" v-if="!isPageInterfaceExist(scope.row.pk)" v-writePermi>
                      <span>添加</span>
                      <el-icon style="margin-left: 5px"><i-ep-d-arrow-right /></el-icon>
                    </el-button>
                  </div>
                </template>
              </el-table-column>
            </el-table>
            <!-- 分页 -->
            <pagination
              v-show="linkInterface.table.total > 0"
              :total="linkInterface.table.total"
              v-model:page="linkInterface.table.queryParams.page"
              v-model:limit="linkInterface.table.queryParams.size"
              layout="prev, pager, next"
              @pagination="linkInterfaceGetInterfaceList"
            />
          </div>
          <div class="linked">
            <div class="title">
              <span>已关联接口：</span>
            </div>
            <el-scrollbar v-loading="linkInterfaceLoading" :height="interfaceBoxTableHeight - 28" :max-height="interfaceBoxTableHeight || 500" :view-style="{ padding: '0 8px' }">
              <div class="list">
                <div class="item" v-for="(item, index) in linkInterface.page.rows" :key="'linked-' + index">
                  <el-card shadow="hover">
                    <div class="item-row">
                      <el-tag :type="interfaceWayTagType(item.interface_way)">{{ item.interface_way }}</el-tag>
                      <span style="margin-left: 8px;" :title="item.interface_name">{{ item.interface_name }}</span>
                    </div>
                    <div class="item-row">
                      <span :title="item.interface_url">{{ item.interface_url }}</span>
                    </div>
                  </el-card>
                  <div class="remove-btn" @click="handleRemovePageInterface(index)" v-writePermi>
                    <el-icon color="#fff"><i-ep-delete /></el-icon>
                  </div>
                </div>
              </div>
            </el-scrollbar>
          </div>
        </div>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="showLinkInterface = false">关 闭</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style lang="scss" scoped>
.link-interface-box {
  display: flex;
  width: 100%;
  .table {
    flex: 1;
    min-height: 500px;
  }
  .linked {
    width: 330px;
    margin-left: 10px;
    .title {
      padding: 0 13px;
      line-height: 28px;
      font-size: 16px;
      color: #606266;
    }
    .list {
      padding: 8px 5px;
      .item {
        position: relative;
        width: 100%;
        color: #606266;
        :deep(.el-card__body) {
          padding: 12px !important;
        }
        .item-row {
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
        }
        .item-row:not(:first-child) {
          margin-top: 5px;
        }
      }
      .item:not(:first-child) {
        margin-top: 8px;
      }
      .item:hover .remove-btn {
        display: flex;
      }
      .remove-btn {
        display: none;
        align-items: center;
        justify-content: center;
        position: absolute;
        top: -5px;
        right: -5px;
        width: 25px;
        height: 25px;
        border-radius: 50%;
        background-color: #F56C6C;
        cursor: pointer;
      }
    }
  }
}
</style>
