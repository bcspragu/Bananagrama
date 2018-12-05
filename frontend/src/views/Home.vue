<template>
  <div class="home">
    <h1>This is the home page</h1>
    <input v-model="gameName">
    <button @click="createGame">New Game</button>
    <div>
      <ol>
        <li v-for="game in games">
          <router-link :to="{ name: 'game', params: { 'id': game.getId() } }">
            {{game.getName()}}
          </router-link>
        </li>
      </ol>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

// Import code-generated data structures.
import {NewGameRequest, ListGamesRequest, Game} from '@/proto/banana_pb';

@Component
export default class Home extends Vue {
  private gameName: string = '';
  private games: Game[] = [];

  private mounted(): void {
    this.loadGames();
  }

  private createGame() {
    const req = new NewGameRequest();
    req.setName(this.gameName);
    this.$client.newGame(req, (err, resp) => {
      if (resp) {
        this.loadGames();
      }
    });
  }

  private loadGames() {
    this.$client.listGames(new ListGamesRequest(), (err, resp) => {
      if (resp) {
        this.games = resp.getGamesList();
      }
    });
  }
}
</script>

<style scoped>
li {
  text-align: left;
}
</style>
