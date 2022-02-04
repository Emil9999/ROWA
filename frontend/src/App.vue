<template>
<div class="bg-accentwhite-light">
  
  
  <FarmingInfoTile :planter="'Hannes'" :plantType="'Mint'" :position="3" :filds="4" :pos="3"></FarmingInfoTile>
  <TopRowHome></TopRowHome>
  <h1 class="h-green-big">Headline</h1>
  <h1 class="h-green-small">Headline</h1>
  <h1 class="h-white-small">Headline</h1>
  <p class="p-grey-small">Text Body</p>
  <p class="p-grey-big">Text Body</p>
  <IdeasCards :title="'Smoothie'"></IdeasCards>
  <IdeasCards :title="'Smoothie'"></IdeasCards>
  <IdeasCards :title="'Smoothie'"></IdeasCards>
  <button class="btn-big-green" @click="changeName('Hannes')">Click Me pls!</button>
  <button class="btn-big-white" @click="changeName('Hannes')">Click Me pls!</button>
  <button class="btn-small-green" @click="changeName('Hannes')">Click Me pls!</button>

</div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import FarmingInfoTile from  './components/farming/FarmingInfoTile.vue'
import IdeasCards from  './components/home/IdeasCards.vue'
import TopRowHome from  './components/home/TopRow.vue'

export default defineComponent({
  name: 'App',
  components: {FarmingInfoTile, IdeasCards, TopRowHome},
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
