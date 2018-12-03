<template>
  <div class="home">
    <h1>This is the home page</h1>
    <ol>
      <li v-for="game in games">
        {{game.Name}}
      </li>
    </ol>
    <input v-model="gameName">
    <button @click="createGame">New Game</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

// Import code-generated data structures.
import {BananaServiceClient} from '@/proto/banana_pb_service';
import {NewGameRequest, ListGamesRequest} from '@/proto/banana_pb';

@Component
export default class Home extends Vue {
  private gameName: string = '';
  private games: string[] = [];
  private client: BananaServiceClient | null = null;

  private mounted(): void {
    this.client = new BananaServiceClient('/api');
    this.client.listGames(new ListGamesRequest(), (resp, err) => {
      console.log('hayyo', resp, err);
    });
  }

  private createGame() {
    const req = new NewGameRequest();
    req.setName(this.gameName);
    this.client!.newGame(req, (resp, err) => {
      console.log('HAYYO', resp, err);
    });
  }
}
</script>

<style scoped>
</style>
