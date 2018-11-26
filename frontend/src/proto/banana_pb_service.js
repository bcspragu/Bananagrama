// package: 
// file: banana.proto

var banana_pb = require("./banana_pb");
var grpc = require("grpc-web-client").grpc;

export var BananaService = (function () {
  function BananaService() {}
  BananaService.serviceName = "BananaService";
  return BananaService;
}());

BananaService.NewGame = {
  methodName: "NewGame",
  service: BananaService,
  requestStream: false,
  responseStream: false,
  requestType: banana_pb.NewGameRequest,
  responseType: banana_pb.NewGameResponse
};

BananaService.JoinGame = {
  methodName: "JoinGame",
  service: BananaService,
  requestStream: false,
  responseStream: true,
  requestType: banana_pb.JoinGameRequest,
  responseType: banana_pb.GameUpdate
};

BananaService.Dump = {
  methodName: "Dump",
  service: BananaService,
  requestStream: false,
  responseStream: false,
  requestType: banana_pb.DumpRequest,
  responseType: banana_pb.DumpResponse
};


export function BananaServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BananaServiceClient.prototype.newGame = function newGame(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BananaService.NewGame, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BananaServiceClient.prototype.joinGame = function joinGame(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(BananaService.JoinGame, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.end.forEach(function (handler) {
        handler();
      });
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

BananaServiceClient.prototype.dump = function dump(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BananaService.Dump, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};


