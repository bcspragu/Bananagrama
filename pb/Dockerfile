FROM node:alpine

WORKDIR /project
ENV GOPATH=/project/go
ENV PATH="/project/go/bin:${PATH}"

RUN apk update && apk add protobuf curl go git musl-dev
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN curl --compressed -o- -L https://yarnpkg.com/install.sh | sh
RUN yarn config set prefix /project && yarn global add ts-protoc-gen
