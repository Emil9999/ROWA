<template>
  <div class="pa-4 secondary text-no-wrap rounded-pill">
      <v-btn-toggle v-model="toggle_exclusive" mandatory>
        <v-btn rounded color="primary">
          {{ firstElement }} </v-btn
        ><v-btn rounded color="primary">
          {{ secondElement }}
        </v-btn>
      </v-btn-toggle>
  </div>
</template>

<script>
export default {
  props: {
    length: Number,
    firstElement: String,
    secondElement: String,
  },
};
</script>

<style lang="scss" scoped></style>
