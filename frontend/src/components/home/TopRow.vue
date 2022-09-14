<template>
    <div class="w-full flex justify-between items-center">
        <div @click="this.$router.push('/admin')" class="bg-white shadow-list rounded-full h-16 w-16"><CogIcon class="h-16 w-16 text-green"></CogIcon></div>

        <div class="bg-white h-green-small flex rounded-xl justify-around">
           <div v-for="item in Selectables" :key="item" @click="clickedRow(item)">
            <div v-if="item != selected"  class="bg-white px-4 py-1 m-2 rounded-xl ">{{item}}</div>
            <div v-if="item == selected" class="bg-green px-4 py-1 m-2 rounded-xl shadow-md text-white">{{item}}</div>
           </div>    
        </div>

        <div @click="$emit('faqOpen')" class="bg-white shadow-list rounded-full text-green text-center font-bold self-center text-6xl h-16 w-16">?</div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import {CogIcon } from '@heroicons/vue/solid'

export default defineComponent({
    components: {CogIcon},
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
    emits: ["ClickedRow", "faqOpen"],
    methods: {
        clickedRow(choose: string) {
        this.selected = choose
        this.$emit("ClickedRow", choose);
        },
  }
   
})
</script>