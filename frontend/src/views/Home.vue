<template>
  <div class="board">
    <Board ref="board"/>
    <UnusedLetters/>
    <input v-model="word">
    <button @click="getSuggestions">Suggest</button>
    <button @click="place">Place</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Board from '@/components/Board.vue'; // @ is an alias to /src
import UnusedLetters from '@/components/UnusedLetters.vue'; // @ is an alias to /src

@Component({
  components: {
    Board,
    UnusedLetters,
  },
})
export default class Home extends Vue {
  private word: string = '';

  private mounted(): void {
    document.addEventListener('keyup', (e) => {
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

      e.stopPropagation();
    });
  }

  private getSuggestions(): void {
    (this.$refs.board as Board).suggestPlacement(this.word);
  }

  private place(): void {
    (this.$refs.board as Board).placeCurrentWord();
  }

  private prev(): void {
    (this.$refs.board as Board).prevSuggestion();
  }

  private next(): void {
    (this.$refs.board as Board).nextSuggestion();
  }
}
</script>

<style scoped>
.board {
  height: 100%;
}
</style>
