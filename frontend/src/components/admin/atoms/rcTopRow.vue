<template>
    <div class="w-full flex justify-between items-center">
        <div @click="$emit('backClicked')" class="bg-white shadow-md rounded-full h-16 w-16"><ArrowLeftIcon class="h-16 w-16 text-green"></ArrowLeftIcon></div>

        <div class="bg-white h-green-small flex rounded-xl justify-around">
          <div class="h-big-green">Reality Check</div>
        </div>

        <div @click="$emit('faqOpen')" class="bg-white shadow-md rounded-full text-green text-center font-bold self-center text-6xl h-16 w-16">?</div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import {ArrowLeftIcon } from '@heroicons/vue/solid'

export default defineComponent({
    components: {ArrowLeftIcon},

})

</script>