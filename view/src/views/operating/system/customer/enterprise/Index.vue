<script lang="ts">
export default { name: 'OperatingSystemCustomerEnterprise' }
</script>
<script lang="ts" setup>
import { getSecurityEnterpriseList, createSecurityEnterprise, updateSecurityEnterprise } from '@/api/operating/system/customer/securityEnterprise'
import type { GeoLocation } from '@/utils/map'
import { objectToUnderscore } from '@/utils/common'

const { proxy } = getCurrentInstance() as any

const list = ref([])
const total = ref(0)
const loading = ref(false)
const showSearch = ref(true)
const formLoading = ref(false)
const showForm = ref(false)
const showSetCorporate = ref(false)
const setCorporateEnterprise = ref({})

const data = reactive({
  accountData: [], //账号关联的账号数据
  queryParams: {
    name: '',
    address_code: '',
    page: 1,
    size: 10
  },
  form: {
    pk: '',
    name: '',
    cover: '',
    staff_size: '',
    country: '',
    province: '',
    city: '',
    district: '',
    county: '',
    site: '',
    address_code: '',
    longitude: '',
    latitude: ''
  },
  files: {
    cover: undefined,
    license: undefined
  },
  rules: {
    name: [{ required: true, message: '请输入单位名称', trigger: 'blur' }],
    site: [{ required: true, message: '请输入单位详细地址', trigger: 'blur' }]
  }
})
const { queryParams, form, files, rules } = toRefs(data)

/** 地图已选点 */
const selectCoordinate = computed({
  get: () => {
    if (form.value.longitude && form.value.latitude) {
      return {
        longitude: parseFloat(form.value.longitude),
        latitude: parseFloat(form.value.latitude)
      }
    } else {
      return {}
    }
  },
  set: (val) => {
    if (val.longitude && val.latitude) {
      form.value.longitude = val.longitude.toString()
      form.value.latitude = val.latitude.toString()
    } else {
      form.value.longitude = ''
      form.value.latitude = ''
    }
  }
})

/** 获取列表 */
function getList() {
  loading.value = true
  getSecurityEnterpriseList(queryParams.value).then((response: any) => {
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
    name: '',
    cover: '',
    staff_size: '',
    country: '',
    province: '',
    city: '',
    district: '',
    county: '',
    site: '',
    address_code: '',
    longitude: '',
    latitude: ''
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
/** 定点 */
function setMapAddress() {
  if (form.value.site) {
    proxy.$refs['formMap'].setAddressAndMark(form.value.site)
  } else {
    ElMessage({ message: '请先输入详细地址', type: 'error' })
  }
}
/** 地图选点 */
function handleMapSelect(location: GeoLocation) {
  const nForm = form.value
  nForm.province = location.province || ''
  nForm.city = location.city || ''
  nForm.district = location.district || ''
  nForm.county = location.township || ''
  nForm.site = location.fullAddress || ''
  nForm.address_code = location.townCode || location.adCode || ''
  form.value = nForm
}
/** 确认表单 */
function confirmForm() {
  proxy.$refs['formRef'].validate((valid: boolean) => {
    if (valid) {
      formLoading.value = true
      if (form.value.pk) {
        updateSecurityEnterprise(form.value)
          .then(() => {
            ElMessage.success('修改成功')
            formLoading.value = false
            showForm.value = false
            getList()
          })
          .catch(() => {
            formLoading.value = false
          })
      } else {
        const data = { ...form.value, file: files.value.cover }
        createSecurityEnterprise(data)
          .then(() => {
            ElMessage.success('新增成功')
            formLoading.value = false
            showForm.value = false
            getList()
          })
          .catch(() => {
            formLoading.value = false
          })
      }
    }
  })
}
/** 打开设置责任人 */
function handleSetCorporate(row: any) {
  setCorporateEnterprise.value = row
  showSetCorporate.value = true
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
      <el-form-item label="公司名称" prop="name">
        <el-input v-model="queryParams.name" placeholder="请输入公司名称" clearable style="width: 220px" @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="所在地区" prop="area">
        <area-select v-model="queryParams.address_code" placeholder="请选择所在地区" clearable style="width: 220px" @change="handleQuery" />
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
      <el-table-column label="序号" align="center" width="80">
        <template #default="scope">
          <span>{{ (queryParams.page - 1) * queryParams.size + scope.$index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column label="名称" align="center" width="220" prop="name" show-overflow-tooltip />
      <el-table-column label="人员规模" align="center" width="120">
        <template #default="scope">
          <select-show v-model:intoValue="scope.row.staff_size" type="show" selectType="staffSize"></select-show>
        </template>
      </el-table-column>
      <el-table-column label="联系人" align="center" width="120" prop="corporateName" show-overflow-tooltip />
      <el-table-column label="联系方式" align="center" width="140" prop="corporatePhone" />
      <el-table-column label="地址" align="center" prop="createAt" show-overflow-tooltip>
        <template #default="scope">
          <span>{{ scope.row.site || '' }}</span>
        </template>
      </el-table-column>
      <el-table-column label="经纬度" align="center" width="150">
        <template #default="scope">
          <span>{{ scope.row.longitude + ' / ' + scope.row.latitude }}</span>
        </template>
      </el-table-column>
      <el-table-column label="入驻时间" align="center" width="150" prop="createAt" />
      <el-table-column label="操作" align="center" width="240" fixed="right" class-name="small-padding fixed-width">
        <template #default="scope">
          <div class="table-flex-cell">
            <el-button link type="primary" @click="handleForm(scope.row)" v-writePermi>
              <el-icon><i-ep-edit /></el-icon>
              <span>修改</span>
            </el-button>
            <el-button link type="primary" @click="handleSetCorporate(scope.row)" v-writePermi>
              <el-icon><i-ep-link /></el-icon>
              <span>设置责任人</span>
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
        <h4>{{ form.pk ? '修改' : '新增' }}保安公司</h4>
      </template>
      <template #default>
        <div class="drawer-content">
          <el-form ref="formRef" :model="form" :rules="rules" label-width="140px">
            <el-form-item label="名称" prop="name">
              <el-input v-model="form.name" placeholder="请输入公司名称" style="width: 270px" />
            </el-form-item>
            <el-form-item label="人员规模" prop="staff_size">
              <select-show v-model:intoValue="form.staff_size" type="select" selectType="staffSize"></select-show>
            </el-form-item>
            <el-form-item label="详细地址" prop="site">
              <el-input v-model="form.site" placeholder="请输入详细地址" style="width: 270px" @keyup.enter="setMapAddress">
                <template #append>
                  <el-button type="primary" @click="setMapAddress">定点</el-button>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label="经度" prop="longitude">
              <el-input v-model="form.longitude" placeholder="地图选点" disabled style="width: 270px" />
            </el-form-item>
            <el-form-item label="纬度" prop="latitude">
              <el-input v-model="form.latitude" placeholder="地图选点" disabled style="width: 270px" />
            </el-form-item>
            <div style="height: 240px; padding: 10px 30px">
              <map-select
                v-if="showForm"
                ref="formMap"
                v-model="selectCoordinate"
                a-key="supervise-enterprise-form"
                show-bottom-tips
                @select="handleMapSelect"
              ></map-select>
            </div>
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
    <!-- 设置责任人对话框 -->
    <operating-system-customer-enterprise-set-corporate v-model="showSetCorporate" :enterprise="setCorporateEnterprise" @close="handleQuery" />
  </div>
</template>

<style lang="scss" scoped></style>
