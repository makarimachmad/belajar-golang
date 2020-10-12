import Vue from 'vue'
import Router from 'vue-router'


import Beranda from './components/Beranda.vue'
import Detail from './components/Detail.vue'
import Detailnya from './components/Detailnya.vue'


Vue.use(Router)

export default new Router ({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'beranda',
            component: Beranda
        },
        {
            path: '/detail/:id',
            name: 'detail',
            component: Detail
        },
        {
            path: '/detailnya',
            name: 'detailnya',
            component: Detailnya
        },
    ]
})