<template>
    <div :class="Focused ? 'bg-green' : 'bg-white'" class="flex-grow flex-col p-3 p-green-small rounded-lg flex justify-center">
        <div :class="Focused ? 'h-realwhite-small' : 'h-green-small'">{{TimeOne?.hours}}:{{TimeOne?.minutes}} - {{TimeTwo?.hours}}:{{TimeTwo?.minutes}}</div>
        <div class="h-white-small" v-if="TimeTwo != undefined && TimeOne != undefined">{{Number(TimeTwo.hours)-Number(TimeOne.hours)}} h</div>
    </div>    
</template>
<script lang="ts">
import { defineComponent, PropType } from 'vue'
import timeString from '../../../types/timeString'

export default defineComponent({
      props: {
        TimeOne:{
            type: Object as PropType<timeString>,
            required: true,
        },
        TimeTwo:{
            type: Object as PropType<timeString>,
            required: true
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
