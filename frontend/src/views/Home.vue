<template>
  <div class="home">
    <div class="header-image">
      <img class="logo" src="@/assets/logo.svg">
      <h2 v-if="playerName" class="has-text-centered is-size-5">Playing as {{playerName}} <sup><a class="is-size-6" @click="changeName">(change name)</a></sup></h2>
    </div>

    <div class="columns is-gapless">

      <div class="column is-two-thirds">
        <div>
          <h1 class="has-text-centered is-size-4">Welcome to B(r)ananagrams!</h1>

          <div class="description">
            <p>
              <strong>What is this?</strong> Brananagrams is my internet version/discount knockoff of
              smash-hit <a href="https://bananagrams.com/">Bananagrams</a>,
              which itself is just <a
              href="https://scrabble.hasbro.com/en-us">Scrabble</a> for people
              with short attention spans.
            </p>

            <p>
              <strong>How do I play?</strong> Create a new game, select it from the
              game list, find at least one other friend (and up to 23 other
              friends), and hit 'Start Game'. Tiles will be distributed. Add
              words to your board by typing them out, entry locations will be
              automatically suggested. If a location isn't suggested
              automatically, you can click tiles in your hand and click to
              place them on the grid. Double-click a tile to dump, and click
              letters on the board to remove them.
            </p>

            <p>
              <strong>Can I contribute?</strong> Totally! If you'd like to
              contribute, the code for this project is <a
                href="https://github.com/bcspragu/Bananagrama">on GitHub</a>.
            </p>

            <p>
              <strong>Can I sue you?</strong> Definitely! If you're Hasbro or
              whatever board game conglomerate owns Bananagrams, and would like
              to send a cease and desist letter, my email address is: 
            </p>

            <p class="has-text-centered">
              <a href="mailto:imshakinginmyboots@merrychristmasfuckers.com">
                ImShakingInMyBoots@MerryChristmasFuckers.com
              </a>
            </p>
            <p class="has-text-centered super-serious">
              (legitimately though, I will receive an email if you send it there)
            </p>
          </div>
        </div>
      </div>

      <div class="vertical-rule column is-narrow"></div>

      <div class="column">
        <h1 class="has-text-centered is-size-4">Create a Game</h1>
        <NewGameForm/>
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
                  @keyup.native.enter="registerPlayer"
                  required>
              </b-input>
          </b-field>
        </section>
        <footer class="modal-card-foot">
          <button type="button" @click="showModal = false" class="button">Close</button>
          <button class="button is-primary" @click="registerPlayer">Set Name</button>
        </footer>
      </div>
    </b-modal>

  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import * as grpcWeb from 'grpc-web';

import NewGameForm from '@/components/NewGameForm.vue';
import {RegisterRequest, NewGameRequest, ListGamesRequest,
        Game as PBGame, GameStatus} from '@/proto/banana_pb';

interface Game {
  id: string;
  name: string;
  playerCount: number;
  stat: GameStatus;
}

@Component({
  components: {
    NewGameForm,
  },
})
export default class Home extends Vue {
  private games: Game[] = [];

  private showModal = false;
  private playerName: string = '';

  private mounted(): void {
    const playerName = this.$cookies.get('player-name');
    if (playerName) {
      this.playerName = playerName;
      this.loadGames();
    } else {
      // Have them set their name.
      this.showModal = true;
    }

  }

  private registerPlayer(): void {
    if (!this.playerName) {
      // If they didn't set a name, just leave the dialog open.
      return;
    }
    const req = new RegisterRequest();
    req.setName(this.playerName);
    this.$client.register(req, {}, (err, resp) => {
      if (err) {
        console.log(err);
        return;
      }
      this.showModal = false;
      this.$cookies.set('player-name', this.playerName);
      this.$cookies.set('player-id', resp.getPlayerId());
      this.$cookies.set('player-token', resp.getToken());
      this.loadGames();
    });
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
    stream.on('error', (err: grpcWeb.Error) => {
      this.$emit('rpcError', err);
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
.columns.is-gapless > .column.vertical-rule {
  background-color: whitesmoke;
  width: 2px;
  margin: 0.5rem 1.5rem 0 1.5rem;
}

.description {
  margin: 0.5rem 1.5rem;
}

.description p {
  margin-top: 0.5rem;
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

.super-serious {
  font-size: 0.5rem;
}
</style>
