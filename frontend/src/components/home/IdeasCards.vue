<template>
<div class="bg-green flex-none w-44 h-80 m-4 relative rounded-2xl drop-shadow-md overflow-hidden ">
   
   <div class="w-full z-20 top-10 h-white-big absolute text-center">{{title}}</div>
   <div class="bg-gradient-to-b from-gradientCol opacity-50 w-full h-1/2 absolute top-0"></div>
    <img class="object-cover w-full h-full"  :src="require('../../assets/img/ideas/'+title.toLowerCase()+'.jpg')" alt="Ideas">
  

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
