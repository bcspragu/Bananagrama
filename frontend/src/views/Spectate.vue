<template>
  <div id="spectator-board">
    <div class="board" v-for="board in boards">
      <BoardView 
        :board="board.board"
        :invalidWords="board.invalidWords"
        :detached="board.detached"/>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import BoardView from '@/components/BoardView.vue'; // @ is an alias to /src
import {Game as PBGame, Board as PBBoard, ListGamesRequest,
        SpectateRequest, SpectateUpdate, CharLocs} from '@/proto/banana_pb';

interface Board {
  playerID: string;
  playerName: string;
  board: PBBoard;
  invalidWords: CharLocs[];
  detached: boolean;
}

@Component({
  components: {
    BoardView,
  },
})
export default class Spectate extends Vue {
  private game: PBGame | null = null;
  private boards: Board[] = [];

  private mounted(): void {
    this.$client.listGames(new ListGamesRequest(), {}, (err, resp) => {
      if (!resp) {
        console.log(err);
        return;
      }
      for (const game of resp.getGamesList()) {
        if (game.getId() === this.$route.params.id) {
          this.game = game;
          break;
        }
      }

      if (!this.game || this.game.getStatus() === PBGame.Status.WAITING_FOR_PLAYERS) {
        this.$router.push({ name: 'home' });
        console.log(`Couldn't find game ID ${this.$route.params.id}, or it hadn't started yet`);
        return;
      }

      this.spectateGame();
    });
  }

  private spectateGame(): void {
    const req = new SpectateRequest();
    req.setId(this.game!.getId());

    const stream = this.$client.spectate(req, {});

    stream.on('data', (resp) => {
      const indexMap: { [s: string]: number; } = {};
      let i = 0;
      for (const board of this.boards) {
        indexMap[board.playerID] = i;
        i++;
      }

      // Game already exists, update it.
      if (indexMap.hasOwnProperty(resp.getPlayerId())) {
        Vue.set(this.boards, indexMap[resp.getPlayerId()], this.toBoard(resp));
        return;
      }

      // Board doesn't exist yet, add it.
      this.boards.push(this.toBoard(resp));
    });
  }

  private toBoard(resp: SpectateUpdate): Board {
    return {
      playerID: resp.getPlayerId(),
      playerName: resp.getPlayerName(),
      board: resp.getBoard(),
      invalidWords: resp.getInvalidWordsList(),
      detached: resp.getDetachedBoard(),
    };
  }
}
</script>

<style scoped>
.board {
  width: 350px;
  height: 350px;
}
</style>
