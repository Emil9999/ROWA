<template>
  
        <div class=" bg-accentwhite h-green-small flex rounded-xl justify-around">
           <div v-for="item in Selectables" :key="item" @click="clickedRow(item)">
            <div v-if="item != selected" class="px-4 py-1 m-2 rounded-xl">{{item}}</div>
            <div v-if="item == selected" class="bg-green px-4 py-1 m-2 rounded-xl shadow-md text-white">{{item}}</div>
           </div>    
        </div>

</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
    components: {},
    setup(props){
        const Selectables = ref(props.menuPoints)
        const selected = ref(props.defaultSelection)

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

        },
        disabledOptions: {
            type: Boolean,
            default: true
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