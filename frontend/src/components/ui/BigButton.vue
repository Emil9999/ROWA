<template>
  <div class="text-center">
      <v-btn v-if="secondary" width="250"
          x-large

      rounded
      color="secondary"
      dark
    > {{Text}}
        </v-btn>
    <v-btn v-else width="250"
    x-large
      rounded
      color="primary"
      dark
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
}
</script>
