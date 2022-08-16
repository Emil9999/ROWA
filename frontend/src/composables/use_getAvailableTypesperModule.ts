import {ref} from 'vue'
import axios from 'axios'
import VarietyData from '@/types/VarietyData'

export default function getAvailTypesperModule() {
    const availTypes = ref<VarietyData[]>([])

    

    const loadTypes = (mNumber: number) => {
       
        axios.get('/admin/planttypes/'+mNumber.toString()).then((r) => {availTypes.value = r.data})
        .catch(error => {  if(global.debug)
            {
                if (mNumber < 3){
                    availTypes.value = [{plant_type: 'Basil', growth_time: 42}, {plant_type: 'Mojo Mint', growth_time: 42}, {plant_type: 'Mint', growth_time: 42}, {plant_type: 'Thai Basil', growth_time: 42}]
                } else {
                    availTypes.value = [{plant_type: 'Oakleaf', growth_time: 42}]
                }
            } else 
            {
                console.log(error)
            }
        })
        

       
    }


    return {availTypes, loadTypes}
}