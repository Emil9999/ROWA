import {ref} from 'vue'
import FarmablePlant from '../types/FarmablePlant'
import axios from 'axios'


export default function finishHarvesting(type: string,data: FarmablePlant) {
    const status = ref('empty')
    const url = type == 'p' ? 'http://localhost:8080/plant'  : 'http://localhost:8080/harvest'
   
    
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