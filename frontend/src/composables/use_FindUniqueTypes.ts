import Plant from "@/types/Plant";
import {ref} from 'vue'


export default function findUniqueTypes(module: Array<Plant>){
    
        const Uniquevarietys= ref<Array<string>>([])
        module.forEach(element => {
            if(!Uniquevarietys.value.includes(element.plant_type) && element.plant_type != ''){Uniquevarietys.value.push(element.plant_type)}
        })
        
        return (Uniquevarietys)
}
