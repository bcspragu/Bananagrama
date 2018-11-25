// package: 
// file: banana.proto

import * as jspb from "google-protobuf";

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

export class NewGameRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getScaleFactor(): number;
  setScaleFactor(value: number): void;

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
    scaleFactor: number,
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
    playerUpdate?: PlayerUpdate.AsObject,
    statusUpdate?: StatusUpdate.AsObject,
    tileUpdate?: TileUpdate.AsObject,
  }

  export enum UpdateCase {
    UPDATE_NOT_SET = 0,
    PLAYER_UPDATE = 1,
    STATUS_UPDATE = 2,
    TILE_UPDATE = 3,
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
    allTiles?: Tiles.AsObject,
  }
}

export class DumpRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

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

