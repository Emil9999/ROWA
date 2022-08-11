import FarmablePlant from '@/types/FarmablePlant'
import {ref} from 'vue'


export default function findFarmablePerModule(farmModules: Array<FarmablePlant>) {
    const plantsInModule = ref<FarmablePlant[]>([])
    const hasMultipleFarmable = ref(false)
    const hasFarmable = ref(true)
    const findforModule = (selectedmodule: number) => {
        plantsInModule.value = []
        hasMultipleFarmable.value = false
        hasFarmable.value = true
        for(const farmModule of farmModules){
                if(farmModule.modulenumber == selectedmodule){
                    plantsInModule.value.push(farmModule)
                }
            }
                if (plantsInModule.value.length > 1) { 
                    plantsInModule.value = plantsInModule.value.sort((a,b) => (a.position < b.position)? 1 : -1)
                    hasMultipleFarmable.value = true
                } else if(plantsInModule.value.length  == 0){hasFarmable.value = false}
    }

    return {plantsInModule, hasMultipleFarmable, hasFarmable, findforModule}
        
}