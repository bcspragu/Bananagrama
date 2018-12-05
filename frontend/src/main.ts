import Vue from 'vue';
import App from './App.vue';
import router from './router';

import {BananaServiceClient} from '@/proto/banana_pb_service';

Vue.config.productionTip = false;

Vue.prototype.$client = new BananaServiceClient('');

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');
