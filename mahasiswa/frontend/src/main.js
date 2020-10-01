import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import VueRouter from 'vue-router'

Vue.use(VueRouter)

import Detail from './components/Detail.vue'
import User from './components/User.vue'
import Post from './components/Posts.vue'

const router = new VueRouter({
  routes :[{
      path: '/',
      name: 'biodata',
      component: User
    },
    {
      path: '/detail/:id',
      name: 'detail',
      component: Detail
    },
    {
      path: '/zulfiqor',
      name: 'zulfiqor',
      component: Post
    }
  ],
  mode: 'history'
})

new Vue({
  router: router,
  render: h => h(App),
}).$mount('#app')
