<template>
<div>
    <div class="h-green-small">{{enteredName}}</div>
    <div class="grid grid-cols-11 gap-5 my-10 justify-self-center">
             
            <div v-for="x in firstRow" :key="x">
                <div class="btn-selector-white" @click="addChar(x+CaseOffSet)">{{asciiToChar(x+CaseOffSet)}}</div>
            </div>      
            <div class="btn-selector-white" @click="removeChar"><ArrowLeftIcon class="h10 w-10"/></div>
    </div>
    <div class="grid grid-cols-10 gap-5 mb-10 justify-self-center">
        <div class="btn-selector-white" @click="upperCase =! upperCase"><ArrowUpIcon class="h-10 w-10"/></div>
            <div v-for="x in secondRow" :key="x">
                <div class="btn-selector-white" @click="addChar(x+CaseOffSet)">{{asciiToChar(x+CaseOffSet)}}</div>
            </div>      
    </div>
    <div class="grid grid-cols-2 gap-5 my-10 justify-self-center">         
        <div class="btn-selector-white" @click="$emit('saveName', enteredName)">Save</div>    
        <div class="btn-selector-white" @click="addChar(32)">Space</div>
    </div>
</div>
</template>

<script lang="ts">
import { ref, defineComponent, computed } from 'vue'
import {ArrowLeftIcon, ArrowUpIcon} from '@heroicons/vue/solid'

export default defineComponent({
    components: {ArrowLeftIcon, ArrowUpIcon},
  
    emits: ['saveName'],

    setup() {

        const firstRow = ref([17,23,5,18,20,26,21,9,15,16])
        const secondRow = ref([1,19,4,6,7,8,10,11,12,25,24,3,22,2,14,13])

        const CaseOffSet = computed(() => {return upperCase.value ? 64: 96})
      
        const enteredName = ref('')

        const addChar = (i: number) =>{enteredName.value += asciiToChar(i); upperCase.value = false}
        const removeChar = () =>{enteredName.value = enteredName.value.slice(0, -1)}

        const upperCase = ref(false)

        const asciiToChar = (i: number) => {
            return String.fromCharCode(i)
        }

    return {asciiToChar, CaseOffSet, enteredName, upperCase, addChar, removeChar, firstRow, secondRow}   
    },
})
</script>
