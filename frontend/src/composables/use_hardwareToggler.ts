import axios from "axios"

export default function hardwareToggler(){

    const toggle = (toggleType: string) => {
        const url = 'http://localhost:8080/admin/toggle-'+ toggleType +'/'
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
    }

    
    return {toggle}

}