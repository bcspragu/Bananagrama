<template>
  <div class="board">

    <Board ref="board"/>
    <UnusedLetters v-if="letters" ref="hand" :letters="letters" />

    <input v-model="word">
    <button @click="getSuggestions">Suggest</button>
    <button @click="place">Place</button>
    <div>{{required}}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Board from '@/components/Board.vue'; // @ is an alias to /src
import UnusedLetters from '@/components/UnusedLetters.vue'; // @ is an alias to /src
import {Letter} from '@/data';

// Import code-generated data structures.
import {BananaServiceClient} from '@/proto/banana_pb_service';
import {NewGameRequest} from '@/proto/banana_pb';

@Component({
  components: {
    Board,
    UnusedLetters,
  },
})
export default class Home extends Vue {
  private word: string = '';
  private board: Board = new Board();

  // The component that renders our hand.
  private hand: UnusedLetters = new UnusedLetters();

  // Letters in our hand.
  private letters: Letter[] = [];
  // Letters required to complete the word on the current board.
  private requiredLetters: string[] = [];

  private mounted(): void {
    this.board = (this.$refs.board as Board);
    this.hand = (this.$refs.hand as UnusedLetters);

    this.letters = this.randomLetters(20);
    this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));

    document.addEventListener('keyup', this.keyup);

    const client = new BananaServiceClient('/api');
    const req = new NewGameRequest();
    req.setName('new game');
    client.newGame(req, (resp, err) => {console.log('HAYYO', resp, err);});
  }

  get required(): string {
    return this.requiredLetters.join(', ') + `, you're missing ` + this.missing.join(', ');
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

  get missing(): string[] {
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
    // Left
    if (e.keyCode === 37) {
      this.next();
      return;
    }

    // Right
    if (e.keyCode === 39) {
      this.prev();
      return;
    }

    // Escape
    if (e.keyCode === 27) {
      this.clearSelected();
      this.word = '';
      return;
    }

    e.stopPropagation();
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
    this.requiredLetters = this.board.suggestPlacement(this.word);
  }

  private place(): void {
    if (this.missing.length > 0) {
      return;
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

    window.setTimeout(() => {
      for (const idx of indicesToRemove) {
        this.letters.splice(idx, 1);
      }
      this.letters.push(...this.randomLetters(indicesToRemove.length));
      this.letters.sort((a, b) => a.letter > b.letter ? 1 : (a.letter < b.letter ? -1 : 0));
      this.hand.renderLetters();
    }, 500);
  }

  private prev(): void {
    this.requiredLetters = this.board.prevSuggestion();
  }

  private next(): void {
    this.requiredLetters = this.board.nextSuggestion();
  }
}
</script>

<style scoped>
.board {
  height: 100%;
}
</style>
