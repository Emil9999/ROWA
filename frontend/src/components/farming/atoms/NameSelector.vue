<template>
<div>
    <div v-if="!selectedName">
        <div class="h-green-big">Let your colleagues <br> know it was you...</div>
        <div class="p-grey-big mt-5">The name will appear next to the plant <br> when harvesting.</div>
        <button @click="openSheet" class="btn-big-white">Add Name</button>

    </div>
    <div v-else>
    <div class="h-green-big">Thanks {{selectedName}}</div>
    <div class=""> 
        <button @click="openSheet" class="btn-small-white">Change Name</button>
        <button @click="removeName" class="btn-small-white">Remove Name</button>
    </div>
    

</div>
<Sheet :isopen="false" ref="mysheet">
    <div>
        <div class="h-green-big">Add your Name</div>
        <div><button @click="enteringName = true" class="btn-big-white">Add new Name</button></div>
        
        <div v-if="enteringName">
            <Keyboard @saveName="addName"/>
        </div>
        
        <div v-else>
            <div class="h-green-mid mb-5">Choose Exisiting Name:</div>
            <div class="grid grid-cols-4 gap-10 my-10">
                <div v-for="(name, index) in names" :key="index"><button @click="addName(name)" class="btn-selector-white">{{name}}</button></div>
            </div>
        </div>
    </div>
</Sheet>
</div>
    
</template>

<script lang="ts">
import Sheet from '../../../bottom-sheet/bottom-sheet.vue'
import Keyboard from './keyboard.vue'
import { defineComponent, ref } from 'vue'

export default defineComponent({
    components:{Sheet, Keyboard},
    emits: ['changeName'],
    setup(_, {emit}) {
        
        const mysheet = ref<InstanceType<typeof Sheet> | null>(null)
        const openSheet = () => {
        mysheet.value?.open()
        
    }
        const closeSheet = () => {
        mysheet.value?.close()
        
    }

        const enteringName = ref(false)
        const selectedName = ref()

        function removeName(){
            selectedName.value = ""
            emit('changeName', selectedName.value)
        }
        function addName(name:string){
            closeSheet()
            selectedName.value = name
            enteringName.value = false
            emit('changeName', selectedName.value)
        }
        const names = ref<string[]>(["Hannes O.","Emil S.","Simon H.","Hannes O.","Emil S.","Simon H.","Hannes O.","Emil S.","Simon H.","Hannes O.","Emil S.","Simon H.","Hannes O.","Emil S.","Simon H."])
        

        return{selectedName, enteringName, removeName, addName, mysheet, openSheet, closeSheet, names}
        
        
    },
    
})
</script>
