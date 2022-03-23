<template>
    <div>
       
        <div class="flex justify-between items-center p-6">
            <ArrowLeftIcon @click="decreaseStep(1)" class="h-16 w-16 p-2 shadow-md rounded-full text-green bg-white"/>
            <StepIndicator :title="stepstext[step+((instrucVid)?-3:0)]" :step="this.step"/>
            <XIcon @click="this.$router.push('/')" class="h-16 w-16 p-2 shadow-md rounded-full text-green bg-white"/>
        </div>
        
        <div v-if="step==1" class="centered-div">
            <SelectorPill :defaultSelection="(farmView)?'Farm': 'List'" :menuPoints="['Farm','List']" @ClickedRow="updateSelector"/>
                <div v-if="farmView">
                    <h1 class="h-green-big my-8"> {{text[farmingType].FarmHeader}}</h1>
                    <div class="flex items-center mb-5 justify-center mt-4 p-grey-small">
                        <div class="bg-brownred mx-4 rounded-full h-7 w-7"><ArrowSmDownIcon class="text-white"/></div> 
                        <div>Available to {{text[farmingType].Word}}</div> 
                    </div>
                    <Sheet :isopen="false"  ref="detailModule"><DetailModule @SelectedPlant="selectedPlantFromModule" :farmModules="selectedModulePlants"/></Sheet>
                    <FarmView @ModuleClicked="(moduleSelected)"/>
                </div>
                <div v-else class="flex flex-col justify-center items-center">
                    <h1 class="h-green-big my-8">{{text[farmingType].ListHeader}}</h1>
                    <div v-for="(plant, index) in farmModules" :key="plant">
                        <farmingInfoTile @click="testSelection(index)" :farmModule="plant"/>
                    </div>
                </div>
        </div>

        <div v-if="step==2" class="centered-div">
            <farmingInfoTile  :farmModule="farmModules[selectedPlant]" :boxtype="farmingType"/>
            <h1 class="h-green-big">Is this your first time?</h1>
            <button @click="instrucVid = true, increaseStep(1)" class="btn-big-green">Watch Instructions</button>
            <div class="info-box-instruc">
                <h2 class="h-green-big">Already know how <br> to {{text[farmingType].Word}}?</h2>
                <p class="p-grey-small">By clicking you confirm that you {{text[farmingType].WordParticip}}.</p>
                <button @click="increaseStep(1)" class="btn-big-white">Start {{text[farmingType].WordProgressive}}</button>
            
            </div>
        </div>

        <div v-if="step==3" class="centered-div">
            <farmingInfoTile  :farmModule="farmModules[selectedPlant]" :boxtype="farmingType"/>
            <instructionsWindow :leafHarvest="leafHarvest" :infoTypeprop="farmingType+'_salad'" @buttonPressed="increaseStep(1)"/>
            <button @click="instrucVid = true" class="btn-big-white">Show instructions again</button>
            <div v-if="instrucVid == true" class="video-overlay">
                <videoFrame @VidEnded="updateinstruc"/>
                <button @click="instrucVid = false" class="btn-small-white">Skip Instructions</button>
            </div>
        </div>
        
        <div v-if="step==4" class="centered-div">
            <farmingInfoTile  :farmModule="farmModules[selectedPlant]" :boxtype="farmingType"/>   
            <RatingPicker class="w-10/12 my-10" @ratingUpdate="ratingUpdate"/>
            <NameSelector />
            <div>Plant?</div>
            <button @click="this.$router.push('/')" class="btn-big-no">Go Home</button>

        <CheckIcon/>
        </div>

    </div> 
</template>



<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import StepIndicator from '../components/farming/atoms/StepIndicator.vue'
import farmingInfoTile from '../components/farming/FarmingInfoTile.vue'
import instructionsWindow from '../components/farming/instructionsWindow.vue'
import videoFrame from '../components/farming/videoFrame.vue'
import { ArrowLeftIcon, XIcon } from '@heroicons/vue/solid'
import FarmablePlant from '../types/FarmablePlant'
import RatingPicker from '../components/farming/atoms/RatingPicker.vue'
import NameSelector from '../components/farming/atoms/NameSelector.vue'
import FarmView from '../components/farming/FarmRep/FarmRepresentation.vue'
import SelectorPill from '../components/general/SelectorPill.vue'
import { CheckIcon, ArrowSmDownIcon } from '@heroicons/vue/solid'
import DetailModule from '../components/farming/FarmRep/DetailModule.vue'
import Sheet from '../bottom-sheet/bottom-sheet.vue'



export default defineComponent({
    name: 'Farming_Stepper',
    components: {StepIndicator, Sheet, DetailModule, farmingInfoTile, FarmView, instructionsWindow, videoFrame, RatingPicker, NameSelector, SelectorPill, XIcon, ArrowLeftIcon, CheckIcon, ArrowSmDownIcon},
    props: {
        farmingType:{
            type: String,
            required: true
        }
    },
    setup(props){
        const dfarmingType = ref(props.farmingType)
        const selectedPlant = ref<number>()
        const farmView = ref((props.farmingType == 'h')?true:false)
        const listView = ref(false)
        const instrucVid = ref(false)
        const selectedModulePlants = ref<FarmablePlant[]>([])
        const step = ref(1)
        const leafHarvest = computed(() => (dfarmingType.value == 'h') ? true : null )
        const rating = ref(0)

        const isMultiplantModule = (selectedmodule: number) => {
                let  plantsInModule: Array<FarmablePlant> = []
                for(const farmModule of farmModules.value){
                    if(farmModule.modulenumber == selectedmodule){
                        plantsInModule.push(farmModule)
                    }
                }
                    if (plantsInModule.length > 1) { 
                        selectedModulePlants.value = plantsInModule.sort((a,b) => (a.position < b.position)? 1 : -1)
                        return true
                    } else {return false}
                    
        }

        const updateSelector = (clicked:string) =>{
            if(clicked == 'Farm'){
                farmView.value = true
            } else {
                farmView.value = false
            }
            
        }

        const selectedPlantFromModule = (selectedModule: FarmablePlant) =>{
                 increaseStep(1)
                 //TODO save plant as selected Plant  
        }
        
        const moduleSelected = (selectedModule:number) =>{
            if(isMultiplantModule(selectedModule) && dfarmingType.value == 'h'){
                openSheet()
             } else {
                 selectedPlant.value = selectedModule-1
                 increaseStep(1)
             }
        }

        const increaseStep = (increaseAmount:number) => (step.value = step.value+increaseAmount)

        const detailModule = ref<InstanceType<typeof Sheet> | null>(null)
        const openSheet = () => {
        detailModule.value?.open()
        }


        //'Debug' Data 
        const stepstext = ["Video","Select a plant","Choose","Instructions","Success"]
        const text = {h:{
            Word: 'harvest',
            WordParticip: 'harvested',
            WordProgressive: 'harvesting',
            ListHeader: 'Select a variety to Harvest',
            FarmHeader: 'Select a Module to Harvest from'

        },
        p:{
            Word: 'plant',
            WordParticip: 'planted',
            WordProgressive: 'planting',
            ListHeader: 'Select a variety to Plant',
            FarmHeader: 'Select a Module to Plant in'

        }

        }
        const farmModules = ref<FarmablePlant[]>([
            {planttype: 'Lollo Bionda', planter: '', modulenumber: 1, position: 0},
             {planttype: 'Basil', planter: 'Emil', modulenumber: 3, position: 3},
              {planttype: 'Pirat', planter: 'Simon', modulenumber: 4, position: 0},
               {planttype: 'Mint', planter: 'Emil', modulenumber: 3, position: 4},

                {planttype: 'Basil', planter: 'Emil', modulenumber: 3, position: 4},
                 {planttype: 'Thai Basil', planter: 'Emil', modulenumber: 3, position: 4},
                 
              {planttype: 'Pirat', planter: 'Simon', modulenumber: 2, position: 0}
        ])
        

        return{step, selectedPlantFromModule, selectedModulePlants, isMultiplantModule, increaseStep, moduleSelected, openSheet, detailModule, farmModules, stepstext,selectedPlant, listView, instrucVid, rating, farmView, dfarmingType, text, updateSelector, leafHarvest}
    },
    methods:{

        updateinstruc(val: boolean){
            this.instrucVid = val

        },
        ratingUpdate(rating: number){
            this.rating = rating
        },
        testSelection(selection: number){
            this.selectedPlant = selection
            this.step = this.step +1
        },

        decreaseStep(stepamount: number){
            this.instrucVid = false
            if(this.step-1 > 0){
           this.step = this.step -stepamount} else {
               this.$router.push('/')
           }
        }
    }
    
})
</script>
