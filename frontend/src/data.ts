export interface Cell {
  x: number;
  y: number;
  row: number;
  column: number;
  letterLoc: LetterLoc;
  suggestion: boolean;
  invalid: boolean;
}

// LetterLoc contains metadata about a cell position on the board.
export interface LetterLoc {
  letter: string;
  // The index of the word(s) on the board that contributed this letter.
  wordIndex: number[];
  // The index of the letter(s) in the board that contributed this letter.
  letterIndex: number[];
}

// Suggestion where to place a word.
export interface Placement {
  x: number;
  y: number;
  orientation: Orientation;
}

// Suggestion where to place a word, including the intersection.
export interface Match {
  x: number;
  y: number;
  intX: number;
  intY: number;
  orientation: Orientation;
}

export interface PlacedWord {
  x: number;
  y: number;
  orientation: Orientation;
  word: string;
  suggestion: boolean;
}

export interface Word {
  x: number;
  y: number;
  orientation: Orientation;
  word: string;
}

export enum Orientation {
  Vertical = 1,
  Horizontal,
}

export interface Letter {
  letter: string;
  deleting: boolean;
  selected: boolean;
}
