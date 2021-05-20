<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4" style="color: #A0918D">
          {{ PlantHarvest }}
        </div>
        <v-list-item-title class="headline mb-1" style="color: #009966"> 
          {{ PlantName }}
        </v-list-item-title>
        <v-list-item-subtitle>{{ Module }}</v-list-item-subtitle>
      </v-list-item-content>

        <v-img :src="getImageUrl(Pic)"></v-img>
    </v-list-item>
  </v-card>
</template>

<script>
export default {
  props: {
    PlantHarvest: Number,
    PlantName: String,
    Module: String,
    Pic: String
  },

  methods: {
     getImgUrl(pic) {
    var images = require.context('@/assets/farm/plants', false, /\.svg$/)
    return images('./' + pic + ".svg")
  }
  },
};
</script>

<style lang="scss" scoped></style>
