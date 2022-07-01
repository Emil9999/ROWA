<template>
<div>
   <div class="centered-div">
       <div v-if="!padActive">
       <div class="h-green-big">Light Schedule</div>
        <timeField :TimeOne="lightTime.StartTime" :TimeTwo="lightTime.EndTime"></timeField>
        <div @click="padActive = true" class="btn-small-green">New Time</div>
       </div>
       <div v-else>
           <div @click="padActive = false" class="btn-small-green">Back</div>
        <timeField :Focused="selectedField==0" :TimeOne="lightTime.nTimes[0].time" :Text="'Change Start Time'" @click="selectedField = 0"></timeField>
        <timeField :Focused="selectedField==1" :TimeOne="lightTime.nTimes[1].time" :Text="'Change End Time'" @click="selectedField = 1"></timeField>
        <numberPad :altText="'Save'" @clickedAlt="lightTime.sendTimes" @numberClicked="lightTime.nTimes[selectedField].addNumber" @backClicked="lightTime.nTimes[selectedField].removeNumber"/>
       </div>
   </div>
</div>
</template>
<script lang="ts">
import { defineComponent, ref, reactive } from 'vue'
import timeField from './atoms/timeField.vue'
import numberPad from './atoms/numberPad.vue'
import {lightTimes} from '../../composables/use_LightTimer'

export default defineComponent({
    components:{timeField, numberPad},
    setup() {
       
        const lightTime = reactive(lightTimes())
        const selectedField = ref(0)
        const padActive = ref(false)
        return{lightTime, selectedField, padActive}
    },
})
</script>