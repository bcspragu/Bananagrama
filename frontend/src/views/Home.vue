<template>
  <div class="home">
    <div class="columns">
      <div class="column is-one-fifth"></div>
      <div class="column is-three-fifths">
        <h1 class="has-text-centered is-size-3">Welcome to Brananagrams</h1>
        <h2 v-if="playerName" class="has-text-centered is-size-5">Playing as {{playerName}} <sup><a class="is-size-6" @click="changeName">(change name)</a></sup></h2>
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
              <a @click="joinGame(game.getId())">{{game.getName()}}</a>
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
          <p class="modal-card-title">Set your name</p>
        </header>
        <section class="modal-card-body">
          <b-field label="Username">
              <b-input
                  type="text"
                  v-model="playerName"
                  placeholder="How Now, Brown Steer?"
                  required>
              </b-input>
          </b-field>
        </section>
        <footer class="modal-card-foot">
          <button type="button" @click="showModal = false" class="button">Close</button>
          <button class="button is-primary" @click="setPlayerName">Set Name</button>
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
  private playerName: string = '';

  private mounted(): void {
    this.loadGames();

    const playerName = this.$cookies.get('player-name');
    if (playerName) {
      this.playerName = playerName;
    } else {
      // Have them set their name.
      this.showModal = true;
    }

  }

  private setPlayerName(): void {
    if (!this.playerName) {
      // If they didn't set a name, just leave the dialog open.
      return;
    }
    this.showModal = false;
    this.setPlayerNameInCookies(this.playerName);
  }

  private changeName(): void {
    this.showModal = true;
  }

  private joinGame(id: string): void {
    const playerName = this.playerNameFromCookies();
    if (playerName) {
      this.$router.push({ name: 'game', params: { id } });
      return;
    }

    // Show a login modal.
    this.showModal = true;
  }

  private playerNameFromCookies(): string | undefined {
    return this.$cookies.get('player-name');
  }

  private setPlayerNameInCookies(name: string) {
    this.$cookies.set('player-name', name);
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
