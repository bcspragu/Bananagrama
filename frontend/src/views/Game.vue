<template>
  <div class="game">
    <div class="columns is-gapless left-side">
      <div class="column is-four-fifths">
        <div class="board">
          <Board ref="board"/>
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
          <UnusedLetters v-if="letters" ref="hand" :letters="letters" />
        </div>
      </div>
      <div class="column is-one-fifth right-side">
        <div class="players"></div>
        <div class="logs"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Board from '@/components/Board.vue'; // @ is an alias to /src
import UnusedLetters from '@/components/UnusedLetters.vue'; // @ is an alias to /src
import ActiveWord from '@/components/ActiveWord.vue'; // @ is an alias to /src
import Notice from '@/components/Notice.vue'; // @ is an alias to /src
import {Letter} from '@/data';
import {Game as PBGame, ListGamesRequest, JoinGameRequest,
        GameUpdate, YouUpdate, PlayerUpdate, TileUpdate,
        StatusUpdate} from '@/proto/banana_pb';

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

  private word: string = '';
  private notice: string[] = [];
  private board: Board = new Board();
  private activeWord: ActiveWord = new ActiveWord();

  // The component that renders our hand.
  private hand: UnusedLetters = new UnusedLetters();

  // Letters in our hand.
  private letters: Letter[] = [];
  // Letters required to complete the word on the current board.
  private requiredLetters: string[] = [];

  private mounted(): void {
    this.$client.listGames(new ListGamesRequest(), (err, resp) => {
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
        console.log(`Couldn't find game ID ${this.$route.params.id}`);
        return;
      }

      this.joinGame();
    });

    this.board = (this.$refs.board as Board);
    this.hand = (this.$refs.hand as UnusedLetters);

    this.letters = this.randomLetters(20);
    this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));

    document.addEventListener('keyup', this.keyup);
  }

  private joinGame(): void {
    if (!this.game) {
      return;
    }

    const id = this.game.getId();
    const name = this.$cookies.get(`game-${id}`);
    const req = new JoinGameRequest();
    console.log('asd', id, name);
    req.setId(id);
    req.setName(name);

    const stream = this.$client.joinGame(req);

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
      }
    });
    stream.on('status', (status) => {
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });
    stream.on('end', () => {
      // Game over.
      console.log('game over');
    });
  }

  private handleYouUpdate(up: YouUpdate): void {
    this.playerID = up.getYourId();
  }

  private handlePlayerUpdate(up: PlayerUpdate): void {
    return;
  }

  private handleStatusUpdate(up: StatusUpdate): void {
    return;
  }

  private handleTileUpdate(up: TileUpdate): void {
    return;
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
      this.next();
      this.getSuggestions();
      return;
    }

    // Right
    if (e.keyCode === 39) {
      this.prev();
      this.getSuggestions();
      return;
    }

    // Escape
    if (e.keyCode === 27) {
      this.clearSelected();
      this.word = '';
      this.notice = [];
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

  private randomLetters(n: number): Letter[] {
    const letters: Letter[] = [];
    for (let i = 0; i < n; i++) {
      letters.push({
        letter: String.fromCharCode(Math.floor(Math.random() * 26) + 65),
        selected: false,
      });
    }
    return letters;
  }

  private selectIfExists(letter: string): void {
    const index = this.letters.map((x) => x.letter + (x.selected ? '1' : '0')).indexOf(letter + '0');
    if (index > -1) {
      this.letters[index].selected = true;
      this.hand.renderLetters();
    }
  }

  private clearSelected(): void {
    for (const letter in this.letters) {
      if (this.letters.hasOwnProperty(letter)) {
        this.letters[letter].selected = false;
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
        this.letters[i].selected = true;
        toRemove[l]--;
      }
    }

    this.hand.renderLetters();
    this.requiredLetters = [];
    this.word = '';
    this.notice = [];

    window.setTimeout(() => {
      for (const idx of indicesToRemove) {
        this.letters.splice(idx, 1);
      }
      this.letters.push(...this.randomLetters(indicesToRemove.length));
      this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));
      this.hand.renderLetters();
    }, 500);

    return true;
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
}
</script>

<style scoped>
.game, .left-side, .right-side {
  height: 100%;
}

.players {
  height: 50%;
  border-left: solid 1px black;
  border-bottom: solid 1px black;
}
.logs {
  height: 50%;
  border-left: solid 1px black;
  border-top: solid 1px black;
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
