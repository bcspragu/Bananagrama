/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./banana_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.BananaServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.BananaServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.RegisterRequest,
 *   !proto.RegisterResponse>}
 */
const methodDescriptor_BananaService_Register = new grpc.web.MethodDescriptor(
  '/BananaService/Register',
  grpc.web.MethodType.UNARY,
  proto.RegisterRequest,
  proto.RegisterResponse,
  /**
   * @param {!proto.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.RegisterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.RegisterRequest,
 *   !proto.RegisterResponse>}
 */
const methodInfo_BananaService_Register = new grpc.web.AbstractClientBase.MethodInfo(
  proto.RegisterResponse,
  /**
   * @param {!proto.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.RegisterResponse.deserializeBinary
);


/**
 * @param {!proto.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.RegisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.RegisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/BananaService/Register',
      request,
      metadata || {},
      methodDescriptor_BananaService_Register,
      callback);
};


/**
 * @param {!proto.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.RegisterResponse>}
 *     A native promise that resolves to the response
 */
proto.BananaServicePromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/BananaService/Register',
      request,
      metadata || {},
      methodDescriptor_BananaService_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.NewGameRequest,
 *   !proto.NewGameResponse>}
 */
const methodDescriptor_BananaService_NewGame = new grpc.web.MethodDescriptor(
  '/BananaService/NewGame',
  grpc.web.MethodType.UNARY,
  proto.NewGameRequest,
  proto.NewGameResponse,
  /**
   * @param {!proto.NewGameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.NewGameResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.NewGameRequest,
 *   !proto.NewGameResponse>}
 */
const methodInfo_BananaService_NewGame = new grpc.web.AbstractClientBase.MethodInfo(
  proto.NewGameResponse,
  /**
   * @param {!proto.NewGameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.NewGameResponse.deserializeBinary
);


/**
 * @param {!proto.NewGameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.NewGameResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.NewGameResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.newGame =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/BananaService/NewGame',
      request,
      metadata || {},
      methodDescriptor_BananaService_NewGame,
      callback);
};


/**
 * @param {!proto.NewGameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.NewGameResponse>}
 *     A native promise that resolves to the response
 */
proto.BananaServicePromiseClient.prototype.newGame =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/BananaService/NewGame',
      request,
      metadata || {},
      methodDescriptor_BananaService_NewGame);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ListGamesRequest,
 *   !proto.ListGamesResponse>}
 */
const methodDescriptor_BananaService_ListGames = new grpc.web.MethodDescriptor(
  '/BananaService/ListGames',
  grpc.web.MethodType.UNARY,
  proto.ListGamesRequest,
  proto.ListGamesResponse,
  /**
   * @param {!proto.ListGamesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ListGamesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ListGamesRequest,
 *   !proto.ListGamesResponse>}
 */
const methodInfo_BananaService_ListGames = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ListGamesResponse,
  /**
   * @param {!proto.ListGamesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ListGamesResponse.deserializeBinary
);


/**
 * @param {!proto.ListGamesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ListGamesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ListGamesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.listGames =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/BananaService/ListGames',
      request,
      metadata || {},
      methodDescriptor_BananaService_ListGames,
      callback);
};


/**
 * @param {!proto.ListGamesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ListGamesResponse>}
 *     A native promise that resolves to the response
 */
proto.BananaServicePromiseClient.prototype.listGames =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/BananaService/ListGames',
      request,
      metadata || {},
      methodDescriptor_BananaService_ListGames);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ListGamesRequest,
 *   !proto.GamesList>}
 */
const methodDescriptor_BananaService_StreamGames = new grpc.web.MethodDescriptor(
  '/BananaService/StreamGames',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.ListGamesRequest,
  proto.GamesList,
  /**
   * @param {!proto.ListGamesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GamesList.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ListGamesRequest,
 *   !proto.GamesList>}
 */
const methodInfo_BananaService_StreamGames = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GamesList,
  /**
   * @param {!proto.ListGamesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GamesList.deserializeBinary
);


/**
 * @param {!proto.ListGamesRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.GamesList>}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.streamGames =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/StreamGames',
      request,
      metadata || {},
      methodDescriptor_BananaService_StreamGames);
};


/**
 * @param {!proto.ListGamesRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.GamesList>}
 *     The XHR Node Readable Stream
 */
proto.BananaServicePromiseClient.prototype.streamGames =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/StreamGames',
      request,
      metadata || {},
      methodDescriptor_BananaService_StreamGames);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.StartGameRequest,
 *   !proto.StartGameResponse>}
 */
const methodDescriptor_BananaService_StartGame = new grpc.web.MethodDescriptor(
  '/BananaService/StartGame',
  grpc.web.MethodType.UNARY,
  proto.StartGameRequest,
  proto.StartGameResponse,
  /**
   * @param {!proto.StartGameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.StartGameResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.StartGameRequest,
 *   !proto.StartGameResponse>}
 */
const methodInfo_BananaService_StartGame = new grpc.web.AbstractClientBase.MethodInfo(
  proto.StartGameResponse,
  /**
   * @param {!proto.StartGameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.StartGameResponse.deserializeBinary
);


/**
 * @param {!proto.StartGameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.StartGameResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.StartGameResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.startGame =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/BananaService/StartGame',
      request,
      metadata || {},
      methodDescriptor_BananaService_StartGame,
      callback);
};


/**
 * @param {!proto.StartGameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.StartGameResponse>}
 *     A native promise that resolves to the response
 */
proto.BananaServicePromiseClient.prototype.startGame =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/BananaService/StartGame',
      request,
      metadata || {},
      methodDescriptor_BananaService_StartGame);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.JoinGameRequest,
 *   !proto.GameUpdate>}
 */
const methodDescriptor_BananaService_JoinGame = new grpc.web.MethodDescriptor(
  '/BananaService/JoinGame',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.JoinGameRequest,
  proto.GameUpdate,
  /**
   * @param {!proto.JoinGameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GameUpdate.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.JoinGameRequest,
 *   !proto.GameUpdate>}
 */
const methodInfo_BananaService_JoinGame = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GameUpdate,
  /**
   * @param {!proto.JoinGameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GameUpdate.deserializeBinary
);


/**
 * @param {!proto.JoinGameRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.GameUpdate>}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.joinGame =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/JoinGame',
      request,
      metadata || {},
      methodDescriptor_BananaService_JoinGame);
};


/**
 * @param {!proto.JoinGameRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.GameUpdate>}
 *     The XHR Node Readable Stream
 */
proto.BananaServicePromiseClient.prototype.joinGame =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/JoinGame',
      request,
      metadata || {},
      methodDescriptor_BananaService_JoinGame);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.StreamLogsRequest,
 *   !proto.LogEntry>}
 */
const methodDescriptor_BananaService_StreamLogs = new grpc.web.MethodDescriptor(
  '/BananaService/StreamLogs',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.StreamLogsRequest,
  proto.LogEntry,
  /**
   * @param {!proto.StreamLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.LogEntry.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.StreamLogsRequest,
 *   !proto.LogEntry>}
 */
const methodInfo_BananaService_StreamLogs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.LogEntry,
  /**
   * @param {!proto.StreamLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.LogEntry.deserializeBinary
);


/**
 * @param {!proto.StreamLogsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.LogEntry>}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.streamLogs =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/StreamLogs',
      request,
      metadata || {},
      methodDescriptor_BananaService_StreamLogs);
};


/**
 * @param {!proto.StreamLogsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.LogEntry>}
 *     The XHR Node Readable Stream
 */
proto.BananaServicePromiseClient.prototype.streamLogs =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/StreamLogs',
      request,
      metadata || {},
      methodDescriptor_BananaService_StreamLogs);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.UpdateBoardRequest,
 *   !proto.UpdateBoardResponse>}
 */
const methodDescriptor_BananaService_UpdateBoard = new grpc.web.MethodDescriptor(
  '/BananaService/UpdateBoard',
  grpc.web.MethodType.UNARY,
  proto.UpdateBoardRequest,
  proto.UpdateBoardResponse,
  /**
   * @param {!proto.UpdateBoardRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateBoardResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.UpdateBoardRequest,
 *   !proto.UpdateBoardResponse>}
 */
const methodInfo_BananaService_UpdateBoard = new grpc.web.AbstractClientBase.MethodInfo(
  proto.UpdateBoardResponse,
  /**
   * @param {!proto.UpdateBoardRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpdateBoardResponse.deserializeBinary
);


/**
 * @param {!proto.UpdateBoardRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.UpdateBoardResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.UpdateBoardResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.updateBoard =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/BananaService/UpdateBoard',
      request,
      metadata || {},
      methodDescriptor_BananaService_UpdateBoard,
      callback);
};


/**
 * @param {!proto.UpdateBoardRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.UpdateBoardResponse>}
 *     A native promise that resolves to the response
 */
proto.BananaServicePromiseClient.prototype.updateBoard =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/BananaService/UpdateBoard',
      request,
      metadata || {},
      methodDescriptor_BananaService_UpdateBoard);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.DumpRequest,
 *   !proto.DumpResponse>}
 */
const methodDescriptor_BananaService_Dump = new grpc.web.MethodDescriptor(
  '/BananaService/Dump',
  grpc.web.MethodType.UNARY,
  proto.DumpRequest,
  proto.DumpResponse,
  /**
   * @param {!proto.DumpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DumpResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DumpRequest,
 *   !proto.DumpResponse>}
 */
const methodInfo_BananaService_Dump = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DumpResponse,
  /**
   * @param {!proto.DumpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DumpResponse.deserializeBinary
);


/**
 * @param {!proto.DumpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DumpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DumpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.dump =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/BananaService/Dump',
      request,
      metadata || {},
      methodDescriptor_BananaService_Dump,
      callback);
};


/**
 * @param {!proto.DumpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DumpResponse>}
 *     A native promise that resolves to the response
 */
proto.BananaServicePromiseClient.prototype.dump =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/BananaService/Dump',
      request,
      metadata || {},
      methodDescriptor_BananaService_Dump);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.SpectateRequest,
 *   !proto.SpectateUpdate>}
 */
const methodDescriptor_BananaService_Spectate = new grpc.web.MethodDescriptor(
  '/BananaService/Spectate',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.SpectateRequest,
  proto.SpectateUpdate,
  /**
   * @param {!proto.SpectateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SpectateUpdate.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.SpectateRequest,
 *   !proto.SpectateUpdate>}
 */
const methodInfo_BananaService_Spectate = new grpc.web.AbstractClientBase.MethodInfo(
  proto.SpectateUpdate,
  /**
   * @param {!proto.SpectateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SpectateUpdate.deserializeBinary
);


/**
 * @param {!proto.SpectateRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.SpectateUpdate>}
 *     The XHR Node Readable Stream
 */
proto.BananaServiceClient.prototype.spectate =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/Spectate',
      request,
      metadata || {},
      methodDescriptor_BananaService_Spectate);
};


/**
 * @param {!proto.SpectateRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.SpectateUpdate>}
 *     The XHR Node Readable Stream
 */
proto.BananaServicePromiseClient.prototype.spectate =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/BananaService/Spectate',
      request,
      metadata || {},
      methodDescriptor_BananaService_Spectate);
};


module.exports = proto;

