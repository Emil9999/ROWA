<template>
<div class="w-full px-10">
   <div>
       <div v-if="!padActive" class="centered-div">
       <div class="h-green-big mb-10">Pump Time</div>
        <timeField :TimeOne="pumpTime.StartTime"></timeField>
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
        <timeField :Focused="true" :TimeOne="pumpTime.nTime.time" :Text="'Change Start Time'"></timeField>
        <numberPad :altText="numberesFullEntered ? 'Save' : ''" @clickedAlt="pumpTime.sendTimes, padActive = false" @numberClicked="pumpTime.nTime.addNumber" @backClicked="pumpTime.nTime.removeNumber"/>
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
        
        

        const numberesFullEntered = computed(() => {
            return (pumpTime.nTime.time.minutes.length == 2 &&
                    pumpTime.nTime.time.hours.length == 2)})

        const timeOptions = ref<number[]>([20,30,40])

        const fastSelect = (index: number) => {
            selectedFast.value = index
            pumpTime.sendTimes(numberesFullEntered.value ? pumpTime.nTime.time : pumpTime.StartTime,selectedFast.value != -1 ? timeOptions.value[selectedFast.value]: pumpTime.PumpDuration)
        }

        return{pumpTime, selectedFast, padActive, timeOptions, fastSelect, numberesFullEntered}
    },
})
</script>