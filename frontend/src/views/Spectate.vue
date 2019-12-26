<template>
  <div class="spectator-board">
    <div class="board" v-for="board in boards"  :style="{'background-color': board.color}">
      <div class="label">{{board.playerName}}</div>
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
  private boards: Board[] = [];

  private mounted(): void {
    this.spectateGame();
  }

  private spectateGame(): void {
    const req = new SpectateRequest();
    req.setId(this.$route.params.id);

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
.spectator-board {
  display: flex;
  flex-wrap: wrap;
  height: 100%;
}

.board {
  max-width: 50%;
  max-height: 50%;

  min-width: 25%;
  min-height: 25%;

  padding: 0.5rem;

  border: 1px solid black;
}

.label {
  position: absolute;
}
</style>
