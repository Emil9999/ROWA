<template>
<div>
   <div>
       <div v-if="!padActive" class="centered-div">
       <div class="h-green-big">Light Schedule</div>
        <timeField :TimeOne="lightTime.StartTime" :TimeTwo="lightTime.EndTime"></timeField>
                <div @click="padActive = true" class="btn-small-green">New Time</div>

        <div class="h-green-small">Quick Select</div>
        <div  class="flex justify-between w-full">
       <div v-for="(time, index) in fastOptions" :key="index">
            <pretiemField :Focused="index == selectedFast ? true : false" :TimeOne="time[0]" :TimeTwo="time[1]" @click="fastSelect(index)"></pretiemField>
       </div>
        </div>
       </div>
       <div v-else>
           <div @click="padActive = false" class="btn-small-green">Back</div>
        <timeField :Focused="selectedField==0" :TimeOne="lightTime.nTimes[0].time" :Text="'Change Start Time'" @click="selectedField = 0"></timeField>
        <timeField :Focused="selectedField==1" :TimeOne="lightTime.nTimes[1].time" :Text="'Change End Time'" @click="selectedField = 1"></timeField>
        <numberPad :altText="numberesFullEntered ? 'Save' : ''" @clickedAlt="lightTime.sendTimes" @numberClicked="lightTime.nTimes[selectedField].addNumber" @backClicked="lightTime.nTimes[selectedField].removeNumber"/>
       </div>
   </div>
</div>
</template>
<script lang="ts">
import { defineComponent, ref, reactive, computed } from 'vue'
import timeField from './atoms/timeField.vue'
import pretiemField from './atoms/pretimeField.vue'
import numberPad from './atoms/numberPad.vue'
import {lightTimes} from '../../composables/use_LightTimer'
import timeString from '../../types/timeString'

export default defineComponent({
    components:{timeField, numberPad, pretiemField},
    setup() {
       
        const lightTime = reactive(lightTimes())
        const selectedField = ref(1)
        const selectedFast = ref(-1)
        const padActive = ref(false)

        const numberesFullEntered = computed(() => {
            return (lightTime.nTimes[0].time.minutes.length == 2 &&
                    lightTime.nTimes[0].time.hours.length == 2 &&
                    lightTime.nTimes[1].time.minutes.length == 2 &&
                    lightTime.nTimes[1].time.hours.length == 2)
        })

        const fastOptions = ref<timeString[][]>([
           [ {
            minutes: '10',
            hours: '12'
            },
            {
            minutes: '10',
            hours: '18'
            }],
            [ {
            minutes: '10',
            hours: '18'
            },
            {
            minutes: '10',
            hours: '23'
            }]
        ])

        const fastSelect = (index: number) => {
            selectedFast.value = index
            lightTime.sendTimes(fastOptions.value[index][0], fastOptions.value[index][1])
        }

        return{lightTime, selectedFast, selectedField, padActive, fastOptions, fastSelect, numberesFullEntered}
    },
})
</script>