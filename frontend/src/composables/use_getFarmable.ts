import FarmablePlant from '../types/FarmablePlant'
import {ref} from 'vue'
import axios from 'axios'


export default function getFarmables (farmingType: string){
    const farmModules = ref<FarmablePlant[]>([])
    const error = ref<string>("")

    
    const url = farmingType == 'p' ? 'http://localhost:8080/plant' : 'http://localhost:8080/harvest'


    const debugDataPlanting = ref<FarmablePlant[]>([
        {planttype: 'Mint', planter: '', modulenumber: 1, position: 3, leafHarvest: false, group: "herb"},
          {planttype: 'Lollo Bionda', planter: '', modulenumber: 4, position: 1, leafHarvest: false, group: "lettuce"},
             
          {planttype: 'Lollo Bionda', planter: '', modulenumber: 5, position: 1, leafHarvest: false, group: "lettuce"}
    ])
    const debugDataHarvesting = ref<FarmablePlant[]>([
        {planttype: 'Mint', planter: '', modulenumber: 1, position: 1, leafHarvest: true, group: "herb"},
         {planttype: 'Thai Basil', planter: 'Emil', modulenumber: 1, position: 2, leafHarvest: true, group: "herb"},
           {planttype: 'Basil', planter: 'Simon', modulenumber: 1, position: 4, leafHarvest: false, group: "herb"},

           {planttype: 'Thai Basil', planter: '', modulenumber: 2, position: 1, leafHarvest: true, group: "herb"},
    
          {planttype: 'Basil', planter: 'Emil', modulenumber: 2, position: 4, leafHarvest: true, group: "herb"},
    
            {planttype: 'Lollo Bionda', planter: 'Hannes', modulenumber: 3, position: 6, leafHarvest: false, group: "lettuce"},
            {planttype: 'Lollo Bionda', planter: 'Hannes', modulenumber: 3, position: 5, leafHarvest: true, group: "lettuce"},
            {planttype: 'Lollo Bionda', planter: 'Emil', modulenumber: 4, position: 6, leafHarvest: true, group: "lettuce"},
             {planttype: 'Lollo Bionda', planter: 'Hannes', modulenumber: 5, position: 6, leafHarvest: true, group: "lettuce"},
             
          {planttype: 'Lollo Bionda', planter: 'Simon', modulenumber: 6, position: 6, leafHarvest: false, group: "lettuce"}
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
