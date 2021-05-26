import MainView from '@/layouts/MainView'
// import RouteView from '@/layouts/RouteView'

/**
 * configurable parameters under the routing '/'
 * hidden: true                        if `hidden:true` will not show in the sidebar, default is false
 * name:'router-name'                  must set and globally unique
 * meta : {
    auths: ['super_admin','admin']     set multiple roles, default is null
    title: 'title'                     the name show in submenu and breadcrumb, must set
    icon: 'icon-name'                  the icon show in the sidebar, must set
    href: 'url'                        redirect url
  }
**/

export const constantRoutes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '@/pages/login'),
    meta: { title: 'Sign in' },
    hidden: true
  },
  {
    path: '/404',
    component: () => import(/* webpackChunkName: "error-page-404" */ '@/pages/error-page/404'),
    meta: { title: '404' },
    hidden: true
  },
  {
    path: '/401',
    component: () => import(/* webpackChunkName: "error-page-401" */ '@/pages/error-page/401'),
    meta: { title: 'No permission' },
    hidden: true
  },
  {
    path: '/500',
    component: () => import(/* webpackChunkName: "error-page-500" */ '@/pages/error-page/500'),
    meta: { title: '500' },
    hidden: true
  },
  {
    path: '/redirect',
    component: MainView,
    hidden: true,
    children: [
      {
        path: '/redirect/:path*',
        component: () => import(/* webpackChunkName: "layouts-redirect" */ '@/layouts/Redirect')
      }
    ]
  }
]

export const syncRoutes = [
  {
    path: '/',
    name: 'Home',
    meta: { title: 'Home', icon: 'md-home' },
    component: MainView,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import(/* webpackChunkName: "dashboard" */ '@/pages/dashboard'),
        meta: { title: 'Dashboard', icon: 'md-speedometer' }
      }
    ]
  },
  {
    path: '/paf',
    name: 'payments and finance',
    component: MainView,
    meta: { title: 'Payments and finance', icon: 'ios-grid' },
    children: [
      {
        path: 'payments',
        name: 'Payments',
        component: () =>
          import(/* webpackChunkName: "count-to" */ '@/pages/paf/payments'),
        meta: { title: 'Payments' }
      },
      {
        path: 'finance',
        name: 'Finance',
        component: () =>
          import(/* webpackChunkName: "count-to" */ '@/pages/paf/finance'),
        meta: { title: 'Finance' }
      },
    ]
  },
  {
    path: '/proxy-management',
    name: 'ProxyManagement',
    meta: { title: 'Proxy management', icon: 'md-settings' },
    component: MainView,
    redirect: '/dashboard',
    children: [
      {
        path: 'active-proxy',
        name: 'ActiveProxy',
        component: () =>
          import(
            /* webpackChunkName: "user-management" */ '@/pages/proxy-manager/active-proxy'
          ),
        meta: { title: 'Active proxy', icon: 'md-person' }
      },
      {
        path: 'dns-rules',
        name: 'DNS Rules',
        component: () =>
          import(
            /* webpackChunkName: "role-management" */ '@/pages/proxy-manager/dns-rule'
          ),
        meta: { title: 'DNS rules', icon: 'md-people' }
      },
      {
        path: 'proxy-rules',
        name: 'ProxyRules',
        component: () =>
          import(
            /* webpackChunkName: "role-management" */ '@/pages/proxy-manager/proxy-rules'
            ),
        meta: { title: 'Proxy rules', icon: 'md-people' }
      },
    ]
  },
  {
    path: '/authentication-method',
    name: 'authentication-method',
    component: MainView,
    meta: { title: 'Authentication method', icon: 'ios-key' },
    children: [
      {
        path: 'basic-authentication',
        name: 'Basic authentication',
        component: () =>
          import(/* webpackChunkName: "count-to" */ '@/pages/authenticationMethod/basicAuthentication'),
        meta: { title: 'Basic authentication' }
      },
      {
        path: 'white-list',
        name: 'White list',
        component: () =>
          import(/* webpackChunkName: "count-to" */ '@/pages/authenticationMethod/whiteList'),
        meta: { title: 'White list' }
      },
    ]
  },
  {
    path: '/stats',
    name: 'Statistics',
    component: MainView,
    meta: { title: 'Statistics', icon: 'ios-grid' },
    redirect: '/stats/stats',
    children: [
      {
        path: 'stats',
        name: 'Statistics',
        component: () =>
          import(/* webpackChunkName: "count-to" */ '@/pages/statistics'),
        meta: { title: 'Stats' }
      },
    ]
  },
]

export const asyncRoutes = [
  ...syncRoutes,
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]
