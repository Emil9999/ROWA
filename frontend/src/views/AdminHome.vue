<template>
    <div>
        <div class="flex text-center justify-between mt-5 mx-3 items-center">
            <div class="w-16"> </div>
            <div class="h-green-big">Admin Panel</div>
            <div @click="this.$router.push('/')" class="bg-white shadow-md rounded-full h-16 w-16"><XIcon class="h-16 w-16 text-green"></XIcon></div>
        </div>
        <div v-if="!correctPin">
        <pinPanel @correctEntry="pinCorrect()"></pinPanel>
        </div>
       
        <div class="centered-div mt-16" v-else>
          
            <div class="centered-div">
            <div v-for="tabs in adminTabs" :key="tabs" class="my-5">
                <div class="p-grey-big mb-10" v-if="tabs.Title != ''">{{tabs.Title}}</div>
                <div class="grid grid-cols-2 gap-10">
                <button v-for="(tab, index) in tabs.Tabs" :key="index" :disabled="!(QuickTabs.includes(tab.replace(' ', '')) || paths.hasOwnProperty(tab))" :class="QuickTabs.includes(tab.replace(' ', '')) || paths.hasOwnProperty(tab)  ? 'btn-admin-white' : 'btn-admin-greyed'" @click="QuickTabs.includes(tab.replace(' ', '')) ? (selectedTab = tab.replace(' ', ''), openSheet()) : $router.push(paths[tab])">{{tab}}
            
                </button>
                </div>
            </div>
            </div>
           
             <Sheet :isopen="isOpen"  ref="adminSheet"> <component :is="selectableTabs[selectedTab]"></component></Sheet>
 

        </div>


    </div>
</template>



/<script lang="ts">
import { defineComponent, ref } from 'vue'
import {XIcon } from '@heroicons/vue/solid'
import pinPanel from '../components/admin/pinPanel.vue'
import Sheet from '../bottom-sheet/bottom-sheet.vue'
import LightTimes from '../components/admin/lightTimes.vue'
import PumpTimes from '../components/admin/pumpTimes.vue'
import QuickActions from '../components/admin/quickActions.vue'

export default defineComponent({
    components: {XIcon, pinPanel, Sheet},
    setup() {

        const isOpen = ref(false)

        const selectableTabs = {LightTimes, QuickActions, PumpTimes}
        const QuickTabs = Object.keys(selectableTabs)
        const paths = {'Reality Check': '/admin/realitycheck', 'Change Plant Varieties' : '/admin/variety'}
        const adminTabs = {
            1:{Title: "",
               Tabs: ["Quick Actions", "Support"]
               },
            2:{Title: "Farming",
               Tabs: ["Reality Check", "Change Plant Varieties"]
               },
            3:{Title: "System",
               Tabs: ["Light Times", "Pump Times", "Change Pass Code"]
               },}
        const selectedTab = ref('LightTimes')


        const adminSheet = ref<InstanceType<typeof Sheet> | null>(null)
        const openSheet = () => {
            if (QuickTabs.includes(selectedTab.value)){
                adminSheet.value?.open()}
        }

        const correctPin = ref(true)


        const pinCorrect = () =>{ correctPin.value = true}

        return{correctPin,isOpen,selectedTab, pinCorrect,paths, selectableTabs, adminSheet, adminTabs, openSheet, QuickTabs}
        
    },
})
</script>
