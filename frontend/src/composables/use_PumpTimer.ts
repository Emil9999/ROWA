import axios from 'axios'
import timeString from '../types/timeString'
import {ref, reactive} from 'vue'
import { TimerFunctions } from './useTimerFunctions'

export function pumpTimes() {

        const StartTime = ref<timeString>({minutes: '12', hours: '10'})
        const EndTime = ref<timeString>({minutes: '32', hours: '20'})
        const PumpDuration = ref<number>(20)

        const getTimes = () =>{
          
            axios.get('/admin/pump/times').then((r) =>
            {
                const onTime = r.data.onTime.slice(-8, -3)

                const offTime = r.data.offTime.slice(-8, -3)

                StartTime.value.hours = onTime.slice(0,2)
                StartTime.value.minutes = onTime.slice(3,5)

                EndTime.value.hours = offTime.slice(0,2)
                EndTime.value.minutes =  offTime.slice(3,5)

                PumpDuration.value = r.data.duration
            })
            
        }
        getTimes()

        const nTimes =  [reactive(TimerFunctions()), reactive(TimerFunctions())]
        const nPumpDuration = ref<number>(PumpDuration.value)

        
        const sendTimes = (sendStartTime: timeString = nTimes[0].time, sendEndTime: timeString = nTimes[1].time, sendPumpDuration: number = nPumpDuration.value) =>{
            axios.post('/admin/pump/times', 
                {"onTime": sendStartTime.hours + ':' + sendStartTime.minutes,
                "offTime": sendEndTime.hours + ':' + sendEndTime.minutes,
                "duration": sendPumpDuration
                }
            ).then(() => getTimes())
         
            
        }




    return{StartTime, EndTime, PumpDuration, sendTimes, nTimes, nPumpDuration}
} 