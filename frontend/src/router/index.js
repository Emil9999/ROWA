import Vue from 'vue'
import VueRouter from 'vue-router'
//import { component } from 'vue/types/umd'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/',
    name: 'Farm',
    component: () => import('@/views/Farm')
  },
  {
    path: '/',
    name: 'Stats',
    component: () => import('@/views/Stats')
  },
  {
    path: '/planting',
    name: 'Planting',
    component: () => import('@/views/Planting')
  }
  
]

const router = new VueRouter({
  routes
})

export default router
