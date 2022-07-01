//import axios from 'axios'
import {ref} from 'vue'

export default function findSinglePlantModule(){

    const group = ref<string>("undefined")
    //const url = 'http://localhost:8080/plant'
    const findGroup = (mNumber: number) => {
        if (mNumber <= 2){
            group.value = "herb"
        } else {
            group.value = "lettuce"
        }
    }   
    //axios.get(url)
    


    return {group, findGroup}
}