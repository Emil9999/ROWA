<template>
    <div class="bg-white items-center text-left shadow-md my-10 rounded-xl grid grid-cols-2 w-96">
        <div class="px-2">
            <p v-if="boxtype == 'h'" class="p-grey-small py-2">You are harvesting:</p>
             <p v-if="boxtype == 'p'" class="p-grey-small py-2">You are planting:</p>
            <h1 class="h-green-small">{{farmModule.planttype}}</h1>
            <p class="p-grey-small">Module {{farmModule.modulenumber}}</p>
            <div v-if="farmModule.position != 0 && farmModule.position <= filds"><PositionIndicator :type="boxtype" :filds="filds" :pos="farmModule.position"></PositionIndicator></div>
            <p v-if="farmModule.planter != '' && boxtype != 'p'" class="p-grey-small">Planted by: {{farmModule.planter}}</p>
        </div>
        <div class="place-self-center">            
                <img class="p-4" :src="require('../../assets/img/plants/'+checkedString)">
        </div>
        
        </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from 'vue'
import PositionIndicator from './atoms/PositonIndicator.vue'
import FarmablePlant from '../../types/FarmablePlant'
import {checkImage} from '../../composables/use_imgChecker'

export default defineComponent({
    components:{PositionIndicator},
    props: {
    farmModule:{
        type: Object as PropType<FarmablePlant>,
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
    setup(props){

        const {cImage} = checkImage("png")
        const checkedString = cImage(props.farmModule.planttype.toLowerCase())

        return{checkedString}
    },


 
    
})
</script>
