import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


let router = new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: resove => require(['@/view/login/index'], resove)
    }]
})


export default router;
