import {ref} from 'vue'
import axios from 'axios'
import Plant from '../types/Plant'


export default function getFarmables(modulenumber: number){
    const modulePlants = ref<Plant[]>([])
    const error = ref<string>("")
    const moduleNumber = ref(modulenumber)

    const plantcountInModule = ref(0)
    
    const url = '/dashboard/farm/' 


    const  debugDataHerb = ref<Plant[]>([
        {plant_type: "Mint", age: 60, position: 1, growth_time: 60},
        {plant_type: "Thai Basil", age: 60, position: 2, growth_time: 60},

        {plant_type: "Basil", age: 100, position: 4, growth_time: 60},
    ])
    const  debugDataTwoHerb = ref<Plant[]>([
        {plant_type: "Pepper Mint", age: 30, position: 1, growth_time: 60},
        {plant_type: "Thai Basil", age: 60, position: 2, growth_time: 60},
        {plant_type: "Red Basil", age: 10, position: 3, growth_time: 60},

        {plant_type: "Genovese Basil", age: 30, position: 4, growth_time: 60},
    ])
    const  debugDataSalad = ref<Plant[]>([
        {plant_type: "Oakleaf", age: 7, position: 2, growth_time: 42},
        {plant_type: "Oakleaf", age: 1, position: 3, growth_time: 42},
        {plant_type: "Oakleaf", age: 21, position: 4, growth_time: 42},
        {plant_type: "Oakleaf", age: 30, position: 5, growth_time: 42},
    ])
    const  debugDataSaladTwo = ref<Plant[]>([
        {plant_type: "Oakleaf", age: 30, position: 5, growth_time: 42}
    ])
    
    const loadModulePlants = (mNumber: number = modulenumber) => {
            console.log(mNumber)
            axios.get(url+mNumber.toString())
            .then(response => {modulePlants.value = response.data.plants
                            plantcountInModule.value = response.data.plant_spots})
            .catch(error => {  if(global.debug)
                {
                    if (mNumber == 1){
                        modulePlants.value = debugDataHerb.value
                        plantcountInModule.value = 4} 
                        else if(mNumber == 2){
                            modulePlants.value = debugDataTwoHerb.value
                            plantcountInModule.value = 4
                        }

                        else if(mNumber == 8){
                            modulePlants.value = debugDataSaladTwo.value
                            plantcountInModule.value = 6
                        }
                        else {  modulePlants.value = debugDataSalad.value
                                plantcountInModule.value = 6
                        }
                } else 
                {
                    console.log(error)
                }
            })
        
    }

    loadModulePlants()

    return {modulePlants, error, moduleNumber, loadModulePlants, plantcountInModule}
}
