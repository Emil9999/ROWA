<template>
    <div class="flex items-center justify-center max-h-10 w-1/2 mb-12">
  
  <label class="flex items-center cursor-pointer">
    <!-- toggle -->
    <div class="relative">
      <!-- input -->
      <input type="checkbox" @click="$emit('stateChange', toggleState)" v-model="toggleState" class="sr-only">
      <!-- line -->
      <div :class="[toggleState ? 'bg-green': 'bg-grey']" class="block  w-14 h-8 rounded-full transition"></div>
      <!-- dot -->
      <div :class="{'left-7': toggleState}" class="absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition"></div>
    </div>
  </label>

</div>

</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'


export default defineComponent({
    props:{

    },
    setup() {
        const toggleState = ref(false)
        return{toggleState}
    },
})
</script>
