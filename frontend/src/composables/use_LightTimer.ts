import axios from 'axios'
import timeString from '../types/timeString'
import {ref, reactive} from 'vue'
import { TimerFunctions } from './useTimerFunctions'

export function lightTimes() {

        const StartTime = ref<timeString>({minutes: '12', hours: '10'})
        const EndTime = ref<timeString>({minutes: '32', hours: '20'})

        const getTimes = () =>{
            return
            axios.get('')
        }
        getTimes()

        const nStartTime =  reactive(TimerFunctions()) 
        const nEndTime =  reactive(TimerFunctions())

        const nTimes = [reactive(TimerFunctions()), reactive(TimerFunctions())]

        const sendTimes = (sendStartTime: timeString = nTimes[0].time, sendEndTime: timeString = nTimes[1].time) =>{
            console.log(sendStartTime)
            console.log(sendEndTime)
            return
            axios.post('')
            getTimes()
        }




    return{StartTime, EndTime, sendTimes, nStartTime, nEndTime, nTimes}
} 