<template>
    <div class="bg-grey text-grey-dark rounded flex w-28 place-self-center ">
    <div class="flex" v-for="n in filds" :key="n">
        <p class="z-10 text-center px-2">{{n}}</p> 
        <div v-if="n  == pos" class="w-4 h-4 m-1 rounded-full bg-green fixed z-0"></div>
        <div v-if="n  != filds" class="w-1 h-7 bg-grey-dark">
        </div>

    </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
      filds: Number,
      pos: Number,

  }
})
</script>
