<template>
  <div id="grid"></div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';

import {Cell, LetterLoc, Orientation, Word, PlacedWord, Placement, Match} from '@/data';

import {Board as PBBoard, Word as PBWord, UpdateBoardRequest} from '@/proto/banana_pb';


interface Suggestion {
  requiredLetters: string[];
  valid: boolean;
}

interface BoardData {
  cellSize: number;
  cellRadius: number;
  suggestion: Suggestion;
  data: Cell[][];
}

@Component
export default class Board extends Vue {
  private placedWords: PlacedWord[] = [];

  // Temporary variables used for suggestions.
  private currentWord = '';
  private lastFits: Placement[] = [];
  private fitIndex = 0;

  private sizeX = 0;
  private sizeY = 0;
  private margin = 2;
  private grid: Selection<any, any, any, any> = d3.select('#grid');

  public addCell(c: Cell): void {
    const ll = c.letterLoc;
    for (let i = 0; i < ll.wordIndex.length; i++) {
      const wordIndex = ll.wordIndex[i];
      const letterIndex = ll.letterIndex[i];

      const word = this.placedWords[wordIndex].word;
      this.placedWords[wordIndex].word = word.substr(0, letterIndex) + ll.letter + word.substr(letterIndex + 1);
    }

    if (ll.wordIndex.length === 0) {
      this.normalizeAndAdd({
        word: ll.letter,
        x: c.row - 1,
        y: c.column - 1,
        orientation: Orientation.Horizontal,
        suggestion: false,
      });
    }

    this.renderBoard();
  }

  // Returns true if we had a fit to place.
  public placeCurrentWord(): boolean {
    const fit = this.currentFit;
    if (!fit) {
      return false;
    }

    this.normalizeAndAdd({
      word: this.currentWord,
      x: fit.x,
      y: fit.y,
      orientation: fit.orientation,
      suggestion: false,
    });

    this.clear();

    return true;
  }

  public clear(): void {
    this.currentWord = '';
    this.lastFits = [];
    this.fitIndex = 0;

    this.renderBoard();
  }

  // Returns the letters required to place the given suggestion.
  public suggestPlacement(word: string): Suggestion {
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
      return {
        requiredLetters: word.split('').map((l) => l.toUpperCase()),
        valid: true,
      };
    }

    this.lastFits = this.findFits(word);
    return this.renderBoard();
  }

  public nextSuggestion(): Suggestion {
    return this.selectSuggestion(1);
  }

  public prevSuggestion(): Suggestion {
    return this.selectSuggestion(this.lastFits.length - 1);
  }

  private selectSuggestion(offset: number): Suggestion {
    if (this.lastFits.length === 0) {
      return {requiredLetters: [], valid: false};
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

    if (!word.suggestion) {
      this.sendBoard();
    }
  }

  private normalize(removeLast: boolean): void {
    let suggestion = false;
    if (removeLast) {
      suggestion = this.placedWords.splice(-1)[0].suggestion;
    }

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

    if (!suggestion) {
      this.sendBoard();
    }
  }

  private sendBoard(): void {
    const words: PBWord[] = [];

    for (const pw of this.placedWords) {
      const word = new PBWord();
      word.setX(pw.x);
      word.setY(pw.y);
      word.setText(pw.word);

      if (pw.orientation === Orientation.Horizontal) {
        word.setOrientation(PBWord.Orientation.HORIZONTAL);
      } else if (pw.orientation === Orientation.Vertical) {
        word.setOrientation(PBWord.Orientation.VERTICAL);
      }

      words.push(word);
    }

    const board = new PBBoard();
    board.setWordsList(words);

    this.$emit('boardUpdated', board);
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
    const board = this.genBoard().board;
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
  private fits(m: Match, word: string, board: LetterLoc[][]): boolean {
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
  private isValid(x: number, y: number, target: string, board: LetterLoc[][]): boolean {
    return x < 0
        || y < 0
        || x >= board.length
        || y >= board[x].length
        || board[x][y].letter === ''
        || board[x][y].letter === ' '
        || board[x][y].letter === target;
  }

  private mounted(): void {
    this.grid = d3.select('#grid').append('svg')
      .attr('width', 0)
      .attr('height', 0);

    this.renderBoard();
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

  private genBoard(): {board: LetterLoc[][], requiredLetters: string[]} {
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
      return {board: [[{letter: '', wordIndex: [], letterIndex: []}]], requiredLetters: []};
    }

    const board: LetterLoc[][] = [];
    for (let i = 0; i < maxX + 1; i++) {
      board.push(new Array<LetterLoc>());
      for (let j = 0; j < maxY + 1; j++) {
        board[i].push({letter: '', wordIndex: [], letterIndex: []});
      }
    }

    let suggestedWord: PlacedWord | null = null;
    for (let i = 0; i < this.placedWords.length; i++) {
      const word = this.placedWords[i];
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
      for (let j = 0; j < word.word.length; j++) {
        board[x][y].letter = word.word.charAt(j);
        board[x][y].wordIndex.push(i);
        board[x][y].letterIndex.push(j);
        inc();
      }
    }

    const reqLetters: string[] = [];
    if (suggestedWord) {
      const poses = this.wordToPositions(suggestedWord);
      for (const pos of poses) {
        if (board[pos.x][pos.y].letter === '') {
          reqLetters.push(pos.letter);
        }
        board[pos.x][pos.y].letter = pos.letter;
      }
    }

    return {board, requiredLetters: reqLetters};
  }

  private renderBoard(): Suggestion {
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
        .attr('class', 'square')
        .on('click', (d: any) => {
          if (d.letterLoc.letter === '' || d.letterLoc.letter === ' ') {
            this.$emit('blankClicked', d);
            return;
          }
          this.removeLetterLoc(d.letterLoc);
        });

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
        .text((d: any) => d.letterLoc.letter)
        .style('font-size', (d: any) => {
          return board.cellSize * 0.7 + 'px';
        });

    return board.suggestion;
  }

  private dataFromBoard(): BoardData {
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

    const boardResp = this.genBoard();
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
        let ll: LetterLoc = {letter: '', wordIndex: [], letterIndex: []};
        if (x > 0 && x < boardX - 1 && y > 0 && y < boardY - 1) {
          ll = board[x - 1][y - 1];
        }

        data[x].push({
          x: xpos,
          y: ypos,
          letterLoc: ll,
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
      this.normalize(true /* removeLast */);
    }

    return {
      cellSize: size,
      cellRadius: 5,
      suggestion: {
        requiredLetters: boardResp.requiredLetters,
        valid: !!fit,
      },
      data,
    };
  }

  private removeLetterLoc(ll: LetterLoc): void {
    const toRemove: number[] = [];
    for (let i = 0; i < ll.wordIndex.length; i++) {
      const wordIndex = ll.wordIndex[i];
      const letterIndex = ll.letterIndex[i];

      let word = this.placedWords[wordIndex].word;
      word = word.substr(0, letterIndex) + ' ' + word.substr(letterIndex + 1);

      const spaceIndex = word.search(/\S/);
      // If there's space at the beginning, move the word forward.
      if (spaceIndex > 0) {
        switch (this.placedWords[wordIndex].orientation) {
          case Orientation.Vertical:
            this.placedWords[wordIndex].y += spaceIndex;
            break;
          case Orientation.Horizontal:
            this.placedWords[wordIndex].x += spaceIndex;
            break;
        }
      }
      word = word.trim();
      // Remove empty words.
      if (word.length === 0) {
        toRemove.push(wordIndex);
      }
      this.placedWords[wordIndex].word = word.trim();
    }

    toRemove.sort().reverse();
    for (const index of toRemove) {
      this.placedWords.splice(index, 1);
    }
    this.normalize(false /* removeLast */);
    this.$emit('removeTile', ll.letter);
    this.renderBoard();
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
  height: 100%;
}
</style>
