<template>
  <div class="home">
    <v-container>
      <v-row>
        <v-col>

        </v-col>
        <v-col>

        </v-col>
        <v-col>
         
        </v-col>
      </v-row>
      <v-row>
      </v-row>
      <v-row >
        <v-col>
          <v-row align="center" justify="center">
            <h4> What do you want to do? </h4>
          </v-row>
          <v-row align="center" justify="center">
            <big-button Text="Harvest"/>
          </v-row>
          <v-row align="center" justify="center">
            <big-button Text="Plant"/>
          </v-row>
        </v-col>
      </v-row>
    </v-container>

    <button-home firstElement="Farm" secondElement="Home"/>

    <plant-card PlantHarvest="Your are planting" PlantName="Basil" Pic="lettuce"/>

    <big-button Text="Primary" />
  </div>
</template>

<script>
import BigButton from '../components/ui/BigButton.vue'
// @ is an alias to /src
//import FarmRepresentation from '@/components/farm/FarmRepresentation.vue'
import ButtonHome from '../components/ui/ButtonHome.vue'
import PlantCard from '../components/ui/PlantCard.vue'

export default {
  name: 'Home',
  components: {
    
    ButtonHome,
    PlantCard,
    BigButton
  }
}
</script>
,
    BigButton
    ButtonHome