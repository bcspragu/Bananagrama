<template>
  <div>
    <div id="letters"></div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';
import {Letter} from '@/data';

@Component
export default class UnusedLetters extends Vue {
  @Prop() private letters!: Letter[];
  private board: Selection<BaseType, any, HTMLElement, any> = d3.select('#letters');

  private sizeX = 650;
  private sizeY = 150;
  private margin = 10;

  public renderLetters(): void {
    this.resizeSVG();

    const perRow = 10;
    const rows = Math.ceil(this.letters.length / perRow);

    const size = Math.min(
      this.sizeX / perRow - this.margin,
      this.sizeY / rows - this.margin);

    const offsetX = this.margin / 2;

    const letters = this.board.selectAll('g')
      .data(this.letters);

    const newLetters = letters.enter().append('g');

    newLetters.append('rect');
    newLetters.append('text');

    letters.merge(newLetters).select('rect')
        .attr('x', (d: any, i: number) => {
          return (i % perRow) * (size + this.margin) + offsetX;
        })
        .attr('y', (d: any, i: number) => {
          return Math.floor(i / perRow) * (size + this.margin);
        })
        .attr('width', size)
        .attr('height', size)
        .style('stroke', '#222')
        .style('fill', (d: any) => d.selected ? '#faa' : '#fff');

    letters.merge(newLetters).select('text')
        .attr('x', (d: any, i: number) => (i % perRow) * (size + this.margin) + size / 2 + offsetX)
        .attr('y', (d: any, i: number) => Math.floor(i / perRow) * (size + this.margin) + size / 2)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d.letter);
  }

  private resizeSVG(): void {
    const board = document.getElementById('letters');
    if (board) {
      this.sizeX = board.offsetWidth;
    }

    this.board
      .attr('width', this.sizeX + 'px');
  }

  private mounted(): void {
    this.board = d3.select('#letters').append('svg');

    this.renderLetters();
  }
}
</script>

<style>
#letters {
  width: 30%;

  margin: 1rem auto;
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}
</style>
