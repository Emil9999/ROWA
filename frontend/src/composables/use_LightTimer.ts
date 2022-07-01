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

        const sendTimes = () =>{
            console.log(nTimes[0].time)
            console.log(nTimes[1].time)
            return
            axios.post('')
        }




    return{StartTime, EndTime, sendTimes, nStartTime, nEndTime, nTimes}
} 