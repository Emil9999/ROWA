import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/Home'
import Plant from '@/components/Plant'
import Harvest from '@/components/Harvest'
import PlantingInstructions from '@/components/PlantingInstructions'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HelloWorld
    },
    {
      path: '/plant',
      name: 'Plant',
      component: Plant
    },
    {
      path: '/harvest',
      name: 'Harvest',
      component: Harvest
    },
    {
      path: '/plant/howto',
      name: 'Planting Instructions',
      component: PlantingInstructions
    }
  ]
})
