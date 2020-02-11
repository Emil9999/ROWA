import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
import test from "../views/test";
import Farming from "../views/Farming.vue"
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: "/module/:id",
    name: "ModuleInfo",
    component: () => import('@/views/Plant')
  },
  {
    path: '/test',
    name: 'test',
    component: test
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
  },
  {
    path: '/farming',
    name: 'Farming',
    component: Farming

  },

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
