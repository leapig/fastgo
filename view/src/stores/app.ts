import Cookies from 'js-cookie'

const operatingSidebarStatusKey = 'OperatingSidebarStatus'
const manageSidebarStatusKey = 'ManageSidebarStatus'

const useAppStore = defineStore('app', {
  state: () => ({
    usePlatform: '',
    operating: {
      sidebar: {
        opened: Cookies.get(operatingSidebarStatusKey) ? Cookies.get(operatingSidebarStatusKey) === '1' : true,
        withoutAnimation: false,
        hide: false
      },
      useTab: ''
    },
    manage: {
      sidebar: {
        opened: Cookies.get(manageSidebarStatusKey) ? Cookies.get(manageSidebarStatusKey) === '1' : true,
        withoutAnimation: false,
        hide: false
      }
    }
  }),
  actions: {
    /**
     * 设置当前使用的平台
     *
     * @param {string} usePlatform 平台key：operating运营平台，manage管理平台
     */
    setUsePlatform(usePlatform: string) {
      this.usePlatform = usePlatform
    },
    /**
     * 清除当前使用的平台
     */
    clearUsePlatform() {
      this.usePlatform = ''
    },
    /**
     * 展开/收起侧边栏
     *
     * @param {boolean} withoutAnimation 是否不适用动画
     * @returns {boolean}
     */
    toggleSideBar(withoutAnimation: boolean) {
      if (this.usePlatform === 'manage') {
        if (this.manage.sidebar.hide) {
          return false
        }
        this.manage.sidebar.opened = !this.manage.sidebar.opened
        this.manage.sidebar.withoutAnimation = withoutAnimation
        if (this.manage.sidebar.opened) {
          Cookies.set(manageSidebarStatusKey, '1')
        } else {
          Cookies.set(manageSidebarStatusKey, '0')
        }
      } else if (this.usePlatform === 'operating') {
        if (this.operating.sidebar.hide) {
          return false
        }
        this.operating.sidebar.opened = !this.operating.sidebar.opened
        this.operating.sidebar.withoutAnimation = withoutAnimation
        if (this.operating.sidebar.opened) {
          Cookies.set(operatingSidebarStatusKey, '1')
        } else {
          Cookies.set(operatingSidebarStatusKey, '0')
        }
      }
      return false
    },
    /**
     * 收起侧边栏
     *
     * @param {any} param
     */
    closeSideBar(param: any) {
      if (this.usePlatform === 'manage') {
        Cookies.set(manageSidebarStatusKey, '1')
        this.manage.sidebar.opened = true
        this.manage.sidebar.withoutAnimation = param.withoutAnimation
      } else if (this.usePlatform === 'operating') {
        Cookies.set(operatingSidebarStatusKey, '0')
        this.operating.sidebar.opened = false
        this.operating.sidebar.withoutAnimation = param.withoutAnimation
      }
    },
    /**
     * 显示/隐藏侧边栏
     *
     * @param {boolean} status true显示，false隐藏
     */
    toggleSideBarHide(status: boolean) {
      if (this.usePlatform === 'manage') {
        this.manage.sidebar.hide = status
      } else if (this.usePlatform === 'operating') {
        this.operating.sidebar.hide = status
      }
    },
    /**
     * 设置运营平台当前使用的tab
     *
     * @param {string} key tab的key
     */
    setOperatingUseTab(key: string) {
      this.operating.useTab = key
    }
  }
})

export default useAppStore
