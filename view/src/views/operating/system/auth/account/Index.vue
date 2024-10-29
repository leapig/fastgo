<script lang="ts">
export default { name: 'OperatingSystemAuthAccount' }
</script>
<script lang="ts" setup>
import {
  getUserAccountList,
  createUser,
  changeUserPhone,
  unbindUserClient,
  getUserRoleGroupList,
  removeUserPermission,
  getUserPermissionList
} from '@/api/operating/system/auth/account'
import usePermissionStore from '@/stores/permission'
import { convertToCamelCase, objectToCamelCase, objectToUnderscore } from '@/utils/common'
import { mobilePhoneValidator } from '@/utils/validate'

const { proxy } = getCurrentInstance() as any

const permissionStore = usePermissionStore()

const list = ref<Array<any>>([])
const total = ref(0)
const loading = ref(false)
const expandRowKeys = ref<Array<string>>([])
const expandLoading = ref<{ [key: string]: boolean }>({})
const showSearch = ref(true)
const addLoading = ref(false)
const showAdd = ref(false)
const editPhoneLoading = ref(false)
const showEditPhone = ref(false)
const permissionLoading = ref(false)
const showPermission = ref(false)
const showAddPermission = ref(false)

const data = reactive({
  queryParams: {
    name: '',
    phone: '',
    page: 1,
    size: 10
  },
  form: {
    name: '',
    phone: ''
  },
  rules: {
    name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
    phone: [
      { required: true, message: '请输入手机号', trigger: 'blur' },
      { validator: mobilePhoneValidator, trigger: 'blur' }
    ]
  },
  editPhone: {
    userName: '',
    form: {
      phone: '',
      user_pk: ''
    },
    rules: {
      phone: [
        { required: true, message: '请输入手机号', trigger: 'blur' },
        { validator: mobilePhoneValidator, trigger: 'blur' }
      ]
    }
  },
  permission: {
    userName: '',
    list: [] as any[]
  },
  addPermission: {
    listIndex: -1,
    type: '' as any,
    user: {
      pk: '',
      name: '',
      role_group_rows: [] as any[],
      role_rows: [] as any[]
    }
  }
})
const { queryParams, form, rules, editPhone, permission, addPermission } = toRefs(data)

/** 获取列表 */
function getList() {
  loading.value = true
  getUserAccountList(queryParams.value).then((response: any) => {
    list.value = response.data.rows
    total.value = response.data.total
    loading.value = false
    if (expandRowKeys.value.length > 0) {
      const keys = [ ...expandRowKeys.value ]
      expandRowKeys.value = []
      nextTick(() => {
        const expandList = list.value.filter((item: any) => keys.indexOf(item.pk) > -1)
        expandList.forEach((item: any) => rowClick(item))
      })
    }
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
  if (!row.roleGroupRows || !row.roleRows) {
    let isExpand = false
    if (expandedRows) {
      for (const index in expandedRows) {
        if (row.pk === expandedRows[index].pk) {
          isExpand = true
          break
        }
      }
    } else {
      isExpand = true
    }
    if (isExpand) {
      const params = {
        user_pk: row.pk,
        enterprise_pk: '1'
      }
      expandLoading.value[row.pk] = true
      getUserRoleGroupList(params).then((response: any) => {
        row.roleGroupRows = response.data.roleGroupRows
        row.roleRows = response.data.roleRows
        expandLoading.value[row.pk] = false
      })
    }
  }
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
/** 重置表单 */
function resetForm() {
  form.value = {
    name: '',
    phone: ''
  }
  const formRef = proxy.$refs['formRef']
  if (formRef) {
    formRef.resetFields()
  }
}
/** 打开新增 */
function handleAdd() {
  resetForm()
  showAdd.value = true
}
/** 新增 */
function confirmAdd() {
  proxy.$refs['formRef'].validate((valid: boolean) => {
    if (valid) {
      addLoading.value = true
      createUser(form.value)
        .then(() => {
          ElMessage.success('新增成功')
          addLoading.value = false
          showAdd.value = false
          getList()
        })
        .catch(() => {
          addLoading.value = false
        })
    }
  })
}
/** 解绑小程序 */
function unbind(command: any) {
  console.log('command',command);
  let typeName = '小程序'
  // let type: number | undefined = 2
 /*  if (command.type === 'mp') {
    typeName = '公众号'
    type = 1
  } else if (command.type === 'ma') {
    typeName = '小程序'
    type = 2
  } */
  
  if (typeName && command.pk) {
    const message = `确定解除 <span style="color: red">${command.name}</span> 的 <span style="color: red">${typeName}</span> 绑定吗？`
    ElMessageBox.confirm(message, '提示', {
      dangerouslyUseHTMLString: true
    })
      .then(() => {
        const data = {
          user_pk: command.pk
        }
        unbindUserClient(data).then(() => {
          ElMessage.success('解绑成功')
          getList()
        })
      })
      .catch(() => {})
  }
}
/** 重置更换手机号码表单 */
function resetEditPhoneForm() {
  editPhone.value.form = {
    phone: '',
    user_pk: ''
  }
  const editPhoneFormRef = proxy.$refs['editPhoneFormRef']
  if (editPhoneFormRef) {
    editPhoneFormRef.resetFields()
  }
}
/** 打开更换手机号码 */
function handleEditPhone(row: any) {
  editPhone.value.userName = ''
  resetEditPhoneForm()
  if (row.pk) {
    editPhone.value.userName = row.name
    editPhone.value.form.user_pk = row.pk
    showEditPhone.value = true
  }
}
/** 更换手机号 */
function confirmEditPhone() {
  proxy.$refs['editPhoneFormRef'].validate((valid: boolean) => {
    if (valid) {
      editPhoneLoading.value = true
      changeUserPhone(editPhone.value.form)
        .then(() => {
          ElMessage.success('更换成功')
          editPhoneLoading.value = false
          showEditPhone.value = false
          getList()
        })
        .catch(() => {
          editPhoneLoading.value = false
        })
    }
  })
}
/** 打开权限 */
function handlePermission(row: any) {
  if (row.pk) {
    permission.value.userName = row.name
    permission.value.list = []
    showPermission.value = true
    const params = {
      user_pk: row.pk,
      enterprise_pk: '1'
    }
    permissionLoading.value = true
    getUserPermissionList(params).then((response: any) => {
      permission.value.list = response.data.permissionRows
      permissionLoading.value = false
    })
  }
}
/** 添加角色/角色组 */
function handleAddPermission(type: string, index: number) {
  addPermission.value.listIndex = index
  addPermission.value.type = type
  addPermission.value.user = objectToUnderscore(list.value[index])
  showAddPermission.value = true
}
/** 列表移除角色/权限提交 */
function tableRemovePermissionSubmit(listIndex: number, type: string, item: any) {
  if (item && item.relationPk) {
    const listItem = list.value[listIndex]
    const userName = listItem.name
    let permissionName = ''
    if (type === 'role_group') {
      permissionName = `角色组（${item.roleGroupName}）`
    } else if (type === 'role') {
      permissionName = `角色（${item.roleName}）`
    }
    ElMessageBox.confirm(`确定移除账号（${userName}）与${permissionName}的关联吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      loading.value = true
      removeUserPermission(item.relationPk).then(() => {
        ElMessage.success('移除成功')
        loading.value = false
        const rows = listItem[convertToCamelCase(type + '_rows')]
        const ind = rows.findIndex((row: any) => row.relationPk === item.relationPk)
        rows.splice(ind, 1)
        // getList()
      }).catch(() => {
        loading.value = false
      })
    }).catch(() => {})
  }
}
/** 角色/权限添加事件处理 */
function handlePermissionAdd(data: any) {
  list.value.splice(addPermission.value.listIndex, 1, objectToCamelCase(data))
}
/** 角色/权限移除事件处理 */
function handlePermissionRemove(data: any) {
  list.value.splice(addPermission.value.listIndex, 1, objectToCamelCase(data))
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
      <el-form-item label="姓名" prop="name">
        <el-input v-model="queryParams.name" placeholder="请输入姓名" clearable style="width: 220px" @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="queryParams.phone" placeholder="请输入手机号" clearable style="width: 220px" @keyup.enter="handleQuery" />
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
        <el-button type="success" plain @click="handleAdd">
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
      <el-table-column label="头像" align="center" width="120">
        <template #default="scope">
          <div class="table-flex-cell">
            <avatar v-model="scope.row.avatar" shape="square"></avatar>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="姓名" align="center" width="120" prop="name" show-overflow-tooltip />
      <el-table-column label="手机" align="center" width="140" prop="phone" />
      <el-table-column label="微信公众号OpenId" align="center" prop="wxOfficialAccountsOpenId" show-overflow-tooltip />
      <el-table-column label="微信小程序OpenId" align="center" prop="wxMiniProgramOpenId" show-overflow-tooltip />
      <el-table-column label="注册时间" align="center" width="160" prop="createAt" />
      <el-table-column label="操作" align="center" width="220" class-name="small-padding fixed-width">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-button link type="primary" @click.stop="handleEditPhone(scope.row)" v-writePermi>
              <el-icon><i-ep-switch /></el-icon>
              <span>更换手机号</span>
            </el-button>
            <el-button link type="primary" @click.stop="unbind(scope.row)" v-if="scope.row.wxMiniProgramOpenId" v-writePermi>
              <el-icon class="ml5"><i-ep-arrow-down /></el-icon>
              <span>解绑小程序</span> <!-- 不需要解绑公众号了 -->
            </el-button>
          <!--   <el-dropdown @command="unbind" v-if="scope.row.wxOfficialAccountsOpenId || scope.row.wxMiniProgramOpenId">
              <el-link type="primary" :underline="false" style="margin-left: 10px">
                <span>解绑</span>
                <el-icon class="ml5"><i-ep-arrow-down /></el-icon>
              </el-link>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item :command="{ type: 'mp', row: scope.row }" v-if="scope.row.wxOfficialAccountsOpenId">
                    <el-icon><i-ep-price-tag /></el-icon>
                    <span>公众号</span>
                  </el-dropdown-item>
                  <el-dropdown-item :command="{ type: 'ma', row: scope.row }" v-if="scope.row.wxMiniProgramOpenId">
                    <el-icon><i-ep-price-tag /></el-icon>
                    <span>小程序</span>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown> -->
          </div>
        </template>
      </el-table-column>
      <el-table-column type="expand">
        <template #default="scope">
          <div class="list-content">
            <el-descriptions :column="1">
              <el-descriptions-item label="角色组：">
                <template v-if="expandLoading[scope.row.pk]">
                  <el-icon class="is-loading"><i-ep-loading /></el-icon>
                </template>
                <template v-else>
                  <template v-for="roleGroup in scope.row.roleGroupRows" :key="roleGroup.pk">
                    <el-tag
                      type="primary"
                      class="desc-tag"
                      :closable="permissionStore.hasWritePermi()"
                      @close="tableRemovePermissionSubmit(scope.$index, 'role_group', roleGroup)"
                    >{{ roleGroup.roleGroupName }}</el-tag>
                  </template>
                  <el-button size="small" @click="handleAddPermission('role_group', scope.$index)" v-writePermi>
                    <el-icon><i-ep-edit /></el-icon>
                    <span>配置</span>
                  </el-button>
                </template>
              </el-descriptions-item>
              <el-descriptions-item label="角色：">
                <template v-if="expandLoading[scope.row.pk]">
                  <el-icon class="is-loading"><i-ep-loading /></el-icon>
                </template>
                <template v-else>
                  <template v-for="role in scope.row.roleRows" :key="role.pk">
                    <el-tag
                      type="primary"
                      class="desc-tag"
                      :closable="permissionStore.hasWritePermi()"
                      @close="tableRemovePermissionSubmit(scope.$index, 'role', role)"
                    >{{ role.roleName }}</el-tag>
                  </template>
                  <el-button size="small" @click="handleAddPermission('role', scope.$index)" v-writePermi>
                    <el-icon><i-ep-edit /></el-icon>
                    <span>配置</span>
                  </el-button>
                </template>
              </el-descriptions-item>
              <el-descriptions-item label="权限：">
                <el-button size="small" @click="handlePermission(scope.row)">
                  <el-icon><i-ep-more /></el-icon>
                  <span>查看</span>
                </el-button>
              </el-descriptions-item>
            </el-descriptions>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.page" v-model:limit="queryParams.size" @pagination="getList" />
    <!-- 新增抽屉 -->
    <el-drawer v-model="showAdd" :size="400" append-to-body>
      <template #header>
        <h4>新增账号</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="90px">
            <el-form-item label="姓名" prop="name">
              <el-input v-model="form.name" placeholder="请输入姓名" style="width: 240px" />
            </el-form-item>
            <el-form-item label="手机号" prop="phone">
              <el-input v-model="form.phone" placeholder="请输入手机号" style="width: 240px" />
            </el-form-item>
          </el-form>
        </div>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="showAdd = false">取 消</el-button>
          <el-button type="primary" :loading="addLoading" style="width: 140px" @click="confirmAdd" v-writePermi>
            <el-icon v-if="!addLoading"><i-ep-edit /></el-icon>
            <span>确 定</span>
          </el-button>
        </div>
      </template>
    </el-drawer>
    <!-- 更换手机号码对话框 -->
    <el-dialog :title="`更换手机号码 - ${editPhone.userName}`" v-model="showEditPhone" width="470px" append-to-body>
      <template #default>
        <el-form ref="editPhoneFormRef" :model="editPhone.form" :rules="editPhone.rules" label-width="90px" style="margin-top: 20px">
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="editPhone.form.phone" placeholder="请输入手机号" style="width: 280px" />
          </el-form-item>
        </el-form>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="showEditPhone = false">取 消</el-button>
          <el-button type="primary" :loading="editPhoneLoading" @click="confirmEditPhone" v-writePermi>
            <el-icon v-if="!editPhoneLoading"><i-ep-edit /></el-icon>
            <span>确 定</span>
          </el-button>
        </div>
      </template>
    </el-dialog>
    <!-- 权限抽屉 -->
    <el-drawer v-model="showPermission" :size="400" append-to-body>
      <template #header>
        <h4>权限 - {{ permission.userName }}</h4>
      </template>
      <template #default>
        <div style="padding-top: 20px">
          <el-scrollbar v-loading="permissionLoading" height="calc(100vh - 68px - 72px - 40px)" :view-style="{ padding: '0 8px' }">
            <div class="permission-list" v-if="permission.list.length > 0">
              <div class="permission-item" v-for="(item, index) in permission.list" :key="'linked-' + index">
                <el-card shadow="hover">
                  <span>{{ item.permissionName }}</span>
                  <el-tag style="margin-left: 10px" type="primary" v-if="item.operationType === 1">只读</el-tag>
                  <el-tag style="margin-left: 10px" type="success" v-else-if="item.operationType === 2">读写</el-tag>
                  <el-tag style="margin-left: 10px" type="success" v-if="item.resourceType === 1">菜单</el-tag>
                  <el-tag style="margin-left: 10px" type="primary" v-else-if="item.resourceType === 2">页面</el-tag>
                </el-card>
              </div>
            </div>
            <el-empty description="暂无数据" v-else />
          </el-scrollbar>
        </div>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="showPermission = false">关 闭</el-button>
        </div>
      </template>
    </el-drawer>
    <!-- 关联角色 -->
    <operating-system-auth-permission-add-permission
      v-model="showAddPermission"
      type="user"
      :add-type="addPermission.type"
      :data="addPermission.user"
      @add="handlePermissionAdd"
      @remove="handlePermissionRemove"
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
.permission-list {
  padding: 8px 5px;
  .permission-item {
    position: relative;
    width: 100%;
    color: #606266;
    :deep(.el-card__body) {
      padding: 12px !important;
    }
  }
  .permission-item:not(:first-child) {
    margin-top: 8px;
  }
}
</style>
