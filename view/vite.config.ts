import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd())
  const { VITE_APP_ENV } = env
  return {
    base: VITE_APP_ENV === 'production' ? '/' : '/',
    plugins: [
      vue(),
      AutoImport({
        imports: ['vue', 'vue-router', 'pinia'],
        dts: fileURLToPath(new URL('./types/auto-imports.d.ts', import.meta.url)),
        resolvers: [
          ElementPlusResolver(),
          IconsResolver({
            prefix: 'Icon'
          })
        ]
      }),
      Components({
        dirs: ['src/views', 'src/layout', 'src/components'],
        dts: fileURLToPath(new URL('./types/components.d.ts', import.meta.url)),
        resolvers: [
          ElementPlusResolver(),
          IconsResolver({
            enabledCollections: ['ep']
          })
        ],
        directoryAsNamespace: true
      }),
      Icons({
        autoInstall: true
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
      extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.vue', '.min.js']
    },
    build: {
      chunkSizeWarningLimit: 1600
    },
    // vite 相关配置
    server: {
      port: 9000,
      host: true,
      open: true,
      proxy: {
        '/open-apis': {
          target: 'https://le.d-cos.com', // 测试
          changeOrigin: true,
          rewrite: (p) => p.replace(/^\/open-apis/, '/open-apis')
        }
      }
    }
  }
})
