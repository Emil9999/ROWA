<template>
    <div class="grid grid-cols-3 gap-10 my-16 place-self-center ">
            <div v-for="x in 9" :key="x">
                <div class="btn-keypad" @click="$emit('numberClicked',x.toString())">{{x}}</div>
            </div>
            <div class="btn-keypad" @click="$emit('backClicked')"><ArrowLeftIcon class="h10 w-10"/></div>
            <div class="btn-keypad" @click="$emit('numberClicked', '0')">0</div>
            <div :class="altText != '' ? 'btn-keypad' : null" @click="$emit('clickedAlt')">{{altText}}</div>
        </div>

</template>

<script lang="ts">
import { defineComponent } from 'vue'
import {ArrowLeftIcon} from '@heroicons/vue/solid'

export default defineComponent({
    components: {ArrowLeftIcon},
    props:{
        altText:{
            type: String,
            default: ''

        }
    },
    emits: ['numberClicked','clickedAlt', 'backClicked', ],

    setup() {
    return    
    },
})
</script>
