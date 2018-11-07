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
  private selected: SVGElement | null = null;

  private created(): void {
    window.addEventListener('keydown', (e) => {
      this.setText(e);
    });
  }

  private mounted(): void {
    const gridData = this.gridData();

    const grid = d3.select('#grid')
      .append('svg')
      .attr('width', '510px')
      .attr('height', '510px');

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
      .attr('width', (d: any) => d.width)
      .attr('height', (d: any) => d.height)
      .style('fill', '#fff')
      .style('stroke', '#222')
      .on('click', (d: any, i: number, n: any) => {
        d.click++;
        let fill = '#fff';
        switch (d.click % 4) {
            case 1:
              fill = '#2C93E8';
              break;
            case 2:
              fill = '#F56C4E';
              break;
            case 3:
              fill = '#838690';
              break;
        }

        this.selected = n[i];
      });

    column.append('text')
      .attr('x', (d: any) => d.x + d.width / 2)
      .attr('y', (d: any) => d.y + d.height / 2)
      .attr('text-anchor', 'middle')
      .attr('alignment-baseline', 'middle');
  }

  private gridData(): Cell[][] {
    const data: Cell[][] = new Array();

    // Start xpos and ypos at 1 so the stroke will show when we make the grid
    // below.
    let xpos = 1;
    let ypos = 1;
    const width = 50;
    const height = 50;
    const click = 0;

    // iterate for rows
    for (let row = 0; row < 10; row++) {
      data.push( new Array() );

      // iterate for cells/columns inside rows
      for (let column = 0; column < 10; column++) {
        data[row].push({
          x: xpos,
          y: ypos,
          row,
          column,
          width,
          height,
          click,
        });
        // increment the x position. I.e. move it over by 50 (width variable)
        xpos += width;
      }
      // Reset the x position after a row is complete.
      xpos = 1;
      // Increment the y position for the next row. Move it down 50 (height variable)
      ypos += height;
    }
    return data;
  }

  private setText(event: any): void {
    if (event.keyCode < 65 || event.keyCode > 90) {
      return;
    }

    if (this.selected) {
      d3.select<any, any>(this.selected.parentNode)
        .selectAll('text')
        .text(event.key.toUpperCase());
    }
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
