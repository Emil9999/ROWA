<template>
    <div class="w-full flex justify-between items-center">
        <div><CogIcon class="h-16 w-16 text-green"></CogIcon></div>

        <div class="bg-white h-green-small flex rounded-xl justify-around">
           <div v-for="item in Selectables" :key="item" @click="clickedRow(item)">
            <div v-if="item != selected"  class="bg-white px-4 py-1 m-2 rounded-xl ">{{item}}</div>
            <div v-if="item == selected" class="bg-green px-4 py-1 m-2 rounded-xl shadow-md text-white">{{item}}</div>
           </div>    
        </div>

        <div><QuestionMarkCircleIcon class="h-16 w-16 text-green"></QuestionMarkCircleIcon></div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { QuestionMarkCircleIcon, CogIcon } from '@heroicons/vue/solid'

export default defineComponent({
    components: {QuestionMarkCircleIcon, CogIcon},
    setup(props){
        const Selectables = ref(props.menuPoints)
        const selected = ref("Home")

        return {Selectables, selected}
    },
    props: {
        defaultSelection: {
        type: String,
        required: true
    },
        menuPoints: {
            type: Array,
            required: true

        }
    },
    emits: ["ClickedRow"],
    methods: {
        clickedRow(choose: string) {
        this.selected = choose
        this.$emit("ClickedRow", choose);
        },
  }
   
})
</script>