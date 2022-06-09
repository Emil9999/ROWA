<template>
    <div class="centered-div">
        <div></div>
        <div v-for="variety in Varieties" :key="variety" class="h-green-big">{{variety}}
        <div class="p-grey-small">{{text.text}}</div>

        <div> 
            <button v-if="modulePlantable.find(e => e.planttype === variety)" class="btn-small-green">Plant</button>
            <button @click="onClick" v-if="moduleHarvestable.find(e => e.planttype === variety)" class="btn-small-green">Harvest</button>
        </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, PropType, onMounted, watch, onUpdated } from 'vue'
import ModuleFarmable from '../../types/ModuleFarmable'
import {useRouter} from 'vue-router'
import getPlantInModule from '../../composables/use_getPlantInModule'
import findUniqueTypes from '../../composables/use_FindUniqueTypes'

export default defineComponent({

    props:{
        moduleNumber:{
            type: Number,
            required: true
        },
        moduleHarvestable:{
            type: Array as PropType<Array<ModuleFarmable>>,
            required: true
        },
        modulePlantable:{
            type: Array as PropType<Array<ModuleFarmable>>,
            required: true
        }
        
    },
    setup(props) {
        const {modulePlants, loadModulePlants, moduleNumber: selectedModule}  = getPlantInModule(props.moduleNumber)
        loadModulePlants()
        const moduleNumberAlias = toRef(props, 'moduleNumber')
        const Varieties = findUniqueTypes(modulePlants.value)
        const router = useRouter()
        const moduleFarmable =  ref<ModuleFarmable[]>([{
            varietyName: 'Lollo Bionda',
            harvestables: {planttype: 'Lollo Bionda', planter: '', modulenumber: 1, position: 1},
            plantables: {planttype: 'Lollo Bionda', planter: '', modulenumber: 1, position: 4}
        }])

        const onClick = () => {
            router.push( {
                name:"directFarming", params: {farmingType: 'p',planttype: 'Lollo Bionda', planter: 'Lollo Bionda', modulenumber: 1, position: 1}})
        }
        
        const text = ref({
            name: 'Lollo',
            text: 'Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.'
        })
        watch(moduleNumberAlias, () =>{ 
                            selectedModule.value = props.moduleNumber
                            loadModulePlants()
                            Varieties.value = findUniqueTypes(modulePlants.value).value
                            })
       

        return {text, moduleFarmable, onClick, Varieties,modulePlants}
    },
})
</script>
