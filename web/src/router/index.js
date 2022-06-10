import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


let router = new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: resove => require(['@/view/login/index'], resove)
    },
    {
      path: '/index',
      name: 'Main',
      component: resove => require(['@/view/main/index'], resove),
      children:[
        {
          path: '/index/manage/role',
          name: 'Role',
          component: resove => require(['@/view/role/index'], resove)
        },
        {
          path: '/index/manage/user',
          name: 'User',
          component: resove => require(['@/view/user/index'], resove)
        }
      ]
    }
  ]
})


export default router;
