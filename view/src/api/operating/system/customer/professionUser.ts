import tenantRequest from '@/utils/tenantRequest'

/** 获取从业人员列表 */
export function getProfessionUserList(params: any) {
  return tenantRequest({
    url: '/personnel/user/profession/platform',
    method: 'get',
    params
  })
}
