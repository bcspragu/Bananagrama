import * as grpcWeb from 'grpc-web';

import {
  DumpRequest,
  DumpResponse,
  GameUpdate,
  JoinGameRequest,
  ListGamesRequest,
  ListGamesResponse,
  NewGameRequest,
  NewGameResponse,
  StartGameRequest,
  StartGameResponse,
  UpdateBoardRequest,
  UpdateBoardResponse} from './banana_pb';

export class BananaServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  newGame(
    request: NewGameRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: NewGameResponse) => void
  ): grpcWeb.ClientReadableStream<NewGameResponse>;

  listGames(
    request: ListGamesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ListGamesResponse) => void
  ): grpcWeb.ClientReadableStream<ListGamesResponse>;

  startGame(
    request: StartGameRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: StartGameResponse) => void
  ): grpcWeb.ClientReadableStream<StartGameResponse>;

  joinGame(
    request: JoinGameRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<GameUpdate>;

  updateBoard(
    request: UpdateBoardRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: UpdateBoardResponse) => void
  ): grpcWeb.ClientReadableStream<UpdateBoardResponse>;

  dump(
    request: DumpRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: DumpResponse) => void
  ): grpcWeb.ClientReadableStream<DumpResponse>;

}

export class BananaServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  newGame(
    request: NewGameRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<NewGameResponse>;

  listGames(
    request: ListGamesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<ListGamesResponse>;

  startGame(
    request: StartGameRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<StartGameResponse>;

  joinGame(
    request: JoinGameRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<GameUpdate>;

  updateBoard(
    request: UpdateBoardRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<UpdateBoardResponse>;

  dump(
    request: DumpRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<DumpResponse>;

}

