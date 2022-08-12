<template>
<div class="centered-div h-full">
    <rcModuleInfo :mNumber="moduleNum" :text="group == 'herb' ? 'Herbs' : selectedType.name"/>
    <div class="centered-div w-full px-10 p-grey-small">
        <div class="h-green-big"></div>
        <div :class="[reverse ? ' left-0' : ' right-0']" class=" z-0 w-10 opacity-70  h-96 bg-grey absolute"></div>
        <div class="w-full mt-10 z-10">
            <div :class="'grid items-end grid-cols-'+count+' gap-4'">
                <div class="mx-auto" :class="[reverse ? 'order-'+ ((count+1)-position) : 'order-'+ position]" v-for="position in count" :key="position"> 
                    <img :width="findPlant(position) ? plantWidth(findPlant(position).age, findPlant(position).growthTime) : 20" :class="[findPlant(position) ? '' : emptySpaceClass ]" :src="require('../../assets/img/plant_svg/'+cImage(findPlant(position) ? findPlant(position).variety : 'default'))"/></div>
            </div>
        <div :class="'grid grid-cols-'+count+' gap-4 bg-gradient-to-b py-3 h-auto from-grey to-accentwhite '">
                <div class="grid gap-3" :class="[reverse ? 'order-'+ ((count+1)-position) : 'order-'+ position]" v-for="position in count" :key="position">
                <div>
                    <div>
                        <div class="bg-green mx-auto rounded-full h-6 w-6 invisible"><CheckIcon class="text-white"/></div> 
                        
                    </div>
                    <div><button  :class="{'text-grey ' : !findPlant(position)}"  class="btn-selector-white" @click="applyType(position)">{{findPlant(position) ? generateButtonText(findPlant(position).age, findPlant(position).growthTime) : 'Empty'}}</button> </div>
                    <div>Position: {{position}}</div>
                    
                    </div>
    
                </div>
            </div>
        </div>
    </div>
    <div v-if="selectedPosition != -1">
        <div class="grid grid-cols-4 mt-28 gap-5" v-if="availTypes.length != 1">
            <div v-for="type in availTypes" :key="type"> 
                <div @click="selectedType = type" :class="[selectedType == type ? 'btn-admin-green' : 'btn-admin-white']">{{type.name}}</div>
            </div>
        </div>
        <div v-if="selectedType.name != ''">
        <div class="h-green-small m-10">Select a corresponding Growth Stage</div>
            <div  class="grid grid-cols-3 gap-5">    
            <div v-for="age in selectableTimes" :key="age"> 
                <div @click="updateAge(age)" class="btn-admin-white">{{generateButtonText(age)}}</div>
            </div>
            </div>
        </div>
    </div>
    <div v-else class="h-green-small">Click on a Plant-Position <br> to change its Growth Stage.</div>
    <button :disabled="!bChanges" :class="[bChanges ? 'btn-big-green' : 'btn-big-dis']" class="fixed bottom-0" @click="sendReality()">Save!</button>
   
</div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, watch, computed} from 'vue'
import { CheckIcon } from '@heroicons/vue/solid' 
import {checkImage} from '../../composables/use_imgChecker'
import getPlantInModule  from '../../composables/use_getPlantInModule'
import axios from 'axios'
import rcModuleInfo from './atoms/rcModuleInfo.vue'

import findModuleGroup from '../../composables/use_findModuleGroup'
import getAvailTypesperModule from '../../composables/use_getAvailableTypesperModule'

export default defineComponent({
    components:{ CheckIcon, rcModuleInfo},
    props: {
    moduleNum: {
        type: Number,
        required: true,
        default: 3
    }

    },
    setup(props){
        const  {modulePlants, loadModulePlants, plantcountInModule: count} = getPlantInModule(props.moduleNum)
        const {group, findGroup} = findModuleGroup()
        const {availTypes, loadTypes} = getAvailTypesperModule()
        loadModulePlants()
        findGroup(props.moduleNum)
        loadTypes(props.moduleNum)
        const moduleNumberAlias = toRef(props, 'moduleNum')
        const emptySpaceClass = ref('filter grayscale opacity-75')
        const inFarmModule = ref(true)
        const selectedType = ref(availTypes.value.length == 1 ? availTypes.value[0] : {name: '', gTime: 0})

        const applyType = (position: number) =>{
            selectedPosition.value = position
            let plant = modulePlants.value.find(e => e.position == selectedPosition.value)
            selectedType.value = availTypes.value.find(e => e.name == plant?.variety) || {name: '', gTime: 0}

        }

        const bChanges = ref(false)

        const reverse = computed(() => {
             if(props.moduleNum%2 != 0){
                return true } else {
                    return false
                }
        })
        
        const sendReality = () => {
            axios.post('/admin/reality-check', {...modulePlants.value, module: moduleNumberAlias})
            .catch(() => {console.log(modulePlants.value)})
            
        }

        const selectedIndex = ref(-1)
        const selectedPosition = ref(-1)
        const selectableTimes = computed(() => {
            if (group.value == 'lettuce' && selectedType.value.gTime != 0){
                return [0,7,14,21,28,35,selectedType.value.gTime]
            } else if (group.value == 'herb' && selectedType.value.gTime != 0){
                return [0,Math.round(selectedType.value.gTime/2),selectedType.value.gTime]
            }
                return [0,7,14,21,28,35,42]
            
            })
        const updateAge = (selectedAge: number) =>{

            let index = modulePlants.value.findIndex(e => e.position == selectedPosition.value)
            if(index == -1){
                if(selectedAge != 0){
                let newPlant = {variety: selectedType.value.name,age: selectedAge, growthTime: selectedType.value.gTime,position: selectedPosition.value}
                modulePlants.value.push(newPlant)}
            } else {
                modulePlants.value[index].age = selectedAge;
                modulePlants.value[index].variety = selectedType.value.name
                modulePlants.value[index].growthTime = selectedType.value.gTime
                if(selectedAge == 0){
                    modulePlants.value.splice(index, 1)
                    }
            }               
            selectedType.value = availTypes.value.length == 1 ? availTypes.value[0] : {name: '', gTime: 0}
            selectedPosition.value = -1
            bChanges.value = true
        }

        const findPlant = (position: number) =>{
            return modulePlants.value.find(e => e.position == position)

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
                           
                            })
        return {generateButtonText, bChanges, updateAge, applyType, selectedType, availTypes,group, reverse, count,selectedPosition, selectableTimes, selectedIndex, modulePlants,emptySpaceClass, inFarmModule, cImage, plantWidth, sendReality, findPlant}

    }
   
})
</script>
