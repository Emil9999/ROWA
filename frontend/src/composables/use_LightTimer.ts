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
                StartTime.value.hours = r.data.lightOn.slice(0,2)
                StartTime.value.minutes =  r.data.lightOn.slice(3,5)
                EndTime.value.hours = r.data.lightOff.slice(0,2)
                EndTime.value.minutes =  r.data.lightOff.slice(3,5)
            })
            
        }
        getTimes()

        const nTimes = [reactive(TimerFunctions()), reactive(TimerFunctions())]

        const sendTimes = (sendStartTime: timeString = nTimes[0].time, sendEndTime: timeString = nTimes[1].time) =>{
            
            axios.post('/admin/light/times', 
                {"lightOn": sendStartTime.hours + ':' + sendStartTime.minutes,
                "lightOff": sendEndTime.hours + ":" + sendEndTime.minutes
                }
            )
            getTimes()
            }
        




    return{StartTime, EndTime, sendTimes, nTimes}
} 