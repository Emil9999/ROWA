<template>
    <div class="w-full flex justify-between items-center">
        <div class="h-20 w-20 "><img src="../../assets/icons/Cog.svg" alt="Ideas"></div>

        <div class="bg-white h-green-small flex rounded-xl justify-around">
            <div class="bg-white px-4 py-1 m-2 rounded-lg ">Farm</div>
            <div class="bg-green px-4 py-1 m-2 rounded-lg shadow-md text-white">Home</div>
            <div class="bg-white px-4 py-1 m-2 rounded-lg ">Stats</div>
        </div>
    
        <div class="h-20 w-20 "><img src="../../assets/icons/Questionmark.svg" alt="Ideas"></div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    props: {
        title: {
        type: String,
        required: true
    },
    },
       methods: {
    
   }
   
})
</script>