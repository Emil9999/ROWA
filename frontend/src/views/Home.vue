<template>
<div class=" max-w-full">
  
 
   <TopRowHome :menuPoints="['Farm','Home','Stats']" :defaultSelection="selectedPage" @ClickedRow="changePage" class="pt-8"></TopRowHome>
  <div v-if="selectedPage == 'Home'" class="centered-div">
   <h1 class="h-green-big">Explore Ideas</h1>
  <div class="flex w-11/12 my-10 overflow-auto flex-nowrap">
    <IdeasCards :title="'Smoothie'"></IdeasCards>
    <IdeasCards :title="'Smoothie'"></IdeasCards>
    <IdeasCards :title="'Smoothie'"></IdeasCards>
    <IdeasCards :title="'Smoothie'"></IdeasCards>
    <IdeasCards :title="'Smoothie'"></IdeasCards>
  </div>
 
 <h1 class="h-green-big">What do you want to do?</h1>
  
    <button class="btn-big-green inline-flex justify-between" @click="this.$router.push('Planting')">
      <span class="w-8"></span>
      <span class="">Planting</span>
      <ArrowRightIcon v-if="bPlanting" class="text-white h-8 w-8"></ArrowRightIcon>
      </button>
  
    <button class="btn-big-green" @click="changeState()">Harvesting</button>
 
 
  <Module class="m-8"></Module>
 
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import IdeasCards from  '../components/home/IdeasCards.vue'
import TopRowHome from  '../components/home/TopRow.vue'
import { ArrowRightIcon } from '@heroicons/vue/solid'

import Module from '../components/farming/FarmRep/Module.vue'

export default defineComponent({
  name: 'Home',
  components: {IdeasCards, TopRowHome, ArrowRightIcon, Module},
  setup() {
    const selectedPage = ref('Home')
    const bPlanting = ref(true)
    const bHarvest = ref(true)
    
    return {selectedPage, bHarvest, bPlanting}
  },
  methods: {
    changeState(){
      this.bHarvest =  !this.bHarvest
    },
    changePage(clickedRow:string){
        this.selectedPage = clickedRow
    }
  }
});
</script>

<style>
</style>
