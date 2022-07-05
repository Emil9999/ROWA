import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'
import ErrrorPage from '../App.vue'
import Home from '../views/Home.vue'
import FarmingStepper from '../views/FarmingStepper.vue'
import Admin from '../views/AdminHome.vue'
import RealityCheck from '../views/RealityCheck.vue'

const routes: Array<RouteRecordRaw> = [
    {path:'/error', 
    name: "App",
    component: ErrrorPage},

    {path:'/', 
    name: "Home",
    component: Home},

    {path: '/admin',
    name: "Admin",
    component: Admin},

    {path: '/farming/:farmingType',
    props: true,
    component: FarmingStepper},

    {path: '/farming/direct',
    name: 'directFarming',
    props: { selectedPlant: {planttype: String, planter: String, modulenumber: Number, position: Number}},
    component: FarmingStepper},

    
    {path: '/admin/realitycheck',
    name: "Reality Check",
    component: RealityCheck},

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router