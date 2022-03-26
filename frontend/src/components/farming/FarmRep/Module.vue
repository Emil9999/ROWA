/<template>
        <div :class="{'opacity-50': (harvestable.length + plantable.length == 0)}" class=" w-72 h-32 flex-col flex">
            <div  :class="[reverseModule,'-mx-'+(7-count)*2+' basis-1/2 inline-flex justify-between items-end']" >
                    <div v-for="plant in plantsInModule" :key="plant" :style="'width:'+ 72/count/4+'rem;'" class="h-12 text-center mx-auto flex items-end justify-center"> <img :width="80*(plant.growthTime*0.03)" :src="require('../../../assets/img/plant_svg/'+cImage(plant.variety))"></div>
            </div>
            
            <div class="border-t-4 border-grey">

                <div :class="reverseModule"  class="h-full basis-1/4 flex justify-between">
                    <div :class="{'invisible': !(harvestable.includes(index) || plantable.includes(index))}" v-for="(i, index) in count" :key="i" class="rounded-full flex flex-none justify-center content-start -mb-7 bg-accentwhite-light w-8">
                       <div v-if="harvestable.includes(index)" class="bg-green mt-0.5 rounded-full h-7 w-7"><CheckIcon class="text-white"/></div> 
                       <div v-if="plantable.includes(index)" class="bg-brownred mt-0.5 rounded-full h-7 w-7"><ArrowSmDownIcon class="text-white"/></div> 
                    </div>
                </div>

    


                <div :class="reverseModule" class="flex-none rounded-full basis-1/4 flex items-baseline h-8 justify-between bg-accentwhite-light">
                    <div class="text-center self-center bg-accentwhite ml-0.5 rounded-full h-7 w-7 p-green-small">{{moduleNumber}}</div>
                    <div class="p-small-grey">{{moduleText}}</div>
                    <div class="h-1/12 w-1/12"></div>
                </div>
            </div>
        
        </div>
</template>
/<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { CheckIcon, ArrowSmDownIcon } from '@heroicons/vue/solid'
import getPlantsInModule from '../../../composables/use_getPlantInModule'
import {checkImage} from '../../../composables/use_imgChecker'

export default defineComponent({
    components:{CheckIcon, ArrowSmDownIcon},
    setup(props){
       
        // Array plantable harvestable pos
      

        const {modulePlants: plantsInModule, error, loadModulePlants} = getPlantsInModule(props.moduleNumber)
        
        const {cImage} = checkImage("svg")
        const harvestable = ref(props.harvestableSpots)
        const plantable = ref(props.plantableSpots)
        const reverseModule = computed(() => ({
             'flex-row-reverse': props.reverse
        }))
        loadModulePlants()
        const count = computed(() => plantsInModule.value.length)

        const moduleText = computed(() => {
            if (harvestable.value.length + plantable.value.length == 0){
                if(plantsInModule.value.length != 0){
                    return 'Unavailable'
                } else { return 'Empty'}
            } else { return 'Herb'}
        })
        

        return{count, cImage, moduleText, plantsInModule, reverseModule, harvestable, plantable}
    },
    props:{
        moduleNumber:{
        type: Number,
        default: 1,
    },
        reverse:{
            type:  Boolean,
            default: false,
            required: false,

        },
        plantableSpots:{
            type: Array,
            default: () => []
        },
        harvestableSpots:{
            type: Array,
            default: () => []
        }
    },
    methods:{
    }
    
})
</script>
<style>
 text{
     font-family: Mulish;
    font-style: normal;
    font-weight: 600;
    font-size: 36px;

    /* UI COLORS/Grey Dark */

    fill: #828282;

 }
</style>