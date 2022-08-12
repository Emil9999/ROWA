import axios from 'axios'
import {ref} from 'vue'

export default function findSinglePlantModule(){
    
    const group = ref<string>("undefined")
    const url = 'http://localhost:5000/modulegroup/'
    const findGroup = (mNumber: number) => {
        if (global.debug)
        {
            if (mNumber <= 2){
                group.value = "herb"
            } else {
                group.value = "lettuce"
            }
        } else
        {
         axios.get(url+mNumber.toString()).then((r)=> group.value = r.data)
        }
    }   

    


    return {group, findGroup}
}