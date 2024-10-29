import route, { independenceRoutes } from '@/router'
import useSettingsStore from '@/stores/settings'
import {
  getMenusApi,
  getPagesApi,
} from '@/api/permission'
import OperatingLayout from '@/layout/operating/Index.vue'
import ManageLayout from '@/layout/manage/Index.vue'

export type Menu = {
  title?: string
  name?: string
  path: string
  icon?: string
  component?: any
  component_name?: string
  is_cache?: boolean
  redirect?: string
  children?: Array<Menu>
  meta?: MenuMeta
}

export type MenuMeta = {
  title?: string
  icon?: string
  component?: string
}

// 匹配views里面所有的.vue文件
const modules = import.meta.glob('./../views/**/*.vue')

const usePermissionStore = defineStore('permission', {
  state: () => ({
    routes: [] as Array<Menu>,
    indRoutes: [] as Array<Menu>,
    pages: [] as Array<any>,
    permiPages: {} as any,
    operating: {
      tabs: [] as Array<any>,
      menus: {} as any
    },
    manage: {
      menus: [] as Array<Menu>
    }
  }),
  actions: {
    /**
     * 动态加载路由
     * @param {string} usePlatform
     */
    generateRoutes(usePlatform: string) {
      return new Promise((resolve: Function, reject: Function) => {
        if (usePlatform === 'manage') {
          getMenusApi()
            .then((response: any) => {
              // 转换路由
              const manageTopMenus = response.data.filter((menu: any) => menu.menuName && menu.path === '/manage')
              if (manageTopMenus.length) {
                const settingsStore = useSettingsStore()
                const manageTopMenu = manageTopMenus[0]
                // if (settingsStore.showManageHomePage) {
                //   const manageMenus = manageTopMenu.childMenus || []
                //   const homeMenu = {
                //     menuName: '首页',
                //     menuType: 1,
                //     icon: 'HomeFilled',
                //     path: 'home',
                //     component: 'manage/Home',
                //     component_name: 'ManageHome',
                //     is_cache: true
                //   }
                //   manageMenus.unshift(homeMenu)
                //   manageTopMenu.childMenus = manageMenus
                // }
                const routes = convertDataToRoutes(manageTopMenu)
                if (routes) {
                  // 整理菜单数据
                  this.manage.menus = routes.children || []
                  // 获取权限页面，成功的话再加载路由
                  this.getPages(usePlatform)
                    .then((pages: any) => {
                      this.permiPages = convertPermiPages(pages)
                      // 加载独立页面路由
                      this.indRoutes = loadIndependenceRoutes(this.permiPages, independenceRoutes)
                      this.indRoutes.forEach((indRoute: any) => { route.addRoute(indRoute) })
                      // 加载菜单组件
                      loadMenuComponent(routes)
                      this.routes = [routes]
                      resolve(routes)
                    })
                    .catch((error: any) => {
                      console.error(error)
                      reject()
                    })
                } else {
                  reject()
                }
              } else {
                reject()
              }
            })
            .catch((error: any) => {
              console.error(error)
              reject()
            })
        } else if (usePlatform === 'operating') {
          getMenusApi()
            .then((response: any) => {
              // 读取顶部Tabs
              this.operating.tabs = getOperatingTabsFromMenusTree(response.data)
              // 转换路由
              const operatingTopMenus = response.data.filter((menu: any) => menu.menuName && menu.path === '/operating')
              if (operatingTopMenus.length) {
                const operatingTopMenu = operatingTopMenus[0]
                const routes = convertDataToRoutes(operatingTopMenu)
                if (routes) {
                  // 整理菜单数据
                  if (routes.children && routes.children.length) {
                    const operatingMenus = {} as any
                    routes.children.forEach((child: Menu) => {
                      if (child.path) {
                        operatingMenus[child.path] = child.children || []
                      }
                    })
                    this.operating.menus = operatingMenus
                  }
                  // 获取权限页面，成功的话再加载路由
                  this.getPages(usePlatform)
                    .then((pages: any) => {
                      this.permiPages = convertPermiPages(pages)
                      // 加载独立页面路由
                      this.indRoutes = loadIndependenceRoutes(this.permiPages, independenceRoutes)
                      this.indRoutes.forEach((indRoute: any) => { route.addRoute(indRoute) })
                      // 加载菜单组件
                      loadMenuComponent(routes)
                      this.routes = [routes]
                      resolve(routes)
                    })
                    .catch((error: any) => {
                      console.error(error)
                      reject()
                    })
                } else {
                  reject()
                }
              } else {
                reject()
              }
            })
            .catch((error: any) => {
              console.error(error)
              reject()
            })
        } else {
          resolve([])
        }
      })
    },
    /**
     * 权限页面
     * @param {string} usePlatform
     */
    getPages (usePlatform?: string) {
      return new Promise((resolve: Function, reject: Function) => {
        if (this.pages && this.pages.length) {
          resolve(this.pages)
        } else if (usePlatform === 'manage') {
          getPagesApi()
            .then((response: any) => {
              this.pages = response.data
              resolve(response.data)
            })
            .catch((error: any) => {
              console.error(error)
              reject()
            })
        } else if (usePlatform === 'operating') {
          getPagesApi()
            .then((response: any) => {
              this.pages = response.data
              resolve(response.data)
            })
            .catch((error: any) => {
              console.error(error)
              reject()
            })
        } else {
          resolve([])
        }
      })
    },
    /**
     * 是否拥有读权限(根据当前路由判断)
     */
    hasReadPermi () {
      let hasPermi = false
      if (this.permiPages && this.permiPages.read && this.permiPages.read.length) {
        const nowRoute = route.currentRoute.value
        if (nowRoute && nowRoute.meta && nowRoute.meta.component) {
          const readPermis = this.permiPages.read
          const routeComponent = nowRoute.meta.component
          const index = readPermis.findIndex((permis: string) => permis === routeComponent)
          hasPermi = index > -1
        }
      }
      return hasPermi
    },
    /**
     * 是否拥有写权限(根据当前路由判断)
     */
    hasWritePermi () {
      let hasPermi = false
      if (this.permiPages && this.permiPages.write && this.permiPages.write.length) {
        const nowRoute = route.currentRoute.value
        if (nowRoute && nowRoute.meta && nowRoute.meta.component) {
          const writePermis = this.permiPages.write
          const routeComponent = nowRoute.meta.component
          const index = writePermis.findIndex((permis: string) => permis === routeComponent)
          hasPermi = index > -1
        }
      }
      return hasPermi
    },
    /**
     * 是否拥有页面权限
     * @param {string} page 页面的组件路径
     */
    hasPagePermi (page: string) {
      let hasPermi = false
      if (page && this.permiPages && this.permiPages.write && this.permiPages.write.length) {
        const writePermis = this.permiPages.write
        const index = writePermis.findIndex((permis: string) => permis === page)
        hasPermi = index > -1
      }
      return hasPermi
    }
  }
})

/**
 * 从菜单树中获取运营平台的顶部Tabs
 * @param menusTree 菜单树
 * @returns
 */
function getOperatingTabsFromMenusTree(menusTree: Array<any>) {
  const tabs = [] as Array<any>
  const operatingTopMenus = menusTree.filter((menu: any) => menu.menuName && menu.path === '/operating')
  if (operatingTopMenus.length) {
    const settingsStore = useSettingsStore()
    const operatingTopMenu = operatingTopMenus[0]
    const operatingTabMenus = operatingTopMenu.childMenus || []
    operatingTabMenus.forEach((tabMenu: any) => {
      if (settingsStore.showOperatingHomePage) {
        tabMenu.component = 'Redirect:/operating/home'
      }
      if (tabMenu.menuName && tabMenu.path && tabMenu.menuType === 3) {
        const tab = { label: tabMenu.menuName, key: tabMenu.path }
        tabs.push(tab)
      }
    })
    if (settingsStore.showOperatingHomePage) {
      const homeMenu = {
        menuName: '首页',
        menuType: 1,
        icon: 'HomeFilled',
        path: 'home',
        component: 'operating/Home',
        component_name: 'OperatingHome',
        is_cache: true
      }
      operatingTabMenus.unshift(homeMenu)
      operatingTopMenu.childMenus = operatingTabMenus
    }
  }
  return tabs
}

/**
 * 接口中获取的菜单数据转换为页面路由
 * @param data 菜单
 * @returns
 */
function convertDataToRoutes(data: any) {
  if (data.menuName && data.path) {
    const menuMeta = {
      title: data.menuName,
      icon: data.icon,
      component: data.component
    } as MenuMeta
    const menu = {
      title: data.menuName,
      path: data.path,
      icon: data.icon,
      component: '',
      component_name: '',
      children: [],
      meta: menuMeta
    } as Menu
    if (data.menuType === 3) {
      if (data.path === '/operating') {
        // operating顶级目录
        menu.component = 'Layout:operating'
        menu.component_name = 'Operating'
      } else if (data.path === '/manage') {
        // manage顶级目录
        menu.component = 'Layout:manage'
        menu.component_name = 'Manage'
      } else {
        // 普通目录
        menu.component = data.component || ''
      }
      if (data.childMenus && data.childMenus.length) {
        // 子菜单
        const children = [] as Array<Menu>
        for (const ind in data.childMenus) {
          const childMenuData = data.childMenus[ind]
          const childMenu = convertDataToRoutes(childMenuData)
          if (childMenu) {
            children.push(childMenu)
          }
        }
        menu.children = children
      }
    } else if (data.menuType === 1) {
      // 页面
      menu.component = data.component || ''
      menu.component_name = data.component_name || ''
      menu.is_cache = data.is_cache
    }
    return menu
  }
  return null
}

/**
 * 转换页面权限
 */
function convertPermiPages(pages: Array<any>) {
  if (pages && pages.length) {
    const read = [] as Array<string>
    const write = [] as Array<string>
    pages.forEach((page: any) => {
      if (page.operation_type === 1) {
        read.push(page.component)
      } else if (page.operation_type === 2) {
        read.push(page.component)
        write.push(page.component)
      }
    })
    return {
      read,
      write
    }
  }
}

/**
 * 加载独立页面路由
 */
function loadIndependenceRoutes(permis: any, routes: Array<any>) {
  const newRoutes = [] as Array<Menu>
  if (permis.write && permis.write.length) {
    routes.forEach((item: any) => {
      if (item.children && item.children.length) {
        item.children = loadIndependenceRoutes(permis, item.children)
      }
      if (item.permission) {
        item.meta = {
          ...item.meta,
          title: item.title,
          icon: item.icon,
          component: item.permission
        }
        const index = permis.write.findIndex((permis: string) => permis === item.permission)
        if (index > -1) {
          newRoutes.push(item)
        }
      } else {
        newRoutes.push(item)
      }
    })
  }
  return newRoutes
}

/**
 * 加载菜单组件
 * @param menu
 */
function loadMenuComponent(menu: Menu) {
  if (menu.component) {
    if (menu.component === 'Layout:operating') {
      menu.component = OperatingLayout
    } else if (menu.component === 'Layout:manage') {
      menu.component = ManageLayout
    } else if (menu.component.startsWith('Redirect:')) {
      const redPath = menu.component.split(':')[1]
      menu.redirect = redPath
      delete menu['component']
    } else {
      menu.component = loadView(menu.component)
    }
  }
  if (menu.children && menu.children.length > 0) {
    menu.children.forEach((menu: Menu) => {
      loadMenuComponent(menu)
    })
  } else {
    delete menu['children']
  }
  if (menu.component_name) {
    menu.name = menu.component_name
    delete menu['component_name']
  }
}

/**
 * 加载组件文件
 * @param view
 * @returns
 */
function loadView(view: string) {
  let res
  for (const path in modules) {
    const dir = path.split('views/')[1].split('.vue')[0]
    if (dir === view) {
      res = () => modules[path]()
    }
  }
  return res
}

export default usePermissionStore
