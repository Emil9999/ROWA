import axios from "axios"
import {ref} from 'vue'

export default function hardwareToggler(){
    const states : any = ref({"pump": true, "light": false, "airstone": true})
    const getState = () => {
        if(global.debug)
        {
            console.log('Got States')
        } else
        {
           Object.keys(states.value).forEach(e => 
            axios.get('http://localhost:8080/admin/'+ e)
            .then((r) => {states.value[e] = r.data})
            
            )
        }
    }
   
   
   
    const toggle = (toggleType: string, desiredState: boolean) => {
        const url = 'http://localhost:8080/admin/'+ toggleType +'/' + desiredState ? 'on' : 'off'
        if(global.debug)
        {
            console.log('toggeled ' + toggleType)
        } else
        {
        axios.get(url)
        .then()
        .catch()
        return
        }
        getState()
    }

    
    getState()
    
    return {toggle, states, getState}

}