import FarmablePlant from '../types/FarmablePlant'
import {ref} from 'vue'
import axios from 'axios'


const getFarmables = () => {
    const farmableModules = ref<FarmablePlant[]>([])
    const error = ref(null)
    

    const loadFarmables = () => {
        axios.get('http://localhost:8080/plant/finish')
        .then(response => {
            farmableModules.value = response.data
        })
    }

    loadFarmables()

    return {farmableModules, error, loadFarmables}
}

export default {getFarmables}