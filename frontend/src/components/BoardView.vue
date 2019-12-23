<template>
  <div ref="grid" class="grid"></div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';

import {LetterLoc} from '@/data';

import {Board as PBBoard, Word as PBWord, UpdateBoardRequest, CharLocs} from '@/proto/banana_pb';

export interface Cell {
  x: number;
  y: number;
  row: number;
  column: number;
  letterLoc: LetterLoc;
  invalid: boolean;
}

interface BoardData {
  cellSize: number;
  cellRadius: number;
  data: Cell[][];
}

@Component
export default class BoardView extends Vue {
  @Prop() private board!: PBBoard;
  @Prop() private invalidWords!: CharLocs[];
  @Prop() private detached!: boolean;

  private sizeX = 0;
  private sizeY = 0;
  private margin = 2;
  private grid: Selection<any, any, any, any> = d3.select('#grid');

  private mounted(): void {
    const grid = (this.$refs.grid as HTMLElement);
    this.grid = d3.select(grid).append('svg')
      .attr('width', 0)
      .attr('height', 0);

    this.renderBoard();
  }

  private resizeSVG(): void {
    const grid = (this.$refs.grid as HTMLElement);
    if (grid) {
      this.sizeX = grid.offsetWidth;
      this.sizeY = grid.offsetHeight;
    }

    this.grid
      .attr('width', this.sizeX + 'px')
      .attr('height', this.sizeY + 'px');
  }

  private genBoard(): LetterLoc[][] {
    let maxX = 0;
    let maxY = 0;

    for (const word of this.board.getWordsList()) {
      let x = word.getX();
      let y = word.getY();
      switch (word.getOrientation()) {
        case PBWord.Orientation.HORIZONTAL:
          x += word.getText().length - 1;
          break;
        case PBWord.Orientation.VERTICAL:
          y += word.getText().length - 1;
          break;
      }
      if (x > maxX) {
        maxX = x;
      }
      if (y > maxY) {
        maxY = y;
      }
    }

    if (maxX === 0 && maxY === 0) {
      return [[{letter: '', wordIndex: [], letterIndex: []}]];
    }

    const board: LetterLoc[][] = [];
    for (let i = 0; i < maxX + 1; i++) {
      board.push(new Array<LetterLoc>());
      for (let j = 0; j < maxY + 1; j++) {
        board[i].push({letter: '', wordIndex: [], letterIndex: []});
      }
    }

    let ii = 0;
    for (const word of this.board.getWordsList()) {
      let inc: () => void = () => { /* Blank until orientation is looked at */};
      let x = word.getX();
      let y = word.getY();
      switch (word.getOrientation()) {
        case PBWord.Orientation.HORIZONTAL:
          inc = () => x++;
          break;
        case PBWord.Orientation.VERTICAL:
          inc = () => y++;
          break;
      }
      for (let j = 0; j < word.getText().length; j++) {
        board[x][y].letter = word.getText().charAt(j);
        board[x][y].wordIndex.push(ii);
        board[x][y].letterIndex.push(j);
        inc();
      }
      ii++;
    }

    return board;
  }

  @Watch('board')
  private renderBoard() {
    this.resizeSVG();
    const board = this.dataFromBoard();

    const htmlGrid = (this.$refs.grid as HTMLElement);
    if (this.detached) {
      htmlGrid.classList.add('detached');
    } else {
      htmlGrid.classList.remove('detached');
    }

    // Set up the rows.
    const row = this.grid.selectAll('.row')
      .data(board.data);

    row.exit().remove();

    const newRows = row.enter().append('g')
        .attr('class', 'row');

    // Now that the rows we need exist, map each data cell to a square.
    const column = row.merge(newRows)
      .data(board.data).selectAll('.square')
      .data((d: any) => d);

    column.exit().remove();

    const cells = column.enter().append('g').attr('class', 'square');

    cells.append('rect');
    cells.append('text');

    column.merge(cells).select('rect')
      .attr('x', (d: any) => d.x)
      .attr('y', (d: any) => d.y)
      .attr('rx', board.cellRadius)
      .attr('rx', board.cellRadius)
      .attr('width', board.cellSize)
      .attr('height', board.cellSize)
      .style('fill', (d: any) => {
        if (d.invalid) {
          return '#fcc';
        }
        return '#fff';
      })
      .style('stroke', (d: any) => {
        if (d.letterLoc.letter === '') {
          return '#fff';
        }
        return '#222';
      });

    column.merge(cells).select('text')
        .attr('x', (d: any) => d.x + board.cellSize / 2)
        .attr('y', (d: any) => d.y + board.cellSize / 1.8)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d.letterLoc.letter)
        .style('font-size', (d: any) => {
          return board.cellSize * 0.7 + 'px';
        });
  }

  private dataFromBoard(): BoardData {
    const data: Cell[][] = new Array();

    const board = this.genBoard();

    const boardX = board.length;
    const boardY = board[0].length;

    const size = Math.min(
      (this.sizeX / boardX) - this.margin,
      (this.sizeY / boardY) - this.margin);

    const paddingX = this.sizeX - ((size + this.margin) * boardX);
    const paddingY = this.sizeY - ((size + this.margin) * boardY);

    // Start xpos and ypos at margin so the stroke will show when we make the grid
    // below.
    let xpos = this.margin / 2 + paddingX / 2;
    let ypos = this.margin / 2 + paddingY / 2;

    // iterate for rows
    for (let x = 0; x < boardX; x++) {
      data.push(new Array());
      for (let y = 0; y < boardY; y++) {
        data[x].push({
          x: xpos,
          y: ypos,
          letterLoc: board[x][y],
          row: x,
          column: y,
          invalid: false,
        });

        // Increment the Y position.
        ypos += size + this.margin;
      }
      // Reset the Y position after a column is complete.
      ypos = this.margin / 2 + paddingY / 2;
      // Increment the X position.
      xpos += size + this.margin;
    }

    for (const iw of this.invalidWords) {
      for (const cl of iw.getLocsList()) {
        if (cl.getX() >= data.length) {
          continue;
        }
        if (cl.getY() >= data[cl.getX()].length) {
          continue;
        }
        data[cl.getX()][cl.getY()].invalid = true;
      }
    }

    return {
      cellSize: size,
      cellRadius: 5,
      data,
    };
  }
}


</script>

<style>
.square {
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}

.grid {
  height: 100%;
}

.grid.detached {
  background-color: #FF6961;
}
</style>
