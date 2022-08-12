import {ref} from 'vue'
import axios from 'axios'
import VarietyData from '@/types/VarietyData'

export default function getAvailTypesperModule() {
    const availTypes = ref<VarietyData[]>([])

    

    const loadTypes = (mNumber: number) => {
        if(global.debug)
        {        if (mNumber < 3){
            availTypes.value = [{name: 'Basil', gTime: 42}, {name: 'Mojo Mint', gTime: 42}, {name: 'Mint', gTime: 42}, {name: 'Thai Basil', gTime: 42}]
        } else {
            availTypes.value = [{name: 'Lollo Bionda', gTime: 42}]
        }
        return
        } else 
        {
            axios.get('http://localhost:5000/admin/planttypes/'+mNumber.toString()).then((r) => {availTypes.value = r.data})
        }

       
    }


    return {availTypes, loadTypes}
}