<template>
    <div class="centered-div w-full px-10 p-grey-small">
        <div class="h-green-big">Select Herb Variety</div>
        <div >Harvest from the corresponding position in the next step.</div>
        <div class="w-full mt-10">
            <div :class="'grid grid-cols-'+farmModules.length+' gap-4'">
                <div class="mx-auto" v-for="farmModule in this.farmModules" :key="farmModule"> <img src="../../../assets/img/plant_svg/herb1.svg"/></div>
            </div>
        <div :class="'grid grid-cols-'+farmModules.length+' gap-4 bg-gradient-to-b py-3 h-auto from-grey to-accentwhite '">
                <div class="grid gap-3" v-for="farmModule in this.farmModules" :key="farmModule">
                    
                    <div>
                        <div class="bg-green mx-auto rounded-full h-6 w-6"><CheckIcon class="text-white"/></div> 
                        
                    </div>
                    <div><button class="btn-selector-white" @click="$emit('SelectedPlant', farmModule)">{{farmModule.planttype}}</button> </div>
                    <div>Position: {{farmModule.position}}</div>
                    
                    </div>

            <div class=" col-span-4">Position 1 is the most inner Plant</div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, PropType} from 'vue'
import FarmablePlant from '../../../types/FarmablePlant'
import { CheckIcon } from '@heroicons/vue/solid' 

export default defineComponent({
    components:{ CheckIcon},
    props: {
    farmModules:{
        type: Array as PropType<Array<FarmablePlant>>,
        required: true
    },
    },
   
})
</script>
