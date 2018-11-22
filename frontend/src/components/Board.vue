<template>
  <div id="grid"></div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';

import {Cell, Orientation, Word, PlacedWord, Placement, Match} from '@/data';


@Component
export default class Board extends Vue {
  private placedWords: PlacedWord[] = [];

  // Temporary variables used for suggestions.
  private currentWord = '';
  private lastFits: Placement[] = [];
  private fitIndex = 0;

  // Specify a sane default size in case we can't get the page size for some
  // reason.
  private sizeX = 750;
  private sizeY = 500;
  private margin = 2;
  private grid: Selection<any, any, any, any> = d3.select('#grid');

  public placeCurrentWord(): void {
    const fit = this.currentFit;
    if (fit) {
      this.normalizeAndAdd({
        word: this.currentWord,
        x: fit.x,
        y: fit.y,
        orientation: fit.orientation,
        suggestion: false,
      });

      this.currentWord = '';
      this.lastFits = [];
      this.fitIndex = 0;

      this.renderBoard();
    }
  }

  public suggestPlacement(word: string): string[] {
    word = word.toUpperCase();
    this.fitIndex = 0;
    this.currentWord = word;

    // If nothing has been placed on the board, just suggest dropping that word
    // in the center of the board.
    if (this.boardEmpty) {
      this.lastFits = [
        {x: 0, y: 0, orientation: Orientation.Horizontal},
        {x: 0, y: 0, orientation: Orientation.Vertical},
      ];
      this.renderBoard();
      return word.split('').map((l) => l.toUpperCase());
    }

    this.lastFits = this.findFits(word);
    return this.renderBoard();
  }

  public nextSuggestion(): string[] {
    return this.selectSuggestion(1);
  }

  public prevSuggestion(): string[] {
    return this.selectSuggestion(this.lastFits.length - 1);
  }

  private selectSuggestion(offset: number): string[] {
    if (this.lastFits.length === 0) {
      return [];
    }

    this.fitIndex = (this.fitIndex + offset) % this.lastFits.length;
    return this.renderBoard();
  }

  private normalizeAndAdd(word: PlacedWord): void {
    let offsetX = 0;
    let offsetY = 0;
    if (word.x < 0) {
      offsetX = -word.x;
    }
    if (word.y < 0) {
      offsetY = -word.y;
    }

    this.placedWords.push(word);
    for (const pWord of this.placedWords) {
      pWord.x += offsetX;
      pWord.y += offsetY;
    }
  }

  private removeLastAndNormalize(): void {
    this.placedWords.splice(-1);
    if (this.placedWords.length === 0) {
      return;
    }

    let offsetX = this.placedWords[0].x;
    let offsetY = this.placedWords[0].y;

    for (const word of this.placedWords) {
      if (word.x < offsetX) {
        offsetX = word.x;
      }
      if (word.y < offsetY) {
        offsetY = word.y;
      }
    }

    for (const word of this.placedWords) {
      word.x -= offsetX;
      word.y -= offsetY;
    }
  }

  get boardEmpty(): boolean {
    return this.placedWords.length === 0;
  }

  get currentFit(): Placement | null {
    if (this.lastFits.length === 0) {
      return null;
    }
    return this.lastFits[this.fitIndex];
  }

  // findFits returns a list of places the given word could fit on the board.
  private findFits(word: string): Placement[] {
    // Build up a mapping of letters in our word to the index(es) they appear
    // at.
    const letters: { [s: string]: number[]; } = {};
    const board = this.board.board;
    for (let i = 0; i < word.length; i++) {
      const c = word.charAt(i);
      if (letters[c] === undefined) {
        letters[c] = [];
      }
      letters[c].push(i);
    }

    const fits: Placement[] = [];
    for (const pWord of this.placedWords) {
      // Look and see if any letters of this word overlap with our words.
      const matches = this.overlaps(letters, pWord);
      for (const match of matches) {
        if (this.fits(match, word, board)) {
          fits.push(match);
        }
      }
    }

    return fits;
  }

  private overlaps(letters: { [s: string]: number[]; }, word: Word): Match[] {
    const o = this.opposite(word.orientation);
    const res: Match[] = [];

    // Go through each letter in our word, see if there are any letter
    // overlaps.
    for (let i = 0; i < word.word.length; i++) {
      const locs = letters[word.word.charAt(i)];
      if (locs === undefined) {
        continue;
      }


      // For each overlap, mark that as a placement.
      for (const offset of locs) {
        let x = word.x;
        let y = word.y;
        let intX = word.x;
        let intY = word.y;

        // The first character's location relative to this word depends on
        // where the overlap occurs.
        switch (o) {
          case Orientation.Horizontal:
            intY = y + i;
            x -= offset;
            y += i;
            break;
          case Orientation.Vertical:
            intX = x + i;
            y -= offset;
            x += i;
            break;
        }

        res.push({x, y, orientation: o, intX, intY});
      }
    }

    return res;
  }

  private opposite(o: Orientation): Orientation {
    if (o === Orientation.Horizontal) {
      return Orientation.Vertical;
    }
    return Orientation.Horizontal;
  }

  // fits returns if a given match fits on the board.
  private fits(m: Match, word: string, board: string[][]): boolean {
    let x = m.x;
    let y = m.y;

    // List of positions to verify have a certain value.
    const check: Array<{x: number, y: number, letter: string}> = [];

    let inc: () => void = () => { /* Blank until orientation is looked at */ };
    let buf: () => Array<{x: number, y: number, letter: string}> = () => [];
    switch (m.orientation) {
      case Orientation.Horizontal:
        inc = () => x++;

        // Check that there's nothing immediately before or after the word.
        check.push({x: x - 1, y, letter: ''});
        check.push({x: x + word.length, y, letter: ''});

        buf = () => [{x, y: y - 1, letter: ''}, {x, y: y + 1, letter: ''}];

        break;
      case Orientation.Vertical:
        inc = () => y++;

        // Check that there's nothing immediately before or after the word.
        check.push({x, y: y - 1, letter: ''});
        check.push({x, y: y + word.length, letter: ''});

        buf = () => [{x: x - 1, y, letter: ''}, {x: x + 1, y, letter: ''}];

        break;
    }

    if (buf === undefined || inc === undefined) {
      return false;
    }

    for (let i = 0; i < word.length; i++) {
      check.push({x, y, letter: word.charAt(i)});
      /* The commented-out code changes the overlap checking behavior.
      if (!(x === m.intX && y === m.intY)) {
        check.push(...buf());
      }
      */
      inc();
    }

    for (const pos of check) {
      if (!this.isValid(pos.x, pos.y, pos.letter, board)) {
        return false;
      }
    }


    return true;
  }

  // isValid returns whether or not the given target is a valid location on the
  // board. A location is valid if it is:
  //   - Out of bounds
  //   - Empty
  //   - The target letter we're looking for
  private isValid(x: number, y: number, target: string, board: string[][]): boolean {
    return x < 0
        || y < 0
        || x >= board.length
        || y >= board[x].length
        || board[x][y] === ''
        || board[x][y] === target;
  }

  private mounted(): void {
    this.grid = d3.select('#grid').append('svg');

    this.renderBoard();
  }

  private randomBoard(): string[][] {
    const x = Math.floor(Math.random() * 25) + 10;
    const y = Math.floor(Math.random() * 15) + 1;
    const board = new Array();
    for (let i = 0; i < y; i++) {
      board.push(new Array());
      for (let j = 0; j < x; j++) {
        if (Math.random() > 0.5) {
          board[i].push(String.fromCharCode(65 + Math.floor(Math.random() * 26)));
        } else {
          board[i].push('');
        }
      }
    }

    return board;
  }

  private resizeSVG(): void {
    const grid = document.getElementById('grid');
    if (grid) {
      this.sizeX = grid.offsetWidth;
      this.sizeY = grid.offsetHeight;
    }

    this.grid
      .attr('width', this.sizeX + 'px')
      .attr('height', this.sizeY + 'px');
  }

  private wordToPositions(word: Word): Array<{x: number, y: number, letter: string}> {
    let inc: () => void = () => { /* Blank until orientation is looked at */};
    let x = word.x;
    let y = word.y;
    let offsetX = 0;
    let offsetY = 0;
    if (x < 0) {
      offsetX = -x;
    }
    if (y < 0) {
      offsetY = -y;
    }
    switch (word.orientation) {
      case Orientation.Horizontal:
        inc = () => x++;
        break;
      case Orientation.Vertical:
        inc = () => y++;
        break;
    }
    const locs: Array<{x: number, y: number, letter: string}> = [];
    for (let i = 0; i < word.word.length; i++) {
      locs.push({x: x + offsetX, y: y + offsetY, letter: word.word.charAt(i)});
      inc();
    }

    return locs;
  }

  get board(): {board: string[][], requiredLetters: string[]} {
    let maxX = 0;
    let maxY = 0;
    for (const word of this.placedWords) {
      let x = word.x;
      let y = word.y;
      switch (word.orientation) {
        case Orientation.Horizontal:
          x += word.word.length - 1;
          break;
        case Orientation.Vertical:
          y += word.word.length - 1;
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
      return {board: [['']], requiredLetters: []};
    }

    const board = new Array();
    for (let i = 0; i < maxX + 1; i++) {
      board.push(new Array());
      for (let j = 0; j < maxY + 1; j++) {
        board[i].push('');
      }
    }

    let suggestedWord: PlacedWord | null = null;
    for (const word of this.placedWords) {
      if (word.suggestion) {
        suggestedWord = word;
        continue;
      }

      let inc: () => void = () => { /* Blank until orientation is looked at */};
      let x = word.x;
      let y = word.y;
      switch (word.orientation) {
        case Orientation.Horizontal:
          inc = () => x++;
          break;
        case Orientation.Vertical:
          inc = () => y++;
          break;
      }
      for (let i = 0; i < word.word.length; i++) {
        board[x][y] = word.word.charAt(i);
        inc();
      }
    }

    const reqLetters: string[] = [];
    if (suggestedWord) {
      const poses = this.wordToPositions(suggestedWord);
      for (const pos of poses) {
        if (board[pos.x][pos.y] === '') {
          reqLetters.push(pos.letter);
        }
        board[pos.x][pos.y] = pos.letter;
      }
    }

    return {board, requiredLetters: reqLetters};
  }

  private renderBoard(): string[] {
    this.resizeSVG();
    const board = this.dataFromBoard();

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

    const cells = column.enter().append('g')
        .attr('class', 'square');

    cells.append('rect');
    cells.append('text');

    column.merge(cells).select('rect')
      .attr('x', (d: any) => d.x)
      .attr('y', (d: any) => d.y)
      .attr('rx', board.cellRadius)
      .attr('rx', board.cellRadius)
      .attr('width', board.cellSize)
      .attr('height', board.cellSize)
      .style('fill', (d: any) => d.suggestion ? '#ffc' : '#fff')
      .style('stroke', '#222');

    column.merge(cells).select('text')
        .attr('x', (d: any) => d.x + board.cellSize / 2)
        .attr('y', (d: any) => d.y + board.cellSize / 1.8)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d.letter)
        .style('font-size', (d: any) => {
          return board.cellSize * 0.7 + 'px';
        });

    return board.requiredLetters;
  }

  private dataFromBoard(): {cellSize: number, cellRadius: number, requiredLetters: string[], data: Cell[][]} {
    const data: Cell[][] = new Array();

    const fit = this.currentFit;
    if (fit) {
      this.normalizeAndAdd({
        word: this.currentWord,
        x: fit.x,
        y: fit.y,
        orientation: fit.orientation,
        suggestion: true,
      });
    }

    const boardResp = this.board;
    const board = boardResp.board;

    const boardX = board.length + 2;
    const boardY = board[0].length + 2;

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
        let letter = '';
        if (x > 0 && x < boardX - 1 && y > 0 && y < boardY - 1) {
          letter = board[x - 1][y - 1];
        }

        data[x].push({
          x: xpos,
          y: ypos,
          letter,
          row: x,
          column: y,
          suggestion: false,
        });

        // Increment the Y position.
        ypos += size + this.margin;
      }
      // Reset the Y position after a column is complete.
      ypos = this.margin / 2 + paddingY / 2;
      // Increment the X position.
      xpos += size + this.margin;
    }

    if (fit) {
      const poses = this.wordToPositions({word: this.currentWord, x: fit.x, y: fit.y, orientation: fit.orientation});
      for (const pos of poses) {
        const x = pos.x + 1;
        const y = pos.y + 1;
        data[x][y].suggestion = true;
      }
      this.removeLastAndNormalize();
    }

    return {cellSize: size, cellRadius: 5, requiredLetters: boardResp.requiredLetters, data};
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

#grid {
  height: 60%;
}
</style>
