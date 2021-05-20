<template>

      <v-btn-toggle :color="active ? 'white' : '#009966'" mandatory borderless >
        <v-btn  color="white"  @click="clickHandler('first')">
         {{ firstElement }} </v-btn
        ><v-btn rounded color="white" @click="clickHandler('second')">
          {{ secondElement }}
        </v-btn>
        <v-btn v-if="thirdElement != null" rounded color="white" @click="clickHandler('third')">
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

<style lang="scss" scoped>

</style>
