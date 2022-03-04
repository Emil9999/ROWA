/<template>
        <div class=" w-80 h-32 flex-col flex flex-none bg-brown">
            <div class=" basis-1/2 flex justify-between items-baseline" >
               <div  v-for="i in count" :key="i" class="ml-8"> <img :width="i*15" src="../../../assets/cat_tree/modules/salad.svg"> </div>
            </div>
            
            <div class=" border-t-4 border-grey">
                <div class="h-full basis-1/4 flex justify-between">
                    <div v-for="(i, index) in count" :key="i" class="rounded-full flex flex-none justify-center content-start -mb-7 bg-accentwhite-light w-8">
                       <div v-if="index == 1" class="bg-green mt-0.5 rounded-full h-7 w-7"><CheckIcon class="text-white"/></div> 
                       <div v-if="index != 1" class="bg-brownred mt-0.5 rounded-full h-7 w-7"><ArrowSmDownIcon class="text-white"/></div> 
                    </div>
                </div>
                <div class="flex-none rounded-full basis-1/4 flex items-baseline h-8 justify-between bg-accentwhite-light">
                    <div class="text-center self-center bg-accentwhite ml-0.5 rounded-full h-7 w-7 p-green-small">1</div>
                    <div class="p-small-grey">{{varietyname}}</div>
                    <div class="h-1/12 w-1/12"></div>
                </div>
            </div>
        

        </div>
</template>
/<script lang="ts">
import { defineComponent, ref } from 'vue'
import { CheckIcon, ArrowSmDownIcon } from '@heroicons/vue/solid'

export default defineComponent({
    components:{CheckIcon, ArrowSmDownIcon},
    setup(){
        const count = ref(6)
        return{count}
    },
    props:{
        varietyname:{
        type: String,
        default: "Lollo Bionda",
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