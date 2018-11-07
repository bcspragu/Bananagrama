<template>
  <div>
    <Board/>
    <UnusedLetters/>
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
export default class Home extends Vue {}
</script>
