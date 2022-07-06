<template>
<div>
    <div class="centered-div w-full px-10 p-grey-small">
        <div class="h-green-big"></div>
        <div ></div>
        <div class="w-full mt-10">
            <div :class="'grid items-end grid-cols-'+modulePlants.length+' gap-4'">
                <div class="mx-auto" v-for="farmModule in modulePlants" :key="farmModule"> <img :width="plantWidth(farmModule.age, farmModule.growthTime)" :class="[farmModule.variety == '' ? emptySpaceClass : '' ]" :src="require('../../assets/img/plant_svg/'+cImage(farmModule.variety))"/></div>
            </div>
        <div :class="'grid grid-cols-'+modulePlants.length+' gap-4 bg-gradient-to-b py-3 h-auto from-grey to-accentwhite '">
                <div class="grid gap-3" v-for="(farmModule, index) in modulePlants" :key="farmModule">
                <div>
                    <div>
                        <div class="bg-green mx-auto rounded-full h-6 w-6 invisible"><CheckIcon class="text-white"/></div> 
                        
                    </div>
                    <div><button  :class="{'text-grey ' : farmModule.variety == ''}"  class="btn-selector-white" @click="selectedIndex = index">{{generateButtonText(farmModule.age, farmModule.growthTime)}}</button> </div>
                    <div>Position: {{farmModule.position}}</div>
                    
                    </div>
    
                </div>
            </div>
        </div>
    </div>
    <div v-if="selectedIndex != -1">
        <div class="grid grid-cols-4" v-if="availTypes.length != 1">
            <div v-for="type in availTypes" :key="type"> 
                <div @click="selectedType = type" class="btn-admin-white">{{type.name}}</div>
            </div>
        </div>
        <div class="grid grid-cols-3" v-if="selectedType.name != ''">
            <div v-for="age in selectableTimes" :key="age"> 
                <div @click="updateAge(age)" class="btn-admin-white">{{generateButtonText(age)}}</div>
            </div>
        </div>
    </div>
    <div :class="{'invisible': !bChanges}" class="btn-big-green" @click="console.log('Success')">Save!</div>
   
</div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, watch, computed} from 'vue'
import { CheckIcon } from '@heroicons/vue/solid' 
import {checkImage} from '../../composables/use_imgChecker'
import getPlantInModule  from '../../composables/use_getPlantInModule'

import findSinglePlantModule from '../../composables/use_findSinglePlantModule'
import getAvailTypesperModule from '../../composables/use_getAvailableTypesperModule'

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
        const {group, findGroup} = findSinglePlantModule()
        const {availTypes, loadTypes} = getAvailTypesperModule()
        loadModulePlants()
        findGroup(props.moduleNum)
        loadTypes(props.moduleNum)
        const moduleNumberAlias = toRef(props, 'moduleNum')
        const emptySpaceClass = ref('filter grayscale opacity-75')
        const inFarmModule = ref(true)
        const selectedType = ref(availTypes.value.length == 1 ? availTypes.value[0] : {name: '', gTime: 0})


        const bChanges = ref(false)

        const selectedIndex = ref(-1)
        const selectableTimes = computed(() => {
            if (group.value == 'lettuce' && selectedType.value.gTime != 0){
                return [0,7,14,21,28,35,selectedType.value.gTime]
            } else if (group.value == 'herb' && selectedType.value.gTime != 0){
                return [0,Math.round(selectedType.value.gTime/2),selectedType.value.gTime]
            }
                return [0,7,14,21,28,35,42]
            
            })
        const updateAge = (selectedAge: number) =>{
            modulePlants.value[selectedIndex.value].age = selectedAge;
            modulePlants.value[selectedIndex.value].variety = selectedType.value.name
            modulePlants.value[selectedIndex.value].growthTime = selectedType.value.gTime
            if(selectedAge == 0){modulePlants.value[selectedIndex.value].variety = ''}
            selectedType.value = availTypes.value.length == 1 ? availTypes.value[0] : {name: '', gTime: 0}
            selectedIndex.value = -1
            bChanges.value = true
        }

        const clamp = (num: number, min: number, max: number) => Math.min(Math.max(num, min), max);

        const plantWidth = (age: number, gTime: number) => {
            gTime = gTime == 0 ? 1 : gTime
            let wnumber = 20 + clamp((100*(age/gTime)), 0, 100)
            return wnumber
            };
        
        const generateButtonText = (age: number, growthT: number = selectedType.value.gTime) =>{
            if (age == 0 || growthT == 0){
                    return 'Empty'
                } 
            if(group.value == 'herb'){
                 if(age >= growthT){
                    return 'Ready to Harvest'
                } else { return 'Planted - Not Ready'}
            } else {
                if(age >= growthT){
                    return 'Ready to Harvest'
                } else {
                    return String(Math.round(age/7)) + ' Weeks'
                }

            }
        }
        const {cImage} = checkImage("svg")
        watch(moduleNumberAlias, () =>{ 
                            loadModulePlants(moduleNumberAlias.value)
                            findGroup(props.moduleNum)
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
        return {generateButtonText, bChanges, updateAge, selectedType, availTypes, selectableTimes, selectedIndex, modulePlants,emptySpaceClass, inFarmModule, cImage, plantWidth}

    }
   
})
</script>
