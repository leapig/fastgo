export {}

declare global {
  interface Window {
    _AMapSecurityConfig: any
  }
  interface ImportMetaEnv {
    readonly VITE_APP_TITLE: string
    readonly VITE_APP_COPYRIGHT: string
    readonly VITE_APP_ENV: string
    readonly VITE_APP_BASE_API: string
    readonly VITE_APP_AMAP_KEY: string
    readonly VITE_APP_AMAP_SECRET: string
  }
  interface ImportMeta {
    readonly env: ImportMetaEnv
  }
}