import Plant from "@/types/Plant";
import {ref} from 'vue'


export default function findUniqueTypes(module: Array<Plant>){
    
        const Uniquevarietys= ref<Array<string>>([])
        module.forEach(element => {
            if(!Uniquevarietys.value.includes(element.variety) && element.variety != ''){Uniquevarietys.value.push(element.variety)}
        })
        
        return (Uniquevarietys)
}
