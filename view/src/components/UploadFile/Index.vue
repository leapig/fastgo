<script lang="ts" setup>
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import { getUserAccessToken, getTenantAccessToken } from '@/utils/auth'
import type { UploadProps } from 'element-plus'

const props = defineProps({
  /**
   * 根据业务逻辑需要，图片上传后该值为接口返回的name
   */
  modelValue: {
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
  /** 自动上传目标 device租户设备相关 general租户其他相关 privacy 没有租户的个人使用 system 运营者系统使用 */
  autoUploadTarget: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'device' | 'general' | 'privacy' | 'system', unknown>>,
    default: 'general'
  },
  /** 是否允许点击全屏查看 */
  preview: {
    type: Boolean,
    default: true
  },
  /**上传文件限制数量 */
  limit: {
    type: Number,
    default: 4
  },
  /** 文件限制类型 */
  accept: {
    type: String,
    default: '.pdf,.png,.jpg'
  },
  /** 展示类型 是否展示展示 是否只是上传 编辑两个都有 */
  type: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'onlyUpload' | 'onlyPreview' | 'edit', unknown>>,
    default: 'onlyUpload'
  },
  /** 返回类型 */
  comeBackType: {
    type: String as PropType<EpPropMergeType<StringConstructor, 'array' | 'string', unknown>>,
    default: 'array'
  },
  /**已经上传得文件列表 格式 [{file_name:'',file_url:''}]*/
  alreadyUploadList: {
    type: Array,
    default: (): any => []
  }
})

const emit = defineEmits(['update:modelValue', 'update:file'])
const fileList = ref<any>([])
const fileAlready = ref<any>([])

/* 监听文件数组改变 */
watch(
  [() => fileList, () => fileAlready],
  () => {
    handleComeBcak()
  },
  { deep: true }
)

watch(
  () => props.alreadyUploadList,
  (newValue) => {
    if (props.type === 'onlyPreview') {
      fileAlready.value = JSON.parse(JSON.stringify(newValue))
    }
  },
  { deep: true }
)

/* 计算上传地址 */
const uploadUrl = computed(() => {
  const baseUrl = import.meta.env.VITE_APP_BASE_API + '/core/tool/file'
  return `${baseUrl}/${props.autoUploadTarget}`
})
/* 计算请求头 */
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
/* 计算编辑模式下还能上传多少个文件 */
const uplpoadLimit: any = computed<number>(() => {
  return props.limit - fileAlready.value.length
})

/* 处理文件删除 */
const handleRemove: UploadProps['onRemove'] = (file: any, uploadFiles) => {
  let name = file.response.data.name
  let list = fileList.value
  fileList.value = list.filter((item: any) => item != name)
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
    return true
  }
}
// 处理上传成功
const handleSuccess = (file: any, uploadFiles: any) => {}

/* 回传组件数据 */
const handleComeBcak = () => {
  let list = fileList.value.map((item: any) => {
    return item.response.data.name
  })
  let aradyList = fileAlready.value.map((item: any) => {
    return item.file_name
  })
  list = aradyList.concat(list)
  if (props.comeBackType === 'string') {
    list = list.join(',')
  }
  emit('update:modelValue', list)
}

/* 打开非图片文件 */
function openPdf(item: any) {
  window.open(item.file_url, '_blank')
}
/* 处理已经上传的文件删除name */
function deleteFileName(index: any) {
  fileAlready.value.splice(index, 1)
}

onMounted(() => {
  console.log('加载')
  console.log(props.alreadyUploadList)
  fileAlready.value = JSON.parse(JSON.stringify(props.alreadyUploadList))
})
</script>

<template>
  <div style="width: 100%">
    <div v-if="type === 'onlyUpload' || type === 'edit'">
      <el-upload
        v-if="uplpoadLimit != 0"
        v-model:file-list="fileList"
        class="upload-demo"
        :action="uploadUrl"
        multiple
        :headers="headers"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :before-remove="beforeRemove"
        :before-upload="handleBeforeUpload"
        :on-success="handleSuccess"
        :limit="uplpoadLimit"
        :on-exceed="handleExceed"
        :accept="accept"
      >
        <el-button type="primary">点击上传</el-button>
        <template #tip>
          <div class="el-upload__tip">只能上传{{ accept }}文件，且最多还可以上传{{ uplpoadLimit }}个文件</div>
        </template>
      </el-upload>
    </div>

    <div v-if="type === 'onlyPreview' || type === 'edit'">
      <el-divider content-position="left" v-if="type === 'edit'">已上传文件：</el-divider>
      <div v-for="(item, index) in fileAlready" :key="item.file_name" style="margin-bottom: 20px">
        <div v-if="item.file_name.includes('.pdf')" style="cursor: pointer">
          <span @click="openPdf(item)"> {{ item.file_name }}</span>
          <el-button type="danger" style="margin-left: 10px" round size="small" @click="deleteFileName(index)" v-if="type === 'edit'">删除</el-button>
        </div>
        <div v-else-if="item.file_name.includes('.mp4')" style="cursor: pointer; display: flex; align-items: center">
          <video style="height: 120px" controls>
            <source :src="item.file_url" type="video/mp4" />
          </video>
          <el-button type="danger" style="margin-left: 10px" round size="small" @click="deleteFileName(index)" v-if="type === 'edit'">删除</el-button>
        </div>
        <div v-else style="display: flex; align-items: center">
          <ms-image v-model="item.file_url" :width="80" :height="80" fit="cover"></ms-image>
          <el-button type="danger" style="margin-left: 10px" round size="small" @click="deleteFileName(index)" v-if="type === 'edit'">删除</el-button>
        </div>
      </div>
      <div style="text-align: center" v-if="alreadyUploadList?.length === 0">暂无已上传文件</div>
    </div>
  </div>
</template>

<style lang="scss" scoped></style>
