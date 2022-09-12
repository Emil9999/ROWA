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
            <div class="h-green-big overflow-clip">{{text.name}}</div>
            <div class="p-grey-small overflow-clip">{{text.text}}</div>
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
            name: 'Taste',
            text: 'Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eumccccccc iriure dolor in he Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet,ndrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.   Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet,' 
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
