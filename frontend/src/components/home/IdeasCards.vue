<template>
<div class="bg-green flex-none w-48 h-96 m-4 relative rounded-lg shadow-lg overflow-hidden">
   
   <div class="w-full top-10 h-white-big absolute text-center">{{title}}</div>
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
