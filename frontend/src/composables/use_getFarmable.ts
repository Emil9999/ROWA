import FarmablePlant from '../types/FarmablePlant'
import {ref} from 'vue'
import axios from 'axios'


export default function getFarmables (farmingType: string){
    const farmModules = ref<FarmablePlant[]>([])
    const error = ref<string>("")

    
    const url = farmingType == 'p' ? 'http://localhost:8080/plant/finish' : 'http://localhost:8080/harvest/finish'


    const debugDataPlanting = ref<FarmablePlant[]>([
        {planttype: 'Mint', planter: '', modulenumber: 2, position: 3},
          {planttype: 'Pirat', planter: '', modulenumber: 4, position: 4},
             
          {planttype: 'Pirat', planter: '', modulenumber: 5, position: 4}
    ])
    const debugDataHarvesting = ref<FarmablePlant[]>([
        {planttype: 'Mint', planter: '', modulenumber: 1, position: 1},
         {planttype: 'Basil', planter: 'Emil', modulenumber: 1, position: 2},
          {planttype: 'Thai Basil', planter: 'Simon', modulenumber: 1, position: 3},
           {planttype: 'Mint', planter: 'Emil', modulenumber: 1, position: 4},

           {planttype: 'Mint', planter: '', modulenumber: 2, position: 1},
         {planttype: 'Basil', planter: 'Emil', modulenumber: 2, position: 2},
          {planttype: 'Mint', planter: 'Emil', modulenumber: 2, position: 4},
    
            {planttype: 'Pirat', planter: 'Emil', modulenumber: 3, position: 1},
             {planttype: 'Lollo Bionda', planter: 'Emil', modulenumber: 4, position: 1},
             {planttype: 'Lollo Bionda', planter: 'Emil', modulenumber: 5, position: 1},
             
          {planttype: 'Pirat', planter: 'Simon', modulenumber: 6, position: 1}
    ])
    farmModules.value = farmingType == 'p' ? debugDataPlanting.value : debugDataHarvesting.value

    const loadFarmables = () => {
        return
        axios.get(url)
        .then(response => {
            farmModules.value = response.data
        })
        .catch(  () => {
            farmModules.value = farmingType == 'p' ? debugDataPlanting.value : debugDataHarvesting.value
        }
           
        )
    }


    return {farmModules, error, loadFarmables}
}
