<template>
  <div>
    <div id="letters"></div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';

interface Letters {
  letter: string;
  selected: boolean;
}

@Component
export default class UnusedLetters extends Vue {
  private letters: Letters[] = [];
  private board: Selection<BaseType, any, HTMLElement, any> = d3.selectAll('div');

  private created(): void {
    for (let i = 0; i < 20; i++) {
      this.letters.push({
        letter: String.fromCharCode(Math.floor(Math.random() * 26) + 65),
        selected: false,
      });
    }

    window.addEventListener('keydown', (e) => {
      if (e.keyCode === 27) {
        this.clearSelected();
        return;
      }
      if (e.keyCode < 65 || e.keyCode > 90) {
        return;
      }
      this.selectIfExists(e.key.toUpperCase());
    });
  }

  private selectIfExists(letter: string): void {
    const index = this.letters.map((x) => x.letter + (x.selected ? '1' : '0')).indexOf(letter + '0');
    if (index > -1) {
      this.letters[index].selected = true;
      this.renderLetters();
    }
  }

  private clearSelected(): void {
    for (const letter in this.letters) {
      if (this.letters.hasOwnProperty(letter)) {
        this.letters[letter].selected = false;
      }
    }
    this.renderLetters();
  }

  private mounted(): void {
    this.board = d3.select('#letters')
      .append('svg')
      .attr('width', '650px')
      .attr('height', '250px');

    this.renderLetters();
  }

  private renderLetters(): void {
    const size = 55;
    const margin = 10;

    const letters = this.board.selectAll('g')
      .data(this.letters)
      .enter().append('g');

    letters.append('rect')
        .attr('x', (d: any, i: number) => {
          return (i % 10) * (size + margin);
        })
        .attr('y', (d: any, i: number) => {
          return Math.floor(i / 10) * (size + margin);
        })
        .attr('width', size)
        .attr('height', size)
        .style('stroke', '#222')
        .style('fill', (d: any) => d.selected ? '#faa' : '#fff');

    letters.append('text')
        .attr('x', (d: any, i: number) => (i % 10) * (size + margin) + size / 2)
        .attr('y', (d: any, i: number) => Math.floor(i / 10) * (size + margin) + size / 2)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d.letter);
  }
}
</script>

<style>
#letters {
  margin-top: 1rem;
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}
</style>
