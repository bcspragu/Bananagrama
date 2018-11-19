<template>
  <div class="board">
    <Board ref="board"/>
    <UnusedLetters/>
    <input v-model="word">
    <button @click="sendToBoard">Send</button>
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

  private sendToBoard(): void {
    (this.$refs.board as Board).placeWord(this.word);
  }
}
</script>

<style scoped>
.board {
  height: 100%;
}
</style>
