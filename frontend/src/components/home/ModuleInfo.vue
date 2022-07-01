<template>
    <div class="centered-div">
        
        <div v-for="plants in modulePlants" :key="plants" class="h-green-big">{{plants?.variety}}
            <div v-if="plants.variety != ''">
            <div class="p-grey-small">{{text.text}}</div>

            <div> 
                <button v-if="bFarmable(plants.position, modulePlantable)" @click="onClick((bFarmable(plants.position, modulePlantable)), 'p')" class="btn-small-green">Plant</button>
                <button v-if="bFarmable(plants.position, moduleHarvestable)" @click="onClick((bFarmable(plants.position, moduleHarvestable)), 'h')"  class="btn-small-green">Harvest</button>
            </div>
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
                            Varieties.value = findUniqueTypes(modulePlants.value).value
                            if (group.value == "lettuce"){
                                modulePlants.value = modulePlants.value?.slice(1, 2)
                            }

                            })
       
        
        return {text, bFarmable, onClick, Varieties,modulePlants}
    },
})
</script>
