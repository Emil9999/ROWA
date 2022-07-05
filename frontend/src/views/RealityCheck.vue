<template>
<div>
    <rcTopRow/> 
    <div v-if="!detailView">
    <FarmRepInfo/>
    <FarmRep @ModuleClicked="moduleSelected"/>
    
    </div>
    <div v-else>
    <DetailModule :moduleNum="selectedModule"/>
    </div>
</div>   
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import rcTopRow from '../components/admin/atoms/rcTopRow.vue'
import FarmRep from '../components/farming/FarmRep/FarmRepresentation.vue'
import FarmRepInfo from '../components/home/atoms/FarmRepInfo.vue'
import DetailModule from '../components/admin/rcDetailModule.vue'

export default defineComponent({
    components: {rcTopRow, FarmRep, FarmRepInfo, DetailModule},
    setup() {

        const detailView = ref(false)
        const selectedModule = ref(1)

        const moduleSelected = (mNumber: number) => {
            detailView.value = !detailView.value
            selectedModule.value = mNumber
        }
        
        return {detailView, selectedModule, moduleSelected}
    },
})
</script>
