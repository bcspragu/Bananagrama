<template>
  <div class="game">
    <div class="columns is-gapless left-side">
      <div class="column is-four-fifths">
        <div class="board">
          <Board ref="board"
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
          <UnusedLetters v-if="letters" ref="hand" :letters="letters" v-on:dumpTile="dumpTile"/>
        </div>
      </div>
      <div class="column is-one-fifth right-side">
        <div class="players">
          <div v-if="tilesInBunch !== null">{{tilesInBunch}} tiles left in bunch</div>
          <ol>
            <li v-for="player in players">{{player.getName()}}: {{player.getTilesInHand()}} - {{player.getTilesInBunch()}}</li>
          </ol>
          <a class="button is-info" @click="startGame">
            Start Game
          </a>
        </div>
        <div class="logs">
          <ul>
            <li v-for="log in logs">{{log}}</li>
          </ul>
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
import {Cell, Letter, PlacedWord, Orientation} from '@/data';
import {Game as PBGame, Board as PBBoard, ListGamesRequest, JoinGameRequest,
        GameUpdate, YouUpdate, PlayerUpdate, TileUpdate,
        StatusUpdate, MoveUpdate, StartGameRequest, Player, DumpRequest,
        UpdateBoardRequest, BoardUpdate, Word} from '@/proto/banana_pb';

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
  private gameOver: boolean = false;

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

    const id = this.game.getId();
    const name = this.$cookies.get('player-name');
    const req = new JoinGameRequest();
    req.setId(id);
    req.setName(name);

    const playerID = this.getPlayerIDFromCookies();
    if (playerID) {
      req.setPlayerId(playerID);
    }

    const stream = this.$client.joinGame(req, {});

    stream.on('data', (resp) => {
      switch (resp.getUpdateCase()) {
        case GameUpdate.UpdateCase.YOU_UPDATE:
          this.handleYouUpdate(resp.getYouUpdate()!);
          break;
        case GameUpdate.UpdateCase.PLAYER_UPDATE:
          this.handlePlayerUpdate(resp.getPlayerUpdate()!);
          break;
        case GameUpdate.UpdateCase.STATUS_UPDATE:
          this.handleStatusUpdate(resp.getStatusUpdate()!);
          break;
        case GameUpdate.UpdateCase.TILE_UPDATE:
          this.handleTileUpdate(resp.getTileUpdate()!);
          break;
        case GameUpdate.UpdateCase.BOARD_UPDATE:
          this.handleBoardUpdate(resp.getBoardUpdate()!);
          break;
        case GameUpdate.UpdateCase.MOVE_UPDATE:
          this.handleMoveUpdate(resp.getMoveUpdate()!);
          break;
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

  private handleYouUpdate(up: YouUpdate): void {
    this.playerID = up.getYourId();
    this.setPlayerIDInCookies(this.playerID);
  }

  private handlePlayerUpdate(up: PlayerUpdate): void {
    this.players = up.getPlayersList();
    this.tilesInBunch = up.getRemainingTiles();
    return;
  }

  private handleStatusUpdate(up: StatusUpdate): void {
    if (up.getStatus() === StatusUpdate.Status.GAME_OVER) {
      const t = new Date();
      const message = `[${t.getHours()}:${t.getMinutes()}:${t.getSeconds()}] GAME OVAH`;
      this.logs.unshift(message);
      this.gameOver = true;
    }
    return;
  }

  private handleTileUpdate(up: TileUpdate): void {
    const letters: Letter[] = [];

    for (const l of up.getAllTiles()!.getLettersList()) {
      letters.push({
        letter: l,
        deleting: false,
        selected: false,
      });
    }

    this.letters = letters;

    const t = new Date();
    let message = `[${t.getHours()}:${t.getMinutes()}:${t.getSeconds()}] `;
    let valid = true;
    switch (up.getEvent()) {
    case TileUpdate.Event.SPLIT:
      message += 'The game has started!';
      break;
    case TileUpdate.Event.PEEL:
      message += `${up.getPlayer()} peeled`;
      break;
    case TileUpdate.Event.DUMP:
      message += `${up.getPlayer()} dumped`;
      break;
    default:
      valid = false;
    }
    if (valid) {
      this.logs.unshift(message);
    }
    return;
  }

  private handleBoardUpdate(up: BoardUpdate): void {
    const pws: PlacedWord[] = [];
    for (const word of up.getBoard()!.getWordsList()) {
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

  private handleMoveUpdate(up: MoveUpdate): void {
    const t = new Date();
    const message = `[${t.getHours()}:${t.getMinutes()}:${t.getSeconds()}] ${up.getPlayer()} played "${up.getWord()}"`;
    this.logs.unshift(message);
    this.gameOver = true;
    return;
  }

  private getPlayerIDFromCookies(): string | undefined {
    if (!this.game) {
      return;
    }

    const id = this.game.getId();
    return this.$cookies.get(`game/${id}/player-id`);
  }

  private setPlayerIDInCookies(playerID: string): void {
    if (!this.game) {
      return;
    }

    const id = this.game.getId();
    this.$cookies.set(`game/${id}/player-id`, playerID);
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
    const suggestion = this.board.suggestPlacement(this.word);
    this.requiredLetters = suggestion.requiredLetters;
    this.setNotice(suggestion.valid);
  }

  private setNotice(fit: boolean): void {
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
    this.letters.push({letter, deleting: false, selected: false});
    this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));
  }

  private dumpTile(letter: string): void {
    const req = new DumpRequest();
    req.setId(this.game!.getId());
    req.setPlayerId(this.playerID!);
    req.setLetter(letter);
    this.$client.dump(req, {}, (err, resp) => {
      if (err) {
        console.log(err);
      }
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
</style>
