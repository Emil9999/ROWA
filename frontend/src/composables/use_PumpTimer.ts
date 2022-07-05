import axios from 'axios'
import timeString from '../types/timeString'
import {ref, reactive} from 'vue'
import { TimerFunctions } from './useTimerFunctions'

export function pumpTimes() {

        const StartTime = ref<timeString>({minutes: '12', hours: '10'})
        const PumpDuration = ref<number>(20)

        const getTimes = () =>{
            return
            axios.get('')
        }
        getTimes()

        const nTime =  ref(TimerFunctions()) 
        const nPumpDuration = ref<number>(PumpDuration.value)

        
        const sendTimes = (sendStartTime: timeString = nTime.value.time, sendPumpDuration: number = nPumpDuration.value) =>{
            console.log(sendStartTime)
            console.log(sendPumpDuration)
            return
            axios.post('')
            getTimes()
        }




    return{StartTime, PumpDuration, sendTimes, nTime, nPumpDuration}
} 