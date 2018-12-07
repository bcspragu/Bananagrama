import Vue from 'vue';
import VueCookies from 'vue-cookies';

import Buefy from 'buefy';
import 'buefy/dist/buefy.css';

import App from './App.vue';
import router from './router';


import {BananaServiceClient} from '@/proto/banana_pb_service';

Vue.use(VueCookies);
Vue.use(Buefy);

Vue.config.productionTip = false;

Vue.prototype.$client = new BananaServiceClient('');
Vue.prototype.$cookies.config(Infinity);

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');
