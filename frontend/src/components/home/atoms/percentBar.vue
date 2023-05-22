<template>
    <div class="w-full h-full flex-row rounded-md">
        <div class=" bg-white" :style="{ 'height': fill + '%' }"></div>
        <div class="bg-red" :style="{ 'height': 100-fill + '%' }"></div>
    </div>
</template>

/<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    props:{
        fill:{
            type: Number,
            default: 30
        }
    },
    setup() {
        return
    },
})
</script>
