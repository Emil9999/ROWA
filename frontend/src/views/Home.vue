<template>
  <div class="home">
    
    <button-home firstElement="Farm" secondElement="Home"/>

    <plant-card PlantHarvest="Your are planting"/>
  </div>
</template>

<script>
// @ is an alias to /src
//import FarmRepresentation from '@/components/farm/FarmRepresentation.vue'
import ButtonHome from '../components/ui/ButtonHome.vue'
import PlantCard from '../components/ui/PlantCard.vue'

export default {
  name: 'Home',
  components: {
    
    ButtonHome,
    PlantCard
  }
}
</script>

    ButtonHome