<template>
    <div>
       
        <div class="flex justify-between items-center p-6">
            <ArrowLeftIcon @click="decreaseStep(1)" class="h-16 w-16 p-2 shadow-md rounded-full text-green bg-white"/>
            <StepIndicator :title="stepstext[step+((instrucVid)?-3:0)]" :step="this.step"/>
            <XIcon @click="isFinalView(), this.$router.push('/')" class="h-16 w-16 p-2 shadow-md rounded-full text-green bg-white"/>
        </div>
        
        <div v-if="!farmModules.length" class="centered-div mt-3 h-green-big"> Loading data... </div>

        <div v-if="step==1 && farmModules.length" class="centered-div">
            <SelectorPill :defaultSelection="(farmView)?'Farm': 'List'" :menuPoints="['Farm','List']" @ClickedRow="updateSelector"/>
                <div v-if="farmView" class="centered-div">
                    <h1 class="h-green-big my-8"> {{text[dfarmingType].FarmHeader}}</h1>
                        <div class="flex items-center mb-5 justify-center mt-4 p-grey-small">
                            <div :class="[((dfarmingType=='h')?'bg-green':'bg-brownred')]" class="mx-4 rounded-full h-7 w-7">
                                <component :is="((dfarmingType == 'h')? 'CheckIcon' : 'ArrowSmDownIcon')" class="text-white"></component>
                            </div> 
                            <div>Available to {{text[dfarmingType].Word}}</div> 
                        </div>
                    <Sheet :isopen="false"  ref="detailModule"><DetailModule @SelectedPlant="selectedPlantFromModule" :moduleNum="iselectedModule" :farmModules="selectedModulePlants"/></Sheet>
                    <FarmView :farmingMode="dfarmingType" @ModuleClicked="moduleSelected"/>
                </div>
                <div v-else class="flex flex-col justify-center items-center">
                    <h1 class="h-green-big my-8">{{text[dfarmingType].ListHeader}}</h1>
                    <div class="overflow-auto" style="height: 800px;">
                        <div v-for="plant in filteredList" :key="plant">
                            <farmingInfoTile @click="testSelection(plant)" :farmModule="plant" :boxtype="dfarmingType"/>
                        </div>
                    </div>
                </div>
        </div>

        <div v-if="step==2" class="centered-div">
            <farmingInfoTile  :farmModule="selectedPlant" :boxtype="dfarmingType"/>
            <h1 class="h-green-big">Is this your first time?</h1>
            <button @click="instrucVid = true, increaseStep(1)" class="btn-big-green">Watch Instructions</button>
            <div class="info-box-instruc">
                <h2 class="h-green-big">Already know how <br> to {{text[dfarmingType].Word}}?</h2>
                <p class="p-grey-small">By clicking you confirm that you {{text[dfarmingType].WordParticip}}.</p>
                <button @click="increaseStep(1)" class="btn-big-white">Start {{text[dfarmingType].WordProgressive}}</button>
            
            </div>
        </div>

        <div v-if="step==3" class="centered-div">
            <farmingInfoTile  :farmModule="selectedPlant" :boxtype="dfarmingType"/>
            <instructionsWindow :leafHarvest="selectedPlant.leafHarvest" :infoFarmingType="dfarmingType" :infoPlantType="selectedPlant.group" @buttonPressed="(n) => {leafsHarvested = n, increaseStep(1)}"/>
            <button @click="instrucVid = true" class="btn-big-white">Show instructions again</button>
            <div v-if="instrucVid == true" class="video-overlay">
                <videoFrame @VidEnded="updateinstruc"/>
                <button @click="instrucVid = false" class="btn-small-white">Skip Instructions</button>
            </div>
        </div>
        
        <div v-if="step==4" class="centered-div">
            <farmingInfoTile   :farmModule="selectedPlant" :boxtype="dfarmingType"/> 
            
                <carousel  :autoplay="5000" :wrap-around="true">
                    <slide :index="1">
                        <div class="w-8/12 inline-flex">
                            <img src="../assets/icons/partycone.svg">
                            <div class="h-green-big">You Successfully {{text[dfarmingType].WordParticip}}</div>
                            <img src="../assets/icons/partycone.svg">
                        </div>
                    </slide> 
                    <slide :index="2">
                        <RatingPicker class="w-10/12 my-10" @ratingUpdate="ratingUpdate"/>
                    </slide>
                    <slide :index="((true)?3:0)">
                        <NameSelector @changeName="(n) => newPlanter= n" />
                    </slide>
                    

                    <template #addons>
                    <pagination />
                    </template>
                </carousel>
            
           
            <div class="info-box-instruc mt-20">
                
                <div class="h-green-mid mb-8">{{text[dfarmingType].ActionCall}}</div>
                <ButtonArrow @click="finishFarming()" :bEnabled="false" :link="'/farming/'+((dfarmingType=='h')? 'p' : 'h')" :buttonText="text[dfarmingType].ActionCallButton" :smbutton="true"/>
            </div>
            <button @click="finishFarming(), this.$router.push('/')" class="btn-big-no">Go Home</button>
        </div>

    </div> 
</template>





<script lang="ts">
// If type Herb -> always Leaf harvest until no more possible !Important
import {defineComponent, onMounted, onBeforeMount, ref, computed } from 'vue'
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
import 'vue3-carousel/dist/carousel.css';
import usegetFarmable from '../composables/use_getFarmable'
import { Carousel, Slide, Pagination, Navigation } from 'vue3-carousel';
import {useRoute} from 'vue-router'
import FindFarmableperModule from '../composables/use_FindFarmableperModule'
import FinishFarming from '../composables/use_finishFarming'
import ButtonArrow from '../components/home/atoms/ButtonArrow.vue'



export default defineComponent({
    name: 'Farming_Stepper',
    components: {StepIndicator, Sheet, DetailModule, farmingInfoTile, FarmView,
                 instructionsWindow, videoFrame, RatingPicker, NameSelector, SelectorPill, 
                 Carousel, Slide, Pagination, Navigation ,
                 XIcon, ArrowLeftIcon, CheckIcon, ArrowSmDownIcon, ButtonArrow},
    props: {
        farmingType:{
            type: String,
            default: 'h'
        },
    },
    
    setup(props){

        const route = useRoute()

       
        

        const isFinalView = () => {
            if(step.value == 4){finishFarming()}
        }
        const dfarmingType = ref(props.farmingType)
        const selectedPlant = ref<FarmablePlant>()
        const farmView = ref((dfarmingType.value == 'h')?true:false)
        const listView = ref(false)
        const instrucVid = ref(false)

        //Data for Finish Planting
        const newPlanter = ref("")
        const leafsHarvested = ref(false)
        const iselectedModule = ref(0)
       
        const step = ref(1)

        const {farmModules, loadFarmables} = usegetFarmable(dfarmingType.value)

        const {plantsInModule: selectedModulePlants, hasMultipleFarmable, findforModule} = FindFarmableperModule(farmModules.value)
 
        const rating = ref(0)
    

        const updateSelector = (clicked:string) =>{
            if(clicked == 'Farm'){
                farmView.value = true
            } else {
                farmView.value = false
            }
            
        }

        const filteredList = computed(() => {
            const filteredPlants = ref<FarmablePlant[]>([])
            
            farmModules.value.forEach(e => {
                if (!(filteredPlants.value.find(f => f.planttype == e.planttype))){
                    filteredPlants.value.push(e)
                }
                
            })
            
            console.log(filteredPlants.value)
            return filteredPlants.value
            
        })

        const selectedPlantFromModule = (selectedModule: FarmablePlant) =>{
                 increaseStep(1)
                 selectedPlant.value = selectedModule
                 //TODO save plant as selected Plant  
        }
        
        const moduleSelected = (selectedModule:number) =>{
            iselectedModule.value = selectedModule
            findforModule(selectedModule)
            if(hasMultipleFarmable.value && dfarmingType.value == 'h'){
                openSheet()
             } else {
                 if (farmModules.value.find((e) => e.modulenumber == selectedModule)){
                    selectedPlant.value = farmModules.value.find((e) => e.modulenumber == selectedModule)
                    increaseStep(1)
                 }
                
             }
        }

        const finishFarming = () =>{
            const result = FinishFarming(dfarmingType.value,{planttype: selectedPlant.value?.planttype ?? 'empty', modulenumber: selectedPlant.value?.modulenumber || 0, position: selectedPlant.value?.position || 0,
             group: selectedPlant.value?.group ?? '', leafHarvest: leafsHarvested.value, planter: newPlanter.value})
            console.log(result)
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
            FarmHeader: 'Select a Module to Harvest from',
            ActionCall: 'Ensure a weekly harvest',
            ActionCallButton: 'Plant'

        },
        p:{
            Word: 'plant',
            WordParticip: 'planted',
            WordProgressive: 'planting',
            ListHeader: 'Select a variety to Plant',
            FarmHeader: 'Select a Module to Plant in',
            ActionCall: 'Are You Hungry?',
            ActionCallButton: 'Harvest'

        }

        }

         onBeforeMount(() => {
            if(route.name == 'directFarming'){
                dfarmingType.value = String(route.params.farmingType)
                selectedPlant.value = {planttype: String(route.params.planttype),leafHarvest: Boolean(route.params.leafHarvest), group: String(route.params.group), planter: String(route.params.planter), modulenumber: Number(route.params.modulenumber), position: Number(route.params.position)}
                
                step.value = 2
            }
        })
        
        onMounted(() => {
            loadFarmables()
        })
        


        return{step, newPlanter, iselectedModule, filteredList, finishFarming,isFinalView, selectedPlantFromModule, leafsHarvested, selectedModulePlants, increaseStep, moduleSelected, openSheet, detailModule, farmModules, stepstext,selectedPlant, listView, instrucVid, rating, farmView, dfarmingType, text, updateSelector}
    },
    methods:{

        updateinstruc(val: boolean){
            this.instrucVid = val

        },
        ratingUpdate(rating: number){
            this.rating = rating
        },
        testSelection(selection: FarmablePlant){
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
