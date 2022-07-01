<template>
    <div :class="Focused ? ' border-green-dark border-2' : ''" class="flex-grow bg-brown p-3 p-green-small rounded-full flex">
        <div>{{Text}}</div>
        <div>{{TimeOne?.hours}}:{{TimeOne?.minutes}}</div>
        <div v-if="TimeTwo != undefined">- {{TimeTwo?.hours}}:{{TimeTwo?.minutes}}</div>
        <div v-if="TimeTwo != undefined && TimeOne != undefined">|{{TimeTwo.hours-TimeOne.hours}} h</div>
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
