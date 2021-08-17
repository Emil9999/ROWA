<template>
<div>
  <button class="btn-Susgreen" @click="changeName('Hannes')">Click Me!</button>
<div class="p-6 max-w-sm mx-auto bg-white rounded-xl shadow-md flex items-center space-x-4">
  <div class="flex-shrink-0">
    <img class="h-12 w-12" src="" alt="ChitChatLogo">
  </div>
  <div>
  <div class="text-xl font-medium text-black">{{name}}</div>
  <p class="text-grey-500">You have a new message</p>
  </div>
  
</div>
</div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';

export default defineComponent({
  name: 'App',
  components: {},
  setup() {
    const state = reactive({
      name: 'Simon',
      age: 25 as number | string
    })

    state.name = 'Emil'

    return {...toRefs(state)}
  },
  methods: {
    changeName(name: string) {
      this.name = name

    }
  }
});
</script>

<style>
</style>
