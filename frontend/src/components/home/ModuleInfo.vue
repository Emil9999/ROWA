<template>
    <div v-if="modulePlants[selectedPlantIndex] != null" class="centered-div flex-auto max-h-full h-full min-h-0">
    <div class="flex-fit min-h-fit justify-self-center">
        <div class="h-green-big my-5">{{modulePlants[selectedPlantIndex].plant_type}}</div>
        <div v-if="group != 'lettuce'" :class="{'flex-row-reverse':reverse}" class="flex gap-10">
            <div v-for="position in plantcountInModule" :key="position" class="h-green-big mb-5">
                    <button @click="(selectedPlantIndex  = (findPlant(position) ? modulePlants.findIndex(e => e.position == position) : selectedPlantIndex)), $emit('selectedPlant', findPlant(position)?.plant_type)" :class="modulePlants[selectedPlantIndex].position != position ? 'btn-selector-white' : 'btn-selector-green'">{{findPlant(position)  ? findPlant(position).plant_type : 'Empty'}}</button>
            </div>
        </div>
    </div>
     <div class="flex-grow h-72 min-h-0 overflow-scroll w-10/12" v-if="modulePlants[selectedPlantIndex].plant_type != ''">
            <div class="h-green-big overflow-clip mb-5">{{text.name}}</div>
            <div class="p-grey-small overflow-clip">{{text.text}}</div>

            <div class="grid grid-cols-3 gap-5 my-10">
                <div v-for="(infopart, index) in text.info" :key="index">
                        <div class="h-green-small">{{text.titles[index]}}</div>
                        <div class="p-grey-small">{{infopart}}</div>
                </div>
            </div>

             <div class="h-green-big overflow-clip my-5">Harvesting Technique</div>
            <div class="p-grey-small overflow-clip">{{text.technique}}</div>
     </div>
   
      <div class="flex-fit min-h-fit my-3" v-if="modulePlants[selectedPlantIndex].plant_type != ''">
            <div class="h-green-big"> 
                <button v-if="bFarmable(modulePlants[selectedPlantIndex].position, modulePlantable)" @click="onClick((bFarmable(modulePlants[selectedPlantIndex].position, modulePlantable)), 'p')" class="btn-small-green">Plant</button>
                <button v-if="bFarmable(modulePlants[selectedPlantIndex].position, moduleHarvestable)" @click="onClick((bFarmable(modulePlants[selectedPlantIndex].position, moduleHarvestable)), 'h')"  class="btn-small-green">Harvest</button>
            </div>
    </div>
    
    
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, watch, computed} from 'vue'
import {useRouter} from 'vue-router'
import getPlantInModule from '../../composables/use_getPlantInModule'
import findUniqueTypes from '../../composables/use_FindUniqueTypes'
import findModuleGroup from '../../composables/use_findModuleGroup'
import FarmablePlant from '../../types/FarmablePlant'

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
        
        const text = ref({
            name: 'Lollo',
            text: 'A vitamin C and antioxidant bomb, red basil is loved for its rich red color and great taste. Stronger than regular basil, red basil packs a punch that will make any salad, sandwich, or pesto you craft a complete knockout.Red Basil likes to be cut! You can start trimming your plants approximately one month after planting. You may remove single leaves occasionally.',
            technique: 'Find two of the largest leaves on the stem and locate the two small leaf nodes just above the larger set. Cut just above the two small nodes, about halfway down the plant. Pinching off the whole tip just above a true leaf pair - it will grow two new shoots in a week. Cutting back your Red Basil stems will prevent the plant from flowering and developing a bitter flavour. But, if you do want to see your Red Basil flowering, stop cutting your plant and let it thrive! Please note that basil is an annual plant. It does not last forever. If it has been over twelve weeks and the plant looks aged, it is time to replace the pod.',
            info: ['Slightly bitter, licorice.','7 - 14 days','5 - 12 weeks'],
            titles: ['Flavor Characteristics','Sprouts in:','Seed to Harvest']
            })
        watch(modulePlants, () =>{emit('selectedPlant', modulePlants.value[0].plant_type || 'empty')
                           })
        watch(moduleNumberAlias, () =>{ 
                            loadModulePlants(props.moduleNumber)
                            findGroup(props.moduleNumber)
                            selectedPlantIndex.value = 0
                            Varieties.value = findUniqueTypes(modulePlants.value).value
                            if (group.value == "lettuce"){
                                modulePlants.value = modulePlants.value?.slice(1, 2)
                            }                
                            
                             })
       
        
        return {text, bFarmable, onClick, Varieties,modulePlants, reverse, findPlant, group, selectedPlantIndex, plantcountInModule}
    },
})
</script>
