<template>
<div>
    <rcTopRow class="p-5" @backClicked="detailView  ? detailView = !detailView : this.$router.push('/admin')"/> 
    <div v-if="!detailView" class="centered-div">
    <FarmRepInfo/>
    <FarmRep @ModuleClicked="moduleSelected"/>
    
    </div>
    <div v-else>
    <DetailModule @saved="reset" :moduleNum="selectedModule"/>
    </div>
</div>   
</template>

<script lang="ts">
import { defineComponent, ref, provide } from 'vue'
import rcTopRow from '../components/admin/atoms/rcTopRow.vue'
import FarmRep from '../components/farming/FarmRep/FarmRepresentation.vue'
import FarmRepInfo from '../components/home/atoms/FarmRepInfo.vue'
import DetailModule from '../components/admin/rcDetailModule.vue'

export default defineComponent({
    components: {rcTopRow, FarmRep, FarmRepInfo, DetailModule},
    setup() {

        provide('showunavailable', true)
        const detailView = ref(false)
        const selectedModule = ref(1)

        const reset = () =>{
                detailView.value = false
        }

        const moduleSelected = (mNumber: number) => {
            detailView.value = !detailView.value
            selectedModule.value = mNumber
        }
        
        return {detailView, reset, selectedModule, moduleSelected}
    },
})
</script>
