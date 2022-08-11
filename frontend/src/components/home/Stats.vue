<template>
    <div class="grid grid-cols-4 gap-10 w-full h-full px-24">
        <div class="col-span-1 stat-box">
            <div class="h-realwhite-small">Total Plants</div>
            <div class="h-realwhite-big">Planted</div>
            <div class="h-realwhite-big mt-5">25</div>
        </div>
        <div class="col-span-1 stat-box">
            <div class="h-realwhite-small">Total Plants</div>
            <div class="h-realwhite-big">Harvested</div>
            <div class="h-realwhite-big mt-5">20</div>
        </div>
        <div class="col-span-2 stat-box">
                <div class="h-realwhite-big">Water Level</div>
                <div class="flex h-2/3">
                    <div class=" w-1/2 p-3"><percentBar/></div>
                    <div class="h-realwhite-big">50%</div>
                </div>
            </div>
        <div class="col-span-4 stat-box">Temperature</div>
        <div class="col-span-4 stat-box">Humidity</div>

    </div>
</template>

/<script lang="ts">
import { defineComponent } from 'vue'
import percentBar from './atoms/percentBar.vue'

export default defineComponent({
    components:{percentBar},
    setup() {
        return
    },
})
</script>
