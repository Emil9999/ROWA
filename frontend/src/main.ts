import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios  from 'axios'
import Vue3TouchEvents from 'vue3-touch-events'
import './assets/global.css'

createApp(App).use(router).use(Vue3TouchEvents, {
	swipeTolerance: 500
}).mount('#app')
global.debug = false

axios.defaults.baseURL = 'http://localhost:5000';

