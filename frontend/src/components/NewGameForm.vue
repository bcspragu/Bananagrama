<template>
  <div>

    <div class="field">
      <label class="label">Game Name</label>
      <div class="control">
        <input class="input" v-model="gameName" type="text" placeholder="New Game">
      </div>
    </div>

    <div class="field">
      <label class="label">Minimum Letters for Valid Word</label>
      <div class="control">
        <div class="select">
          <b-select v-model="minLetters">
            <option value="2">2</option>
            <option value="3">3</option>
            <option value="4">4</option>
            <option value="5">5</option>
          </b-select>
        </div>
      </div>
    </div>


    <div class="field">
      <div class="control has-text-centered">
        <button class="button is-info" @click="createGame">Create Game</button>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import {NewGameRequest} from '@/proto/banana_pb';

@Component
export default class NewGameForm extends Vue {
  private gameName: string = '';
  private minLetters: string = '2';
  private advancedOpen = false;

  private createGame() {
    const req = new NewGameRequest();
    req.setName(this.gameName);
    req.setMinLettersInWord(parseInt(this.minLetters));
    this.$client.newGame(req, {}, (err, resp) => {
      if (err) {
        console.log(err);
      }
    });
  }
}
</script>

<style scoped>
.new-game-input {
  justify-content: center;
  margin-bottom: 1rem;
}
</style>
