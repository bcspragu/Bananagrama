<template>
  <div class="home">
    <div class="columns">
      <div class="column is-one-fifth"></div>
      <div class="column is-three-fifths">
        <div>
          <h1 class="has-text-centered is-size-3">Welcome to B(r)ananagrams</h1>
          <img class="logo" src="@/assets/logo.svg">
        </div>
        <hr>
        <h2 v-if="playerName" class="has-text-centered is-size-5">Playing as {{playerName}} <sup><a class="is-size-6" @click="changeName">(change name)</a></sup></h2>
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
        <hr>
        <div>
          <h2 class="has-text-centered is-size-3">Game List</h2>
          <table class="games">
            <tr>
              <td>Game Name</td>
              <td>Status</td>
              <td>Number of Players</td>
              <td></td>
            </tr>
            <tr class="game-row" v-for="game in games">
              <td><a @click="joinGame(game.id)">{{game.name}}</a></td>
              <td>{{gameStatus(game.stat)}}</td>
              <td>{{game.playerCount}} players</td>
              <td><a @click="spectate(game.id)">Spectate</a></td>
            </tr>
          </table>
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
                  placeholder="e.g. &quot;Cave Johnson&quot;"
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
import {NewGameRequest, ListGamesRequest, Game as PBGame, GameStatus} from '@/proto/banana_pb';

interface Game {
  id: string;
  name: string;
  playerCount: number;
  stat: GameStatus;
}

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

  private spectate(id: string): void {
    this.$router.push({ name: 'spectate', params: { id } });
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
    this.$client.newGame(req, {}, (err, resp) => {
      if (err) {
        console.log(err);
      }
    });
  }

  private loadGames() {
    const stream = this.$client.streamGames(new ListGamesRequest(), {});

    stream.on('data', (resp) => {
      const indexMap: { [s: string]: number; } = {};
      let i = 0;
      for (const game of this.games) {
        indexMap[game.id] = i;
        i++;
      }

      const gameUpdates = resp.getGamesList();
      for (const game of gameUpdates) {
        // Game already exists, update it.
        if (indexMap.hasOwnProperty(game.getId())) {
          Vue.set(this.games, indexMap[game.getId()], this.toGame(game));
          continue;
        }

        // Game doesn't exist yet, add it.
        this.games.push(this.toGame(game));
      }
    });
  }

  private gameStatus(stat: GameStatus): string {
    switch (stat) {
      case GameStatus.WAITING_FOR_PLAYERS:
        return 'Waiting for players';
      case GameStatus.IN_PROGRESS:
        return 'In progress';
      case GameStatus.FINISHED:
        return 'Finished';
      default:
        return 'Unknown';
    }
  }

  private toGame(g: PBGame): Game {
    return {
      id: g.getId(),
      name: g.getName(),
      playerCount: g.getPlayerCount(),
      stat: g.getStatus(),
    };
  }
}
</script>

<style scoped>
.new-game-input {
  justify-content: center;
  margin-bottom: 1rem;
}
.games {
  border-collapse: collapse;
  width: 100%;
}
.games tr { 
  border: solid;
  border-width: 1px 0;
  padding-bottom: 2rem;
}
.games tr:first-child {
  border-top: none;
}
.games tr:last-child {
  border-bottom: none;
}
.logo {
  display: block;
  margin: 0 auto;
}
</style>
