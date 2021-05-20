<template>
  <v-slide-group
      v-model="model"
      class="pa-4"
      active-class="success"
      show-arrows
    >
      <v-slide-item
        v-for="n in 10"
        :key="n"
        v-slot="{ active, toggle }"
      >
        <v-card
          :color="active ? undefined : 'white'"
          class="ma-8"
          height="300"
          width="150"
          @click="toggle"
        >
        <v-card-title >
           Title
        </v-card-title>
        <v-img
            src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="100px"
            >       </v-img>
        <v-card-subtitle>
      1,000 miles of wonder
    </v-card-subtitle>
        </v-card>
      </v-slide-item>
    </v-slide-group>
</template>

<script>
export default {
    props: {
        Ideas: [{title:"Smoothie"}]
    }
}
</script>

<style>

</style>