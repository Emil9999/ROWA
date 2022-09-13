import axios from 'axios'
import timeString from '../types/timeString'
import {ref} from 'vue'
import { TimerFunctions } from './useTimerFunctions'

export function pumpTimes() {

        const StartTime = ref<timeString>({minutes: '12', hours: '10'})
        const PumpDuration = ref<number>(20)

        const getTimes = () =>{
          
            axios.get('/admin/pump/times').then((r) =>
            {
                const onTime = r.data.onTime.slice(-8, -3)

                StartTime.value.hours = onTime.slice(0,2)
                StartTime.value.minutes = onTime.slice(3,5)
                PumpDuration.value = r.data.duration
            })
            
        }
        getTimes()

        const nTime =  ref(TimerFunctions()) 
        const nPumpDuration = ref<number>(PumpDuration.value)

        
        const sendTimes = (sendStartTime: timeString = nTime.value.time, sendPumpDuration: number = nPumpDuration.value) =>{
            axios.post('/admin/pump/times', 
                {"onTime": sendStartTime.hours + ':' + sendStartTime.minutes,
                "duration": sendPumpDuration
                }
            ).then(() => getTimes())
         
            
        }




    return{StartTime, PumpDuration, sendTimes, nTime, nPumpDuration}
} 