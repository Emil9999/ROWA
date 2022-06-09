<template>
<div class="info-box-instruc">
  <h1 class="h-green-big">{{farmingInstruc[infoType].headline}}</h1>
  <SelectorPill class="my-5" :defaultSelection="'Lettuce'" :menuPoints="['Lettuce','Leafs']" @ClickedRow="changeInstruc" v-if="leafHarvest != null"/>
  <div v-if="leafHarvest == false" class="p-grey-small">Leafs have already been harvested.</div>
  <ol class="h-green-mid m-6">
    <li v-for="(step, index) in farmingInstruc[infoType].steps" :key="step.key" class="flex m-8"><div class="bg-accentwhite rounded-full mr-4 py-0.5 px-3">{{index}}</div>{{step}}</li>
  </ol>
  <button @click="buttonPresse" class="btn-big-green">{{farmingInstruc[infoType].buttonText}}</button>

</div>
    
</template>

<script lang="ts">
import { defineComponent, ref} from 'vue';
import SelectorPill from '../general/SelectorPill.vue'

export default defineComponent({
  components: {SelectorPill},
  name: 'instructionsWindow',
  setup(props){
    const infoType = ref(props.infoTypeprop)
    const farmingInstruc =  { 
                              p_salad:{
                                        headline: 'Planting Instructions',
                                        buttonText: 'I Planted',
                                        steps: {
                                        1: "Move all plants to the outside.",
                                        2: "Take a Seed and a Pot.",
                                        3: "Put the Seed into the Pot.",
                                        4: "Plant it into the Module."
                                        }
                                        },
                              h_salad:{
                                        headline: 'Harvesting Instructions',
                                        buttonText: 'I harvested a lettuce',
                                        steps: {
                                        1: "Get a Drip-Pot (for the step 3).",
                                        2: "Move the plant to the outside.",
                                        3: "Take Plant & place into Drip-Pot.",
                                        4: "Remove Plant and return Pots."
                                        }
                                        },
                              h_salad_leaf:{
                                        headline: 'Harvesting Instructions',
                                        buttonText: 'I harvested leafs',
                                        steps: {
                                        1: "Get the scissors.",
                                        2: "Choose the most outer lettuce.",
                                        3: "Try to leave young leafs intact.",
                                        4: "Carefully cut big, outer leafs."
                                        }
                                        },
                             }
    
    

    return {farmingInstruc, infoType}
  },
  props:{ 
      infoTypeprop:{
          type: String,
          default: "p_salad"

      },
      leafHarvest:{
        type: Boolean,
        default: null,
      }
  },
  emits: ["buttonPressed"],
  methods: {
    buttonPresse() {
      this.$emit("buttonPressed");
    },
     changeInstruc(clickedRow:string){
        if (clickedRow == 'Leafs'){
          this.infoType = 'h_salad_leaf'
        } else {
          this.infoType = 'h_salad'
        }
        
    }
  }
  
});
</script>