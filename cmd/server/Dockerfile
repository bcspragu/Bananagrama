FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
ADD ./server /app/server
ADD ./dict.txt /app/dict.txt
CMD ["./server", "--dict", "dict.txt"]

