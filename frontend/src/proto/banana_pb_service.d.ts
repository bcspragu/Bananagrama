// package: 
// file: banana.proto

import * as banana_pb from "./banana_pb";
import {grpc} from "grpc-web-client";

type BananaServiceNewGame = {
  readonly methodName: string;
  readonly service: typeof BananaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof banana_pb.NewGameRequest;
  readonly responseType: typeof banana_pb.NewGameResponse;
};

type BananaServiceListGames = {
  readonly methodName: string;
  readonly service: typeof BananaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof banana_pb.ListGamesRequest;
  readonly responseType: typeof banana_pb.ListGamesResponse;
};

type BananaServiceStartGame = {
  readonly methodName: string;
  readonly service: typeof BananaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof banana_pb.StartGameRequest;
  readonly responseType: typeof banana_pb.StartGameResponse;
};

type BananaServiceJoinGame = {
  readonly methodName: string;
  readonly service: typeof BananaService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof banana_pb.JoinGameRequest;
  readonly responseType: typeof banana_pb.GameUpdate;
};

type BananaServiceUpdateBoard = {
  readonly methodName: string;
  readonly service: typeof BananaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof banana_pb.UpdateBoardRequest;
  readonly responseType: typeof banana_pb.UpdateBoardResponse;
};

type BananaServiceDump = {
  readonly methodName: string;
  readonly service: typeof BananaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof banana_pb.DumpRequest;
  readonly responseType: typeof banana_pb.DumpResponse;
};

export class BananaService {
  static readonly serviceName: string;
  static readonly NewGame: BananaServiceNewGame;
  static readonly ListGames: BananaServiceListGames;
  static readonly StartGame: BananaServiceStartGame;
  static readonly JoinGame: BananaServiceJoinGame;
  static readonly UpdateBoard: BananaServiceUpdateBoard;
  static readonly Dump: BananaServiceDump;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class BananaServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  newGame(
    requestMessage: banana_pb.NewGameRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: banana_pb.NewGameResponse|null) => void
  ): UnaryResponse;
  newGame(
    requestMessage: banana_pb.NewGameRequest,
    callback: (error: ServiceError|null, responseMessage: banana_pb.NewGameResponse|null) => void
  ): UnaryResponse;
  listGames(
    requestMessage: banana_pb.ListGamesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: banana_pb.ListGamesResponse|null) => void
  ): UnaryResponse;
  listGames(
    requestMessage: banana_pb.ListGamesRequest,
    callback: (error: ServiceError|null, responseMessage: banana_pb.ListGamesResponse|null) => void
  ): UnaryResponse;
  startGame(
    requestMessage: banana_pb.StartGameRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: banana_pb.StartGameResponse|null) => void
  ): UnaryResponse;
  startGame(
    requestMessage: banana_pb.StartGameRequest,
    callback: (error: ServiceError|null, responseMessage: banana_pb.StartGameResponse|null) => void
  ): UnaryResponse;
  joinGame(requestMessage: banana_pb.JoinGameRequest, metadata?: grpc.Metadata): ResponseStream<banana_pb.GameUpdate>;
  updateBoard(
    requestMessage: banana_pb.UpdateBoardRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: banana_pb.UpdateBoardResponse|null) => void
  ): UnaryResponse;
  updateBoard(
    requestMessage: banana_pb.UpdateBoardRequest,
    callback: (error: ServiceError|null, responseMessage: banana_pb.UpdateBoardResponse|null) => void
  ): UnaryResponse;
  dump(
    requestMessage: banana_pb.DumpRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: banana_pb.DumpResponse|null) => void
  ): UnaryResponse;
  dump(
    requestMessage: banana_pb.DumpRequest,
    callback: (error: ServiceError|null, responseMessage: banana_pb.DumpResponse|null) => void
  ): UnaryResponse;
}
