<template>
  <div class="board">
    <Board ref="board"/>
    <UnusedLetters v-if="letters" ref="hand" :letters="letters" />
    <input v-model="word">
    <button @click="getSuggestions">Suggest</button>
    <button @click="place">Place</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Board from '@/components/Board.vue'; // @ is an alias to /src
import UnusedLetters from '@/components/UnusedLetters.vue'; // @ is an alias to /src
import {Letter} from '@/data';

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

    document.addEventListener('keyup', this.keyup);
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
      return;
    }

    // Any non-character.
    if (e.keyCode < 65 || e.keyCode > 90) {
      return;
    }

    this.selectIfExists(e.key.toUpperCase());

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
    this.board.placeCurrentWord();

    this.requiredLetters = [];
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
