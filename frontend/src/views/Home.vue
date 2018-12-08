<template>
  <div class="home">
    <div class="columns">
      <div class="column is-one-fifth"></div>
      <div class="column is-three-fifths">
        <h1 class="has-text-centered is-size-3">Welcome to Brananagrams</h1>
        <div class="columns is-centered">
          <div class="column is-two-fifths">
            <img src="@/assets/bananagrams.jpg">
          </div>
        </div>
        <div class="field has-addons new-game-input">
          <div class="control">
            <input class="input" v-model="gameName" type="text" placeholder="New Game">
          </div>
          <div class="control">
            <a class="button is-info" @click="createGame">
              Create Game
            </a>
          </div>
        </div>
        <div>
          <h2 class="has-text-centered is-size-3">Game List</h2>
          <ol>
            <li v-for="game in games">
              <a @click="checkLogin(game.getId())">{{game.getName()}}</a>
              <span> ({{gameStatus(game)}}, {{game.getPlayerCount()}} joined)</span>
            </li>
          </ol>
        </div>
      </div>
      <div class="column is-one-fifth"></div>
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
                  v-model="username"
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
    this.$cookies.set(`game-${this.modalGameID}`, this.username);
    this.$router.push({ name: 'game', params: { id: this.modalGameID } });
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
.new-game-input {
  justify-content: center;
  margin-bottom: 1rem;
}
</style>
