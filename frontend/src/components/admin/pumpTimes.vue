<template>
<div class="w-full px-10">
   <div>
       <div v-if="!padActive" class="centered-div">
       <div class="h-green-big mb-10">Pump Time</div>
        <timeField :TimeOne="pumpTime.StartTime" :TimeTwo="pumpTime.EndTime"></timeField>
                <div @click="padActive = true" class="btn-small-green my-5">New Time</div>

        <div class="h-green-small mt-10">Pump Duration</div>
        <div  class="flex justify-around w-1/2 my-5">
       <div v-for="(time, index) in timeOptions" :key="index" >
        <div :class="time == pumpTime.PumpDuration ? 'bg-green text-white' : 'bg-white text-green'" class="py-3 px-8 p-green-small text-2xl rounded-lg flex justify-center">
            <div @click="fastSelect(index)">{{time}}</div>
        </div>
       </div>
        </div>
       </div>
       <div v-else class="centered-div">
           <div @click="padActive = false" class="rounded-full h-10 w-10 mr-auto mb-5"><ArrowLeftIcon class="text-green"/></div>
        <timeField class="my-5" :Focused="selectedField==0" :TimeOne="pumpTime.nTimes[0].time" :Text="'Start Time:  '" @click="selectedField = 0"></timeField>
        <timeField :Focused="selectedField==1" :TimeOne="pumpTime.nTimes[1].time" :Text="'End Time:  '" @click="selectedField = 1"></timeField>
        <numberPad :altText="numberesFullEntered ? 'Save' : ''" @clickedAlt="pumpTime.sendTimes(), padActive = false" @numberClicked="pumpTime.nTimes[selectedField].addNumber" @backClicked="pumpTime.nTimes[selectedField].removeNumber"/>
       </div>
   </div>
</div>
</template>
<script lang="ts">
import { defineComponent, onUpdated, ref, reactive, computed } from 'vue'
import timeField from './atoms/timeField.vue'
import numberPad from './atoms/numberPad.vue'
import {ArrowLeftIcon} from '@heroicons/vue/solid'
import {pumpTimes} from '../../composables/use_PumpTimer'

export default defineComponent({
    components:{timeField,ArrowLeftIcon, numberPad},
    setup() {
       
        const pumpTime = reactive(pumpTimes())
        const selectedFast = ref(-1)
        const padActive = ref(false)

        const selectedField = ref(1)
        
        

        const numberesFullEntered = computed(() => {
            return (pumpTime.nTimes[0].time.minutes.length == 2 &&
                pumpTime.nTimes[0].time.hours.length == 2 &&
                pumpTime.nTimes[1].time.minutes.length == 2 &&
                pumpTime.nTimes[1].time.hours.length == 2)})

        const timeOptions = ref<number[]>([20,30,40])

        const fastSelect = (index: number) => {
            selectedFast.value = index
            pumpTime.sendTimes(numberesFullEntered.value ? pumpTime.nTimes[0].time : pumpTime.StartTime, 
                                numberesFullEntered.value ? pumpTime.nTimes[1].time : pumpTime.EndTime,
                                selectedFast.value != -1 ? timeOptions.value[selectedFast.value]: pumpTime.PumpDuration)
        }

        return{pumpTime, selectedFast, padActive, timeOptions, selectedField, fastSelect, numberesFullEntered}
    },
})
</script>