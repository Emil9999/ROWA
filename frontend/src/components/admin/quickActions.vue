<template>
   <div class="central-div">
       <div class="h-green-big my-8">Quick Actions</div>
       <div class="flex justify-evenly justify-items-center" v-for="option in toggleables" :key="option">
            <div class="h-green-small ">{{fLCapital(option)}}</div>
            <toggleSwitch class="" @stateChange="(n) => callToggle(n, option)" />
           
       </div>
       
   </div>
</template>
<script lang="ts">
import { defineComponent, ref } from 'vue'

import hardwareToggler from '../../composables/use_hardwareToggler'
import toggleSwitch from '../general/ToggleSwitch.vue'

export default defineComponent({
    components:{toggleSwitch},
    setup() {


        const fLCapital = (s: string) => s = s.charAt(0).toUpperCase() + s.slice(1)

        const callToggle = (state:boolean, type: string, ) =>{
            //Add support for correct toggle state
            if (state){
                toggle(type)
            } else { toggle(type)}

        }
        const toggleables = ref(['pump','light','airstone'])
        const {toggle} = hardwareToggler()
        return{callToggle,fLCapital, toggleables}
    },
})
</script>