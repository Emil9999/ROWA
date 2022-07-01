import {ref} from 'vue'
import axios from 'axios'
import Plant from '../types/Plant'


export default function getFarmables(modulenumber: number){
    const modulePlants = ref<Plant[]>([])
    const error = ref<string>("")
    const moduleNumber = ref(modulenumber)
    
    const url = 'http://localhost:8080/dashboard/farm/' 


    const  debugDataHerb = ref<Plant[]>([
        {variety: "Mint", plantDate: '12-03-2022', position: 1, growthTime: 42},
        {variety: "Thai Basil", plantDate: '12-03-2022', position: 2, growthTime: 35},
        {variety: "", plantDate: '12-03-2022', position: 3, growthTime: 0},

        {variety: "Basil", plantDate: '12-03-2022', position: 4, growthTime: 60},
    ])
    const  debugDataTwoHerb = ref<Plant[]>([
        {variety: "Thai Basil", plantDate: '12-03-2022', position: 1, growthTime: 42},
        {variety: "Thai Basil", plantDate: '12-03-2022', position: 2, growthTime: 35},
        {variety: "Mint", plantDate: '12-03-2022', position: 3, growthTime: 20},

        {variety: "Basil", plantDate: '12-03-2022', position: 4, growthTime: 60},
    ])
    const  debugDataSalad = ref<Plant[]>([
        {variety: "", plantDate: '12-03-2022', position: 1, growthTime: 0},
        {variety: "Lollo Bionda", plantDate: '12-03-2022', position: 2, growthTime: 7},
        {variety: "Lollo Bionda", plantDate: '12-03-2022', position: 3, growthTime: 14},
        {variety: "Lollo Bionda", plantDate: '12-03-2022', position: 4, growthTime: 21},
        {variety: "Lollo Bionda", plantDate: '12-03-2022', position: 5, growthTime: 28},
        {variety: "Lollo Bionda", plantDate: '12-03-2022', position: 6, growthTime: 35},
    ])
    
    const loadModulePlants = (mNumber: number = modulenumber) => {
        if (mNumber == 1){
            modulePlants.value = debugDataHerb.value} else if(mNumber == 2){
                modulePlants.value = debugDataTwoHerb.value
            }
            else { modulePlants.value = debugDataSalad.value}
        
       return
        axios.get(url+mNumber.toString)
        .then(response => {
            modulePlants.value = response.data
        })
        .catch(  () => {
            if (mNumber < 3){
                modulePlants.value = debugDataHerb.value} else { modulePlants.value = debugDataSalad.value}
            
        }
           
        )
    }

    loadModulePlants()

    return {modulePlants, error, moduleNumber, loadModulePlants}
}
