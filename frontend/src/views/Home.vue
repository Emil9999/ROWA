<template>
<div class=" max-w-full">

   <TopRowHome  :menuPoints="['Farm','Home','Stats']" :defaultSelection="selectedPage" @ClickedRow="changePage" class="pt-10 px-5"></TopRowHome>
  <div v-if="selectedPage == 'Home'" class="centered-div">
   <h1 class="h-green-big mt-24">Explore Ideas</h1>
  <div class="flex w-11/12 my-10 overflow-auto flex-nowrap">
    <IdeasCards v-for="idea in ideas" :key="idea" :title="idea"></IdeasCards>
   
  </div>
 
 <h1 class="h-green-big mt-20 mb-12">What do you want to do?</h1>

  <ButtonArrow :buttonText="'Harvest'" :bEnabled="!bHarvesting" :link="'/farming/h'"/>
  <div class="p-grey-small" v-if="!bHarvesting"> Everything has been planted. <br> Please come back  in in 5 days.</div>

  <ButtonArrow :buttonText="'Plant'" :bEnabled="!bPlanting" :link="'/farming/p'"/>
   <div class="p-grey-small" v-if="!bPlanting"> Everything has been planted. <br> Please come back  in in 5 days.</div>

  
    
  </div>
  <div v-if="selectedPage == 'Farm'" class="centered-div">
  <FarmRepInfo/>
  <Sheet :isopen="isOpen"  ref="moduleSheet"><ModuleInfo :moduleNumber="ModuleInfo" :moduleHarvestable="harvestableInModule" :modulePlantable="plantableInModule"/></Sheet>
  <FarmRepresentation @ModuleClicked="(clickedModule)"  class="m-8"></FarmRepresentation>
 </div>
 </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue';
import IdeasCards from  '../components/home/IdeasCards.vue'
import TopRowHome from  '../components/home/TopRow.vue'
import FarmRepInfo from  '../components/home/atoms/FarmRepInfo.vue'
import ButtonArrow from  '../components/home/atoms/ButtonArrow.vue'
import FarmRepresentation from '../components/farming/FarmRep/FarmRepresentation.vue'
import Sheet from '../bottom-sheet/bottom-sheet.vue'
import ModuleInfo from '../components/home/ModuleInfo.vue'
import usegetFarmable from '../composables/use_getFarmable'
import useFindFarmableperModule from '../composables/use_FindFarmableperModule'



export default defineComponent({
  name: 'Home',
  components: {IdeasCards, TopRowHome, ModuleInfo, ButtonArrow, FarmRepresentation, FarmRepInfo, Sheet},
  setup() {
    const selectedPage = ref('Home')
    
    
    
    const {farmModules: harvestables,loadFarmables: loadharvestables} = usegetFarmable('h')
    const {farmModules: plantables,loadFarmables: loadplantables} = usegetFarmable('p')
    loadharvestables()
    loadplantables()
    const {plantsInModule: harvestableInModule,findforModule: findHarvestable} = useFindFarmableperModule(harvestables.value)
    const {plantsInModule: plantableInModule,findforModule: findPlantable} = useFindFarmableperModule(plantables.value)
    
    const moduleSheet = ref<InstanceType<typeof Sheet> | null>(null)
    const openSheet = () => {
      moduleSheet.value?.open()
    }

    const bPlanting = computed(() => plantables.value.length >= 1)
    const bHarvesting = computed(() => harvestables.value.length >= 1)

    const ModuleInfo = ref(0)
    const clickedModule = (event:number) => {
      ModuleInfo.value = event
      findHarvestable(event)
      findPlantable(event)
      openSheet()
    }

    const isOpen = ref(false)
    const ideas = ref(['Garnish','Salad','Smoothie','Tea'])
    return {selectedPage, bHarvesting, ModuleInfo, bPlanting,isOpen,plantableInModule, harvestableInModule, ideas,moduleSheet, openSheet, clickedModule}
  },
  methods: {
    
   
    changePage(clickedRow:string){
        this.selectedPage = clickedRow
    }
  }
});
</script>

<style>
</style>
