import Vue from 'vue';
import Router from 'vue-router';
import Game from './views/Game.vue';
import Home from './views/Home.vue';
import Spectate from './views/Spectate.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/game/:id',
      name: 'game',
      component: Game,
    },
    {
      path: '/spectate/:id',
      name: 'spectate',
      component: Spectate,
    },
  ],
});
