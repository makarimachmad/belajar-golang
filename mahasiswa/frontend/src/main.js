
import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import VueRouter from 'vue-router'

Vue.use(VueRouter)

import Detail from './components/Detail.vue'
import Biodata from './components/Biodata.vue'
import Posts from './components/Posts.vue'

const router = new VueRouter({
  routes :[{
      path: '/',
      name: 'biodata',
      component: Biodata
    },
    {
      path: '/detail/:id',
      name: 'detail',
      component: Detail
    },
    {
      path: '/zulfiqor',
      name: 'zulfiqor',
      component: Posts
    }
  ],
  mode: 'history'
})

new Vue({
  router: router,
  render: h => h(App),
}).$mount('#app')
