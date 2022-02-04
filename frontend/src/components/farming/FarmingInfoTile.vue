<template>
    <div class="bg-white rounded-xl grid grid-cols-2 w-96">
        <div class="px-2">
            <p class="p-grey-small py-2">You are harvesting:</p>
            <h1 class="h-green-small">{{plantType}}</h1>
            <p class="p-grey-small">Module {{position}}</p>
            <div v-if="filds != null"><PositionIndicator :filds="filds" :pos="pos"></PositionIndicator></div>
            <p v-if="planter != ''" class="p-grey-small">Planted by: {{planter}}</p>
        </div>
        <div class="place-self-center">
            <img class="p-4" src="../../assets/img/plants/mint.png">
        </div>
        
        </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import PositionIndicator from './PositonIndicator.vue'

export default defineComponent({
    props: {
    planter:{
        type: String,
        required: false
    },
    plantType: {
        type: String,
        required: true
    },
    position: {
        type: Number,
        required: true
    },
    filds: {
        type: Number,
        default: 2,
        required: true
    },
    pos: {
        type: Number,
         default: 1,
        required: true
    },
    },

 
    components:{PositionIndicator},
})
</script>
