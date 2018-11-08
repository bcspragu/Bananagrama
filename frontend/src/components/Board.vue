<template>
  <div id="grid"></div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';

import Cell from '@/data';

@Component
export default class Board extends Vue {
  private board: string[][] = [['']];
  private sizeX = 750;
  private sizeY = 500;
  private margin = 5;

  private mounted(): void {
    const x = Math.floor(Math.random() * 25) + 10;
    const y = Math.floor(Math.random() * 15) + 1;
    this.board = new Array();
    for (let i = 0; i < y; i++) {
      this.board.push(new Array());
      for (let j = 0; j < x; j++) {
        if (Math.random() > 0.5) {
          this.board[i].push(String.fromCharCode(65 + Math.floor(Math.random() * 26)));
        } else {
          this.board[i].push('');
        }
      }
    }

    const gridData = this.dataFromBoard();

    const grid = d3.select('#grid')
      .append('svg')
      .attr('width', this.sizeX + 'px')
      .attr('height', this.sizeY + 'px');

    const row = grid.selectAll('.row')
      .data(gridData)
      .enter().append('g')
      .attr('class', 'row');

    const column = row.selectAll('.square')
      .data((d: any) => d)
      .enter().append('g')
      .attr('class', 'square');

    column.append('rect')
      .attr('x', (d: any) => d.x)
      .attr('y', (d: any) => d.y)
      .attr('width', (d: any) => d.size)
      .attr('height', (d: any) => d.size)
      .style('fill', '#fff')
      .style('stroke', '#222');

    column.append('text')
        .attr('x', (d: any) => d.x + d.size / 2)
        .attr('y', (d: any) => d.y + d.size / 2)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d.letter)
        .style('font-size', (d: any) => {
          return d.size * 0.6 + 'px';
        });
  }

  private dataFromBoard(): Cell[][] {
    const data: Cell[][] = new Array();

    const boardX = this.board.length + 2;
    const boardY = this.board[0].length + 2;

    const size = Math.min(
      (this.sizeX / boardY) - this.margin,
      (this.sizeY / boardX) - this.margin);

    const paddingX = this.sizeX - ((size + this.margin) * boardY);
    const paddingY = this.sizeY - ((size + this.margin) * boardX);

    // Start xpos and ypos at margin so the stroke will show when we make the grid
    // below.
    let xpos = this.margin / 2 + paddingX / 2;
    let ypos = this.margin / 2 + paddingY / 2;

    // iterate for rows
    for (let row = 0; row < boardX; row++) {
      data.push(new Array());

      // iterate for cells/columns inside rows
      for (let column = 0; column < boardY; column++) {
        let letter = '';
        if (row > 0 && row < boardX - 1 && column > 0 && column < boardY - 1) {
          letter = this.board[row - 1][column - 1];
        }
        data[row].push({
          x: xpos,
          y: ypos,
          letter,
          row,
          column,
          size,
        });
        // increment the x position. I.e. move it over by 50 (width variable)
        xpos += size + this.margin;
      }
      // Reset the x position after a row is complete.
      xpos = this.margin / 2 + paddingX / 2;
      // Increment the y position for the next row. Move it down 50 (height variable)
      ypos += size + this.margin;
    }
    return data;
  }
}


</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
