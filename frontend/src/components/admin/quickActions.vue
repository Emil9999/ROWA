<template>
   <div class="central-div">
       <div class="h-green-big my-8">Quick Actions</div>
       <div class="flex justify-evenly justify-items-center" v-for="(value, name) in states" :key="name">
            <div class="h-green-small ">{{fLCapital(name)}}</div>
            <toggleSwitch :state="value" @stateChange="(n) => callToggle(n, name)" />
           
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
           
            toggle(type, state)
           

        }


        const {toggle, states} = hardwareToggler()
        return{callToggle,fLCapital, states}
    },
})
</script>