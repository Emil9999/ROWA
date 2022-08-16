import FarmablePlant from '../types/FarmablePlant'
import {ref} from 'vue'
import axios from 'axios'


export default function getFarmables (farmingType: string){
    const farmModules = ref<FarmablePlant[]>([])
    const error = ref<string>("")

    
    const url = farmingType == 'p' ? '/plant' : '/harvest'


    const debugDataPlanting = ref<FarmablePlant[]>([
        {plant_type: 'Mint',  module: 1, position: 3, group: "herb"},
          {plant_type: 'Lollo Bionda',  module: 4, position: 1, group: "lettuce"},
             
          {plant_type: 'Lollo Bionda', module: 5, position: 1, group: "lettuce"}
    ])
    const debugDataHarvesting = ref<FarmablePlant[]>([
        {plant_type: 'Mint', planter: '', module: 1, position: 1, leaf_harvest: true, group: "herb"},
         {plant_type: 'Thai Basil', planter: 'Emil', module: 1, position: 2, leaf_harvest: true, group: "herb"},
           {plant_type: 'Basil', planter: 'Simon', module: 1, position: 4, leaf_harvest: false, group: "herb"},

           {plant_type: 'Thai Basil', planter: '', module: 2, position: 1, leaf_harvest: true, group: "herb"},
    
          {plant_type: 'Basil', planter: 'Emil', module: 2, position: 4, leaf_harvest: true, group: "herb"},
    
            {plant_type: 'Lollo Bionda', planter: 'Hannes', module: 3, position: 6, leaf_harvest: false, group: "lettuce"},
            {plant_type: 'Lollo Bionda', planter: 'Hannes', module: 3, position: 5, leaf_harvest: true, group: "lettuce"},
            {plant_type: 'Lollo Bionda', planter: 'Emil', module: 4, position: 6, leaf_harvest: true, group: "lettuce"},
             {plant_type: 'Lollo Bionda', planter: 'Hannes', module: 5, position: 6, leaf_harvest: true, group: "lettuce"},
             
          {plant_type: 'Lollo Bionda', planter: 'Simon', module: 6, position: 6, leaf_harvest: false, group: "lettuce"}
    ])
    const loadFarmables = () => {
        axios.get(url).then(response => {farmModules.value = response.data})
        .catch(error => {  if(global.debug)
            {
                farmModules.value = farmingType == 'p' ? debugDataPlanting.value : debugDataHarvesting.value
        
            } else 
            {
                console.log(error)
            }
        })
    }
    loadFarmables()

    return {farmModules, error, loadFarmables}
}
