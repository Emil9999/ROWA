<template>
<div class="bg-green w-72 h-135 m-10 relative rounded-lg shadow-lg overflow-hidden">
   
   <div class="w-full top-3 h-white-big absolute text-center">{{title}}</div>
    <img class="object-cover w-full h-full"  src="../../assets/img/ideas/smoothie.jpg" alt="Ideas">
  

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
