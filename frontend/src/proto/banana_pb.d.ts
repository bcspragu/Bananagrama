import * as jspb from "google-protobuf"

export class NewGameRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewGameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NewGameRequest): NewGameRequest.AsObject;
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
  static serializeBinaryToWriter(message: ListGamesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListGamesRequest;
  static deserializeBinaryFromReader(message: ListGamesRequest, reader: jspb.BinaryReader): ListGamesRequest;
}

export namespace ListGamesRequest {
  export type AsObject = {
  }
}

export class ListGamesResponse extends jspb.Message {
  getGamesList(): Array<Game>;
  setGamesList(value: Array<Game>): void;
  clearGamesList(): void;
  addGames(value?: Game, index?: number): Game;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListGamesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListGamesResponse): ListGamesResponse.AsObject;
  static serializeBinaryToWriter(message: ListGamesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListGamesResponse;
  static deserializeBinaryFromReader(message: ListGamesResponse, reader: jspb.BinaryReader): ListGamesResponse;
}

export namespace ListGamesResponse {
  export type AsObject = {
    gamesList: Array<Game.AsObject>,
  }
}

export class GamesList extends jspb.Message {
  getGamesList(): Array<Game>;
  setGamesList(value: Array<Game>): void;
  clearGamesList(): void;
  addGames(value?: Game, index?: number): Game;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GamesList.AsObject;
  static toObject(includeInstance: boolean, msg: GamesList): GamesList.AsObject;
  static serializeBinaryToWriter(message: GamesList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GamesList;
  static deserializeBinaryFromReader(message: GamesList, reader: jspb.BinaryReader): GamesList;
}

export namespace GamesList {
  export type AsObject = {
    gamesList: Array<Game.AsObject>,
  }
}

export class Game extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getStatus(): GameStatus;
  setStatus(value: GameStatus): void;

  getPlayerCount(): number;
  setPlayerCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Game.AsObject;
  static toObject(includeInstance: boolean, msg: Game): Game.AsObject;
  static serializeBinaryToWriter(message: Game, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Game;
  static deserializeBinaryFromReader(message: Game, reader: jspb.BinaryReader): Game;
}

export namespace Game {
  export type AsObject = {
    id: string,
    name: string,
    status: GameStatus,
    playerCount: number,
  }
}

export class StartGameRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartGameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StartGameRequest): StartGameRequest.AsObject;
  static serializeBinaryToWriter(message: StartGameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartGameRequest;
  static deserializeBinaryFromReader(message: StartGameRequest, reader: jspb.BinaryReader): StartGameRequest;
}

export namespace StartGameRequest {
  export type AsObject = {
    id: string,
  }
}

export class StartGameResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartGameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StartGameResponse): StartGameResponse.AsObject;
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

  getPlayerId(): string;
  setPlayerId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinGameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinGameRequest): JoinGameRequest.AsObject;
  static serializeBinaryToWriter(message: JoinGameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinGameRequest;
  static deserializeBinaryFromReader(message: JoinGameRequest, reader: jspb.BinaryReader): JoinGameRequest;
}

export namespace JoinGameRequest {
  export type AsObject = {
    id: string,
    name: string,
    playerId: string,
  }
}

export class GameUpdate extends jspb.Message {
  getCurrentStatus(): CurrentStatus | undefined;
  setCurrentStatus(value?: CurrentStatus): void;
  hasCurrentStatus(): boolean;
  clearCurrentStatus(): void;

  getPlayerUpdate(): PlayerUpdate | undefined;
  setPlayerUpdate(value?: PlayerUpdate): void;
  hasPlayerUpdate(): boolean;
  clearPlayerUpdate(): void;

  getGameStarted(): GameStarted | undefined;
  setGameStarted(value?: GameStarted): void;
  hasGameStarted(): boolean;
  clearGameStarted(): void;

  getGameEnded(): GameEnded | undefined;
  setGameEnded(value?: GameEnded): void;
  hasGameEnded(): boolean;
  clearGameEnded(): void;

  getSelfTileUpdate(): SelfTileUpdate | undefined;
  setSelfTileUpdate(value?: SelfTileUpdate): void;
  hasSelfTileUpdate(): boolean;
  clearSelfTileUpdate(): void;

  getOtherTileUpdate(): OtherTileUpdates | undefined;
  setOtherTileUpdate(value?: OtherTileUpdates): void;
  hasOtherTileUpdate(): boolean;
  clearOtherTileUpdate(): void;

  getUpdateCase(): GameUpdate.UpdateCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: GameUpdate): GameUpdate.AsObject;
  static serializeBinaryToWriter(message: GameUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameUpdate;
  static deserializeBinaryFromReader(message: GameUpdate, reader: jspb.BinaryReader): GameUpdate;
}

export namespace GameUpdate {
  export type AsObject = {
    currentStatus?: CurrentStatus.AsObject,
    playerUpdate?: PlayerUpdate.AsObject,
    gameStarted?: GameStarted.AsObject,
    gameEnded?: GameEnded.AsObject,
    selfTileUpdate?: SelfTileUpdate.AsObject,
    otherTileUpdate?: OtherTileUpdates.AsObject,
  }

  export enum UpdateCase { 
    UPDATE_NOT_SET = 0,
    CURRENT_STATUS = 1,
    PLAYER_UPDATE = 2,
    GAME_STARTED = 3,
    GAME_ENDED = 4,
    SELF_TILE_UPDATE = 5,
    OTHER_TILE_UPDATE = 6,
  }
}

export class CurrentStatus extends jspb.Message {
  getYourId(): string;
  setYourId(value: string): void;

  getPlayersList(): Array<Player>;
  setPlayersList(value: Array<Player>): void;
  clearPlayersList(): void;
  addPlayers(value?: Player, index?: number): Player;

  getRemainingTiles(): number;
  setRemainingTiles(value: number): void;

  getBoard(): Board | undefined;
  setBoard(value?: Board): void;
  hasBoard(): boolean;
  clearBoard(): void;

  getAllTiles(): Tiles | undefined;
  setAllTiles(value?: Tiles): void;
  hasAllTiles(): boolean;
  clearAllTiles(): void;

  getStatus(): GameStatus;
  setStatus(value: GameStatus): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CurrentStatus.AsObject;
  static toObject(includeInstance: boolean, msg: CurrentStatus): CurrentStatus.AsObject;
  static serializeBinaryToWriter(message: CurrentStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CurrentStatus;
  static deserializeBinaryFromReader(message: CurrentStatus, reader: jspb.BinaryReader): CurrentStatus;
}

export namespace CurrentStatus {
  export type AsObject = {
    yourId: string,
    playersList: Array<Player.AsObject>,
    remainingTiles: number,
    board?: Board.AsObject,
    allTiles?: Tiles.AsObject,
    status: GameStatus,
  }
}

export class Tiles extends jspb.Message {
  getLettersList(): Array<string>;
  setLettersList(value: Array<string>): void;
  clearLettersList(): void;
  addLetters(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Tiles.AsObject;
  static toObject(includeInstance: boolean, msg: Tiles): Tiles.AsObject;
  static serializeBinaryToWriter(message: Tiles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Tiles;
  static deserializeBinaryFromReader(message: Tiles, reader: jspb.BinaryReader): Tiles;
}

export namespace Tiles {
  export type AsObject = {
    lettersList: Array<string>,
  }
}

export class PlayerUpdate extends jspb.Message {
  getPlayer(): Player | undefined;
  setPlayer(value?: Player): void;
  hasPlayer(): boolean;
  clearPlayer(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerUpdate): PlayerUpdate.AsObject;
  static serializeBinaryToWriter(message: PlayerUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerUpdate;
  static deserializeBinaryFromReader(message: PlayerUpdate, reader: jspb.BinaryReader): PlayerUpdate;
}

export namespace PlayerUpdate {
  export type AsObject = {
    player?: Player.AsObject,
  }
}

export class Player extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getTilesInHand(): number;
  setTilesInHand(value: number): void;

  getTilesInBoard(): number;
  setTilesInBoard(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Player.AsObject;
  static toObject(includeInstance: boolean, msg: Player): Player.AsObject;
  static serializeBinaryToWriter(message: Player, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Player;
  static deserializeBinaryFromReader(message: Player, reader: jspb.BinaryReader): Player;
}

export namespace Player {
  export type AsObject = {
    id: string,
    name: string,
    tilesInHand: number,
    tilesInBoard: number,
  }
}

export class GameStarted extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameStarted.AsObject;
  static toObject(includeInstance: boolean, msg: GameStarted): GameStarted.AsObject;
  static serializeBinaryToWriter(message: GameStarted, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameStarted;
  static deserializeBinaryFromReader(message: GameStarted, reader: jspb.BinaryReader): GameStarted;
}

export namespace GameStarted {
  export type AsObject = {
  }
}

export class GameEnded extends jspb.Message {
  getStandingsList(): Array<Player>;
  setStandingsList(value: Array<Player>): void;
  clearStandingsList(): void;
  addStandings(value?: Player, index?: number): Player;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameEnded.AsObject;
  static toObject(includeInstance: boolean, msg: GameEnded): GameEnded.AsObject;
  static serializeBinaryToWriter(message: GameEnded, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameEnded;
  static deserializeBinaryFromReader(message: GameEnded, reader: jspb.BinaryReader): GameEnded;
}

export namespace GameEnded {
  export type AsObject = {
    standingsList: Array<Player.AsObject>,
  }
}

export class SelfTileUpdate extends jspb.Message {
  getAllTiles(): Tiles | undefined;
  setAllTiles(value?: Tiles): void;
  hasAllTiles(): boolean;
  clearAllTiles(): void;

  getFromOtherPeel(): boolean;
  setFromOtherPeel(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SelfTileUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: SelfTileUpdate): SelfTileUpdate.AsObject;
  static serializeBinaryToWriter(message: SelfTileUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SelfTileUpdate;
  static deserializeBinaryFromReader(message: SelfTileUpdate, reader: jspb.BinaryReader): SelfTileUpdate;
}

export namespace SelfTileUpdate {
  export type AsObject = {
    allTiles?: Tiles.AsObject,
    fromOtherPeel: boolean,
  }
}

export class OtherTileUpdates extends jspb.Message {
  getUpdatesList(): Array<OtherTileUpdate>;
  setUpdatesList(value: Array<OtherTileUpdate>): void;
  clearUpdatesList(): void;
  addUpdates(value?: OtherTileUpdate, index?: number): OtherTileUpdate;

  getRemainingTiles(): number;
  setRemainingTiles(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtherTileUpdates.AsObject;
  static toObject(includeInstance: boolean, msg: OtherTileUpdates): OtherTileUpdates.AsObject;
  static serializeBinaryToWriter(message: OtherTileUpdates, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtherTileUpdates;
  static deserializeBinaryFromReader(message: OtherTileUpdates, reader: jspb.BinaryReader): OtherTileUpdates;
}

export namespace OtherTileUpdates {
  export type AsObject = {
    updatesList: Array<OtherTileUpdate.AsObject>,
    remainingTiles: number,
  }
}

export class OtherTileUpdate extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  getTilesInHand(): number;
  setTilesInHand(value: number): void;

  getTilesInBoard(): number;
  setTilesInBoard(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtherTileUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: OtherTileUpdate): OtherTileUpdate.AsObject;
  static serializeBinaryToWriter(message: OtherTileUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtherTileUpdate;
  static deserializeBinaryFromReader(message: OtherTileUpdate, reader: jspb.BinaryReader): OtherTileUpdate;
}

export namespace OtherTileUpdate {
  export type AsObject = {
    playerId: string,
    tilesInHand: number,
    tilesInBoard: number,
  }
}

export class StreamLogsRequest extends jspb.Message {
  getGameId(): string;
  setGameId(value: string): void;

  getPlayerId(): string;
  setPlayerId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamLogsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StreamLogsRequest): StreamLogsRequest.AsObject;
  static serializeBinaryToWriter(message: StreamLogsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamLogsRequest;
  static deserializeBinaryFromReader(message: StreamLogsRequest, reader: jspb.BinaryReader): StreamLogsRequest;
}

export namespace StreamLogsRequest {
  export type AsObject = {
    gameId: string,
    playerId: string,
  }
}

export class LogEntry extends jspb.Message {
  getPlayerMove(): PlayerMove | undefined;
  setPlayerMove(value?: PlayerMove): void;
  hasPlayerMove(): boolean;
  clearPlayerMove(): void;

  getPlayerDump(): PlayerDump | undefined;
  setPlayerDump(value?: PlayerDump): void;
  hasPlayerDump(): boolean;
  clearPlayerDump(): void;

  getPlayerPeel(): PlayerPeel | undefined;
  setPlayerPeel(value?: PlayerPeel): void;
  hasPlayerPeel(): boolean;
  clearPlayerPeel(): void;

  getEventCase(): LogEntry.EventCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LogEntry.AsObject;
  static toObject(includeInstance: boolean, msg: LogEntry): LogEntry.AsObject;
  static serializeBinaryToWriter(message: LogEntry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LogEntry;
  static deserializeBinaryFromReader(message: LogEntry, reader: jspb.BinaryReader): LogEntry;
}

export namespace LogEntry {
  export type AsObject = {
    playerMove?: PlayerMove.AsObject,
    playerDump?: PlayerDump.AsObject,
    playerPeel?: PlayerPeel.AsObject,
  }

  export enum EventCase { 
    EVENT_NOT_SET = 0,
    PLAYER_MOVE = 1,
    PLAYER_DUMP = 2,
    PLAYER_PEEL = 3,
  }
}

export class PlayerMove extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  getPlayerName(): string;
  setPlayerName(value: string): void;

  getWord(): string;
  setWord(value: string): void;

  getWordValid(): boolean;
  setWordValid(value: boolean): void;

  getBoardConnected(): boolean;
  setBoardConnected(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerMove.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerMove): PlayerMove.AsObject;
  static serializeBinaryToWriter(message: PlayerMove, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerMove;
  static deserializeBinaryFromReader(message: PlayerMove, reader: jspb.BinaryReader): PlayerMove;
}

export namespace PlayerMove {
  export type AsObject = {
    playerId: string,
    playerName: string,
    word: string,
    wordValid: boolean,
    boardConnected: boolean,
  }
}

export class PlayerDump extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  getPlayerName(): string;
  setPlayerName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerDump.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerDump): PlayerDump.AsObject;
  static serializeBinaryToWriter(message: PlayerDump, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerDump;
  static deserializeBinaryFromReader(message: PlayerDump, reader: jspb.BinaryReader): PlayerDump;
}

export namespace PlayerDump {
  export type AsObject = {
    playerId: string,
    playerName: string,
  }
}

export class PlayerPeel extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  getPlayerName(): string;
  setPlayerName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerPeel.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerPeel): PlayerPeel.AsObject;
  static serializeBinaryToWriter(message: PlayerPeel, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerPeel;
  static deserializeBinaryFromReader(message: PlayerPeel, reader: jspb.BinaryReader): PlayerPeel;
}

export namespace PlayerPeel {
  export type AsObject = {
    playerId: string,
    playerName: string,
  }
}

export class SpectateRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpectateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpectateRequest): SpectateRequest.AsObject;
  static serializeBinaryToWriter(message: SpectateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpectateRequest;
  static deserializeBinaryFromReader(message: SpectateRequest, reader: jspb.BinaryReader): SpectateRequest;
}

export namespace SpectateRequest {
  export type AsObject = {
    id: string,
  }
}

export class SpectateUpdate extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  getPlayerName(): string;
  setPlayerName(value: string): void;

  getBoard(): Board | undefined;
  setBoard(value?: Board): void;
  hasBoard(): boolean;
  clearBoard(): void;

  getInvalidWordsList(): Array<CharLocs>;
  setInvalidWordsList(value: Array<CharLocs>): void;
  clearInvalidWordsList(): void;
  addInvalidWords(value?: CharLocs, index?: number): CharLocs;

  getDetachedBoard(): boolean;
  setDetachedBoard(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpectateUpdate.AsObject;
  static toObject(includeInstance: boolean, msg: SpectateUpdate): SpectateUpdate.AsObject;
  static serializeBinaryToWriter(message: SpectateUpdate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpectateUpdate;
  static deserializeBinaryFromReader(message: SpectateUpdate, reader: jspb.BinaryReader): SpectateUpdate;
}

export namespace SpectateUpdate {
  export type AsObject = {
    playerId: string,
    playerName: string,
    board?: Board.AsObject,
    invalidWordsList: Array<CharLocs.AsObject>,
    detachedBoard: boolean,
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
  getAllTiles(): Tiles | undefined;
  setAllTiles(value?: Tiles): void;
  hasAllTiles(): boolean;
  clearAllTiles(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DumpResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DumpResponse): DumpResponse.AsObject;
  static serializeBinaryToWriter(message: DumpResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DumpResponse;
  static deserializeBinaryFromReader(message: DumpResponse, reader: jspb.BinaryReader): DumpResponse;
}

export namespace DumpResponse {
  export type AsObject = {
    allTiles?: Tiles.AsObject,
  }
}

export class UpdateBoardRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getPlayerId(): string;
  setPlayerId(value: string): void;

  getBoard(): Board | undefined;
  setBoard(value?: Board): void;
  hasBoard(): boolean;
  clearBoard(): void;

  getLatestWord(): Word | undefined;
  setLatestWord(value?: Word): void;
  hasLatestWord(): boolean;
  clearLatestWord(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateBoardRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateBoardRequest): UpdateBoardRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateBoardRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateBoardRequest;
  static deserializeBinaryFromReader(message: UpdateBoardRequest, reader: jspb.BinaryReader): UpdateBoardRequest;
}

export namespace UpdateBoardRequest {
  export type AsObject = {
    id: string,
    playerId: string,
    board?: Board.AsObject,
    latestWord?: Word.AsObject,
  }
}

export class Board extends jspb.Message {
  getWordsList(): Array<Word>;
  setWordsList(value: Array<Word>): void;
  clearWordsList(): void;
  addWords(value?: Word, index?: number): Word;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Board.AsObject;
  static toObject(includeInstance: boolean, msg: Board): Board.AsObject;
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

export class CharLocs extends jspb.Message {
  getText(): string;
  setText(value: string): void;

  getLocsList(): Array<CharLoc>;
  setLocsList(value: Array<CharLoc>): void;
  clearLocsList(): void;
  addLocs(value?: CharLoc, index?: number): CharLoc;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharLocs.AsObject;
  static toObject(includeInstance: boolean, msg: CharLocs): CharLocs.AsObject;
  static serializeBinaryToWriter(message: CharLocs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharLocs;
  static deserializeBinaryFromReader(message: CharLocs, reader: jspb.BinaryReader): CharLocs;
}

export namespace CharLocs {
  export type AsObject = {
    text: string,
    locsList: Array<CharLoc.AsObject>,
  }
}

export class CharLoc extends jspb.Message {
  getLetter(): string;
  setLetter(value: string): void;

  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharLoc.AsObject;
  static toObject(includeInstance: boolean, msg: CharLoc): CharLoc.AsObject;
  static serializeBinaryToWriter(message: CharLoc, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharLoc;
  static deserializeBinaryFromReader(message: CharLoc, reader: jspb.BinaryReader): CharLoc;
}

export namespace CharLoc {
  export type AsObject = {
    letter: string,
    x: number,
    y: number,
  }
}

export class UpdateBoardResponse extends jspb.Message {
  getInvalidWordsList(): Array<CharLocs>;
  setInvalidWordsList(value: Array<CharLocs>): void;
  clearInvalidWordsList(): void;
  addInvalidWords(value?: CharLocs, index?: number): CharLocs;

  getUnusedLettersList(): Array<string>;
  setUnusedLettersList(value: Array<string>): void;
  clearUnusedLettersList(): void;
  addUnusedLetters(value: string, index?: number): void;

  getDetachedBoard(): boolean;
  setDetachedBoard(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateBoardResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateBoardResponse): UpdateBoardResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateBoardResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateBoardResponse;
  static deserializeBinaryFromReader(message: UpdateBoardResponse, reader: jspb.BinaryReader): UpdateBoardResponse;
}

export namespace UpdateBoardResponse {
  export type AsObject = {
    invalidWordsList: Array<CharLocs.AsObject>,
    unusedLettersList: Array<string>,
    detachedBoard: boolean,
  }
}

export enum GameStatus { 
  UNKNOWN = 0,
  WAITING_FOR_PLAYERS = 1,
  IN_PROGRESS = 2,
  FINISHED = 3,
}
