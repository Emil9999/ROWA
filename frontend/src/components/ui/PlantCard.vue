<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4">
          {{ PlantHarvest }}
        </div>
        <v-list-item-title class="headline mb-1">
          {{ PlantName }}
        </v-list-item-title>
        <v-list-item-subtitle>{{ Module }}</v-list-item-subtitle>
      </v-list-item-content>

      <v-list-item-avatar>
        <v-img></v-img>
      </v-list-item-avatar>
    </v-list-item>
  </v-card>
</template>

<script>
export default {
  props: {
    plantHarvest: Number,
    PlantName: String,
    Module: String,
  },
};
</script>

<style lang="scss" scoped></style>
