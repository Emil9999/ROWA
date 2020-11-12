import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import Vue2TouchEvents from 'vue2-touch-events'
import * as FullStory from '@fullstory/browser';

Vue.use(Vue2TouchEvents)

FullStory.init({  orgId: 'ZAA45'});
Vue.prototype.$FullStory =  FullStory;

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
