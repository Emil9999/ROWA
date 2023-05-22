import timeString from '../types/timeString'
import {ref} from 'vue'

export function TimerFunctions(Stime: timeString = {minutes: '', hours: ''}) {

    const time = ref(Stime)

    const addNumber = (enteredNumber: string) =>{
        if( time.value.hours.length < 2){
            if(parseInt(time.value.hours + enteredNumber) < 24){
               if(time.value.hours.length == 0 && parseInt(enteredNumber) > 2){
                  console.log("Added 0")
                  time.value.hours += "0"
                  time.value.hours += enteredNumber} else {
                  time.value.hours += enteredNumber   
                  }
            }
         }
         else if( time.value.minutes.length < 2){
            if(parseInt(time.value.minutes + enteredNumber) < 60){
               if(time.value.minutes.length == 0 && parseInt(enteredNumber) > 5){
                  console.log("Added 0")
                  time.value.minutes += "0"
                  time.value.minutes += enteredNumber} else {
                  time.value.minutes += enteredNumber   
                  }
            }
            
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