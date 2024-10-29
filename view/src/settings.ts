export default {
  /**
   * 网页标题
   */
  title: import.meta.env.VITE_APP_TITLE,
  /**
   * 版权信息
   */
  copyright: import.meta.env.VITE_APP_COPYRIGHT,
  /**
   * 导航栏默认颜色
   */
  navbarDefaultColor: '',
  /**
   * 显示运营平台通用首页
   */
  showOperatingHomePage: false,
  /**
   * 运营平台默认首页
   */
  operatingIndexPage: '/operating/system/customer/user',
  /**
   * 静默请求的url(不出现进度条，不弹框提示)
   */
  silentRequestUrls: [
    '/core/jwt/token/user_token',
    '/core/tool/file'
  ]
}
