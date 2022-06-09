<template>
    <div>
  
        <div class="text-center h-white-small">
            <div class="mb-2">{{title}}</div>
            <div class="flex pl-12">
                <div v-for="n in 4" :key="n">
                    <svg width="68" height="13" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <circle v-if="n>this.step" r="6.5" transform="matrix(1 0 0 -1 6.5 6.5)" fill="#A0918D"/>
                    <circle v-if="n<=this.step" r="6.5" transform="matrix(1 0 0 -1 6.5 6.5)" fill="#009966"/>
                    <rect v-if="n!=4" x="13" y="5.5" width="53.9921" height="2" fill="#A0918D"/>
                    </svg>

                </div>
            </div>
        </div>

    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    props: {
        title: {
        type: String,
        default: "Error",
         },
        step: {
            type: Number,
        }
    },
})
</script>
