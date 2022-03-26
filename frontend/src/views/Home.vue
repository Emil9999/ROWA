<template>
<div class=" max-w-full">

   <TopRowHome  :menuPoints="['Farm','Home','Stats']" :defaultSelection="selectedPage" @ClickedRow="changePage" class="pt-10 px-5"></TopRowHome>
  <div v-if="selectedPage == 'Home'" class="centered-div">
   <h1 class="h-green-big">Explore Ideas</h1>
  <div class="flex w-11/12 my-10 overflow-auto flex-nowrap">
    <IdeasCards v-for="idea in ideas" :key="idea" :title="idea"></IdeasCards>
   
  </div>
 
 <h1 class="h-green-big">What do you want to do?</h1>
  
    <button class="btn-big-green inline-flex justify-between" @click="this.$router.push('/farming/p')">
      <span class="w-8"></span>
      <span class="">Planting</span>
      <ArrowRightIcon v-if="bPlanting" class="text-white h-8 w-8"></ArrowRightIcon>
      </button>
  
    <button class="btn-big-green" @click="this.$router.push('/farming/h')">Harvesting</button>
 
  </div>
  <div v-if="selectedPage == 'Farm'" class="centered-div">
  <FarmRepInfo/>
  <Sheet :isopen="isOpen"  ref="moduleSheet"><ModuleInfo :moduleNumber="ModuleInfo"/></Sheet>
  <FarmRepresentation @ModuleClicked="(clickedModule)"  class="m-8"></FarmRepresentation>
 </div>
 </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import IdeasCards from  '../components/home/IdeasCards.vue'
import TopRowHome from  '../components/home/TopRow.vue'
import FarmRepInfo from  '../components/home/atoms/FarmRepInfo.vue'
import { ArrowRightIcon } from '@heroicons/vue/solid' 
import FarmRepresentation from '../components/farming/FarmRep/FarmRepresentation.vue'
import Sheet from '../bottom-sheet/bottom-sheet.vue'
import ModuleInfo from '../components/home/ModuleInfo.vue'



export default defineComponent({
  name: 'Home',
  components: {IdeasCards, TopRowHome, ModuleInfo, ArrowRightIcon, FarmRepresentation, FarmRepInfo, Sheet},
  setup() {
    const selectedPage = ref('Home')
    const bPlanting = ref(true)
    
    
    
    
    
    const moduleSheet = ref<InstanceType<typeof Sheet> | null>(null)
    const openSheet = () => {
      moduleSheet.value?.open()
    }

    const ModuleInfo = ref(0)
    const clickedModule = (event:number) => {
      ModuleInfo.value = event
      openSheet()
    }
    
    const bHarvest = ref(true)
    const isOpen = ref(false)
    const ideas = ref(['Garnish','Salad','Smoothie','Tea'])
    return {selectedPage, bHarvest, ModuleInfo, bPlanting,isOpen, ideas,moduleSheet, openSheet, clickedModule}
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
