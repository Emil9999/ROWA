<template>

      <v-btn-toggle v-model="toggle_exclusive" mandatory rounded color="white" :style="padding= 30">
        <v-btn rounded color="primary">
          {{ firstElement }} </v-btn
        ><v-btn rounded color="primary">
          {{ secondElement }}
        </v-btn>
        <v-btn v-if="thirdElement != null" rounded color="primary">
          {{ thirdElement }}
        </v-btn>
        
      </v-btn-toggle>

</template>

<script>
export default {
  props: {
    length: Number,
    firstElement: String,
    secondElement: String,
    thirdElement: String,
  },
};
</script>

<style lang="scss" scoped>

</style>
