// package: 
// file: banana.proto

import * as jspb from "google-protobuf";

export class NewGameRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewGameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NewGameRequest): NewGameRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NewGameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewGameRequest;
  static deserializeBinaryFromReader(message: NewGameRequest, reader: jspb.BinaryReader): NewGameRequest;
}

export namespace NewGameRequest {
  export type AsObject = {
    name: string,
  }
}

export class NewGameResponse extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewGameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NewGameResponse): NewGameResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NewGameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewGameResponse;
  static deserializeBinaryFromReader(message: NewGameResponse, reader: jspb.BinaryReader): NewGameResponse;
}

export namespace NewGameResponse {
  export type AsObject = {
    id: string,
  }
}

export class ListGamesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListGamesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListGamesRequest): ListGamesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListGamesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListGamesRequest;
  static deserializeBinaryFromReader(message: ListGamesRequest, reader: jspb.BinaryReader): ListGamesRequest;
}

export namespace ListGamesRequest {
  export type AsObject = {
  }
}

export class ListGamesResponse extends jspb.Message {
  clearGamesList(): void;
  getGamesList(): Array<Game>;
  setGamesList(value: Array<Game>): void;
  addGames(value?: Game, index?: number): Game;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListGamesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListGamesResponse): ListGamesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListGamesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListGamesResponse;
  static deserializeBinaryFromReader(message: ListGamesResponse, reader: jspb.BinaryReader): ListGamesResponse;
}

export namespace ListGamesResponse {
  export type AsObject = {
    gamesList: Array<Game.AsObject>,
  }
}

export class Game extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Game.AsObject;
  static toObject(includeInstance: boolean, msg: Game): Game.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Game, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Game;
  static deserializeBinaryFromReader(message: Game, reader: jspb.BinaryReader): Game;
}

export namespace Game {
  export type AsObject = {
    id: string,
    name: string,
  }
}

export class StartGameRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getScaleFactor(): number;
  setScaleFactor(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartGameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StartGameRequest): StartGameRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartGameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartGameRequest;
  static deserializeBinaryFromReader(message: StartGameRequest, reader: jspb.BinaryReader): StartGameRequest;
}

export namespace StartGameRequest {
  export type AsObject = {
    id: string,
    scaleFactor: number,
  }
}

export class StartGameResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartGameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StartGameResponse): StartGameResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartGameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartGameResponse;
  static deserializeBinaryFromReader(message: StartGameResponse, reader: jspb.BinaryReader): StartGameResponse;
}

export namespace StartGameResponse {
  export type AsObject = {
  }
}

export class JoinGameRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinGameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinGameRequest): JoinGameRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JoinGameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinGameRequest;
  static deserializeBinaryFromReader(message: JoinGameRequest, reader: jspb.BinaryReader): JoinGameRequest;
}

export namespace JoinGameRequest {
  export type AsObject = {
    id: string,
    name: string,
  }
}

export class GameUpdate extends jspb.Message {
  hasYouUpdate(): boolean;
  clearYouUpdate(): void;
  getYouUpdate(): YouUpdate | undefined;
  setYouUpdate(value?: YouUpdate): void;

  hasPlayerUpdate(): boolean;
  clearPlayerUpdate(): void;
  getPlayerUpdate(): PlayerUpdate | undefined;
  setPlayerUpdate(value?: PlayerUpdate): void;

  hasStatusUpdate(): boolean;
  clearStatusUpdate(): void;
  getStatusUpdate(): StatusUpdate | undefined;
  setStatusUpdate(value?: StatusUpdate): void;

  hasTileUpdate(): boolean;
  clearTileUpdate(): void;
  getTileUpdate(): TileUpdate | undefined;
  setTileUpdate(value?: TileUpdate): void;

  getUpdateCase(): GameUpdate.UpdateCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: GameUpdate): GameUpdate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameUpdate;
  static deserializeBinaryFromReader(message: GameUpdate, reader: jspb.BinaryReader): GameUpdate;
}

export namespace GameUpdate {
  export type AsObject = {
    youUpdate?: YouUpdate.AsObject,
    playerUpdate?: PlayerUpdate.AsObject,
    statusUpdate?: StatusUpdate.AsObject,
    tileUpdate?: TileUpdate.AsObject,
  }

  export enum UpdateCase {
    UPDATE_NOT_SET = 0,
    YOU_UPDATE = 1,
    PLAYER_UPDATE = 2,
    STATUS_UPDATE = 3,
    TILE_UPDATE = 4,
  }
}

export class YouUpdate extends jspb.Message {
  getYourId(): string;
  setYourId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YouUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: YouUpdate): YouUpdate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: YouUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YouUpdate;
  static deserializeBinaryFromReader(message: YouUpdate, reader: jspb.BinaryReader): YouUpdate;
}

export namespace YouUpdate {
  export type AsObject = {
    yourId: string,
  }
}

export class PlayerUpdate extends jspb.Message {
  clearPlayersList(): void;
  getPlayersList(): Array<Player>;
  setPlayersList(value: Array<Player>): void;
  addPlayers(value?: Player, index?: number): Player;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerUpdate): PlayerUpdate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerUpdate;
  static deserializeBinaryFromReader(message: PlayerUpdate, reader: jspb.BinaryReader): PlayerUpdate;
}

export namespace PlayerUpdate {
  export type AsObject = {
    playersList: Array<Player.AsObject>,
  }
}

export class Player extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Player.AsObject;
  static toObject(includeInstance: boolean, msg: Player): Player.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Player, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Player;
  static deserializeBinaryFromReader(message: Player, reader: jspb.BinaryReader): Player;
}

export namespace Player {
  export type AsObject = {
    name: string,
  }
}

export class StatusUpdate extends jspb.Message {
  getStatus(): StatusUpdate.Status;
  setStatus(value: StatusUpdate.Status): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatusUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: StatusUpdate): StatusUpdate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatusUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatusUpdate;
  static deserializeBinaryFromReader(message: StatusUpdate, reader: jspb.BinaryReader): StatusUpdate;
}

export namespace StatusUpdate {
  export type AsObject = {
    status: StatusUpdate.Status,
  }

  export enum Status {
    UNKNOWN = 0,
    GAME_STARTED = 1,
    GAME_OVER = 2,
  }
}

export class TileUpdate extends jspb.Message {
  getEvent(): TileUpdate.Event;
  setEvent(value: TileUpdate.Event): void;

  getPlayer(): string;
  setPlayer(value: string): void;

  hasAllTiles(): boolean;
  clearAllTiles(): void;
  getAllTiles(): Tiles | undefined;
  setAllTiles(value?: Tiles): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TileUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: TileUpdate): TileUpdate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TileUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TileUpdate;
  static deserializeBinaryFromReader(message: TileUpdate, reader: jspb.BinaryReader): TileUpdate;
}

export namespace TileUpdate {
  export type AsObject = {
    event: TileUpdate.Event,
    player: string,
    allTiles?: Tiles.AsObject,
  }

  export enum Event {
    UNKNOWN = 0,
    PEEL = 1,
    DUMP = 2,
  }
}

export class Tiles extends jspb.Message {
  clearLettersList(): void;
  getLettersList(): Array<string>;
  setLettersList(value: Array<string>): void;
  addLetters(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Tiles.AsObject;
  static toObject(includeInstance: boolean, msg: Tiles): Tiles.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Tiles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Tiles;
  static deserializeBinaryFromReader(message: Tiles, reader: jspb.BinaryReader): Tiles;
}

export namespace Tiles {
  export type AsObject = {
    lettersList: Array<string>,
  }
}

export class DumpRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getPlayerId(): string;
  setPlayerId(value: string): void;

  getLetter(): string;
  setLetter(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DumpRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DumpRequest): DumpRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DumpRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DumpRequest;
  static deserializeBinaryFromReader(message: DumpRequest, reader: jspb.BinaryReader): DumpRequest;
}

export namespace DumpRequest {
  export type AsObject = {
    id: string,
    playerId: string,
    letter: string,
  }
}

export class DumpResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DumpResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DumpResponse): DumpResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DumpResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DumpResponse;
  static deserializeBinaryFromReader(message: DumpResponse, reader: jspb.BinaryReader): DumpResponse;
}

export namespace DumpResponse {
  export type AsObject = {
  }
}

export class PeelRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getPlayerId(): string;
  setPlayerId(value: string): void;

  hasBoard(): boolean;
  clearBoard(): void;
  getBoard(): Board | undefined;
  setBoard(value?: Board): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeelRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PeelRequest): PeelRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PeelRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeelRequest;
  static deserializeBinaryFromReader(message: PeelRequest, reader: jspb.BinaryReader): PeelRequest;
}

export namespace PeelRequest {
  export type AsObject = {
    id: string,
    playerId: string,
    board?: Board.AsObject,
  }
}

export class Board extends jspb.Message {
  clearWordsList(): void;
  getWordsList(): Array<Word>;
  setWordsList(value: Array<Word>): void;
  addWords(value?: Word, index?: number): Word;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Board.AsObject;
  static toObject(includeInstance: boolean, msg: Board): Board.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Board, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Board;
  static deserializeBinaryFromReader(message: Board, reader: jspb.BinaryReader): Board;
}

export namespace Board {
  export type AsObject = {
    wordsList: Array<Word.AsObject>,
  }
}

export class Word extends jspb.Message {
  getText(): string;
  setText(value: string): void;

  getOrientation(): Word.Orientation;
  setOrientation(value: Word.Orientation): void;

  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Word.AsObject;
  static toObject(includeInstance: boolean, msg: Word): Word.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Word, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Word;
  static deserializeBinaryFromReader(message: Word, reader: jspb.BinaryReader): Word;
}

export namespace Word {
  export type AsObject = {
    text: string,
    orientation: Word.Orientation,
    x: number,
    y: number,
  }

  export enum Orientation {
    UNKNOWN = 0,
    HORIZONTAL = 1,
    VERTICAL = 2,
  }
}

export class PeelResponse extends jspb.Message {
  getStatus(): PeelResponse.Status;
  setStatus(value: PeelResponse.Status): void;

  clearErrorsList(): void;
  getErrorsList(): Array<string>;
  setErrorsList(value: Array<string>): void;
  addErrors(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeelResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PeelResponse): PeelResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PeelResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeelResponse;
  static deserializeBinaryFromReader(message: PeelResponse, reader: jspb.BinaryReader): PeelResponse;
}

export namespace PeelResponse {
  export type AsObject = {
    status: PeelResponse.Status,
    errorsList: Array<string>,
  }

  export enum Status {
    UNKNOWN = 0,
    SUCCESS = 1,
    INVALID_WORD = 2,
    DETACHED_BOARD = 3,
    NOT_ALL_LETTERS = 4,
    EXTRA_LETTERS = 5,
  }
}

