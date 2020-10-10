import Vue from 'vue'
import Router from 'vue-router'

// import HelloWorld from './components/HelloWorld';
// import Hover from './components/Hover.vue'
import Beranda from './components/Beranda.vue'
// import Health from './components/Health.vue'
// import Business from './components/Business.vue'
// import Entertainment from './components/Entertainment.vue'
import Detail from './components/Detail.vue'
//import Footer from './components/Footer.vue'

Vue.use(Router)

export default new Router ({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'beranda',
            component: Beranda
        },
        // {
        //     path: '/health',
        //     name: 'health',
        //     component: Health
        // },
        // {
        //     path: '/business',
        //     name: 'business',
        //     component: Business
        // },
        // {
        //     path: '/entertainment',
        //     name: 'entertainment',
        //     component: Entertainment
        // },
        {
            path: '/detail',
            name: 'detail',
            component: Detail
        },
    ]
})