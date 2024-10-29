<script lang="ts" setup>
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import type { UploadProps } from 'element-plus'
import { getUserAccessToken, getTenantAccessToken } from '@/utils/auth'
import { isHttp, isBase64Image } from '@/utils/validate'

const props = defineProps({
  /**
   * 当前显示头像地址/Base64绑定值，不包括手动上传所选但未上传的头像
   * 根据业务逻辑需要，图片上传后该值为接口返回的name
   */
  modelValue: {
    type: String
  },
  /** 初始化显示图片地址/Base64，为空时图片地址使用modelValue绑定值 */
  src: {
    type: String
  },
  /** 当前手动上传所选头像 */
  file: {
    type: File
  },
  /** 头像显示宽高度 */
  size: {
    type: Number,
    default: 40
  },
  /** 形状样式 */
  shape: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'circle' | 'square', unknown>>,
    default: 'circle'
  },
  /** 是否开启上传 */
  enableUpload: {
    type: Boolean,
    default: false
  },
  /** 是否自动上传 */
  autoUpload: {
    type: Boolean,
    default: true
  },
  /** 自动上传目标 */
  autoUploadTarget: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'device' | 'general' | 'privacy' | 'system', unknown>>,
    default: 'general'
  }
})

const emit = defineEmits(['update:modelValue', 'update:file'])

const { proxy } = getCurrentInstance() as any

const loading = ref(false)
const url = ref('')
const fileList = ref<any>([])

const uploadUrl = computed(() => {
  const baseUrl = import.meta.env.VITE_APP_BASE_API + '/core/tool/file'
  return `${baseUrl}/${props.autoUploadTarget}`
})

const headers = computed<any>(() => {
  const header = {} as any
  if (props.autoUploadTarget === 'device' || props.autoUploadTarget === 'general') {
    const tenantToken = getTenantAccessToken()
    if (tenantToken) {
      header.Authorization = `Bearer ${tenantToken}`
    }
  } else if (props.autoUploadTarget === 'privacy' || props.autoUploadTarget === 'system') {
    const userToken = getUserAccessToken()
    if (userToken) {
      header.Authorization = `Bearer ${userToken}`
    }
  }
  return header
})

const sizeCom = computed(() => {
  return props.size > 20 ? props.size : 20
})

const shapeStyle = computed(() => {
  if (props.shape === 'circle') {
    return '50%'
  } else if (props.shape === 'square') {
    return 'var(--el-avatar-border-radius)'
  } else {
    return '0'
  }
})

/** 加载头像 */
function loadAvatar() {
  if (props.src) {
    if (isHttp(props.src) || isBase64Image(props.src)) {
      url.value = props.src
    }
  }
  if (!url.value && props.modelValue) {
    if (isHttp(props.modelValue) || isBase64Image(props.modelValue)) {
      url.value = props.modelValue
    }
  }
}
/** 头像上传前的回调 */
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (rawFile.type !== 'image/jpeg' && rawFile.type !== 'image/png') {
    ElMessage.error('头像只支持 jpeg、png 格式')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('头像不能超过 2MB')
    return false
  }
  loading.value = true
  return true
}
/** 头像上传成功回调 */
const handleAvatarSuccess: UploadProps['onSuccess'] = (response: any) => {
  loading.value = false
  url.value = response.data.url
  emit('update:modelValue', response.data.name)
}
/** 头像上传失败回调 */
const handleAvatarError: UploadProps['onError'] = () => {
  loading.value = false
  ElMessage.error('图片上传失败')
}

/** 手动上传文件 */
function submit() {
  if (props.enableUpload && !props.autoUpload && fileList.value.length > 0) {
    const uploadRef = proxy.$refs['uploadRef']
    uploadRef.submit()
    // uploadRef.clearFiles()
  }
}

watch(() => props.modelValue, () => {
  loadAvatar()
})

watch(fileList, (value) => {
  if (value.length > 0) {
    const file = value[value.length - 1].raw
    url.value = URL.createObjectURL(file!)
    emit('update:file', file)
  } else {
    emit('update:file', undefined)
  }
})

defineExpose({
  submit
})

loadAvatar()
</script>

<template>
  <el-upload
    v-if="enableUpload"
    ref="uploadRef"
    accept="image/jpeg,image/png"
    v-model:file-list="fileList"
    :auto-upload="autoUpload"
    :disabled="loading"
    :action="uploadUrl"
    :headers="headers"
    :show-file-list="false"
    :before-upload="beforeAvatarUpload"
    :on-success="handleAvatarSuccess"
    :on-error="handleAvatarError"
  >
    <div class="upload-avatar" :style="{ width: sizeCom + 'px', height: sizeCom + 'px' }">
      <div class="cover" :style="{ 'border-radius': shapeStyle }">
        <el-icon :size="sizeCom / 2.5" color="#fff">
          <i-ep-edit />
        </el-icon>
      </div>
      <el-avatar v-loading="loading" :shape="shape" :size="sizeCom">
        <el-image v-if="modelValue || (!autoUpload && url)" :style="{ width: sizeCom + 'px', height: sizeCom + 'px' }" :src="url" fit="cover">
          <template #error>
            <div class="image-slot">
              <el-icon :size="sizeCom / 2" color="var(--el-text-color-secondary)">
                <svg-icon icon-name="icon-picture-error"/>
              </el-icon>
            </div>
          </template>
        </el-image>
        <el-icon v-else :size="sizeCom / 2.5" color="#fff">
          <i-ep-user-filled />
        </el-icon>
      </el-avatar>
    </div>
  </el-upload>
  <el-avatar v-else v-loading="loading" :shape="shape" :size="sizeCom" :class="{ 'mini-loading': sizeCom < 42 }" :style="{ '--ms-avatar-size': sizeCom + 'px' }">
    <el-image v-if="url" :style="{ width: sizeCom + 'px', height: sizeCom + 'px' }" :src="url" fit="cover">
      <template #error>
        <div class="image-slot">
          <el-icon :size="sizeCom / 2" color="var(--el-text-color-secondary)">
            <svg-icon icon-name="icon-picture-error"/>
          </el-icon>
        </div>
      </template>
    </el-image>
    <el-icon v-else :size="sizeCom / 2.5" color="#fff">
      <i-ep-user-filled />
    </el-icon>
  </el-avatar>
</template>

<style lang="scss" scoped>
.upload-avatar {
  position: relative;
  .cover {
    display: none;
    position: absolute;
    width: 100%;
    height: 100%;
    z-index: 101;
    align-items: center;
    justify-content: center;
    background: rgba(0, 0, 0, 0.3);
  }
  &:hover .cover {
    display: flex;
  }
}
:deep(.image-slot) {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
}
.mini-loading {
  :deep(.el-loading-spinner) {
    margin: 0;
    top: 0;
  }
  :deep(.circular) {
    height: var(--ms-avatar-size);
    width: var(--ms-avatar-size);
  }
}
</style>
