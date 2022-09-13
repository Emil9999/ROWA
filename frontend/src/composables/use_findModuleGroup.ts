import axios from 'axios'
import {ref} from 'vue'

export default function findModuleGroup(){
    
    const group = ref<string>("undefined")
    const url = '/modulegroup/'
    const findGroup = (mNumber: number) => {
        
         axios.get(url+mNumber.toString()).then((r)=> group.value = r.data.group)
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