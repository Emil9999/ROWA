<template>
    <div class="bg-grey text-grey-dark relative rounded w-30 flex justify-evenly">
        <div class="flex items-center" v-for="n in filds" :key="n">
            <div class="z-10 pr-2"><p  v-if="n == pos" class="text-center text-accentwhite px-2">{{n}}</p> 
            <p  v-else class="text-center px-2">{{n}}</p> </div>

            <div v-if="n  == pos" :class="[type == 'h' ? 'bg-green' : 'bg-brownred']" class="w-4 h-4 m-1 rounded-full z-0 absolute"></div>
            <div v-if="n  != filds" class="w-1 h-7 bg-grey-dark"></div>
            <div v-if="n  != filds" class="w-1 h-7 bg-grey"></div>


        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
      filds: Number,
      pos: Number,
      type: String

  }
})
</script>
