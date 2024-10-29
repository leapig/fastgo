import usePermissionStore from '@/stores/permission'

export default {
  mounted(el: HTMLElement) {
    if (!usePermissionStore().hasReadPermi()) {
      el.parentNode && el.parentNode.removeChild(el)
    }
  }
}
