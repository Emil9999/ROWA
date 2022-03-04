<template>
    <div>
        <div class="flex justify-between items-center p-6">
            <ArrowLeftIcon @click="decreaseStep(1)" class="h-16 w-16 p-2 shadow-md rounded-full text-green bg-white"/>
            <StepIndicator :title="'Step'" :step="this.step"/>
            <XIcon @click="this.$router.push('/')" class="h-16 w-16 p-2 shadow-md rounded-full text-green bg-white"/>
        </div>
        
        <div v-if="step==1" class="flex flex-col justify-center items-center">
            <h1 class="h-green-big my-8">Select a Variety to Plant</h1>
            <div v-for="(plant, index) in farmModules" :key="plant">
                <farmingInfoTile @click="testSelection(index)" :farmModule="plant"/>
            </div>
        </div>

        <div v-if="step==2" class="centered-div">
            <farmingInfoTile  :farmModule="farmModules[selectedPlant]" :boxtype="'p'"/>
            <h1 class="h-green-big">Is this your first time?</h1>
            <button @click="instrucVid = true, increaseStep(1)" class="btn-big-green">Watch Instructions</button>
            <div class="info-box-instruc">
                <h2 class="h-green-big">Already know how <br> to harvest?</h2>
                <p class="p-grey-small">By clicking you confirm that you harvested.</p>
                <button @click="increaseStep(1)" class="btn-big-white">Start Planting</button>
            
            </div>
        </div>

        <div v-if="step==3" class="centered-div">
            <farmingInfoTile  :farmModule="farmModules[selectedPlant]" :boxtype="'p'"/>
            <instructionsWindow @buttonPressed="increaseStep(1)"/>
            <button @click="instrucVid = true" class="btn-big-white">Show instructions again</button>
            <div v-if="instrucVid == true" class="video-overlay">
                <videoFrame @VidEnded="updateinstruc"/>
                <button @click="instrucVid = false" class="btn-small-white">Skip Instructions</button>
            </div>
        </div>
        
        <div v-if="step==4" class="centered-div">
            <farmingInfoTile  :farmModule="farmModules[selectedPlant]" :boxtype="'p'"/>   
            <div>Success and aditional Features</div>
            <div>Hungry?</div>
            <button @click="this.$router.push('/')" class="btn-big-no">Go Home</button>
        </div>

    </div> 
</template>



<script lang="ts">
import { defineComponent, ref } from 'vue'
import StepIndicator from '../components/farming/atoms/StepIndicator.vue'
import farmingInfoTile from '../components/farming/farmingInfoTile.vue'
import instructionsWindow from '../components/farming/instructionsWindow.vue'
import videoFrame from '../components/farming/videoFrame.vue'
import { ArrowLeftIcon, XIcon } from '@heroicons/vue/solid'
import FarmModule from '../types/Module'



export default defineComponent({
    name: 'Planting',
    components: {StepIndicator, farmingInfoTile, instructionsWindow, videoFrame, XIcon, ArrowLeftIcon},
    setup(){
        const selectedPlant = ref<number>()
        const listView = ref(false)
        const instrucVid = ref(false)
        const step = ref(1)
        const farmModules = ref<FarmModule[]>([
            {planttype: 'Lolo Bionda', planter: '', modulenumber: 1, position: 0},
             {planttype: 'Basil', planter: 'Emil', modulenumber: 3, position: 3},
              {planttype: 'Pirat', planter: 'Simon', modulenumber: 2, position: 0}
        ])
        

        return{step, farmModules, selectedPlant, listView, instrucVid}
    },
    methods:{
        updateinstruc(val: boolean){
            this.instrucVid = val
            console.log(val)

        },
        testSelection(selection: number){
            this.selectedPlant = selection
            this.step = this.step +1
        },

        increaseStep(stepamount: number){
           this.step = this.step +stepamount
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
