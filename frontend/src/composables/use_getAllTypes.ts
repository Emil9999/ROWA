
import axios from 'axios'
import {ref} from 'vue'
import AvailTypes from '../types/AvailVariety'


export default function getAllTypes(){
    const url = 'http://localhost:5000/admin/planttypes' 
    const availTypes = ref<AvailTypes[]>([])
   
        axios.get(url).then((r) => {availTypes.value = r.data})
        .catch(error => {  if(global.debug)
            {
                availTypes.value = [{variety: 'Thai Basil', group: 'herb'},
                {variety: 'Basil', group: 'herb'},
                {variety: 'Mint', group: 'herb'},
                {variety: 'Thyme', group: 'herb'},
                {variety: 'Cilantro', group: 'herb'},
                {variety: 'Lollo Bionda', group: 'lettuce'},
                {variety: 'Neckar Giant', group: 'lettuce'},
                {variety: 'Pirat Lettuce', group: 'lettuce'}]
            } else {
                console.log(error)
            }
})
    




    return {availTypes}
}