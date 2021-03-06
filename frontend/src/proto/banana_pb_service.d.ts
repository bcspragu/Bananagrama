import * as grpcWeb from 'grpc-web';

import {
  DumpRequest,
  DumpResponse,
  GameUpdate,
  GamesList,
  JoinGameRequest,
  ListGamesRequest,
  ListGamesResponse,
  LogEntry,
  NewGameRequest,
  NewGameResponse,
  RegisterRequest,
  RegisterResponse,
  SpectateRequest,
  SpectateUpdate,
  StartGameRequest,
  StartGameResponse,
  StreamLogsRequest,
  UpdateBoardRequest,
  UpdateBoardResponse} from './banana_pb';

export class BananaServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  register(
    request: RegisterRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: RegisterResponse) => void
  ): grpcWeb.ClientReadableStream<RegisterResponse>;

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

  streamGames(
    request: ListGamesRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<GamesList>;

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

  streamLogs(
    request: StreamLogsRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<LogEntry>;

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

  spectate(
    request: SpectateRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<SpectateUpdate>;

}

export class BananaServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  register(
    request: RegisterRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<RegisterResponse>;

  newGame(
    request: NewGameRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<NewGameResponse>;

  listGames(
    request: ListGamesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<ListGamesResponse>;

  streamGames(
    request: ListGamesRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<GamesList>;

  startGame(
    request: StartGameRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<StartGameResponse>;

  joinGame(
    request: JoinGameRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<GameUpdate>;

  streamLogs(
    request: StreamLogsRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<LogEntry>;

  updateBoard(
    request: UpdateBoardRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<UpdateBoardResponse>;

  dump(
    request: DumpRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<DumpResponse>;

  spectate(
    request: SpectateRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<SpectateUpdate>;

}

