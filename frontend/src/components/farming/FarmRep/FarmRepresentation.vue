<template>
    <div class="relative w-farmW h-farmH flex-none">
        <img width="720" height="auto" src="../../../assets/cat_tree/FarmFrame.png"/>
        <div @click="$emit('ModuleClicked', i)" :style="'top: '+(77+((index)*165))+'px;'" class="z-10 absolute left-14" v-for="(i, index) in oddNumbers" :key="i">
            <Module :harvestableSpots="(this.farmingMode == 'p')? [] : harvestableSpots[i-1]" :plantableSpots="(this.farmingMode == 'h')? [] : plantableSpots[i-1]" :moduleNumber="i" :reverse="true"/>
   
             </div>
         <div @click="$emit('ModuleClicked', i)"  :style="'top: '+(77+(index)*165)+'px;'" class="z-10 absolute right-14" v-for="(i, index) in evenNumbers" :key="i">
            <Module :harvestableSpots="(this.farmingMode == 'p')? [] :harvestableSpots[i-1]" :plantableSpots="(this.farmingMode == 'h')? [] : plantableSpots[i-1]" :moduleNumber="i" />
          
             </div>
    </div>
</template>

/<script lang="ts">
import { defineComponent,computed, ref } from 'vue'

import Module from './Module.vue'
import getFarmablePosbyIndex from '../../../composables/use_getFarmablePosbyIndex'


export default defineComponent({
    components: {Module},
    props: {
        farmingMode:{
            type: String,
            required: false,
        }
    },
    setup() {

        const{harvestableSpots, plantableSpots} = getFarmablePosbyIndex()
        const numbers = ref([1, 2, 3, 4, 5, 6, 7, 8])

        const evenNumbers = computed(() => {
        return numbers.value.filter((n) => n % 2 === 0)
        })
        const oddNumbers = computed(() => {
        return numbers.value.filter((n) => n % 2 !== 0)
        })



         

        
        return {evenNumbers, oddNumbers, harvestableSpots, plantableSpots}
        
    },
    
})
</script>
