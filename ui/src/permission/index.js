import router from '@/router'
import store from '@/store'
import { LoadingBar, Message } from 'view-design'

const whiteList = ['Login'] // no redirect whitelist

// const hasPermission = (rights, rigth) => {
//   return rigth !== undefined ? rights.includes(rigth) : 1
// }
// const canVisitor = visitor => {
//   return visitor !== undefined ? visitor : 0
// }

router.beforeEach(async (to, from, next) => {
  LoadingBar.start()
  next()

  const token = store.getters.token

  if (token) {
    if (to.name === 'Login') {
      // 已登录且要跳转的页面是登录页 跳转到home页
      next({ name: '/' })
      LoadingBar.finish()
    } else {
      // const roles = store.getters.roles
      // const HasRoles = roles && roles.length > 0
      const hasUser = store.getters.user.userName
      if (hasUser) {
        // if (hasPermission(roles, to.meta.auths)) {
        //   next()
        // } else {
        //   next({ path: '/401', replace: true })
        // }
        next()
      } else {
        try {
          // Get user information
          await store.dispatch('user/getUserInfo')
          // No need to dynamically register the route, just comment the following two lines of code
          // generate accessible routes map based
          const accessRoutes = await store.dispatch('routes/getUserMenutree')
          console.log(accessRoutes)
          // dynamically add accessible routes
          accessRoutes.forEach(route => router.addRoute(route))
          // hack method to ensure that addRoutes is complete
          // set the replace: true, so the navigation will not leave a history record
          next({ ...to, replace: true })
        } catch (error) {
          console.log(error) // eslint-disable-line
          await store.dispatch('user/resetToken')
          next({ name: 'Login', query: { redirectPath: to.path } })
          LoadingBar.finish()
        }
      }
    }
  } else {
    if (whiteList.indexOf(to.name) !== -1) {
      // The page you are not logged in and you want to jump to is the login page
      next()
    } else {
      // 未登录且要跳转的页面不是登录页 跳转到登录页
      next({ name: 'Login', query: { redirectPath: to.path } })
      LoadingBar.finish()
    }
  }
})

router.afterEach((to, from) => {
  // const title = to.meta && to.meta.title
  // if (title) document.title = title
  // finish progress bar
  LoadingBar.finish()
})
