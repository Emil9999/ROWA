import {ref} from 'vue'
import getFarmables from './use_getFarmable'
import FarmablePlant from '@/types/FarmablePlant'

export default function getFarmablePosbyIndex(){

    const  harvestableSpots = ref<number[][]>([[], [], [], [], [],  [], [], []])
    const  plantableSpots = ref<number[][]>([[], [], [], [], [],  [], [], []])


    const {farmModules: harvestablePlants, loadFarmables: loadHarvestable} = getFarmables('h')
    const {farmModules: plantablePlants, loadFarmables: loadPlantable} = getFarmables('p')



    loadHarvestable()
    loadPlantable()
   

    const findFarmableSpotsperIndex = () => {

      
       
        for(let i = 0; i <  8; i++){
            for(const farmablePlant of harvestablePlants.value){
                if(farmablePlant.modulenumber == i+1){
                    harvestableSpots.value[i].push(farmablePlant.position-1)
                }
            }
        }
        for(let i = 0; i <  8; i++){
            for(const farmablePlant of plantablePlants.value){
                if(farmablePlant.modulenumber == i+1){
                    plantableSpots.value[i].push(farmablePlant.position-1)
                }
            }
        }
    }
    

    findFarmableSpotsperIndex()
    
    
    return{harvestableSpots, plantableSpots}

}