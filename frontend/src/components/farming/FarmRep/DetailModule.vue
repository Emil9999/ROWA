<template>
    <div class="centered-div w-full px-10 p-grey-small">
        <div class="h-green-big">Select Herb Variety</div>
        <div >Harvest from the corresponding position in the next step.</div>
        <div class="w-full mt-10">
            <div :class="'grid items-end grid-cols-'+modulePlants.length+' gap-4'">
                <div class="mx-auto" v-for="farmModule in modulePlants" :key="farmModule"> <img :width="farmModule.variety == '' ? 30 : 120" :class="[farmModule.variety == '' ? emptySpaceClass : '' ]" :src="require('../../../assets/img/plant_svg/'+cImage(farmModule.variety))"/></div>
            </div>
        <div :class="'grid grid-cols-'+modulePlants.length+' gap-4 bg-gradient-to-b py-3 h-auto from-grey to-accentwhite '">
                <div class="grid gap-3" v-for="farmModule in modulePlants" :key="farmModule">
                <div>
                    <div>
                        <div :class="{'invisible': !inFarmModule(farmModule)}" class="bg-green mx-auto rounded-full h-6 w-6"><CheckIcon class="text-white"/></div> 
                        
                    </div>
                    <div><button :disabled="!inFarmModule(farmModule)" :class="{'text-opacity-80': !inFarmModule(farmModule), 'text-grey ' : farmModule.variety == ''}"  class="btn-selector-white" @click="$emit('SelectedPlant', findPosInPlantable(farmModule.position))">{{(farmModule.variety != '') ? farmModule.variety: 'Empty'}}</button> </div>
                    <div>Position: {{farmModule.position}}</div>
                    
                    </div>
    
                </div>
            <div class=" col-span-4">Position 1 is the most inner Plant</div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, PropType} from 'vue'
import FarmablePlant from '../../../types/FarmablePlant'
import { CheckIcon } from '@heroicons/vue/solid' 
import {checkImage} from '../../../composables/use_imgChecker'
import getPlantInModule  from '../../../composables/use_getPlantInModule'
import Plant from '../../../types/Plant'

export default defineComponent({
    components:{ CheckIcon},
    props: {
    farmModules:{
        type: Array as PropType<Array<FarmablePlant>>,
        required: true
    },

    },
    setup(props){
        const  {modulePlants, loadModulePlants} = getPlantInModule(1)

        loadModulePlants
        const emptySpaceClass = ref('filter grayscale opacity-75')
        const inFarmModule = (farmModule: Plant) => {
            return props.farmModules.find((e) => e.position === farmModule.position)
        }
        const findPosInPlantable = (plantPos: number) => {
           const i = props.farmModules.findIndex((e) => e.position === plantPos)
           return props.farmModules[i]
        }
        const {cImage} = checkImage("svg")

        return {modulePlants,emptySpaceClass, findPosInPlantable, inFarmModule, cImage}

    }
   
})
</script>
