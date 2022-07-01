import timeString from '../types/timeString'
import {ref} from 'vue'

export function TimerFunctions(Stime: timeString = {minutes: '', hours: ''}) {

    const time = ref(Stime)

    const addNumber = (enteredNumber: string) =>{
        if( time.value.hours.length < 2){
            time.value.hours += enteredNumber
         }
         else if( time.value.minutes.length < 2){
            time.value.minutes += enteredNumber
         }
     }

     const overrideTime = (override: timeString) =>{

        time.value = override
     }

     const removeNumber = () => {
         if (time.value.minutes.length > 0){
            time.value.minutes = time.value.minutes.slice(0, -1)
         }
         else if (time.value.hours.length > 0){
            time.value.hours = time.value.hours.slice(0, -1)
         }

     }

    return{time, addNumber, removeNumber, overrideTime}
}