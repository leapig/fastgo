<script setup lang="ts">
import useUserStore from '@/stores/user'

const props = defineProps({
  limitNum: {
    default: 10,
    type: Number
  },
  fileList: {
    default: []
  }
})

const token = ref<any>('')
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
let waitFileList = ref<any>([])
waitFileList.value = props.fileList

watch(
  () => props.fileList,
  () => {
    waitFileList.value = props.fileList
  }
)

const emits = defineEmits(['uploadSuccess', 'updateFile'])

// 删除
const handleRemove = (file: any, uploadFiles: any) => {
  if (file.size < 4194304) {
    emits('updateFile', file, uploadFiles)
  }
}

// 预览
const handlePreview = (uploadFile: any) => {
  dialogImageUrl.value = uploadFile.url!
  dialogVisible.value = true
}

// 超出限制
const handleExceed = () => {
  ElMessage.warning(`上传限制 ${props.limitNum} 个文件`)
}

// 删除之前的确定
const beforeRemove = (uploadFile: any, uploadFiles: any) => {
  // 过大文件直接删除
  if (uploadFile.size > 4194304) {
    return true
  } else {
    return ElMessageBox.confirm(`确定移除 ${uploadFile.name}文件吗? 此操作不可逆!!!`, '温馨提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(
      () => true,
      () => false
    )
  }
}

// 处理上传之前
const handleBeforeUpload = (uploadFile: any) => {
  if (uploadFile.size > 4194304) {
    ElMessage.error('文件过大，超出4MB')
    return false
  } else {
    return true
  }
}

// 处理上传成功
const handleSuccess = (file: any, uploadFiles: any) => {
  emits('uploadSuccess', file.data)
  console.log(file, uploadFiles)
}

onMounted(async () => {
  token.value = await useUserStore().getTenantToken()
})
</script>

<template>
  <el-upload
    v-model:file-list="waitFileList"
    class="upload-demo"
    multiple
    action="/open-apis/core/tool/file/general"
    :headers="{
      Authorization: `Bearer ${token}`
    }"
    :on-preview="handlePreview"
    :before-upload="handleBeforeUpload"
    :on-remove="handleRemove"
    :on-success="handleSuccess"
    :before-remove="beforeRemove"
    :limit="limitNum"
    :on-exceed="handleExceed"
    list-type="picture-card"
  >
    <el-button type="primary">上传文件</el-button>
    <template #tip>
      <div class="el-upload__tip">支持jpg/jpeg/png；图片大小不能超过4M；最多上传{{ limitNum }}个文件</div>
    </template>
  </el-upload>

  <el-dialog v-model="dialogVisible">
    <img w-full :src="dialogImageUrl" alt="Preview Image" style="width: 600px" />
  </el-dialog>
</template>

<style lang="scss" scoped>
.previewDownload {
  margin-top: 80px;
  margin-left: 50px;
  display: flex;
  flex-direction: column;

  .el-link {
    height: 40px;
  }
}
</style>
