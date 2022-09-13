import axios from 'axios'
import timeString from '../types/timeString'
import {ref, reactive} from 'vue'
import { TimerFunctions } from './useTimerFunctions'

export function lightTimes() {

        const StartTime = ref<timeString>({minutes: '12', hours: '10'})
        const EndTime = ref<timeString>({minutes: '32', hours: '20'})

        const getTimes = () =>{

            axios.get('/admin/light/times').then((r) =>
            {   
                const onTime = r.data.onTime.slice(-8, -3)
                const offTime = r.data.offTime.slice(-8, -3)

                StartTime.value.hours = onTime.slice(0,2)
                StartTime.value.minutes =  onTime.slice(3,5)
                EndTime.value.hours = offTime.slice(0,2)
                EndTime.value.minutes =  offTime.slice(3,5)
            })
            
        }
        getTimes()

        const nTimes = [reactive(TimerFunctions()), reactive(TimerFunctions())]

        const sendTimes = (sendStartTime: timeString = nTimes[0].time, sendEndTime: timeString = nTimes[1].time) =>{
            
            axios.post('/admin/light/times', 
                {"onTime": sendStartTime.hours + ':' + sendStartTime.minutes,
                "offTime": sendEndTime.hours + ":" + sendEndTime.minutes
                }
            ).then(() => getTimes())
            
            }
        




    return{StartTime, EndTime, sendTimes, nTimes}
} 