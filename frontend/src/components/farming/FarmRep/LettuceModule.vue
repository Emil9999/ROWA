/<template>
        <div class=" w-72 h-32 flex-col flex">
            <div  :class="[reverseModule,'-mx-'+(7-count)*2+' basis-1/2 inline-flex justify-between items-end']" >
                    <div v-for="i in count" :key="i" :style="'width:'+ 72/count/4+'rem;'" class="h-12 text-center flex items-end justify-center"> <img :width="288/100*(100/count*(0.2+0.15*i))" src="../../../assets/cat_tree/modules/salad.svg"></div>
            </div>
            
            <div class="border-t-4 border-grey">
                <div :class="reverseModule" class="h-full basis-1/4 flex justify-between">
                    <div :class="{'invisible': false}" class="rounded-full flex flex-none justify-center content-start -mb-7 bg-accentwhite-light w-8">
                       <div class="bg-brownred mt-0.5 rounded-full h-7 w-7"><ArrowSmDownIcon class="text-white"/></div> 
                    </div>
                    
                    <div :class="{'invisible': false}" class="rounded-full flex flex-none justify-center content-start -mb-7 bg-accentwhite-light w-8">
                       <div  class="bg-green mt-0.5 rounded-full h-7 w-7"><CheckIcon class="text-white"/></div> 
                    </div>
                    
                </div>
                <div :class="reverseModule" class="flex-none rounded-full basis-1/4 flex items-baseline h-8 justify-between bg-accentwhite-light">
                    <div class="text-center self-center bg-accentwhite ml-0.5 rounded-full h-7 w-7 p-green-small">1</div>
                    <div class="p-small-grey">{{varietyname}}</div>
                    <div class="h-1/12 w-1/12"></div>
                </div>
            </div>
        
        </div>
</template>
/<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import { CheckIcon, ArrowSmDownIcon } from '@heroicons/vue/solid'

export default defineComponent({
    components:{CheckIcon, ArrowSmDownIcon},
    setup(props){
        const count = ref(6)

         const reverseModule = computed(() => ({
             'flex-row-reverse': props.reverse
        }))
        return{count, reverseModule}
    },
    props:{
        varietyname:{
        type: String,
        default: "Lollo Bionda",
    },
          reverse:{
            type:  Boolean,
            default: false,
            required: false,

        },
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