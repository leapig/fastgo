<script lang="ts" setup>
import useUserStore from '@/stores/user'
import type { UploadProps } from 'element-plus'
const props = defineProps({
  type: {
    type: String,
    default: ''
  }
})

const emits=defineEmits(['uploadSuccess'])
const token = ref<any>('')
const fileList = ref<any>([])
const loading = ref(false)
const uploadMap = ref<any>({
  insure: '/open-apis/personnel/insure/import',//保险
  social:'/open-apis/personnel/social/security/import',//导入社保信息(tenant_access_token)
  socialDetail:'/open-apis/personnel/social/security/detail/import',//导入社保明细(tenant_access_token)
})

/* 处理文件删除 */
const handleRemove: UploadProps['onRemove'] = (file: any) => {
  console.log(file)
}
/* 处理文件预览 */
const handlePreview: UploadProps['onPreview'] = (uploadFile: any) => {}
/* 处理超出限制 */
const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
  ElMessage.error('上传超出限制')
}
/* 处理文件删除之前的确认 */
const beforeRemove: UploadProps['beforeRemove'] = (uploadFile, uploadFiles) => {
  return true
}
/* 文件上传之前的处理 */
const handleBeforeUpload = (uploadFile: any) => {
  if (uploadFile.size > 4194304) {
    ElMessage.error('文件过大，超出4MB')
    return false
  } else {
    loading.value = true
    return true
  }
}

/* 处理上传成功 */
const handleSuccess = (file: any, uploadFiles: any) => {
  loading.value = false
  ElMessage.success('上传成功')
  emits('uploadSuccess')
}

/* 处理文件上传失败 */
const handleError = (error: any) => {
  let errorData = JSON.parse(error.message)
  console.log('aa', JSON.parse(error.message))
  ElMessage.error(errorData.errorMessage)
}
/* 下载模板 */
const downloadTemplate = () => {
  console.log('下载模板')
}

onMounted(async () => {
  token.value = await useUserStore().getTenantToken()
})
</script>

<template>
  <div>
    <el-upload
      v-model:file-list="fileList"
      class="upload-demo"
      :action="uploadMap[type]"
      :headers="{
        Authorization: `Bearer ${token}`
      }"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :before-remove="beforeRemove"
      :before-upload="handleBeforeUpload"
      :on-success="handleSuccess"
      :limit="1"
      :on-exceed="handleExceed"
      :on-error="handleError"
      accept=".xltx,.xlsx"
    >
      <el-button type="primary">点击上传</el-button>
      <template #tip>
        <div class="el-upload__tip">只允许上传一个.xltx，.xlsx类型的文件</div>
      </template>
    </el-upload>
  </div>
  <div @click="downloadTemplate" style="cursor: pointer;">请先下载模板之后在上传</div>
</template>

<style lang="scss" scoped></style>
