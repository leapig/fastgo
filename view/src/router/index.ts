import { createRouter, createWebHistory } from 'vue-router'
import OperatingLayout from '@/layout/operating/Index.vue'
import ManageLayout from '@/layout/manage/Index.vue'

// 公共路由
export const constantRoutes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue')
  },
  {
    path: '',
    redirect: '/login',
    children: [
      {
        path: '/index',
        name: 'Index',
        component: () => import('@/views/Index.vue')
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    component: () => import('@/views/error/404.vue')
  }
]

// 独立页面路由
export const independenceRoutes = [
  {
    path: '/manage',
    component: markRaw(ManageLayout),
    children: [
      // {
      //   path: 'project-operate/manage',
      //   name: 'ManageProjectOperateManageTabs',
      //   component: markRaw(() => import('@/views/manage/project/operate/ManageTabs.vue')),
      //   permission: 'manage/project/operate/ManageTabs',
      //   children: [
      //     {
      //       path: 'member/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateMember',
      //       component: () => import('@/views/manage/project/operate/member/Index.vue'),
      //       permission: 'manage/project/operate/member/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'post/:project_pk(\\d+)',
      //       name: 'ManageProjectOperatePost',
      //       component: () => import('@/views/manage/project/operate/post/Index.vue'),
      //       permission: 'manage/project/operate/post/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'work-schedule/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateWorkSchedule',
      //       component: () => import('@/views/manage/project/operate/workSchedule/Index.vue'),
      //       permission: 'manage/project/operate/workSchedule/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'classes/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateClasses',
      //       component: () => import('@/views/manage/project/operate/classes/Index.vue'),
      //       permission: 'manage/project/operate/classes/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'patrol/:project_pk(\\d+)',
      //       name: 'ManageProjectOperatePatrol',
      //       component: () => import('@/views/manage/project/operate/patrol/Index.vue'),
      //       permission: 'manage/project/operate/patrol/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'inspection/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateInspection',
      //       component: () => import('@/views/manage/project/operate/inspection/Index.vue'),
      //       permission: 'manage/project/operate/inspection/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'supervision/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateSupervision',
      //       component: () => import('@/views/manage/project/operate/supervision/Index.vue'),
      //       permission: 'manage/project/operate/supervision/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'attendance/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateAttendance',
      //       component: () => import('@/views/manage/project/operate/attendance/Index.vue'),
      //       permission: 'manage/project/operate/attendance/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'apply-job/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateApplyJob',
      //       component: () => import('@/views/manage/project/operate/applyJob/Index.vue'),
      //       permission: 'manage/project/operate/applyJob/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'leave-job/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateLeaveJob',
      //       component: () => import('@/views/manage/project/operate/leaveJob/Index.vue'),
      //       permission: 'manage/project/operate/leaveJob/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'group/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateGroup',
      //       component: () => import('@/views/manage/project/operate/group/Index.vue'),
      //       permission: 'manage/project/operate/group/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'talk/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateTalk',
      //       component: () => import('@/views/manage/project/operate/talk/Index.vue'),
      //       permission: 'manage/project/operate/talk/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'property/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateProperty',
      //       component: () => import('@/views/manage/project/operate/property/Index.vue'),
      //       permission: 'manage/project/operate/property/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'device/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateDevice',
      //       component: () => import('@/views/manage/project/operate/device/Index.vue'),
      //       permission: 'manage/project/operate/device/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       path: 'authorisation/:project_pk(\\d+)',
      //       name: 'ManageProjectOperateAuthorisation',
      //       component: () => import('@/views/manage/project/operate/authorisation/Index.vue'),
      //       permission: 'manage/project/operate/authorisation/Index',
      //       meta: { fade: false }
      //     }
      //   ]
      // },

      // {
      //   // 客户维护-客户列表
      //   path: 'customer-hold/business/manage',
      //   name: 'ManageCustomerHoldBusinessManageTabs',
      //   component: markRaw(() => import('@/views/manage/customerHold/business/ManageTabs.vue')),
      //   permission: 'manage/customerHold/business/ManageTabs',
      //   children: [
      //     {
      //       // 签约项目
      //       path: 'project/:enterprise_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessProject',
      //       component: () => import('@/views/manage/customerHold/business/project/Index.vue'),
      //       permission: 'manage/customerHold/business/project/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 劳务派遣
      //       path: 'dispatch/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessDispatch',
      //       component: () => import('@/views/manage/customerHold/business/dispatch/Index.vue'),
      //       permission: 'manage/customerHold/business/dispatch/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 物品资产
      //       path: 'asset/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessAsset',
      //       component: () => import('@/views/manage/customerHold/business/asset/Index.vue'),
      //       permission: 'manage/customerHold/business/asset/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 客户评价
      //       path: 'appraise/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessAppraise',
      //       component: () => import('@/views/manage/customerHold/business/appraise/Index.vue'),
      //       permission: 'manage/customerHold/business/appraise/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 合同管理
      //       path: 'contract/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessContract',
      //       component: () => import('@/views/manage/customerHold/business/contract/Index.vue'),
      //       permission: 'manage/customerHold/business/contract/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 账单明细
      //       path: 'billing/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessBilling',
      //       component: () => import('@/views/manage/customerHold/business/billing/Index.vue'),
      //       permission: 'manage/customerHold/business/billing/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 开票记录
      //       path: 'invoice/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessInvoice',
      //       component: () => import('@/views/manage/customerHold/business/invoice/Index.vue'),
      //       permission: 'manage/customerHold/business/invoice/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 回款状态
      //       path: 'payment/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessPayment',
      //       component: () => import('@/views/manage/customerHold/business/payment/Index.vue'),
      //       permission: 'manage/customerHold/business/payment/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 经营简报
      //       path: 'briefing/:project_pk(\\d+)',
      //       name: 'ManageCustomerHoldBusinessBriefing',
      //       component: () => import('@/views/manage/customerHold/business/briefing/Index.vue'),
      //       permission: 'manage/customerHold/business/briefing/Index',
      //       meta: { fade: false }
      //     }
      //   ]
      // },

      // {
      //   // 客户维护-签约项目
      //   path: 'customer-hold/project/manage',
      //   name: 'ManageCustomerHoldProjectManageTabs',
      //   component: markRaw(() => import('@/views/manage/customerHold/project/ManageTabs.vue')),
      //   permission: 'manage/customerHold/project/ManageTabs',
      //   children: [
      //     {
      //       // 劳务派遣
      //       path: 'dispatch/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldDispatch',
      //       component: () => import('@/views/manage/customerHold/project/dispatch/Index.vue'),
      //       permission: 'manage/customerHold/project/dispatch/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 合同管理
      //       path: 'contract/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldContract',
      //       component: () => import('@/views/manage/customerHold/project/contract/Index.vue'),
      //       permission: 'manage/customerHold/project/contract/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 账单明细
      //       path: 'bill/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldBill',
      //       component: () => import('@/views/manage/customerHold/project/bill/Index.vue'),
      //       permission: 'manage/customerHold/project/bill/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 开票记录
      //       path: 'invoice/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldInvoice',
      //       component: () => import('@/views/manage/customerHold/project/invoice/Index.vue'),
      //       permission: 'manage/customerHold/project/invoice/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 回款状态
      //       path: 'payment/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldPayment',
      //       component: () => import('@/views/manage/customerHold/project/payment/Index.vue'),
      //       permission: 'manage/customerHold/business/project/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 月度报告
      //       path: 'reports/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldreports',
      //       component: () => import('@/views/manage/customerHold/project/reports/Index.vue'),
      //       permission: 'manage/customerHold/project/reports/Index',
      //       meta: { fade: false }
      //     },
      //     {
      //       // 经营简报
      //       path: 'briefing/:project_pk(\\d+)',
      //       name: 'ManageProjectCustomerHoldBriefing',
      //       component: () => import('@/views/manage/customerHold/project/briefing/Index.vue'),
      //       permission: 'manage/customerHold/project/briefing/Index',
      //       meta: { fade: false }
      //     }
      //   ]
      // }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes: constantRoutes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

export default router
