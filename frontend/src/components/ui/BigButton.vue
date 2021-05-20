<template>
  <div class="text-center">
      <v-btn v-if="secondary" width="250"
          x-large

      rounded
      color="secondary"
      dark
      @click="buttonClicked()"
    > {{Text}}
        </v-btn>
    <v-btn v-else width="250"
    x-large
      rounded
      color="primary"
      dark
      @click="buttonClicked()"
    >
    {{Text}}
    </v-btn>

    

  </div>
</template>

<script>
export default {
  props: {
    arrow: Boolean,
    secondary: Boolean,
    Text: String,
    
  },
  methods: {
    buttonClicked(){
      this.$emit('clicked')
    }
  }
}
</script>
