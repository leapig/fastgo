<script lang="ts" setup>
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import { getRoleGroupUserList } from '@/api/operating/system/auth/roleGroup'
import { getRoleUserList } from '@/api/operating/system/auth/role'
import { getUserAccountList, addUserPermission, removeUserPermission } from '@/api/operating/system/auth/account'

const props = defineProps({
  /** 对话框显示状态 */
  modelValue: {
    type: Boolean,
    default: false
  },
  /** 需要添加用户的目标类型 */
  type: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'role_group' | 'role' | '', unknown>>,
    default: ''
  },
  /** 需要添加用户的目标关联数据 */
  data: {
    required: true,
    type: Object
  },
  /** 关联数据中名称的字段名 */
  nameKey: {
    type: String,
    default: 'name'
  }
})

const emit = defineEmits(['update:modelValue', 'update:data', 'add', 'remove', 'close'])

const tableList = ref<Array<any>>([])
const tableTotal = ref(0)
const tableLoading = ref(false)
const addLoading = ref(false)
const showAdd = ref(false)

const sData = reactive({
  queryParams: {
    page: 1,
    size: 10
  },
  add: {
    queryParams: {
      name: '',
      page: 1,
      size: 10
    },
    table: {
      list: [] as any,
      total: 0,
      loading: false
    }
  }
})
const { queryParams, add } = toRefs(sData)

const show = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const dataComp = computed({
  get: () => props.data,
  set: (val) => emit('update:data', val)
})

const typeName = computed(() => {
  if (props.type === 'role_group') {
    return '角色组'
  } else if (props.type === 'role') {
    return '角色'
  }
  return ''
})

/** 获取列表 */
function tableGetList() {
  if (props.type === 'role_group') {
    tableGetRoleGroupUserList()
  } else if (props.type === 'role') {
    tableGetRoleUserList()
  }
}
/** 获取角色组列表 */
function tableGetRoleGroupUserList() {
  const params = {
    role_group_pk: dataComp.value.pk,
    enterprise_pk: dataComp.value.enterprise_pk,
    page: queryParams.value.page,
    size: queryParams.value.size
  }
  tableLoading.value = true
  getRoleGroupUserList(params).then((response: any) => {
    tableList.value = response.data.rows
    tableTotal.value = response.data.total
    tableLoading.value = false
  })
}
/** 获取角色列表 */
function tableGetRoleUserList() {
  const params = {
    role_pk: dataComp.value.pk,
    enterprise_pk: dataComp.value.enterprise_pk,
    page: queryParams.value.page,
    size: queryParams.value.size
  }
  tableLoading.value = true
  getRoleUserList(params).then((response: any) => {
    tableList.value = response.data.rows
    tableTotal.value = response.data.total
    tableLoading.value = false
  })
}
/** 打开添加用户对话框 */
function openAddUser() {
  add.value.queryParams.name = ''
  add.value.table.list = []
  add.value.table.total = 0
  showAdd.value = true
}
/** 获取用户列表 */
function getUserList() {
  if (add.value.queryParams.name) {
    add.value.table.loading = true
    getUserAccountList(add.value.queryParams).then((response: any) => {
      add.value.table.list = response.data.rows
      add.value.table.total = response.data.total
      add.value.table.loading = false
    })
  } else {
    add.value.table.list = []
    add.value.table.total = 0
  }
}
/** 查询用户列表 */
function handleAddQuery() {
  add.value.queryParams.page = 1
  getUserList()
}
/** 添加用户 */
function handleAddUser(row: any) {
  if (row && row.pk) {
    const data = {
      permission_pk: dataComp.value.pk,
      permission_type: 0,
      user_pk: row.pk,
      enterprise_pk: dataComp.value.enterprise_pk || '1'
    }
    if (props.type === 'role_group') {
      data.permission_type = 1
    } else if (props.type === 'role') {
      data.permission_type = 2
    }
    if (data.permission_type) {
      addLoading.value = true
      addUserPermission(data).then(() => {
        ElMessage.success('添加成功')
        addLoading.value = false
        tableGetList()
      }).catch(() => {
        addLoading.value = false
      })
    }
  }
}
/** 移除用户 */
function handleRemoveUser(index: number) {
  const item = tableList.value[index]
  if (item && item.relationPk) {
    const permissionName = `${typeName.value}（${dataComp.value[props.nameKey]}）`
    ElMessageBox.confirm(`确定移除用户（${item.name}）与${permissionName}的关联吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      tableLoading.value = true
      removeUserPermission(item.relationPk).then(() => {
        ElMessage.success('移除成功')
        tableGetList()
      }).catch(() => {
        tableLoading.value = false
      })
    })
  }
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
  <el-drawer v-model="show" :size="500" append-to-body @close="emit('close')">
    <template #header>
      <h4>用户列表 - {{ dataComp[props.nameKey] }}</h4>
    </template>
    <template #default>
      <div class="drawer-content">
        <!-- 操作栏 -->
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5" v-writePermi>
            <el-button type="success" plain @click="openAddUser">
              <el-icon><i-ep-plus /></el-icon>
              <span>添加用户</span>
            </el-button>
          </el-col>
          <right-toolbar :search="false" @queryTable="tableGetList"></right-toolbar>
        </el-row>
        <!-- 表格栏 -->
        <el-table v-loading="tableLoading" :data="tableList">
          <el-table-column label="序号" align="center" width="80">
            <template #default="scope">
              <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column label="头像" align="center" width="80">
            <template #default="scope">
              <div class="table-flex-cell">
                <avatar v-model="scope.row.avatar" shape="square"></avatar>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="姓名" prop="name" show-overflow-tooltip />
          <el-table-column label="操作" align="center" width="120" class-name="small-padding fixed-width">
            <template #default="scope">
              <div class="table-flex-cell">
                <el-button link type="danger" @click.stop="handleRemoveUser(scope.$index)" v-writePermi>
                  <el-icon><i-ep-delete /></el-icon>
                  <span>移除</span>
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
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="show = false">关 闭</el-button>
      </div>
    </template>
  </el-drawer>
  <!-- 添加用户对话框 -->
  <el-dialog :title="`添加用户 - ${dataComp[props.nameKey]}`" v-model="showAdd" width="40%" top="10vh" append-to-body>
    <template #default>
      <div v-loading="addLoading">
        <div style="min-height: 500px">
          <!-- 搜索表单 -->
          <el-form :model="queryParams" ref="addQueryRef" inline label-width="60px" @submit.prevent>
            <el-form-item label="姓名" prop="name">
              <el-input v-model="add.queryParams.name" :placeholder="`请输入姓名`" clearable style="width: 180px" @keyup.enter="handleAddQuery" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleAddQuery">
                <el-icon><i-ep-search /></el-icon>
                <span>搜索</span>
              </el-button>
            </el-form-item>
          </el-form>
          <!-- 表格栏 -->
          <el-table v-loading="add.table.loading" :data="add.table.list" row-key="pk">
            <el-table-column label="序号" align="center" width="90">
              <template #default="scope">
                <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
              </template>
            </el-table-column>
            <el-table-column label="头像" align="center" width="80">
              <template #default="scope">
                <div class="table-flex-cell">
                  <avatar v-model="scope.row.avatar" shape="square"></avatar>
                </div>
              </template>
            </el-table-column>
            <el-table-column label="姓名" align="center" prop="name" />
            <el-table-column label="手机" align="center" prop="phone" />
            <el-table-column label="操作" align="center" width="90" class-name="small-padding fixed-width">
              <template #default="scope">
                <div class="table-flex-cell">
                  <el-button link type="primary" @click="handleAddUser(scope.row)" v-writePermi>
                    <el-icon><i-ep-plus /></el-icon>
                    <span>添加</span>
                  </el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
          <!-- 分页 -->
          <pagination
            v-show="add.table.total > 0"
            :total="add.table.total"
            v-model:page="add.queryParams.page"
            v-model:limit="add.queryParams.size"
            layout="prev, pager, next"
            @pagination="getUserList"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="showAdd = false">关 闭</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped></style>