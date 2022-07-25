
import axios from 'axios'
import {ref} from 'vue'
import AvailTypes from '../types/AvailVariety'


export default function getAllTypes(){

    const availTypes = ref<AvailTypes[]>([])
    if (global.debug)
    {
        availTypes.value = [{variety: 'Thai Basil', group: 'herb'},
                            {variety: 'Basil', group: 'herb'},
                            {variety: 'Mint', group: 'herb'},
                            {variety: 'Thyme', group: 'herb'},
                            {variety: 'Cilantro', group: 'herb'},
                            {variety: 'Lollo Bionda', group: 'lettuce'},
                            {variety: 'Neckar Giant', group: 'lettuce'},
                            {variety: 'Pirat Lettuce', group: 'lettuce'}]
    } else
    {
        axios.get('route').then()
    }




    return {availTypes}
}