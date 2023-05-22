/<template>
    <div class="info-box-instruc">
    <div class="h-green-small">Click on a Module to learn more.</div>
    <div class="flex mt-4 p-grey-small">
        <div  class="bg-green mx-4 rounded-full h-7 w-7"><CheckIcon class="text-white"/></div> 
        <div>Available to harvest</div>           
        <div class="bg-brownred mx-4 rounded-full h-7 w-7"><ArrowSmDownIcon class="text-white"/></div> 
        <div>Available to plant</div>  
    </div>
  </div>
</template>

/<script lang="ts">
import { defineComponent } from 'vue'
import { CheckIcon, ArrowSmDownIcon } from '@heroicons/vue/solid'

export default defineComponent({
    components:{CheckIcon, ArrowSmDownIcon},
   
})
</script>
