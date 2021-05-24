<template>
  <v-container class="farm">
 <v-row class="ma-10">
        <navbar/>
      </v-row>
      <v-row align="center" justify="center">
<v-card elevation="1">
    <h1> Click on a Module <br/> to learn more. </h1>
</v-card>
      </v-row>

        <v-row align="center" justify="center" class="my-15">
          <farm-representation/>
      </v-row>
  </v-container>
</template>

<script>
import FarmRepresentation from '../components/farm/FarmRepresentation.vue'
import Navbar from '../components/ui/NavbarMain.vue'

export default {
    components: {
        Navbar,
        FarmRepresentation
    }
}
</script>

<style scoped>
   
</style>