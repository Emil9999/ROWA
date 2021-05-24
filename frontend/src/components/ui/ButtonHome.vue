<template>

      <v-btn-toggle v-model="toggle_exclusive" mandatory rounded color="white" :style="padding= 30">
        <v-btn rounded color="primary" @click="clickHandler('first')">
          {{ firstElement }} </v-btn
        ><v-btn rounded color="primary" @click="clickHandler('second')">
          {{ secondElement }}
        </v-btn>
        <v-btn v-if="thirdElement != null" rounded color="primary" @click="clickHandler('third')">
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
  methods: {
    clickHandler(num){
      this.$emit(num)
    }
  }
};
</script>

