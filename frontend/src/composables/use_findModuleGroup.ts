import axios from 'axios'
import {ref} from 'vue'

export default function findModuleGroup(){
    
    const group = ref<string>("undefined")
    const url = 'http://localhost:5000/modulegroup/'
    const findGroup = (mNumber: number) => {
        
         axios.get(url+mNumber.toString()).then((r)=> group.value = r.data)
         .catch(error => {  if(global.debug)
                                {
                                    if (mNumber <= 2){
                                        group.value = "herb"
                                    } else {
                                        group.value = "lettuce"
                                    }
                                } else {
                                    console.log(error)
                                }
        })
        
    }   

    


    return {group, findGroup}
}