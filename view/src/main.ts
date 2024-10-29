import '@/assets/styles/main.scss'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import * as Icons from '@element-plus/icons-vue'
import App from './App.vue'
import router from './router'
import directive from './directive' // directive
import '@/assets/fonts/iconfont/iconfont.js'

import './permission' // permission control

declare module 'vue' {
  export interface ComponentCustomProperties {
    $icon: any
  }
}

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.config.globalProperties.$icon = Icons

directive(app)

app.mount('#app')
