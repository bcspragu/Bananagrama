<template>
  <div class="home">
    <h1>This is the home page</h1>
    <input v-model="gameName">
    <button @click="createGame">New Game</button>
    <div>
      <ol>
        <li v-for="game in games">
          <a @click="checkLogin(game.getId())">{{game.getName()}}</a>
          <span> ({{gameStatus(game)}}, {{game.getPlayerCount()}} joined)</span>
        </li>
      </ol>
    </div>
    <b-modal :active.sync="showModal">
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Join Game</p>
        </header>
        <section class="modal-card-body">
          <b-field label="Username">
              <b-input
                  type="text"
                  :value="username"
                  placeholder="How Now, Brown Steer?"
                  required>
              </b-input>
          </b-field>
        </section>
        <footer class="modal-card-foot">
          <button type="button" @click="showModal = false" class="button">Close</button>
          <button class="button is-primary" @click="joinGame">Join</button>
        </footer>
      </div>
    </b-modal>
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

  private showModal = false;
  private modalGameID: string = '';
  private username: string = '';

  private mounted(): void {
    this.loadGames();
  }

  private checkLogin(id: string): void {
    const name = this.$cookies.get(`game-${id}`);
    if (name) {
      this.$router.push({ name: 'game', params: { id } });
      return;
    }

    // Show a login modal.
    this.modalGameID = id;
    this.showModal = true;
  }

  private joinGame(): void {
    this.showModal = false;
    const req = new JoinGameRequest();
    req.setId(this.modalGameID);
    req.setName(this.username);

    this.$client.joinGame(req, (err, resp) => {
      if (!resp) {
        console.log(err);
        return;
      }

      this.$router.push({ name: 'game', params: { id: this.modalGameID } });
    });
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
      if (!resp) {
        console.log(err);
        return;
      }
      this.games = resp.getGamesList();
    });
  }

  private gameStatus(game: Game): string {
    switch (game.getStatus()) {
      case Game.Status.WAITING_FOR_PLAYERS:
        return 'waiting for players';
      case Game.Status.IN_PROGRESS:
        return 'in progress';
      case Game.Status.FINISHED:
        return 'finished';
      default:
        return 'unknown';
    }
  }
}
</script>

<style scoped>
li {
  text-align: left;
}
</style>
