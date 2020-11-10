import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import Vue2TouchEvents from 'vue2-touch-events'

Vue.use(Vue2TouchEvents)
export const eventBus = new Vue();
Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
