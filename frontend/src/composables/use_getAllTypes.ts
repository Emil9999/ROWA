
import axios from 'axios'
import {ref} from 'vue'
import AvailTypes from '../types/AvailVariety'


export default function getAllTypes(){
    const url = '/admin/planttypes' 
    const availTypes = ref<AvailTypes[]>([])
   
        axios.get(url).then((r) => {availTypes.value = r.data})
        .catch(error => {  if(global.debug)
            {
                availTypes.value = [{plant_type: 'Thai Basil', group: 'herb'},
                {plant_type: 'Basil', group: 'herb'},
                {plant_type: 'Mint', group: 'herb'},
                {plant_type: 'Thyme', group: 'herb'},
                {plant_type: 'Cilantro', group: 'herb'},
                {plant_type: 'Lollo Bionda', group: 'lettuce'},
                {plant_type: 'Neckar Giant', group: 'lettuce'},
                {plant_type: 'Pirat Lettuce', group: 'lettuce'}]
            } else {
                console.log(error)
            }
})
    




    return {availTypes}
}