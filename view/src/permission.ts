import router from './router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import useAppStore from '@/stores/app'
import usePermissionStore from '@/stores/permission'

NProgress.configure({ showSpinner: false })

router.beforeEach((to, from, next) => {
  NProgress.start()

  if (!useAppStore().usePlatform) {
    if (to.path.startsWith('/operating')) {
      useAppStore().setUsePlatform('operating')
      usePermissionStore().generateRoutes(useAppStore().usePlatform).then((routes: any) => {
        router.addRoute(routes)
        next({ ...to, replace: true })
      }).catch(() => {
        next({ ...to, replace: true })
      })
    }  else {
      useAppStore().clearUsePlatform()
      next()
      NProgress.done()
    }
  } else {
    if (to.meta.title) {
      document.title = `${import.meta.env.VITE_APP_TITLE} | ${to.meta.title}`
    } else {
      document.title = `${import.meta.env.VITE_APP_TITLE}`
    }
    next()
    NProgress.done()
  }
})

router.afterEach(() => {
  NProgress.done()
})
