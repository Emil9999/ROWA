<template>
    <div class="flex w-135 justify-between m-auto items-center px-10 flex-grow rounded-md bg-almostwhite h-green-small">
        <div>{{text}}:</div>
        <div>Module:{{mNumber}}</div>
        <img src="../../../assets/img/plants/default.png"/>

    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    props:{
        text: {
            type: String,
            default: 'default'
        },
        mNumber:{
            type: Number,
            default: 0
        }
    },
    
    setup() {
        return
    },
})
</script>
