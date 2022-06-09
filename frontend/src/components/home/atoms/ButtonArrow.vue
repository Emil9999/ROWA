<template>
    <button :class="[buttonClass]" :disabled="bEnabled" @click="this.$router.push(link)">
            <span v-if="!bEnabled" class="w-10"></span>
            <span class="">{{buttonText}}</span>
            <ArrowRightIcon v-if="!bEnabled" class="text-white h-10 w-10"></ArrowRightIcon>
      </button>

</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import {ArrowRightIcon} from '@heroicons/vue/solid'



export default defineComponent({
    components:{ArrowRightIcon},
    props:{
        buttonText:{
            type: String,
            required: true
        },
        bEnabled:{
            type: Boolean,
            required: true
        },
        link:{
            type: String,
            required: true
        },
        smbutton:{
            type: Boolean,
            default: false
        }
    },
    setup(props) {
      const activatedButton = ref('btn-big-green inline-flex items-center justify-between px-5')  
      const disabledButton = ref('btn-big-dis justify-center') 
        const buttonClass =  computed(() =>  ({
            'btn-big-green': !props.smbutton && !props.bEnabled,
            'btn-big-dis': !props.smbutton && props.bEnabled,
            'btn-small-green': props.smbutton && !props.bEnabled,
            'btn-small-dis': props.smbutton && props.bEnabled,
            'inline-flex items-center justify-between px-5': !props.bEnabled,
            'justify-center': props.bEnabled
        }))
      return {activatedButton, disabledButton, buttonClass}
    },
})
</script>
