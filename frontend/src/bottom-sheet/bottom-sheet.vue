<template>
    <Teleport to="body">
        <div :class="{'invisible': !(isOpen && Background)}" class="absolute top-0">
            <div>
                <div @click="isOpen = !isOpen" class="bg-white text-green absolute right-5">Exit</div>
            <img class="object-cover h-96 w-screen" src="../assets/img/ideas/garnish.jpg"><img>
            
            </div>
        </div>
        <div :class="{'invisible': !isOpen}" class="absolute bottom-0 w-full overflow-auto max-h-screen z-50 left-0 rounded-t-2xl bg-white">
            <div class="flex justify-center mt-5 h-green-mid">
             <button @click="isOpen = !isOpen">Close</button>
            </div>
            <div class="">
            <slot name="scrollable"/>
            </div>
            <div class="w-full centered-div">
            <slot> </slot>
            </div>
        </div>
    </Teleport>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'



export default defineComponent({
    props:{
        isopen:{
            type: Boolean,
            default: true
        },
        background:{
            type: Boolean,
            default: false
        },
},
    
    setup(props) {
        const sheetP = ref(0)
        const isOpen = ref(props.isopen)
        const Background = ref(props.background)

        const open = () => (isOpen.value = true)
        const close = () => (isOpen.value = false)
        
        
       
        return{sheetP, isOpen, open, close, Background}

        
    },
    
    
})

</script>

<style scoped>

.bottomsheet{
    z-index: 99999;
}
.closed {
opacity: 0;
visibility: hidden;
}
.opened {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}

</style>>

