<template>
    <div class="centered-div">
        <div>{{this.moduleNumber}}</div>
        <div class="h-green-big">{{text.name}}</div>
        <div class="p-grey-small">{{text.text}}</div>

        <div> 
            <button v-if="moduleFarmable[0].plantables" class="btn-small-green">Plant</button>
            <button v-if="moduleFarmable[0].harvestables" class="btn-small-green">Harvest</button>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import ModuleFarmable from '../../types/ModuleFarmable'
import {useRouter} from 'vue-router'

export default defineComponent({

    props:{
        moduleNumber:{
            type: Number,
            required: true
        }
    },
    setup() {

        const router = useRouter()
        const moduleFarmable =  ref<ModuleFarmable[]>([{
            varietyName: 'Lollo Bionda',
            harvestables: {planttype: 'Lollo Bionda', planter: '', modulenumber: 1, position: 1},
            plantables: {planttype: 'Lollo Bionda', planter: '', modulenumber: 1, position: 4}
        }])

        const onClick = () => {
            router.push( {
                name:"dataList", params: {}})
        }
        
        const text = ref({
            name: 'Lollo',
            text: 'Officia proident minim ad eu ea fugiat deserunt cupidatat dolor aute est Lorem sit. Consequat nisi officia minim culpa magna ea ea enim reprehenderit velit. Anim cillum laboris aute ea voluptate. Laborum irure exercitation qui amet ad sunt sunt culpa commodo laborum. Exercitation incididunt ipsum do reprehenderit et occaecat eu adipisicing. Ad non exercitation ullamco duis quis dolore reprehenderit adipisicing ea labore voluptate cupidatat in. Ut fugiat deserunt consectetur ipsum exercitation consequat ex veniam.'
        })

        return {text, moduleFarmable}
    },
})
</script>
