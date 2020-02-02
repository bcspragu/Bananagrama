<template>
  <div id="app">
    <router-view
        v-on:rpcError="handleError"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import * as grpcWeb from 'grpc-web';

import {RegisterRequest} from '@/proto/banana_pb';

@Component
export default class App extends Vue {
  private handleError(err: any): void {
    if (err.code === 16) {
      // Clear our user info, since it's clearly trash.
      this.$cookies.remove('player-id');
      this.$cookies.remove('player-token');

      const playerName = this.$cookies.get('player-name');

      if (playerName) {
        this.registerPlayer(playerName);
      } else {
        if (this.$router.currentRoute.name !== 'home') {
          this.$router.push({ name: 'home' });
        }
        console.log(`user unauthenticated`);
      }
    }
  }

    private registerPlayer(playerName: string): void {
      if (!playerName) {
        return;
      }

      const req = new RegisterRequest();
      req.setName(playerName);
      this.$client.register(req, {}, (err, resp) => {
        if (err) {
          console.log(err);
          return;
        }
        this.$cookies.set('player-name', playerName);
        this.$cookies.set('player-id', resp.getPlayerId());
        this.$cookies.set('player-token', resp.getToken());

        if (this.$router.currentRoute.name !== 'home') {
          this.$router.push({ name: 'home' });
        }
      });
    }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

html, body, #app {
  margin: 0;
  height: 100%;
}
</style>
