<template>
    
    <div class="centered-div">
        <div class=" mt-auto">Please Enter Admin Pin</div>
        <div class="centered-div w-full">
            <div class="my-10 flex w-1/2 justify-evenly">
                <div v-for="x in 4" :key="x">
                    <div :class="{'bg-green': enteredCode.length >= x, 'animate-errorwiggle border-red bg-none': animWrongCode}" class="w-6 h-6 rounded-full border-green border-2"></div>
                    
                </div>
            </div>
            <div v-if="animWrongCode" class="text-red">Wrong admin pin, please try again</div>
        </div> 

        <div class="grid grid-cols-3 gap-16 my-10 justify-self-center">
            <div v-for="x in 9" :key="x">
                <div class="btn-selector-white" @click="addNumber(x.toString())">{{x}}</div>
            </div>
            <div class="btn-selector-white" @click="removeNumber()"><ArrowLeftIcon class="h10 w-10"/></div>
            <div class="btn-selector-white" @click="addNumber('0')">0</div>
            <div class="btn-selector-white"></div>
        </div>

        <div v-if="enteredCode.length > 0" @click="resetNumber()" class="h-green-small">Reset Pin</div>

    </div>
</template>


<script lang="ts">
import { defineComponent, ref, } from 'vue'
import {ArrowLeftIcon} from '@heroicons/vue/solid'

export default defineComponent({
    components: {ArrowLeftIcon},
    emits:[ 'correctEntry'],
    setup(props, ctx) {

        const pin = ref("0815")

        const animWrongCode = ref(false)
        
        const enteredCode = ref("")

        const addNumber = (enteredNumber: string) =>{
           if( enteredCode.value.length < pin.value.length){
               enteredCode.value += enteredNumber
                compareEntrys()
            }
        }

        function animateWrongCode(){
            animWrongCode.value = true
            setTimeout(() => {
                animWrongCode.value = false
            }, 500)
        }

        const removeNumber = () => {
            if (enteredCode.value.length > 0){
               enteredCode.value = enteredCode.value.slice(0, -1)
            }

        }

        const resetNumber = () => {
            enteredCode.value = ""
        }

        const compareEntrys = () => {
            if(enteredCode.value == pin.value){
                ctx.emit('correctEntry')
                return true
            } else { 
                    if(enteredCode.value.length == pin.value.length) { 
                        resetNumber()
                        animateWrongCode()}
                    return false}
        }

        return{enteredCode, animWrongCode, addNumber, compareEntrys, removeNumber, resetNumber}
    },
})
</script>
