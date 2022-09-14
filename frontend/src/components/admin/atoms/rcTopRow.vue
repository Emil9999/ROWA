<template>
    <div class="w-full flex justify-between items-center">
         <Sheet :isopen="infoOpen" ref="FAQSheet"><InfoQR/></Sheet>
        <div @click="$emit('backClicked')" class="bg-white shadow-md rounded-full h-16 w-16"><ArrowLeftIcon class="h-16 w-16 text-green"></ArrowLeftIcon></div>

        <div class=" h-green-small">
          <div class="h-big-green">{{text}}</div>
        </div>

        <div @click="FAQSheet?.open" class="bg-white shadow-md rounded-full text-green text-center font-bold self-center text-6xl h-16 w-16">?</div>
   
   
   </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import {ArrowLeftIcon } from '@heroicons/vue/solid'
import Sheet from '../../../bottom-sheet/bottom-sheet.vue'
import InfoQR from '../../home/atoms/InfoQR.vue'

export default defineComponent({
    components: {ArrowLeftIcon, Sheet, InfoQR},
    props:{
        text:{
            type:String,
            default: 'Reality Check'

        }
    },
    setup(){
        const FAQSheet = ref<InstanceType<typeof Sheet> | null>(null)
        
        const infoOpen = ref(false)


        return{FAQSheet, infoOpen}
    }
})

</script>