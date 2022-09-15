<template>
    <div v-if="modulePlants[selectedPlantIndex] != null" class="centered-div flex-auto max-h-full h-full min-h-0">
    <div class="flex-fit min-h-fit justify-self-center">
        <div class="h-green-big my-5">{{modulePlants[selectedPlantIndex].plant_type}}</div>
        <div v-if="group != 'lettuce'" :class="{'flex-row-reverse':reverse}" class="flex gap-10">
            <div v-for="position in plantcountInModule" :key="position" class="h-green-big mb-5">
                    <button @click="(selectedPlantIndex  = (findPlant(position) ? modulePlants.findIndex(e => e.position == position) : selectedPlantIndex)), $emit('selectedPlant', group)" :class="modulePlants[selectedPlantIndex].position != position ? 'btn-selector-white' : 'btn-selector-green'">{{findPlant(position)  ? findPlant(position).plant_type : 'Empty'}}</button>
            </div>
        </div>
    </div>
     <div class="flex-grow h-72 min-h-0 overflow-scroll mx-5 w-10/12" v-if="modulePlants[selectedPlantIndex].plant_type != ''">
            
            <div class="h-grey-small italic my-3">({{text.altname}})</div>
            <div class="p-grey-big overflow-clip">{{text.text}}</div>

            <div class="grid grid-cols-3 gap-5 my-10">
                <div v-for="(infopart, index) in text.info" :key="index">
                        <div class="h-green-small">{{categoryNames[index]}}</div>
                        <div class="p-grey-big">{{infopart}}</div>
                </div>
            </div>

             <div class="h-green-big overflow-clip my-5">Harvesting Technique</div>
            <div class="p-grey-big overflow-clip">{{text.technique}}</div>
     </div>
   
      <div class="flex-fit min-h-fit my-3" v-if="modulePlants[selectedPlantIndex].plant_type != ''">
            <div class="h-green-big"> 
                <button v-if="bFarmable(modulePlants[selectedPlantIndex].position, modulePlantable)" @click="onClick((bFarmable(modulePlants[selectedPlantIndex].position, modulePlantable)), 'p')" class="btn-small-green">Plant</button>
                <button v-if="bFarmable(modulePlants[selectedPlantIndex].position, moduleHarvestable)" @click="onClick((bFarmable(modulePlants[selectedPlantIndex].position, moduleHarvestable)), 'h')"  class="btn-small-green">Harvest</button>
            </div>
    </div>
    
    
    </div>
    <div v-else class="centered-div">
            <div class="h-green-big my-10">Nothing planted here yet!</div>
            

    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, watch, computed} from 'vue'
import {useRouter} from 'vue-router'
import getPlantInModule from '../../composables/use_getPlantInModule'
import findUniqueTypes from '../../composables/use_FindUniqueTypes'
import findModuleGroup from '../../composables/use_findModuleGroup'
import FarmablePlant from '../../types/FarmablePlant'
import plantInfo from '../../assets/text/plantInfo.json'

export default defineComponent({
    emits:['selectedPlant'],
    props:{
        moduleNumber:{
            type: Number,
            required: true
        },
        moduleHarvestable:{
            type: Array as PropType<Array<FarmablePlant>>,
            required: true
        },
        modulePlantable:{
            type: Array as PropType<Array<FarmablePlant>>,
            required: true
        }
        
    },
    setup(props, {emit}) {
    
        const {modulePlants, loadModulePlants, plantcountInModule}  = getPlantInModule(props.moduleNumber)
        loadModulePlants()
        
        const {group, findGroup} = findModuleGroup()
        const moduleNumberAlias = toRef(props, 'moduleNumber')

        const text = computed(() => {
            if(Object.keys(plantInfo).find(e=> e == modulePlants.value[selectedPlantIndex.value]?.plant_type)){
                return plantInfo[modulePlants.value[selectedPlantIndex.value].plant_type as keyof typeof plantInfo]
                //return plantInfo[modulePlants.value[selectedPlantIndex.value].plant_type]
            } else {
                return plantInfo.default
            }
        })

        const categoryNames =  ref(plantInfo.titles)


        const Varieties = ref(findUniqueTypes(modulePlants.value))
        const router = useRouter()
        const reverse = computed(() => {
             if(props.moduleNumber%2 != 0){
                return true } else {
                    return false
                }
        })
        const selectedPlantIndex = ref(0)

          
        const findPlant = (position: number) =>{
            return modulePlants.value.find(e => e.position == position)

        }
        
        const onClick = (selected: any , farmingType: string) => {
          
            if(typeof selected == "object")
            router.push( {
                name:"directFarming", params: {farmingType: farmingType,planttype: selected.plant_type, leafHarvest: selected.leaf_harvest == true ? 1 : 0, group: selected.group,planter: selected.planter, modulenumber: selected.module, position: selected.position }})
                
        
        }

        const bFarmable = (position: number, list: FarmablePlant[]) => {
            if (group.value == "lettuce"){
                return (list.length > 0  ? list[0] : null)
            } else {
                return (list.find(e => e.position === position))
            }

        }
        
        
        watch(group, () =>{ console.log(group.value)
            emit('selectedPlant', group.value || 'empty')
                           })
        watch(moduleNumberAlias, () =>{ 
                            loadModulePlants(props.moduleNumber)
                            findGroup(props.moduleNumber)
                            selectedPlantIndex.value = 0
                            Varieties.value = findUniqueTypes(modulePlants.value).value
                            emit('selectedPlant', group.value || 'empty')
                            if (group.value == "lettuce"){
                                modulePlants.value = modulePlants.value?.slice(1, 2)
                            }                
                            
                             })
       
        
        return {text, bFarmable, categoryNames, onClick, Varieties,modulePlants, reverse, findPlant, group, selectedPlantIndex, plantcountInModule}
    },
})
</script>
