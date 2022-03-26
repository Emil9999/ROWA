import {ref} from 'vue'
import axios from 'axios'
import Plant from '../types/Plant'


export default function getFarmables (moduleNumber: number){
    const modulePlants = ref<Plant[]>([])
    const error = ref<string>("")

    
    const url = 'http://localhost:8080/dashboard/farm/' 


    const  debugDataHerb = ref<Plant[]>([
        {variety: "Mint", plantDate: new Date(2022, 2, 20), position: 1, growthTime: 42},
        {variety: "Thai Basil", plantDate: new Date(2022, 2, 20), position: 2, growthTime: 35},
        {variety: "Basil", plantDate: new Date(2022, 2, 20), position: 3, growthTime: 35},
        {variety: "Mint", plantDate: new Date(2022, 2, 20), position: 4, growthTime: 20},
    ])

    modulePlants.value = debugDataHerb.value

    const loadModulePlants = () => {
        return
        axios.get(url+moduleNumber.toString)
        .then(response => {
            modulePlants.value = response.data
        })
        .catch(  () => {
            modulePlants.value = debugDataHerb.value 
        }
           
        )
    }

    loadModulePlants()

    return {modulePlants, error, loadModulePlants}
}
