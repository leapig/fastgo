<script lang="ts">
export default { name: 'OperatingSystemProductShelf' }
</script>
<script lang="ts" setup>
import { getUserList } from '@/api/operating/system/customer/user'
import { getSexFromIdcard } from '@/utils/idcard'

const { proxy } = getCurrentInstance() as any

const list = ref([])
const total = ref(0)
const loading = ref(false)
const showSearch = ref(true)

const data = reactive({
  queryParams: {
    name: '',
    phone: '',
    page: 1,
    size: 10
  }
})
const { queryParams } = toRefs(data)

const getSex = computed(() => {
  return function (idCard: string) {
    const sex = getSexFromIdcard(idCard)
    let str = ''
    switch (sex) {
      case 1:
        str = '男'
        break
      case 2:
        str = '女'
        break
    }
    return str
  }
})

/** 获取列表 */
function getList() {
  loading.value = true
  getUserList(queryParams.value).then((response: any) => {
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

onMounted(() => {
  getList()
})
</script>

<template>
  <div class="app-container">
    <!-- 页头 -->
    <page-header></page-header>
    <!-- 搜索表单 -->
    <el-form :model="queryParams" ref="queryRef" inline v-show="showSearch" label-width="70px">
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
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>
    <!-- 表格栏 -->
    <el-table v-loading="loading" :data="list">
      <el-table-column label="序号" align="center" width="90">
        <template #default="scope">
          <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column label="头像" align="center" width="100">
        <template #default="scope">
          <div class="table-flex-cell">
            <avatar v-model="scope.row.avatar" shape="square"></avatar>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="姓名" align="center" width="120" prop="name" show-overflow-tooltip />
      <el-table-column label="手机" align="center" width="140" prop="phone" />
      <el-table-column label="身份证号" align="center" prop="idCard" />
      <el-table-column label="性别" align="center" width="90">
        <template #default="scope">
          <div class="table-flex-cell">
            <span>{{ getSex(scope.row.idCard) }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="生日" align="center" width="120" prop="birthday" />
      <el-table-column label="实名状态" align="center" width="120">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-tag type="success" v-if="scope.row.isRealName === 1">已实名</el-tag>
            <el-tag type="info" v-else>未实名</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="所在城市" align="center" width="120" prop="liveCity" />
      <el-table-column label="注册时间" align="center" width="160" prop="createAt" />
      <el-table-column label="最后活跃" align="center" width="160" prop="lastLiveAt" />
    </el-table>
    <!-- 分页 -->
    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.page" v-model:limit="queryParams.size" @pagination="getList" />
  </div>
</template>

<style lang="scss" scoped></style>
