<template>
    <div :class="Focused ? ' border-green border-2' : ''" class="flex-grow bg-accentwhite px-14 py-3 h-green-small min-w-64 rounded-2xl flex justify-center">
        <div class=" mr-6">{{Text}}</div>
        <div> {{TimeOne?.hours}}{{TimeOne?.hours.length == 2 ? ':':''}}{{TimeOne?.minutes}}<span v-if="Focused && !(TimeOne?.hours.length == 2 && TimeOne?.minutes.length == 2)" class="animate-writepulse">|</span></div>
        <div v-if="TimeTwo != undefined">- {{TimeTwo?.hours}}:{{TimeTwo?.minutes}}</div>
        <div v-if="TimeTwo != undefined && TimeOne != undefined">|{{Number(TimeTwo.hours)-Number(TimeOne.hours)}} h</div>
    </div>    
</template>
<script lang="ts">
import { defineComponent, PropType } from 'vue'
import timeString from '../../../types/timeString'

export default defineComponent({
      props: {
        TimeOne:{
            type: Object as PropType<timeString>,
            required: false,
        },
        TimeTwo:{
            type: Object as PropType<timeString>,
            required: false
        },
        Text: {
            type: String,
            default: 'Time set:'
        },
        Focused: {
            type: Boolean,
            default: false
        }
    },
    setup() {
        return
    },
})
</script>
