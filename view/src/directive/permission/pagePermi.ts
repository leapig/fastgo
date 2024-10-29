import type { DirectiveBinding } from 'vue'
import usePermissionStore from '@/stores/permission'

export default {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const { value } = binding
    if (!usePermissionStore().hasPagePermi(value)) {
      el.parentNode && el.parentNode.removeChild(el)
    }
  }
}
