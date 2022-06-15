// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
//引入路由
import router from "./router"

// Vue.use(Router)
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  //注入框架
  router,
  components: { App },
  template: '<App/>'
})
