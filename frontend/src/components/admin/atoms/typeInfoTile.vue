<template>
    <div class="bg-white items-center text-left shadow-md my-10 rounded-xl grid grid-cols-2 w-96">
        <div class="px-2">
          
            <h1 class="h-green-small">{{ptype.name }}</h1>
            <div v-if="selected" class="p-grey-big">Currently planted</div>
        </div>
        <div class="place-self-center">            
                <img class="p-4" :src="require('../../../assets/img/plants/'+checkedString)">
        </div>
        
        </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from 'vue'
import pType from '../../../types/VarietyData'

import {checkImage} from '../../../composables/use_imgChecker'

export default defineComponent({
    components:{},
    props: {
        ptype:{
        type: Object as PropType<pType>,
        required: true
    },
        selected:{
            type: Boolean,
            default: false
        }

    },
    setup(props){

        
        const {cImage} = checkImage("png")
        const checkedString = cImage(props.ptype.name.toLowerCase())

        return{checkedString}
    },


 
    
})
</script>
