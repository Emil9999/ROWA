<template>
<div class="info-box-instruc">
  <h1 class="h-green-big">Planting Instructions</h1>
  <ol class="h-green-mid m-6">
    <li v-for="(step, index) in p_salad" :key="step.key" class="flex m-8"><div class="bg-accentwhite rounded-full mr-4 px-2">{{index}}</div>{{step}}</li>
  </ol>
  <button @click="buttonPresse" class="btn-big-green">I Planted</button>

</div>
    
</template>

<script lang="ts">
import { defineComponent} from 'vue';

export default defineComponent({
  components: {  },
  name: 'instructionsWindow',
  setup(){
    const p_salad =  {
                          1: "Move all plants to the outside.",
                          2: "Take a Seed and a Pot.",
                          3: "Put the Seed into the Pot.",
                          4: "Plant it into the Module."
                          }
      
    return {p_salad}
  },
  props:{ 
      infoType:{
          type: String,
          default: "p_salad"

      }
  },
  emits: ["buttonPressed"],
  methods: {
    buttonPresse() {
      this.$emit("buttonPressed");
    },
  }
  
});
</script>