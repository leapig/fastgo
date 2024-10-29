<script lang="ts" setup>
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import type { UploadProps } from 'element-plus'
import { getUserAccessToken, getTenantAccessToken } from '@/utils/auth'
import { isHttp, isBase64Image } from '@/utils/validate'

const props = defineProps({
  /**
   * 当前显示图片地址/Base64绑定值，不包括手动上传所选但未上传的图片
   * 根据业务逻辑需要，图片上传后该值为接口返回的name
   */
  modelValue: {
    type: String
  },
  /** 初始化显示图片地址/Base64，为空时图片地址使用modelValue绑定值 */
  src: {
    type: String
  },
  /** 当前手动上传所选图片 */
  file: {
    type: File
  },
  /** 图片显示宽度 */
  width: {
    type: Number,
    default: 80
  },
  /** 图片显示高度 */
  height: {
    type: Number,
    default: 80
  },
  /** 图片的容器适应方式 */
  fit: {
    type: String as PropType<EpPropMergeType<StringConstructor, '' | 'fill' | 'cover' | 'none' | 'contain' | 'scale-down', unknown>>,
    default: ''
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
  /** 自动上传目标 device租户设备相关 general租户其他相关 privacy 没有租户的个人使用 system 运营者系统使用 */
  autoUploadTarget: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'device' | 'general' | 'privacy' | 'system', unknown>>,
    default: 'general'
  },
  /** 是否允许点击全屏查看 */
  preview: {
    type: Boolean,
    default: true
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

/** 加载图片 */
function loadImage() {
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
/** 图片上传前的回调 */
const beforeImageUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (rawFile.type !== 'image/jpeg' && rawFile.type !== 'image/png') {
    ElMessage.error('图片只支持 jpeg、png 格式')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('图片不能超过 2MB')
    return false
  }
  loading.value = true
  return true
}
/** 图片上传成功回调 */
const handleImageSuccess: UploadProps['onSuccess'] = (response: any) => {
  loading.value = false
  url.value = response.data.url
  emit('update:modelValue', response.data.name)
}
/** 图片上传失败回调 */
const handleImageError: UploadProps['onError'] = () => {
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
  loadImage()
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

loadImage()
</script>

<template>
  <div class="ms-image" :style="{ width: width + 'px', height: height + 'px' }">
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
      :before-upload="beforeImageUpload"
      :on-success="handleImageSuccess"
      :on-error="handleImageError"
    >
      <div class="upload-image" :style="{ width: width + 'px', height: height + 'px' }">
        <div class="cover" v-if="modelValue && !loading">
          <el-icon size="30" color="#fff">
            <i-ep-edit />
          </el-icon>
        </div>
        <el-image
          v-if="modelValue || (!autoUpload && url)"
          v-loading="loading"
          :style="{ width: width + 'px', height: height + 'px' }"
          :src="url"
          :fit="fit"
        >
          <template #error>
            <div class="image-slot">
              <el-icon size="30" color="var(--el-text-color-secondary)">
                <svg-icon icon-name="icon-picture-error"/>
              </el-icon>
            </div>
          </template>
        </el-image>
        <div v-else v-loading="loading" class="image">
          <el-icon size="30" color="#dcdfe6">
            <i-ep-plus />
          </el-icon>
        </div>
      </div>
    </el-upload>
    <el-image
      v-else
      v-loading="loading"
      :src="url"
      :style="{ width: width + 'px', height: height + 'px' }"
      :fit="fit"
      :preview-src-list="preview ? [url] : []"
      preview-teleported
    >
      <template #error>
        <div class="image-slot">
          <el-icon size="30" color="var(--el-text-color-secondary)">
            <svg-icon icon-name="icon-picture-error"/>
          </el-icon>
        </div>
      </template>
    </el-image>
  </div>
</template>

<style lang="scss" scoped>
.ms-image {
  & > div:not(.el-image) {
    display: flex;
  }
  .upload-image {
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
    .image {
      display: flex;
      width: 100%;
      height: 100%;
      align-items: center;
      justify-content: center;
      border: 1px dashed var(--el-border-color);
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
}
</style>
