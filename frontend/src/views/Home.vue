<template>
  <div class="home">
    <v-container>
      <v-row class="ma-10">
        <navbar/>
      </v-row>
      <v-row class="mt-10">
        <v-row align="center" justify="center">
          <h1 style="color: #006E5F">Explore Ideas</h1>
        </v-row>
        <v-row>
          <idea-carousel/>
        </v-row>
      </v-row>
      <v-row class="mt-10">
        <v-col>
          <v-row align="center" justify="center" >
            <h1 style="color: #006E5F"> What do you want to do? </h1>
          </v-row>
          <v-row align="center" justify="center" class="ma-10">
            <big-button Text="Harvest"/>
          </v-row>
          <v-row align="center" justify="center">
            <big-button Text="Plant"/>
          </v-row>
        </v-col>
      </v-row>
    </v-container>

  </div>
</template>

<script>
import BigButton from '../components/ui/BigButton.vue'
import IdeaCarousel from '../components/ui/IdeaCarousel.vue'
import Navbar from '../components/ui/Navbar.vue'
// @ is an alias to /src


export default {
  name: 'Home',
  components: {
    
    IdeaCarousel,
    BigButton,
    Navbar,
  }
}
</script>
