import {ref} from 'vue'
import FinishFarmingData from '../types/FinishFarmingData'
import axios from 'axios'


export default function finishHarvesting(type: string,data: FinishFarmingData) {
    const status = ref('empty')
    const url = type == 'p' ? 'http://localhost:8080/plant/finish'  : 'http://localhost:8080/harvest/finish'
   
    
    axios.post(url, {...data})
        .then(response => {
           status.value = response.statusText
        })
        .catch(error => {
            status.value = error
        }
           
        )



    return {status}

} 