import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/plant',
    name: 'Plant',
    component: () => import('@/views/Plant')
  },
  {
    path: '/harvest',
    name: 'Harvest',
    component: () => import('@/views/Harvest')
  },
  {
    path: '/plant/:pos',
    name: 'PlantingInstructions',
    component: () => import('@/views/PlantingInstructions')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
