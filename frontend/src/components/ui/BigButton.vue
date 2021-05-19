<template>
  <div class="text-center">
      <v-btn v-if="secondary"
      rounded
      color="secondary"
      dark
    > {{Text}}
        </v-btn>
    <v-btn
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
