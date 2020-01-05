<template>
  <div class="game">
    <div class="columns is-gapless left-side">
      <div class="column is-four-fifths">
        <div class="board">
          <Board ref="board"
            :gameOver="gameOver"
            v-on:removeTile="addTileToHand"
            v-on:blankClicked="addSelected"
            v-on:boardUpdated="sendBoard"/>
        </div>
        <div class="active-word columns is-gapless">
          <div class="column is-3 active-display">
            <ActiveWord ref="activeWord" :word="word"/>
          </div>
          <div class="column is-9 notice-msg">
            <Notice ref="notice" :notice="notice"/>
          </div>
        </div>
        <div class="letters">
          <UnusedLetters v-if="letters" ref="hand"
            :letters="letters"
            :gameOver="gameOver"
            v-on:dumpTile="dumpTile"/>
        </div>
      </div>
      <div class="column is-one-fifth right-side">
        <div class="players">
          <div v-if="tilesInBunch !== null"><strong>Tiles left in bunch: {{tilesInBunch}}</strong></div>
          <hr/>
          <table class="standings">
            <tr class="standing-header">
              <td>Player</td>
              <td class="player-scores">Board</td>
              <td class="player-scores">Hand</td>
            </tr>
            <tr v-for="player in players">
              <td>{{player.name}}</td>
              <td class="player-scores">{{player.tilesInBoard}}</td>
              <td class="player-scores">{{player.tilesInHand}}</td>
            </tr>
          </table>
          <hr/>
          <div class="start-game-button" v-if="waitingForPlayers">
            <a class="button is-info" @click="startGame">
              Start Game
            </a>
          </div>
        </div>
        <div class="logs">
          <div v-for="log in logs">
            {{log}}
            <br>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import * as grpcWeb from 'grpc-web';
import { Component, Vue } from 'vue-property-decorator';

import Board from '@/components/Board.vue'; // @ is an alias to /src
import UnusedLetters from '@/components/UnusedLetters.vue'; // @ is an alias to /src
import ActiveWord from '@/components/ActiveWord.vue'; // @ is an alias to /src
import Notice from '@/components/Notice.vue'; // @ is an alias to /src
import {Cell, Letter, PlacedWord, Orientation, Player} from '@/data';
import {Game as PBGame, Board as PBBoard, ListGamesRequest, JoinGameRequest,
        GameUpdate, CurrentStatus, PlayerUpdate, GameStarted, GameEnded, StartGameRequest,
        Player as PBPlayer, DumpRequest, UpdateBoardRequest, Word, Tiles, GameStatus,
        StreamLogsRequest, LogEntry, OtherTileUpdates, OtherTileUpdate} from '@/proto/banana_pb';

@Component({
  components: {
    Board,
    ActiveWord,
    UnusedLetters,
    Notice,
  },
})
export default class Game extends Vue {
  private game: PBGame | null = null;
  private playerID: string | null = null;
  private players: Player[] = [];
  private tilesInBunch: number | null = null;
  private logs: string[] = [];
  private gameStatus: GameStatus = GameStatus.UNKNOWN;

  private word: string = '';
  private notice: string[] = [];
  private board: Board = new Board();
  private activeWord: ActiveWord = new ActiveWord();
  private detachedBoard: boolean = false;

  // The component that renders our hand.
  private hand: UnusedLetters = new UnusedLetters();

  // Letters in our hand.
  private letters: Letter[] = [];
  // Letters required to complete the word on the current board.
  private requiredLetters: string[] = [];

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

      if (!this.game) {
        this.$router.push({ name: 'home' });
        console.log(`Couldn't find game ID ${this.$route.params.id}`);
        return;
      }

      this.joinGame();
    });

    this.board = (this.$refs.board as Board);
    this.hand = (this.$refs.hand as UnusedLetters);

    document.addEventListener('keyup', this.keyup);
  }

  private startGame(): void {
    const req = new StartGameRequest();
    req.setId(this.game!.getId());

    this.$client.startGame(req, {}, (err, resp) => {
      if (!resp) {
        console.log(err);
        return;
      }
    });
  }

  private joinGame(): void {
    if (!this.game) {
      return;
    }

    const gameID = this.game.getId();
    const playerID = this.$cookies.get('player-id');
    const req = new JoinGameRequest();
    req.setGameId(gameID);
    req.setPlayerId(playerID);

    const stream = this.$client.joinGame(req, {});

    stream.on('data', (resp) => {
      switch (resp.getUpdateCase()) {
        case GameUpdate.UpdateCase.CURRENT_STATUS:
          this.handleInitialGameState(resp.getCurrentStatus()!);
          break;
        case GameUpdate.UpdateCase.PLAYER_UPDATE:
          this.handlePlayerUpdate(resp.getPlayerUpdate()!);
          break;
        case GameUpdate.UpdateCase.GAME_STARTED:
          this.handleGameStarted(resp.getGameStarted()!);
          break;
        case GameUpdate.UpdateCase.GAME_ENDED:
          this.handleGameOver(resp.getGameEnded()!);
          break;
        case GameUpdate.UpdateCase.SELF_TILE_UPDATE:
          const tileUpdate = resp.getSelfTileUpdate()!;
          if (tileUpdate.getFromOtherPeel()) {
            this.hand.flash();
          }
          this.updateTiles(tileUpdate.getAllTiles());
          break;
        case GameUpdate.UpdateCase.OTHER_TILE_UPDATE:
          this.updateTileCounts(resp.getOtherTileUpdate());
      }
    });
    stream.on('status', (status: grpcWeb.Status) => {
      console.log('status');
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });
    stream.on('error', (err: grpcWeb.Error) => {
      console.log('ERROR', err);
    });
    stream.on('end', () => {
      console.log('closed');
      if (!this.gameOver) {
        // Wait a second, then try joining again.
        window.setTimeout(() => this.joinGame(), 1000);
      }
    });
  }

  private streamLogs(pID: string): void {
    if (!this.game) {
      return;
    }

    const req = new StreamLogsRequest();
    req.setGameId(this.game.getId());
    req.setPlayerId(pID);

    const stream = this.$client.streamLogs(req, {});

    stream.on('data', (resp: LogEntry) => {
      switch (resp.getEventCase()) {
        case LogEntry.EventCase.PLAYER_MOVE:
          const move = resp.getPlayerMove();
          let isWord = 'word';
          if (!move.getWordValid()) {
            isWord = 'not-a-word';
          }
          if (move.getWord().length > 0) {
            this.logs.unshift(`${move.getPlayerName()} played ${isWord} ${move.getWord()}.`);
            if (!move.getBoardConnected()) {
              this.logs.unshift(`${move.getPlayerName()}'s board is in shambles.`);
            }
          }
          break;
        case LogEntry.EventCase.PLAYER_DUMP:
          const dump = resp.getPlayerDump();
          this.logs.unshift(`${dump.getPlayerName()} dumped.`);
          break;
        case LogEntry.EventCase.PLAYER_PEEL:
          const peel = resp.getPlayerPeel();
          this.logs.unshift(`${peel.getPlayerName()} peeled.`);
      }
    });
  }

  private handleInitialGameState(up: CurrentStatus): void {
    this.playerID = up.getYourId();
    for (const p of up.getPlayersList()) {
      this.players.push({
        id: p.getId(),
        name: p.getName(),
        tilesInHand: p.getTilesInHand(),
        tilesInBoard: p.getTilesInBoard(),
      });
    }
    this.tilesInBunch = up.getRemainingTiles();
    this.updateBoard(up.getBoard()!);
    this.updateTiles(up.getAllTiles()!);
    this.gameStatus = up.getStatus();
    if (up.getStatus() === GameStatus.FINISHED) {
      this.hand.clearSelected();
      const ge = new GameEnded();
      this.handleGameOver(ge);
    }
    this.streamLogs(this.playerID);
  }

  private handleGameStarted(gameStarted: GameStarted): void {
    this.gameStatus = GameStatus.IN_PROGRESS;
  }

  private handleGameOver(gameEnded: GameEnded): void {
    this.notice = ['Game Over'];
    this.gameStatus = GameStatus.FINISHED;
  }

  private updateBoard(board: PBBoard): void {
    const pws: PlacedWord[] = [];
    for (const word of board.getWordsList()) {
      let o = Orientation.Vertical;
      if (word.getOrientation() === Word.Orientation.HORIZONTAL) {
        o = Orientation.Horizontal;
      }
      pws.push({
        x: word.getX(),
        y: word.getY(),
        orientation: o,
        word: word.getText(),
        suggestion: false,
      });
    }
    this.board.setWords(pws);
  }

  private updateTiles(tiles: Tiles): void {
    const letters: Letter[] = [];

    for (const l of tiles.getLettersList()) {
      letters.push({
        letter: l,
        deleting: false,
        selected: false,
      });
    }

    this.letters = letters;
  }

  private updateTileCounts(otus: OtherTileUpdates): void {
    this.tilesInBunch = otus.getRemainingTiles();
    const idMap: { [s: string]: OtherTileUpdate; } = {};
    for (const up of otus.getUpdatesList()) {
      idMap[up.getPlayerId()] = up;
    }

    for (const p of this.players) {
      const otu = idMap[p.id];
      p.tilesInHand = otu.getTilesInHand();
      p.tilesInBoard = otu.getTilesInBoard();
    }
  }

  private handlePlayerUpdate(pu: PlayerUpdate): void {
    const indexMap: { [s: string]: number; } = {};
    let i = 0;
    for (const p of this.players) {
      indexMap[p.id] = i;
      i++;
    }

    const pbPlayer = pu.getPlayer();
    const player = {
      id: pbPlayer.getId(),
      name: pbPlayer.getName(),
      tilesInHand: pbPlayer.getTilesInHand(),
      tilesInBoard: pbPlayer.getTilesInBoard(),
    };
    // Player already exists, update them.
    if (indexMap.hasOwnProperty(player.id)) {
      Vue.set(this.players, indexMap[player.id], player);
      return;
    }

    // Board doesn't exist yet, add it.
    this.players.push(player);
  }

  private requiredCount(): { [s: string]: number; } {
    const need: { [s: string]: number; } = {};
    for (const letter of this.requiredLetters) {
      if (need[letter] === undefined) {
        need[letter] = 0;
      }
      need[letter]++;
    }
    return need;
  }

  private missing(): string[] {
    const need = this.requiredCount();

    for (const letter of this.letters) {
      if (need[letter.letter] !== undefined) {
        need[letter.letter]--;
      }
    }

    const missing: string[] = [];
    Object.entries(need).forEach(([letter, count]) => {
      if (count <= 0) {
        return;
      }

      for (let i = 0; i < count; i++) {
        missing.push(letter);
      }
    });

    return missing;
  }

  private destroy(): void {
    document.removeEventListener('keyup', this.keyup);
  }

  private keyup(e: KeyboardEvent): void {
    if (this.gameOver) {
      return;
    }

    e.stopPropagation();

    // Left
    if (e.keyCode === 37) {
      this.prev();
      return;
    }

    // Right
    if (e.keyCode === 39) {
      this.next();
      return;
    }

    // Escape
    if (e.keyCode === 27) {
      this.word = '';
      this.notice = [];
      this.requiredLetters = [];
      this.clearDeleting();
      this.board.clear();
      return;
    }

    // Backspace
    if (e.keyCode === 8) {
      this.word = this.word.slice(0, -1);
      this.getSuggestions();
    }

    // Enter
    if (e.keyCode === 13) {
      this.place();
    }

    // A letter character.
    if (e.keyCode >= 65 && e.keyCode <= 90) {
      this.word += String.fromCharCode(e.keyCode);
      this.getSuggestions();
    }
  }

  private selectIfExists(letter: string): void {
    const index = this.letters.map((x) => x.letter + (x.deleting ? '1' : '0')).indexOf(letter + '0');
    if (index > -1) {
      this.letters[index].deleting = true;
      this.hand.renderLetters();
    }
  }

  private clearDeleting(): void {
    for (const letter in this.letters) {
      if (this.letters.hasOwnProperty(letter)) {
        this.letters[letter].deleting = false;
      }
    }
    this.hand.renderLetters();
  }

  private getSuggestions(): void {
    const suggestion = this.board.suggestPlacement(this.word, this.lettersAsStringArray());
    this.requiredLetters = suggestion.requiredLetters;
    this.setNotice(suggestion.valid);
  }

  private setNotice(fit: boolean): void {
    if (this.gameOver) {
      return;
    }

    if (this.detachedBoard) {
      this.notice = ['Board is not all connected'];
      return;
    }

    const missing = this.missing();
    if (!fit) {
      this.notice = [`Couldn't find a place to put the word`];
      return;
    }

    if (missing.length > 0) {
      this.notice = [
        `Needed: ${this.requiredLetters.join(', ')}`,
        `Missing: ${missing.join(', ')}`,
      ];
    } else if (this.word.length > 0) {
      this.notice = [
        'Have all needed letters:',
        this.requiredLetters.join(', '),
      ];
    } else {
      this.notice = [];
    }
  }

  private place(): boolean {
    if (this.missing().length > 0) {
      return false;
    }

    this.board.placeCurrentWord();

    const toRemove = this.requiredCount();
    const indicesToRemove: number[] = [];

    for (let i = this.letters.length - 1; i >= 0; i--) {
      const l = this.letters[i].letter;
      if (toRemove[l] !== undefined && toRemove[l] > 0) {
        indicesToRemove.push(i);
        toRemove[l]--;
      }
    }

    this.requiredLetters = [];
    this.word = '';
    this.notice = [];

    for (const idx of indicesToRemove) {
      this.letters.splice(idx, 1);
    }
    this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));
    this.hand.renderLetters();

    return true;
  }

  private addTileToHand(letter: string): void {
    if (this.gameOver) {
      return;
    }
    this.letters.push({letter, deleting: false, selected: false});
    this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));
  }

  private dumpTile(letter: string): void {
    if (this.gameOver) {
      return;
    }
    const req = new DumpRequest();
    req.setId(this.game!.getId());
    req.setPlayerId(this.playerID!);
    req.setLetter(letter);
    this.$client.dump(req, {}, (err, resp) => {
      if (err) {
        console.log(err);
        return;
      }
      this.updateTiles(resp.getAllTiles());
    });
  }

  private prev(): void {
    const suggestion = this.board.prevSuggestion();
    this.requiredLetters = suggestion.requiredLetters;
    this.setNotice(suggestion.valid);
  }

  private next(): void {
    const suggestion = this.board.nextSuggestion();
    this.requiredLetters = suggestion.requiredLetters;
    this.setNotice(suggestion.valid);
  }

  private addSelected(c: Cell): void {
    let selected = '';
    let idx = 0;
    for (let i = 0; i < this.letters.length; i++) {
      const l = this.letters[i];
      if (l.selected) {
        selected = l.letter;
        idx = i;
        break;
      }
    }

    if (selected === '') {
      return;
    }

    c.letterLoc.letter = selected;
    this.board.addCell(c);
    this.hand.clearSelected();
    this.letters.splice(idx, 1);
    this.word = '';
    this.notice = [];
    this.requiredLetters = [];
    this.clearDeleting();
    this.board.clear();
  }

  private sendBoard(b: {board: PBBoard, latest: Word | null}): void {
    const req = new UpdateBoardRequest();
    req.setId(this.game!.getId());
    req.setPlayerId(this.playerID!);
    req.setBoard(b.board);
    if (b.latest) {
      req.setLatestWord(b.latest);
    }
    this.$client.updateBoard(req, {}, (err, resp) => {
      if (!resp) {
        console.log(err);
        return;
      }
      this.board.setInvalidWordsAndDetached(
        resp.getInvalidWordsList(),
        resp.getDetachedBoard());
      this.detachedBoard = resp.getDetachedBoard();
      this.setNotice(true);
    });
  }

  get gameOver(): boolean {
    return this.gameStatus === GameStatus.FINISHED;
  }

  get waitingForPlayers(): boolean {
    return this.gameStatus === GameStatus.WAITING_FOR_PLAYERS;
  }

  private lettersAsStringArray(): string[] {
    const out: string[] = [];
    for (const l of this.letters) {
      out.push(l.letter);
    }
    return out;
  }
}
</script>

<style scoped>
.game, .left-side, .right-side {
  /* Just because borders cause annoyingness. */
  height: 99.5%;
}

.players {
  height: 50%;
  overflow-y: auto;
  border-left: solid 1px black;
  border-bottom: solid 1px black;
}
.logs {
  height: 50%;
  border-left: solid 1px black;
  border-top: solid 1px black;
  overflow: auto;
  font-size: 12px;
}
.board {
  height: 70%;
  border-right: solid 1px black;
  border-bottom: solid 1px black;
}
.active-word {
  height: 10%;
  border-top: solid 1px black;
  border-bottom: solid 1px black;
  border-right: solid 1px black;
}
.columns.is-gapless:not(:last-child).active-word {
  padding: 0;
  margin: 0;
}
.active-display {
  border-right: solid 1px black;
}
.notice-msg {
  border-left: solid 1px black;
}
.letters {
  height: 20%;
  border-right: solid 1px black;
  border-top: solid 1px black;
}
.standings {
  width: 100%;
}
.player-scores {
  text-align: center;
}
.standing-header {
  font-weight: bold;
}
.start-game-button {
  text-align: center;
}
</style>
