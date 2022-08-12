import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import axios  from 'axios'
import './assets/global.css'

createApp(App).use(store).use(router).mount('#app')
global.debug = true

axios.defaults.baseURL = 'http://localhost:5000';

