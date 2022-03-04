<template>
    <div class="bg-white items-center text-left shadow-md m-4 rounded-xl grid grid-cols-2 w-96">
        <div class="px-2">
            <p v-if="boxtype == 'h'" class="p-grey-small py-2">You are harvesting:</p>
             <p v-if="boxtype == 'p'" class="p-grey-small py-2">You are planting:</p>
            <h1 class="h-green-small">{{farmModule.planttype}}</h1>
            <p class="p-grey-small">Module {{farmModule.modulenumber}}</p>
            <div v-if="farmModule.position != 0"><PositionIndicator :filds="filds" :pos="farmModule.position"></PositionIndicator></div>
            <p v-if="farmModule.planter != ''" class="p-grey-small">Planted by: {{farmModule.planter}}</p>
        </div>
        <div class="place-self-center">
            <img class="p-4" src="../../assets/img/plants/mint.png">
        </div>
        
        </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from 'vue'
import PositionIndicator from './atoms/PositonIndicator.vue'
import FarmModule from '../../types/Module'

export default defineComponent({
    props: {
    farmModule:{
        type: Object as PropType<FarmModule>,
        required: true
    },
    filds: {
        type: Number,
        default: 4,
       
    },
    boxtype: {
        type: String,
        default: '',
        required: false
    }
    },

 
    components:{PositionIndicator},
})
</script>
