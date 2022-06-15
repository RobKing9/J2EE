//引入Vue框架
import Vue from "vue";
//引入vue-router
import Router from 'vue-router'

import HelloWorld from "../components/HelloWorld";
import Home from "../components/Home";
import Index from "../components";
import Login from "../components/Login";
import Register from "../components/Register";

//全局使用router
Vue.use(Router)

//定义路由配置
export default new Router ({
  mode:'history',
  routes:[
    {
      path:'/hello',
      name:'HelloWorld',
      component:HelloWorld
    },
    {
      path: '/',
      name: "Home",
      component: Home
    },
    {
      path:'/index',
      name:'Index',
      component:Index,
      children:[
        {
          path:'login',
          name:'Login',
          component:Login
        },
        {
          path:'register',
          name:'Register',
          component:Register
        }
      ]
    }
  ]
})

