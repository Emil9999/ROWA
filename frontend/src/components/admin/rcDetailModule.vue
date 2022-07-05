<template>
    <div class="centered-div w-full px-10 p-grey-small">
        <div class="h-green-big"></div>
        <div ></div>
        <div class="w-full mt-10">
            <div :class="'grid items-end grid-cols-'+modulePlants.length+' gap-4'">
                <div class="mx-auto" v-for="farmModule in modulePlants" :key="farmModule"> <img :width="plantWidth(farmModule.age, farmModule.growthTime)" :class="[farmModule.variety == '' ? emptySpaceClass : '' ]" :src="require('../../assets/img/plant_svg/'+cImage(farmModule.variety))"/></div>
            </div>
        <div :class="'grid grid-cols-'+modulePlants.length+' gap-4 bg-gradient-to-b py-3 h-auto from-grey to-accentwhite '">
                <div class="grid gap-3" v-for="farmModule in modulePlants" :key="farmModule">
                <div>
                    <div>
                        <div class="bg-green mx-auto rounded-full h-6 w-6 invisible"><CheckIcon class="text-white"/></div> 
                        
                    </div>
                    <div><button  :class="{'text-grey ' : farmModule.variety == ''}"  class="btn-selector-white" @click="$emit('SelectedPlant')">{{(farmModule.variety != '') ? farmModule.variety : 'Empty'}}</button> </div>
                    <div>Position: {{farmModule.position}}</div>
                    
                    </div>
    
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, watch, computed} from 'vue'
import FarmablePlant from '../../types/FarmablePlant'
import { CheckIcon } from '@heroicons/vue/solid' 
import {checkImage} from '../../composables/use_imgChecker'
import getPlantInModule  from '../../composables/use_getPlantInModule'
import Plant from '../../types/Plant'

export default defineComponent({
    components:{ CheckIcon},
    props: {
    moduleNum: {
        type: Number,
        required: true,
        default: 3
    }

    },
    setup(props){
        const  {modulePlants, loadModulePlants} = getPlantInModule(props.moduleNum)

        loadModulePlants()
        const moduleNumberAlias = toRef(props, 'moduleNum')
        const emptySpaceClass = ref('filter grayscale opacity-75')
        const inFarmModule = ref(true)


        const clamp = (num: number, min: number, max: number) => Math.min(Math.max(num, min), max);

        const plantWidth = (age: number, gTime: number) => {
            gTime = gTime == 0 ? 1 : gTime
            let wnumber = 20 + clamp((100*(age/gTime)), 0, 100)
            return wnumber
            };
        
        const {cImage} = checkImage("svg")
        watch(moduleNumberAlias, () =>{ 
                            loadModulePlants(moduleNumberAlias.value)
                            if(props.moduleNum%2 != 0){
                                modulePlants.value.sort((a,b) =>{
                                    return b.position - a.position 
                               } )
                            } else {
                                modulePlants.value.sort((a,b) =>{
                                    return a.position - b.position 
                               } )
                            }
                           
                            })
        return {modulePlants,emptySpaceClass, inFarmModule, cImage, plantWidth}

    }
   
})
</script>
