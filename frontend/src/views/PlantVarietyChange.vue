<template>
<div>
    <rcTopRow class="p-5" @backClicked="detailView  ? detailView = !detailView : this.$router.push('/admin')"/> 
    <div v-if="!detailView" class="centered-div">
    <div class="info-box-instruc h-green-small">Click on a Moudle <br> to change its Plant Variety.</div>
    <FarmRep @ModuleClicked="moduleSelected"/>
    
    </div>
    <div v-else class="centered-div">
        <div class="p-grey-big">Module {{selectedModule}}</div>
        <div>
            <div v-for="type in debugData" :key="type">
                <typeInfoTile :ptype="type" :selected="true"/>
            </div>
        </div>
        <div>
            Selectables
        </div>
        <div>
            <button class="btn-big-green">Save Selection</button>
        </div>
    </div>
</div>   
</template>

<script lang="ts">
import { defineComponent, ref, provide } from 'vue'
import rcTopRow from '../components/admin/atoms/rcTopRow.vue'
import FarmRep from '../components/farming/FarmRep/FarmRepresentation.vue'
import typeInfoTile from '../components/admin/atoms/typeInfoTile.vue'

export default defineComponent({
    components: {rcTopRow, FarmRep, typeInfoTile},
    setup() {

        provide('showunavailable', true)
        const detailView = ref(false)
        const selectedModule = ref(1)

        const debugData = [{name: 'Lollo Bionda', gTime: 42}, {name: 'Thai Basil', gTime: 60}, {name:'Basil', gTime: 70}, {name: 'Mint', gTime: 56}]

        const moduleSelected = (mNumber: number) => {
            detailView.value = !detailView.value
            selectedModule.value = mNumber
        }
        
        return {detailView, selectedModule, debugData, moduleSelected}
    },
})
</script>
