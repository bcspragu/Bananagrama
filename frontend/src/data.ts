export interface Cell {
  x: number;
  y: number;
  row: number;
  column: number;
  letter: string;
  suggestion: boolean;
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
