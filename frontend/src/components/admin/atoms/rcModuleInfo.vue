<template>
    <div class="flex w-135 justify-between m-auto items-center p-5 flex-grow rounded-md bg-almostwhite h-green-small">
        <div>{{text}}:</div>
        <div>Module:{{mNumber}}</div>
        <div class="bg-white">
            <img class=" max-h-20"  :src="require('../../../assets/img/plants/'+cImage(text))"/>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import {checkImage} from '../../../composables/use_imgChecker'

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
    
    setup(){

        
        const {cImage} = checkImage("png")

        return {cImage}
    },
})
</script>
