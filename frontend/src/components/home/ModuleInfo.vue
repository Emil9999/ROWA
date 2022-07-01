<template>
    <div class="centered-div">
        
        <div v-if="group != 'lettuce'" class="grid grid-flow-col gap-10">
            <div v-for="(plants, index) in modulePlants" :key="plants" class="h-green-big">
                    <button @click="(plants.variety != '' ? selectedPlantIndex = index : null)" :class="selectedPlantIndex != index ? 'btn-selector-white' : 'btn-selector-green'">{{plants.variety != '' ? plants.variety : 'Empty'}}</button>
            </div>
        </div>
            <div v-if="modulePlants[selectedPlantIndex].variety != ''">
            <div class="p-grey-small">{{text.text}}</div>

            <div> 
                <button v-if="bFarmable(modulePlants[selectedPlantIndex].position, modulePlantable)" @click="onClick((bFarmable(modulePlants[selectedPlantIndex].position, modulePlantable)), 'p')" class="btn-small-green">Plant</button>
                <button v-if="bFarmable(modulePlants[selectedPlantIndex].position, moduleHarvestable)" @click="onClick((bFarmable(modulePlants[selectedPlantIndex].position, moduleHarvestable)), 'h')"  class="btn-small-green">Harvest</button>
            </div>
          </div>
        
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, watch} from 'vue'
import {useRouter} from 'vue-router'
import getPlantInModule from '../../composables/use_getPlantInModule'
import findUniqueTypes from '../../composables/use_FindUniqueTypes'
import findSinglePlantModule from '../../composables/use_findSinglePlantModule'
import FarmablePlant from '../../types/FarmablePlant'

export default defineComponent({

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
    setup(props) {
    
        const {modulePlants, loadModulePlants}  = getPlantInModule(props.moduleNumber)
        loadModulePlants()
        const {group, findGroup} = findSinglePlantModule()
        const moduleNumberAlias = toRef(props, 'moduleNumber')
        const Varieties = ref(findUniqueTypes(modulePlants.value))
        const router = useRouter()

        const selectedPlantIndex = ref(0)

        const onClick = (selected: FarmablePlant, farmingType: string) => {
            console.log(selected)
            router.push( {
                name:"directFarming", params: {farmingType: farmingType,planttype: selected.planttype, leafHarvest: selected.leafHarvest == true ? 1 : 0, group: selected.group,planter: selected.planter, modulenumber: selected.modulenumber, position: selected.position }})
                
        
        }

        const bFarmable = (position: number, list: FarmablePlant[]) => {
            if (group.value == "lettuce"){
                return (list.length > 0  ? list[0] : false)
            } else {
                return (list.find(e => e.position === position))
            }

        }
        
        const text = ref({
            name: 'Lollo',
            text: 'Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.'
        })
        watch(moduleNumberAlias, () =>{ 
                            loadModulePlants(props.moduleNumber)
                            findGroup(props.moduleNumber)
                            selectedPlantIndex.value = 0
                            Varieties.value = findUniqueTypes(modulePlants.value).value
                            if (group.value == "lettuce"){
                                modulePlants.value = modulePlants.value?.slice(1, 2)
                            }

                            if(props.moduleNumber%2 != 0){
                                modulePlants.value.sort((a,b) =>{
                                    return b.position - a.position 
                               } )
                            } else {
                                modulePlants.value.sort((a,b) =>{
                                    return a.position - b.position 
                               } )
                            }                    



                            })
       
        
        return {text, bFarmable, onClick, Varieties,modulePlants, group, selectedPlantIndex}
    },
})
</script>
