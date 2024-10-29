import usePermissionStore from '@/stores/permission'

export default {
  mounted(el: HTMLElement) {
    if (!usePermissionStore().hasWritePermi()) {
      el.parentNode && el.parentNode.removeChild(el)
    }
  }
}
