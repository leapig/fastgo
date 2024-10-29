<script lang="ts" setup>
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import { addUserPermission, removeUserPermission} from '@/api/operating/system/auth/account'
import { getRoleGroupList, addRoleGroupPermission, removeRoleGroupPermission } from '@/api/operating/system/auth/roleGroup'
import { getRoleList, addRolePermission, removeRolePermission } from '@/api/operating/system/auth/role'
import { getPermissionGroupList, addPermissionGroupPermission, removePermissionGroupPermission } from '@/api/operating/system/auth/permissionGroup'
import { getPermissionList } from '@/api/operating/system/auth/permission'
import usePermissionStore from '@/stores/permission'
import { objectToUnderscore } from '@/utils/common'

const permissionStore = usePermissionStore()

const props = defineProps({
  /** 对话框显示状态 */
  modelValue: {
    type: Boolean,
    default: false
  },
  /** 需要添加权限的目标类型 */
  type: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'user' | 'role_group' | 'role' | 'permission_group' | '', unknown>>,
    default: ''
  },
  /** 需要添加权限的目标关联数据 */
  data: {
    required: true,
    type: Object
  },
  /** 关联数据中名称的字段名 */
  nameKey: {
    type: String,
    default: 'name'
  },
  /** 所添加权限的类型 */
  addType: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'role_group' | 'role' | 'permission_group' | 'permission' | '', unknown>>,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'update:data', 'add', 'remove', 'close'])

const { proxy } = getCurrentInstance() as any

const queryParams = ref({
  name: '',
  group_type: undefined,
  visibility: undefined,
  enterprise_pk: '',
  page: 1,
  size: 10
})
const tableList = ref<Array<any>>([])
const tableTotal = ref(0)
const tableLoading = ref(false)
const addLoading = ref(false)

const platformOptions = ref([
  { label: '运营平台', value: 1 },
  { label: '管理平台', value: 2 },
  { label: '用户平台', value: 3 },
  { label: '监管平台', value: 4 }
])

const show = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const dataComp = computed({
  get: () => props.data,
  set: (val) => emit('update:data', val)
})

const typeName = computed(() => {
  if (props.type === 'user') {
    return '账号'
  } else if (props.type === 'role_group') {
    return '角色组'
  } else if (props.type === 'role') {
    return '角色'
  } else if (props.type === 'permission_group') {
    return '权限组'
  }
  return ''
})

const addTypeName = computed(() => {
  if (props.addType === 'role_group') {
    return '角色组'
  } else if (props.addType === 'role') {
    return '角色'
  } else if (props.addType === 'permission_group') {
    return '权限组'
  } else if (props.addType === 'permission') {
    return '权限'
  }
  return ''
})

const addPermissionTableItemExist = computed(() => {
  return function(pk: string) {
    const rows = dataComp.value[props.addType + '_rows']
    let exist = false
    for (const index in rows) {
      if (rows[index].pk === pk) {
        exist = true
        break
      }
    }
    return exist
  }
})

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

/** 添加权限到data.*_rows */
function dataCompRowPush(row: any) {
  const dataNew = { ...dataComp.value }
  dataNew[props.addType + '_rows'].push(row)
  dataComp.value = dataNew
  emit('add', dataNew, row)
}
/** 移除权限自data.*_rows */
function dataCompRowSplice(index: number) {
  const dataNew = { ...dataComp.value }
  const dels = dataNew[props.addType + '_rows'].splice(index, 1)
  dataComp.value = dataNew
  emit('remove', dataNew, dels[0])
}
/** 获取列表 */
function tableGetList() {
  if (props.addType === 'role_group') {
    tableGetRoleGroupList()
  } else if (props.addType === 'role') {
    tableGetRoleList()
  } else if (props.addType === 'permission_group') {
    tableGetPermissionGroupList()
  } else if (props.addType === 'permission') {
    tableGetPermissionList()
  }
}
/** 获取角色组列表 */
function tableGetRoleGroupList() {
  const params = {
    role_group_name: queryParams.value.name,
    enterprise_pk: '1',
    page: queryParams.value.page,
    size: queryParams.value.size
  }
  tableLoading.value = true
  getRoleGroupList(params).then((response: any) => {
    tableList.value = response.data.rows
    tableTotal.value = response.data.total
    tableLoading.value = false
  })
}
/** 获取角色列表 */
function tableGetRoleList() {
  const params = {
    role_name: queryParams.value.name,
    enterprise_pk: '1',
    page: queryParams.value.page,
    size: queryParams.value.size
  }
  tableLoading.value = true
  getRoleList(params).then((response: any) => {
    tableList.value = response.data.rows
    tableTotal.value = response.data.total
    tableLoading.value = false
  })
}
/** 获取权限组列表 */
function tableGetPermissionGroupList() {
  const params = {
    group_name: queryParams.value.name,
    group_type: queryParams.value.group_type,
    enterprise_pk: '1',
    page: queryParams.value.page,
    size: queryParams.value.size
  }
  tableLoading.value = true
  getPermissionGroupList(params).then((response: any) => {
    tableList.value = response.data.rows
    tableTotal.value = response.data.total
    tableLoading.value = false
  })
}
/** 获取权限列表 */
function tableGetPermissionList() {
  const params = {
    permission_name: queryParams.value.name,
    visibility: queryParams.value.visibility,
    enterprise_pk: '1',
    page: queryParams.value.page,
    size: queryParams.value.size
  }
  tableLoading.value = true
  getPermissionList(params).then((response: any) => {
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
/** 添加权限提交 */
function addPermissionSubmit(row: any) {
  if (row && row.pk) {
    if (props.type === 'user') {
      addUserPermissionSubmit(row)
    } else if (props.type === 'role_group') {
      addRoleGroupPermissionSubmit(row)
    } else if (props.type === 'role') {
      addRolePermissionSubmit(row)
    } else if (props.type === 'permission_group') {
      addPermissionGroupPermissionSubmit(row)
    }
  }
}
/** 账号添加角色组/角色提交 */
function addUserPermissionSubmit(row: any) {
  const data = {
    permission_pk: row.pk,
    permission_type: 0,
    user_pk: dataComp.value.pk,
    enterprise_pk: '1'
  }
  if (props.addType === 'role_group') {
    data.permission_type = 1
  } else if (props.addType === 'role') {
    data.permission_type = 2
  }
  if (data.permission_type) {
    addLoading.value = true
    addUserPermission(data).then((response) => {
      ElMessage.success('添加成功')
      row.relationPk = response.data.pk
      dataCompRowPush(objectToUnderscore(row))
      addLoading.value = false
    }).catch(() => {
      addLoading.value = false
    })
  }
}
/** 角色组添加角色/权限组/权限提交 */
function addRoleGroupPermissionSubmit(row: any) {
  const data = {
    permission_pk: row.pk,
    permission_type: 0,
    role_group_pk: dataComp.value.pk
  } as any
  if (props.addType === 'role') {
    data.permission_type = 1
  } else if (props.addType === 'permission_group') {
    data.permission_type = 3
  } else if (props.addType === 'permission') {
    data.permission_type = 2
  }
  if (data.permission_type) {
    addLoading.value = true
    addRoleGroupPermission(data).then((response) => {
      ElMessage.success('添加成功')
      row.relationPk = response.data.pk
      dataCompRowPush(objectToUnderscore(row))
      addLoading.value = false
    }).catch(() => {
      addLoading.value = false
    })
  }
}
/** 角色添加权限组/权限提交 */
function addRolePermissionSubmit(row: any) {
  const data = {
    permission_pk: row.pk,
    permission_type: 0,
    role_pk: dataComp.value.pk
  } as any
  if (props.addType === 'permission_group') {
    data.permission_type = 2
  } else if (props.addType === 'permission') {
    data.permission_type = 1
  }
  if (data.permission_type) {
    addLoading.value = true
    addRolePermission(data).then((response) => {
      ElMessage.success('添加成功')
      row.relationPk = response.data.pk
      dataCompRowPush(objectToUnderscore(row))
      addLoading.value = false
    }).catch(() => {
      addLoading.value = false
    })
  }
}
/** 权限组添加权限提交 */
function addPermissionGroupPermissionSubmit(row: any) {
  const data = {
    permission_pk: row.pk,
    permission_group_pk: dataComp.value.pk
  }
  if (props.addType === 'permission') {
    addLoading.value = true
    addPermissionGroupPermission(data).then((response) => {
      ElMessage.success('添加成功')
      row.relationPk = response.data.pk
      dataCompRowPush(objectToUnderscore(row))
      addLoading.value = false
    }).catch(() => {
      addLoading.value = false
    })
  }
}
/** 移除权限提交 */
function removePermissionSubmit(index: number) {
  const rows = dataComp.value[props.addType + '_rows']
  const item = rows[index]
  const permissionName = item ? `${addTypeName.value}（${item.role_group_name || item.role_name || item.group_name || item.permission_name}）` : ''
  if (item && item.relation_pk) {
    const name = dataComp.value[props.nameKey]
    ElMessageBox.confirm(`确定移除${typeName.value}（${name}）与${permissionName}的关联吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      if (props.type === 'user') {
        removeUserPermissionSubmit(index, item.relation_pk)
      } else if (props.type === 'role_group') {
        removeRoleGroupPermissionSubmit(index, item.relation_pk)
      } else if (props.type === 'role') {
        removeRolePermissionSubmit(index, item.relation_pk)
      } else if (props.type === 'permission_group') {
        removePermissionGroupPermissionSubmit(index, item.relation_pk)
      }
    }).catch(() => {})
  }
}
/** 账号移除角色组/角色提交 */
function removeUserPermissionSubmit(index: number, pk: string) {
  addLoading.value = true
  removeUserPermission(pk).then(() => {
    ElMessage.success('移除成功')
    dataCompRowSplice(index)
    addLoading.value = false
  }).catch(() => {
    addLoading.value = false
  })
}
/** 角色组移除角色/权限组/权限提交 */
function removeRoleGroupPermissionSubmit(index: number, pk: string) {
  addLoading.value = true
  removeRoleGroupPermission(pk).then(() => {
    ElMessage.success('移除成功')
    dataCompRowSplice(index)
    addLoading.value = false
  }).catch(() => {
    addLoading.value = false
  })
}
/** 角色移除权限组/权限提交 */
function removeRolePermissionSubmit(index: number, pk: string) {
  addLoading.value = true
  removeRolePermission(pk).then(() => {
    ElMessage.success('移除成功')
    dataCompRowSplice(index)
    addLoading.value = false
  }).catch(() => {
    addLoading.value = false
  })
}
/** 权限组移除权限提交 */
function removePermissionGroupPermissionSubmit(index: number, pk: string) {
  addLoading.value = true
  removePermissionGroupPermission(pk).then(() => {
    ElMessage.success('移除成功')
    dataCompRowSplice(index)
    addLoading.value = false
  }).catch(() => {
    addLoading.value = false
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
  <!-- 关联权限对话框 -->
  <el-dialog :title="`关联${addTypeName} - ${dataComp[props.nameKey]}`" v-model="show" width="60%" top="10vh" append-to-body @close="close">
    <template #default>
      <div v-loading="addLoading">
        <div style="padding: 15px 10px">
          <el-descriptions :column="1">
            <el-descriptions-item label="已关联：" v-if="props.addType === 'role_group'">
              <template v-for="(roleGroup, index) in dataComp.role_group_rows" :key="roleGroup.pk">
                <el-tag type="primary" class="desc-tag" :closable="permissionStore.hasWritePermi()" @close="removePermissionSubmit(index)">{{ roleGroup.role_group_name }}</el-tag>
              </template>
            </el-descriptions-item>
            <el-descriptions-item label="已关联：" v-if="props.addType === 'role'">
              <template v-for="(role, index) in dataComp.role_rows" :key="role.pk">
                <el-tag type="primary" class="desc-tag" :closable="permissionStore.hasWritePermi()" @close="removePermissionSubmit(index)">{{ role.role_name }}</el-tag>
              </template>
            </el-descriptions-item>
            <el-descriptions-item label="已关联：" v-if="props.addType === 'permission_group'">
              <template v-for="(permissionGroup, index) in dataComp.permission_group_rows" :key="permissionGroup.pk">
                <el-tag type="primary" class="desc-tag" :closable="permissionStore.hasWritePermi()" @close="removePermissionSubmit(index)">{{ permissionGroup.group_name }}</el-tag>
              </template>
            </el-descriptions-item>
            <el-descriptions-item label="已关联：" v-if="props.addType === 'permission'">
              <template v-for="(permission, index) in dataComp.permission_rows" :key="permission.pk">
                <el-tag type="primary" class="desc-tag" :closable="permissionStore.hasWritePermi()" @close="removePermissionSubmit(index)">{{ permission.permission_name + permissionOperationTypeStr(permission.operation_type) }}</el-tag>
              </template>
            </el-descriptions-item>
          </el-descriptions>
        </div>
        <div style="min-height: 500px">
          <!-- 搜索表单 -->
          <el-form :model="queryParams" ref="queryRef" inline label-width="60px" @submit.prevent>
            <el-form-item label="名称" prop="name">
              <el-input v-model="queryParams.name" :placeholder="`请输入${addTypeName}名称`" clearable style="width: 180px" @keyup.enter="handleQuery" />
            </el-form-item>
            <el-form-item label="所属平台" prop="group_type" label-width="80px" v-if="props.addType === 'permission_group'">
              <el-select v-model="queryParams.group_type" placeholder="请选择所属平台" clearable style="width: 220px" @change="handleQuery">
                <el-option v-for="item in platformOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="所属平台" prop="visibility" label-width="80px" v-if="props.addType === 'permission'">
              <el-select v-model="queryParams.visibility" placeholder="请选择所属平台" clearable style="width: 220px" @change="handleQuery">
                <el-option v-for="item in platformOptions" :key="item.value" :label="item.label" :value="item.value" />
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
          <!-- 表格栏 -->
          <el-table v-loading="tableLoading" :data="tableList">
            <el-table-column label="序号" align="center" width="80">
              <template #default="scope">
                <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
              </template>
            </el-table-column>
            <el-table-column label="角色组名称" align="center" prop="roleGroupName" show-overflow-tooltip v-if="props.addType === 'role_group'" />
            <el-table-column label="角色名称" align="center" prop="roleName" show-overflow-tooltip v-if="props.addType === 'role'" />
            <el-table-column label="权限组名称" align="center" prop="groupName" show-overflow-tooltip v-if="props.addType === 'permission_group'" />
            <el-table-column label="权限名称" align="center" prop="permissionName" show-overflow-tooltip v-if="props.addType === 'permission'" />
            <el-table-column label="所属平台" align="center" width="120" show-overflow-tooltip v-if="props.addType === 'permission_group'">
              <template #default="scope">
                <el-tag type="success" v-if="scope.row.group_type === 1">运营平台</el-tag>
                <el-tag type="primary" v-else-if="scope.row.group_type === 2">管理平台</el-tag>
                <el-tag type="primary" v-else-if="scope.row.group_type === 3">用户平台</el-tag>
                <el-tag type="primary" v-else-if="scope.row.group_type === 4">监管平台</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="所属平台" align="center" width="120" show-overflow-tooltip v-if="props.addType === 'permission'">
              <template #default="scope">
                <el-tag type="success" v-if="scope.row.visibility === 1">运营平台</el-tag>
                <el-tag type="primary" v-else-if="scope.row.visibility === 2">管理平台</el-tag>
                <el-tag type="primary" v-else-if="scope.row.visibility === 3">用户平台</el-tag>
                <el-tag type="primary" v-else-if="scope.row.visibility === 4">监管平台</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作类型" align="center" width="90" show-overflow-tooltip v-if="props.addType === 'permission'">
              <template #default="scope">
                <el-tag type="primary" v-if="scope.row.operationType === 1">只读</el-tag>
                <el-tag type="success" v-else-if="scope.row.operationType === 2">读写</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="资源类型" align="center" width="90" show-overflow-tooltip v-if="props.addType === 'permission'">
              <template #default="scope">
                <el-tag type="success" v-if="scope.row.resourceType === 1">菜单</el-tag>
                <el-tag type="primary" v-else-if="scope.row.resourceType === 2">页面</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="90" class-name="small-padding fixed-width">
              <template #default="scope">
                <div class="table-flex-cell">
                  <el-button link type="primary" @click="addPermissionSubmit(scope.row)" v-if="!addPermissionTableItemExist(scope.row.pk)" v-writePermi>
                    <el-icon><i-ep-plus /></el-icon>
                    <span>添加</span>
                  </el-button>
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
.desc-tag {
  margin: 3px 10px 3px 0;
}
</style>