import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'
import ErrrorPage from '../App.vue'
import Home from '../views/Home.vue'
import Planting from '../views/Planting.vue'

const routes: Array<RouteRecordRaw> = [
    {path:'/error', 
    name: "App",
    component: ErrrorPage},

    {path:'/', 
    name: "Home",
    component: Home},

    {path: '/planting',
    name: 'Planting',
    component: Planting},
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router