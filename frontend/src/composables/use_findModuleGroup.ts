import axios from 'axios'
import {ref} from 'vue'

export default function findModuleGroup(){
    
    const group = ref<string>("undefined")
    const url = '/modulegroup/'
    const findGroup = (mNumber: number) => {
        
         axios.get(url+mNumber.toString()).then((r)=> {
                                    if(r.data.length == 1){
                                    group.value = r.data[0]
                                    console.log(r.data[0])}
                                    else if(r.data.length > 1){
                                        console.log('Unexpected length of Group Array, using first Element')
                                        group.value = r.data[0]
                                    } else{
                                        console.log('couldnt determine length of Group array or Empty Array')
                                    }
                                
                                })
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