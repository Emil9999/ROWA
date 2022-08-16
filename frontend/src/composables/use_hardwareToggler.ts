import axios from "axios"
import {ref} from 'vue'

export default function hardwareToggler(){
    const states : any = ref({"pump": true, "light": false, "airstone": true})
    const getState = () => {
      
           Object.keys(states.value).forEach(e => 
            axios.get('http://localhost:5000/admin/'+ e)
            .then((r) => {states.value[e] = r.data})
            
            )
        
    }
   
   
   
    const toggle = (toggleType: string, desiredState: boolean) => {
        const url = '/admin/'+ toggleType +'/' + (desiredState ? 'on' : 'off')
        axios.get(url)
        .then()
        .catch()
        
        getState()
    }

    
    getState()
    
    return {toggle, states, getState}

}