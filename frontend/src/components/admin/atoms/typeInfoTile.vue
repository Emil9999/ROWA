<template>
    <div :class="[greyedOut ? 'bg-grey h-grey-small grayscale filter' : 'bg-white h-green-small']" class=" overflow-hidden relative items-center text-left shadow-md my-10 rounded-xl max-h-44 grid grid-cols-2 w-96">
        <div class="px-2">
          
            <h1 class="">{{ptype}}</h1>
            <div v-if="previousS" class="p-grey-big">Currently planted</div>
        </div>
        <div class="place-self-center">            
                <img class="p-4" :src="require('../../../assets/img/plants/'+checkedString)">
        </div>
        <div v-if="selected" class="bg-green h-full w-4 right-0 absolute"></div>
        </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'

import {checkImage} from '../../../composables/use_imgChecker'

export default defineComponent({
    components:{},
    props: {
        ptype:{
        type: String,
        required: true
    },
        selected:{
            type: Boolean,
            default: false
        },
        previousS:{
            type: Boolean,
            default: false
        },
        greyedOut:{
            type: Boolean,
            default: false
        }
    

    },
    setup(props){

        
        const {cImage} = checkImage("png")
        const checkedString = cImage(props.ptype.toLowerCase())

        return{checkedString}
    },


 
    
})
</script>
