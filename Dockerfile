FROM golang:1.20 AS build

RUN apt-get update && apt-get install -y libssl-dev pkg-config protobuf-compiler

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o spinoza

FROM alpine:latest AS runtime

RUN apk update && apk add --no-cache openssl opencv

WORKDIR /app

COPY --from=build /app/spinoza /app/spinoza

EXPOSE 28717
CMD ["/app/spinoza"]
