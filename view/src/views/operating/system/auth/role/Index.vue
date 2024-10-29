<script lang="ts">
export default { name: 'OperatingSystemAuthRole' }
</script>
<script lang="ts" setup>
import {
  getRoleList,
  createRole,
  updateRole,
  removeRole,
  removeRolePermission
} from '@/api/operating/system/auth/role'
import useUserStore from '@/stores/user'
import usePermissionStore from '@/stores/permission'
import { objectToCamelCase, objectToUnderscore } from '@/utils/common'

const { proxy } = getCurrentInstance() as any

const userStore = useUserStore()
const permissionStore = usePermissionStore()

const enterprise = ref<{ [key: string]: string }>({})
const list = ref<Array<any>>([])
const total = ref(0)
const loading = ref(false)
const expandRowKeys = ref<Array<string>>([])
const showSearch = ref(true)
const formLoading = ref(false)
const showForm = ref(false)
const showAddPermission = ref(false)
const showAddUser = ref(false)

const data = reactive({
  queryParams: {
    role_name: '',
    enterprise_pk: '',
    page: 1,
    size: 10
  },
  form: {
    pk: '',
    role_name: '',
    remark: '',
    enterprise_pk: ''
  },
  rules: {
    role_name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }]
  },
  addPermission: {
    listIndex: -1,
    type: '' as any,
    role: {
      pk: '',
      role_name: '',
      permission_group_rows: [] as any[],
      permission_rows: [] as any[]
    }
  },
  addUser: {
    role: {
      pk: '',
      role_name: ''
    }
  }
})
const { queryParams, form, rules, addPermission, addUser } = toRefs(data)

const permissionOperationTypeStr = computed(() => {
  return function(type: number) {
    if (type === 1) {
      return '（只读）'
    } else if (type === 2) {
      return '（读写）'
    }
    return ''
  }
})

/** 获取单位信息 */
function getEnterprise() {
  enterprise.value = userStore.enterprise
  getList()
}
/** 获取列表 */
function getList() {
  loading.value = true
  queryParams.value.enterprise_pk = enterprise.value.pk || '1'
  getRoleList(queryParams.value).then((response: any) => {
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
/** 列表行点击 */
function rowClick(row: any) {
  const index = expandRowKeys.value.indexOf(row.pk)
  if (index > -1) {
    expandRowKeys.value.splice(index, 1)
  } else {
    expandListRow(row)
  }
}
/** 展开列表行 */
function expandListRow(row: any, expandedRows?: any[]) {
  expandRowKeys.value.push(row.pk)
}
/** 展开/折叠操作 */
function toggleExpandAll() {
  if (expandRowKeys.value.length < list.value.length) {
    const noExpandList = list.value.filter((item: any) => expandRowKeys.value.indexOf(item.pk) < 0)
    noExpandList.forEach((item: any) => rowClick(item))
  } else {
    expandRowKeys.value = []
  }
}
/** 打开查询更多用户 */
function handleMoreUsers(row: any) {
  addUser.value.role = objectToUnderscore(row)
  showAddUser.value = true
}
/** 重置表单 */
function resetForm() {
  form.value = {
    pk: '',
    role_name: '',
    remark: '',
    enterprise_pk: enterprise.value.pk || '1'
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
        updateRole(form.value).then(() => {
          ElMessage.success('修改成功')
          formLoading.value = false
          showForm.value = false
          getList()
        }).catch(() => {
          formLoading.value = false
        })
      } else {
        createRole(form.value).then(() => {
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
    ElMessageBox.confirm(`确定删除角色（${row.roleName}）吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removeRole(row.pk).then(() => {
        ElMessage.success('删除成功')
        getList()
      }).catch(() => {
        loading.value = false
      })
    }).catch(() => {})
  }
}
/** 添加权限 */
function handleAddPermission(type: string, index: number) {
  addPermission.value.listIndex = index
  addPermission.value.type = type
  addPermission.value.role = objectToUnderscore(list.value[index])
  showAddPermission.value = true
}
/** 列表移除权限提交 */
function tableRemovePermissionSubmit(roleName: string, type: string, item: any) {
  if (item && item.relationPk) {
    let permissionName = ''
    if (type === 'permission_group') {
      permissionName = `权限组（${item.groupName}）`
    } else if (type === 'permission') {
      permissionName = `权限（${item.permissionName}）`
    }
    ElMessageBox.confirm(`确定移除角色（${roleName}）与${permissionName}的关联吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removeRolePermission(item.relationPk).then(() => {
        ElMessage.success('移除成功')
        loading.value = false
        getList()
      }).catch(() => {
        loading.value = false
      })
    }).catch(() => {})
  }
}
/** 权限添加事件处理 */
function handlePermissionAdd(data: any) {
  list.value.splice(addPermission.value.listIndex, 1, objectToCamelCase(data))
}
/** 权限移除事件处理 */
function handlePermissionRemove(data: any) {
  list.value.splice(addPermission.value.listIndex, 1, objectToCamelCase(data))
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
      <el-form-item label="名称" prop="role_name">
        <el-input v-model="queryParams.role_name" placeholder="请输入角色名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
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
    <el-table v-loading="loading" :data="list" row-key="pk" :expand-row-keys="expandRowKeys" @row-click="rowClick" @expand-change="expandListRow">
      <el-table-column label="序号" align="center" width="90">
        <template #default="scope">
          <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column label="名称" width="260" prop="roleName" show-overflow-tooltip />
      <el-table-column label="用户">
        <template #default="scope">
          <template v-for="user in scope.row.userRows" :key="user.pk">
            <el-button type="primary" round class="avatar-name-button" @click.stop="">
              <avatar v-model="user.avatar" :size="28"></avatar>
              <span class="label" v-if="user.name">{{ user.name }}</span>
            </el-button>
          </template>
          <el-button type="primary" circle :dark="true" @click.stop="handleMoreUsers(scope.row)">
            <el-icon><i-ep-more-filled /></el-icon>
          </el-button>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="190" class-name="small-padding fixed-width">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-button link type="primary" @click.stop="handleForm(scope.row)" v-writePermi>
              <el-icon><i-ep-edit /></el-icon>
              <span>修改</span>
            </el-button>
            <el-button link type="danger" @click.stop="handleRemove(scope.row)" v-writePermi>
              <el-icon><i-ep-delete /></el-icon>
              <span>删除</span>
            </el-button>
          </div>
        </template>
      </el-table-column>
      <el-table-column type="expand">
        <template #default="scope">
          <div class="list-content">
            <el-descriptions :column="1">
              <el-descriptions-item label="备注：">{{ scope.row.remark }}</el-descriptions-item>
              <el-descriptions-item label="权限组：">
                <template v-for="permissionGroup in scope.row.permissionGroupRows" :key="permissionGroup.pk">
                  <el-tag
                    type="primary"
                    class="desc-tag"
                    :closable="permissionStore.hasWritePermi()"
                    @close="tableRemovePermissionSubmit(scope.row.roleName, 'permission_group', permissionGroup)"
                  >{{ permissionGroup.groupName }}</el-tag>
                </template>
                <el-button size="small" @click="handleAddPermission('permission_group', scope.$index)" v-writePermi>
                  <el-icon><i-ep-plus /></el-icon>
                  <span>添加</span>
                </el-button>
              </el-descriptions-item>
              <el-descriptions-item label="权限：">
                <template v-for="permission in scope.row.permissionRows" :key="permission.pk">
                  <el-tag
                    type="primary"
                    class="desc-tag"
                    :closable="permissionStore.hasWritePermi()"
                    @close="tableRemovePermissionSubmit(scope.row.roleName, 'permission', permission)"
                  >{{ permission.permissionName + permissionOperationTypeStr(permission.operationType) }}</el-tag>
                </template>
                <el-button size="small" @click="handleAddPermission('permission', scope.$index)" v-writePermi>
                  <el-icon><i-ep-plus /></el-icon>
                  <span>添加</span>
                </el-button>
              </el-descriptions-item>
            </el-descriptions>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.page" v-model:limit="queryParams.size" @pagination="getList" />
    <!-- 新增/修改抽屉 -->
    <el-drawer v-model="showForm" :size="500" append-to-body>
      <template #header>
        <h4>{{ form.pk ? '修改' : '新增' }}角色</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="110px">
            <el-form-item label="名称" prop="role_name">
              <el-input v-model="form.role_name" placeholder="请输入角色名称" style="width: 300px" />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input type="textarea" v-model="form.remark" placeholder="请输入备注" :rows="4" style="width: 300px" />
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
    <!-- 关联权限对话框 -->
    <operating-system-auth-permission-add-permission
      v-model="showAddPermission"
      type="role"
      :add-type="addPermission.type"
      :data="addPermission.role"
      name-key="role_name"
      @add="handlePermissionAdd"
      @remove="handlePermissionRemove"
      @close="getList"
    />
    <!-- 添加人员对话框 -->
    <operating-system-auth-role-add-user
      v-model="showAddUser"
      type="role"
      :data="addUser.role"
      name-key="role_name"
      @close="getList"
    />
  </div>
</template>

<style lang="scss" scoped>
.list-content {
  padding: 10px 30px;
  .desc-tag {
    margin: 2px 10px 2px 0;
  }
}
</style>
