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
        {variety: "Mint", age: 60, position: 1, growthTime: 60},
        {variety: "Thai Basil", age: 60, position: 2, growthTime: 60},

        {variety: "Basil", age: 100, position: 4, growthTime: 60},
    ])
    const  debugDataTwoHerb = ref<Plant[]>([
        {variety: "Thai Basil", age: 30, position: 1, growthTime: 60},
        {variety: "Thai Basil", age: 60, position: 2, growthTime: 60},
        {variety: "Mint", age: 10, position: 3, growthTime: 60},

        {variety: "Basil", age: 30, position: 4, growthTime: 60},
    ])
    const  debugDataSalad = ref<Plant[]>([
        {variety: "Lollo Bionda", age: 7, position: 2, growthTime: 42},
        {variety: "Lollo Bionda", age: 14, position: 3, growthTime: 42},
        {variety: "Lollo Bionda", age: 21, position: 4, growthTime: 42},
        {variety: "Lollo Bionda", age: 30, position: 5, growthTime: 42},
        {variety: "Lollo Bionda", age: 42, position: 6, growthTime: 42},
    ])
    
    const loadModulePlants = (mNumber: number = modulenumber) => {
        if(global.debug)
        {
            if (mNumber == 1){
                modulePlants.value = debugDataHerb.value
                plantcountInModule.value = 4} 
                else if(mNumber == 2){
                    modulePlants.value = debugDataTwoHerb.value
                    plantcountInModule.value = 4
                }
                else {  modulePlants.value = debugDataSalad.value
                        plantcountInModule.value = 6
                }
        } else
        {
            axios.get(url+mNumber.toString)
            .then(response => {modulePlants.value = response.data})
            .catch(error => {  if(global.debug)
                {
                    if (mNumber == 1){
                        modulePlants.value = debugDataHerb.value
                        plantcountInModule.value = 4} 
                        else if(mNumber == 2){
                            modulePlants.value = debugDataTwoHerb.value
                            plantcountInModule.value = 4
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
    }

    loadModulePlants()

    return {modulePlants, error, moduleNumber, loadModulePlants, plantcountInModule}
}
