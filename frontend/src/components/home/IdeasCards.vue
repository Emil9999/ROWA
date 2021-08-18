<template>
<div class="rounded-xl bg-green">
   
    <img class=""  src="../../assets/img/ideas/smoothie.jpg" alt="Ideas">
     <h1 class="text-white text-6xl font-extrabold">Text</h1>

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
     getImg(img) {
       return require(`../assets/${img}`);
     }
   }
   
})
</script>
