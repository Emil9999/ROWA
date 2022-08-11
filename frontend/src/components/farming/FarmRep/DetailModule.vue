<template>
    <div class="centered-div w-full px-10 p-grey-small">
        <div class="h-green-big">Select Herb Variety</div>
        <div >Harvest from the corresponding position in the next step.</div>
        <div class="w-full mt-10">
            <div :class="'grid items-end grid-cols-'+count+' gap-4'">
                <div class="mx-auto" :class="[reverse ? 'order-'+ ((count+1)-position) : 'order-'+ position]" v-for="position in count" :key="position"> 
                    <img :width="findPlant(position) ? 120 :  30" :class="[findPlant(position) ? '' : emptySpaceClass ]" :src="require('../../../assets/img/plant_svg/'+cImage(findPlant(position) ? findPlant(position).variety : 'default'))"/>
                  
                    </div>
            </div>
        <div :class="'grid grid-cols-'+count+' gap-4 bg-gradient-to-b py-3 h-auto from-grey to-accentwhite '">
                <div class="grid gap-3" :class="[reverse ? 'order-'+ ((count+1)-position) : 'order-'+ position]" v-for="position in count" :key="position">
                <div>
                    <div>
                        <div :class="{'invisible': !inFarmModule(position)}" class="bg-green mx-auto rounded-full h-6 w-6"><CheckIcon class="text-white"/></div> 
                        
                    </div>
                    <div><button :disabled="!inFarmModule(position)" :class="{'text-opacity-80': !inFarmModule(position), 'text-grey ' : !findPlant(position)}"  class="btn-selector-white" @click="$emit('SelectedPlant', inFarmModule(position))">{{findPlant(position) ? findPlant(position).variety : 'Empty'}}</button> </div>
                    <div>Position: {{position}}</div>
                    
                    </div>
    
                </div>
            <div :class="'order-'+(count+1) +' col-span-4'">Position 1 is the most inner Plant</div>
            </div>
            
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, watch, computed} from 'vue'
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
    moduleNum: {
        type: Number,
        required: true,
        default: 3
    }

    },
    setup(props){
        const  {modulePlants, loadModulePlants, plantcountInModule: count} = getPlantInModule(props.moduleNum)

        loadModulePlants()
        const moduleNumberAlias = toRef(props, 'moduleNum')
        const emptySpaceClass = ref('filter grayscale opacity-75')
        const inFarmModule = (position: number) => {
            return props.farmModules.find((e) => e.position === position)
        }
        const reverse = computed(() => {
             if(props.moduleNum%2 != 0){
                return true } else {
                    return false
                }
        })

        const findPlant = (position: number) =>{
            return modulePlants.value.find(e => e.position == position)

        }

        const findPosInPlantable = (plantPos: number) => {
           const i = props.farmModules.findIndex((e) => e.position === plantPos)
           return props.farmModules[i]
        }
        const {cImage} = checkImage("svg")
        watch(moduleNumberAlias, () =>{ 
                            loadModulePlants(moduleNumberAlias.value)
                            
                           
                            })
        return {modulePlants,emptySpaceClass, findPosInPlantable, inFarmModule, cImage, reverse, count, findPlant}

    }
   
})
</script>
