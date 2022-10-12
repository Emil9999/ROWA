<template>
<div>
    <rcTopRow class="p-5" :text="'Change Variety'" @backClicked="detailView  ? detailView = !detailView : this.$router.push('/admin')"/> 
    <div v-if="!detailView" class="centered-div">
    <div class="info-box-instruc h-green-small">Click on a Moudle <br> to change its Plant Variety.</div>
    <FarmRep @ModuleClicked="moduleSelected"/>
    
    </div>
    <div v-else class="centered-div">
        <div class="p-grey-big my-5">Module {{selectedModule}}</div>
        <div class="grid grid-cols-2 gap-7">
            <button :disabled="selectedType.length > 0" @click="filterby == 'herb' ? filterby = '' : filterby = 'herb'" :class="filterby == 'herb' ? 'btn-admin-green' : 'btn-admin-white'">Herb</button>
            <button :disabled="selectedType.length > 0" @click="filterby == 'lettuce'? filterby = '' : filterby = 'lettuce'" :class="filterby == 'lettuce' ? 'btn-admin-green' : 'btn-admin-white'">Lettuce</button>
        </div>
        <div class=" overflow-scroll" style="height: 850px;">
            <div v-for="atype in filteredTypes" :key="atype.plant_type">
                <typeInfoTile  @click="clickType(atype)" :ptype="atype.plant_type" :previousS="currentTypes.findIndex(e => e.plant_type == atype.plant_type)>-1" 
                    :selected="selectedType.findIndex(e => e == atype.plant_type)>-1 "
                    :greyedOut="(selectedType.length >= (atype.group.toString() == 'herb' ? 4 : 1)) && !(selectedType.find(e => e == atype.plant_type))"/>
            </div>
        </div>
        <div>
        </div>
        <div class="w-full">
            <button class="btn-small-green invisible" @click="openSheet()">Add new plant type</button>
            <button class="btn-big-green" @click="sendVarietyChange">Save Selection</button>
        </div>
    </div>
    
    <Sheet :isopen="isOpen"  ref="newTypeSheet"><plantTypeConfigVue @saveNewType="addNewType"></plantTypeConfigVue></Sheet>
 

</div>   
</template>

<script lang="ts">
import { defineComponent, ref, provide, computed } from 'vue'
import rcTopRow from '../components/admin/atoms/rcTopRow.vue'
import axios from 'axios'
import FarmRep from '../components/farming/FarmRep/FarmRepresentation.vue'
import typeInfoTile from '../components/admin/atoms/typeInfoTile.vue'
import getAllTypes from '../composables/use_getAllTypes'
import AvailVariety from '../types/AvailVariety'
import getAvailTypesperModule from '../composables/use_getAvailableTypesperModule'
import plantTypeConfigVue from '@/components/admin/plantTypeConfig.vue'

import Sheet from '../bottom-sheet/bottom-sheet.vue'

export default defineComponent({
    components: {rcTopRow, FarmRep, typeInfoTile, Sheet, plantTypeConfigVue},
    setup() {

        provide('showunavailable', true)
        const detailView = ref(false)
        const selectedModule = ref(1)
        const {availTypes} = getAllTypes() 
        const filterby = ref('')
        const selectedType = ref<string[]>([])
        const {availTypes: currentTypes, loadTypes} = getAvailTypesperModule()
       
        const sendVarietyChange = () => {
            axios.post('/admin/change-planttype', {varieties: selectedType.value, modulenum: selectedModule.value})
            .catch(() => {console.log(selectedType.value)})
            selectedType.value = []
            detailView.value = false
            
        }
        
        function addNewType(name: string, group: string){
            console.log("Saved new Type")
            availTypes.value.push({plant_type: name, group: group})
            newTypeSheet.value?.close()
        }


        const filteredTypes = computed(() => {

            if (filterby.value == ''){
                return availTypes.value
            }
            return (availTypes.value.filter(e => e.group == filterby.value))
        })
        const clickType = (type: AvailVariety) =>{
            let sIndex = selectedType.value.findIndex(e => e == type.plant_type)
            if (sIndex >= 0){
                selectedType.value = selectedType.value.filter(e => e !== type.plant_type)
               // if (selectedType.value.length == 0){filterby.value = ''}
            } else {
                if (selectedType.value.length < (type.group.toString() == 'herb' ? 4 : 1)){selectedType.value.push(type.plant_type.toString())}
                if (filterby.value == ''){filterby.value = type.group.toString()}
            }
        }
        const moduleSelected = (mNumber: number) => {
            detailView.value = !detailView.value
            selectedModule.value = mNumber
            loadTypes(mNumber)
        }
        
        //Sheet
        const isOpen = ref(false)
        const newTypeSheet = ref<InstanceType<typeof Sheet> | null>(null)
        const openSheet = () => {
          
                newTypeSheet.value?.open()
        }



        return {detailView, isOpen, addNewType, newTypeSheet, openSheet, sendVarietyChange, selectedModule,currentTypes, selectedType, clickType, availTypes, moduleSelected,filterby, filteredTypes}
    },
})
</script>
