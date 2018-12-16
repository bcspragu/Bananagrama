<template>
  <div id="letters"></div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';
import {Letter} from '@/data';

enum Fit {
  TOO_SMALL,
  JUST_RIGHT,
  TOO_LARGE,
}

interface BoardSizing {
  perRow: number;
  rows: number;
  fit: Fit;
}

@Component
export default class UnusedLetters extends Vue {
  @Prop() private letters!: Letter[];
  private board: Selection<BaseType, any, HTMLElement, any> = d3.select('#letters');

  private sizeX = 0;
  private sizeY = 0;
  private margin = 10;

  @Watch('letters')
  public renderLetters(): void {
    this.resizeSVG();

    // Start at a sane default.
    let size = Math.min(65, this.sizeY);
    let margin = 0.1 * size;

    // Determine if we can fit all of our tiles into this space.
    let lowerSize = 0;
    let upperSize = 2 * size;
    let perRow = 0;
    let rows = 0;
    for (let i = 0; i < 10; i++) {
      const fit = this.calculateFit(size + margin);
      perRow = fit.perRow;
      rows = fit.rows;
      if (fit.fit === Fit.JUST_RIGHT) {
        break;
      }

      if (fit.fit === Fit.TOO_LARGE) {
        upperSize -= (upperSize - lowerSize) / 4;
      }

      if (fit.fit === Fit.TOO_SMALL) {
        lowerSize += (upperSize - lowerSize) / 4;
      }

      size = (upperSize + lowerSize) / 2;
      margin = 0.1 * size;
    }

    const offsetX = ((margin + (this.sizeX - (size + margin) * perRow)) / 2);
    const offsetY = ((margin + (this.sizeY - (size + margin) * rows)) / 2);

    const letters = this.board.selectAll('g')
      .data(this.letters);

    letters.exit().remove();

    const newLetters = letters.enter().append('g');

    newLetters.append('rect');
    newLetters.append('text');

    letters.merge(newLetters).select('rect')
        .attr('x', (d: any, i: number) => {
          return (i % perRow) * (size + margin) + offsetX;
        })
        .attr('y', (d: any, i: number) => {
          return Math.floor(i / perRow) * (size + margin) + offsetY;
        })
        .attr('width', size)
        .attr('height', size)
        .style('stroke', '#222')
        .style('fill', (d: any) => d.selected ? '#faa' : '#fff');

    letters.merge(newLetters).select('text')
        .attr('x', (d: any, i: number) => (i % perRow) * (size + margin) + size / 2 + offsetX)
        .attr('y', (d: any, i: number) => Math.floor(i / perRow) * (size + margin) + size / 2 + offsetY)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d.letter)
        .style('font-size', (d: any) => {
          return size * 0.5 + 'px';
        });
  }

  // Size includes the margin here.
  private calculateFit(size: number): BoardSizing {
    const perRow = Math.floor(this.sizeX / size);
    const rows = Math.ceil(this.letters.length / perRow);
    // This variable name is garbage, but it's a weird concept to express. Now
    // that we know how many letters we can fit per row and the number of rows
    // required to fit all the letters, we need to know how many rows are
    // actually available at the current size.
    const numRowsWeCanFit = Math.floor(this.sizeY / size);
    let fit = Fit.JUST_RIGHT;
    if (numRowsWeCanFit > rows) {
      fit = Fit.TOO_SMALL;
    }
    if (numRowsWeCanFit < rows) {
      fit = Fit.TOO_LARGE;
    }
    return {
      perRow,
      rows,
      fit,
    };
  }

  private resizeSVG(): void {
    const board = document.getElementById('letters');
    if (board) {
      this.sizeX = board.offsetWidth;
      this.sizeY = board.offsetHeight;
    }

    this.board
      .attr('width', this.sizeX + 'px')
      .attr('height', this.sizeY + 'px');
  }

  private mounted(): void {
    this.board = d3.select('#letters').append('svg')
      .attr('width', 0)
      .attr('height', 0);
    this.resizeSVG();
  }
}
</script>

<style>
#letters {
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
  height: 100%;
}
</style>
