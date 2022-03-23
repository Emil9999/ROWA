<template>
    <div class="relative w-farmW h-farmH flex-none">
        <img width="720" height="auto" src="../../../assets/cat_tree/FarmFrame.png"/>
        <div @click="$emit('ModuleClicked', i)" :style="'top: '+(77+((i-1)*165))+'px;'" class="z-10 absolute left-14" v-for="i in ModuleCount" :key="i">
            <HerbModule v-if="i%2 == 0"/>
            <LettuceModule v-else/>
             </div>
         <div @click="$emit('ModuleClicked', i)"  :style="'top: '+(77+(i-1)*165)+'px;'" class="z-10 absolute right-14" v-for="i in ModuleCount" :key="i">
            <HerbModule :reverse="true" v-if="i%2 != 0"/>
            <LettuceModule :reverse="true" v-else/>
             </div>
    </div>
</template>

/<script lang="ts">
import { defineComponent, ref } from 'vue'

import HerbModule from './HerbModule.vue'
import LettuceModule from './LettuceModule.vue'


export default defineComponent({
    components: {HerbModule, LettuceModule},
    setup() {
        const ModuleCount = ref(4)
    
        return {ModuleCount}
        
    },
    
})
</script>
