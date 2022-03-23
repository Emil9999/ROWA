import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'
import ErrrorPage from '../App.vue'
import Home from '../views/Home.vue'
import FarmingStepper from '../views/FarmingStepper.vue'

const routes: Array<RouteRecordRaw> = [
    {path:'/error', 
    name: "App",
    component: ErrrorPage},

    {path:'/', 
    name: "Home",
    component: Home},

    {path: '/farming/:farmingType',
    props: true,
    component: FarmingStepper},

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router