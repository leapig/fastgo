import defaultSettings from '@/settings'

const {
  title,
  copyright,
  navbarDefaultColor,
  showOperatingHomePage,
  operatingIndexPage,
  silentRequestUrls
} = defaultSettings

const useSettingsStore = defineStore('settings', {
  state: () => ({
    title: title,
    copyright: copyright,
    navbarColor: navbarDefaultColor,
    showOperatingHomePage: showOperatingHomePage,
    operatingIndexPage: operatingIndexPage,
    silentRequestUrls: silentRequestUrls
  }),
  actions: {
    /**
     * 是否静默请求的url
     * @param {string} url 请求地址
     * @returns 
     */
    isSilentRequestUrl(url?: string) {
      if (url) {
        return this.silentRequestUrls.findIndex((one: string) => url.startsWith(one)) > -1
      } else {
        return false
      }
    }
  }
})

export default useSettingsStore
