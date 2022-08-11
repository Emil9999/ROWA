<template>
    <div class="flex w-135 justify-between m-auto items-center p-5 flex-grow rounded-md bg-almostwhite h-green-small">
        <div>{{text}}:</div>
        <div>Module:{{mNumber}}</div>
        <div class="bg-white">
            <img src="../../../assets/img/plants/default.png"/>
        </div>
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
