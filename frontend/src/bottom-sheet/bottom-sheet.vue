<template>
    <Teleport to="body">
        
        <div  :class="{'invisible': !fullScreen}" class="absolute top-0">
            <div>
                <div @click="closeAll" class="bg-white text-green absolute right-5">{{background}}</div>
            <img class="object-cover h-96 w-screen" :src="require('../assets/img/plants/'+cImage(background))"><img>
            
            </div>
        </div>
        <span v-touch:swipe.bottom="down" v-touch:swipe.top="up">
        <div v-on:click="closeAll" :class="{'invisible': !isOpen}" class=" bg-opacity-40 bg-grey z-50 w-full h-full absolute top-0" >
            <div v-on:click.stop  class="absolute bottom-0 w-full z-50 left-0 rounded-t-bottom-sheet bg-white">
               <div class="flex justify-center mt-5">
                <div class="w-28 h-2 rounded-full bg-grey mb-5"></div>
                </div>
              
                <div :class="[background == '' ? 'max-h-farmH' : fullScreen ? ' h-farmH' : ' max-h-200']" class="w-full overflow-hidden mb-5">
                <slot> </slot>
                </div>
            </div>
        </div>
        </span>
    </Teleport>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { checkImage } from '../composables/use_imgChecker'



export default defineComponent({
    props:{
        isopen:{
            type: Boolean,
            default: true
        },
        background:{
            type: String,
            default: ""
            
        },
},
    
    setup(props) {
        const sheetP = ref(0)
        const isOpen = ref(props.isopen)
        const fullScreen = ref(false)

        const {cImage} = checkImage('png')
        const open = () => (isOpen.value = true)
        const close = () => (isOpen.value = false)

        const up = () =>{
            if (props.background != ""){
                fullScreen.value = true
            }
        }

        const down = () =>{
            if (fullScreen.value == false){
                isOpen.value = false} else{fullScreen.value = false}
        }
        const closeAll = () => {
             isOpen.value = false
             fullScreen.value = false
        }
        
        
       
        return{sheetP,cImage, isOpen, closeAll, open, close, fullScreen, up, down}

        
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

