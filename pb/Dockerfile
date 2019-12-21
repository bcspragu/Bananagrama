FROM node:alpine

WORKDIR /project
ENV GOPATH=/project/go
ENV PATH="/project/go/bin:${PATH}"

RUN apk update && apk add protobuf curl go git musl-dev tar
RUN wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz && \
  tar -xvf go1.13.3.linux-amd64.tar.gz && \
  mv go /usr/local
RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip -o protoc.zip && \
  unzip -qq protoc.zip && \
  cp ./bin/protoc /usr/local/bin/protoc
RUN go get -u github.com/golang/protobuf/protoc-gen-go && \
  cd /project/go/src/github.com/golang/protobuf/protoc-gen-go && \
  git checkout v1.3.2 && \
  go install
RUN curl -sSL https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-linux-x86_64 -o /usr/local/bin/protoc-gen-grpc-web && \
  chmod +x /usr/local/bin/protoc-gen-grpc-web
