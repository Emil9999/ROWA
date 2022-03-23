<template>
<div>
    <h1 v-if="rating == 0" class="h-green-big mb-10">How was your farming experience?</h1>
    <h1 v-if="rating != 0" class="h-green-big mb-10">Thanks for the Feedback</h1>
<div v-if="rating == 0" class="flex w-full justify-around">
    <div @click="rating = index+1, RatingUpdate()"  v-for="(emoji, index) in emojis" :key="index" class="bg-green text-6xl items-center rounded-full w-20 h-20">
     <p class="my-3">  {{emoji}} </p>
    </div>
</div>
<div v-if="rating != 0">
    <button @click="rating = 0, RatingUpdate()" class="btn-big-white">Undo</button>
</div>
</div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
    setup() {
        const emojis = ref([
    "🙁", "😕", "🙂", "😃", "😍",
        
  ])
        
        const rating = ref(0)
  return{emojis, rating}
    },
    emits: ["ratingUpdate"],
    methods: {
    RatingUpdate() {
      this.$emit("ratingUpdate", this.rating);
    },
  }
})
</script>
