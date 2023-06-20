FROM golang:1.20-alpine3.18 AS build

RUN apk update && apk add --no-cache \
    build-base \
    cmake \
    git \
    openssl-dev \
    protobuf-dev

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o spinoza

FROM alpine:3.18 AS runtime

RUN apk update && apk add --no-cache \
    libstdc++ \
    libgcc \
    tini \
    curl

COPY --from=build /app/spinoza /bin/spinoza

EXPOSE 28717
ENTRYPOINT ["tini", "--"]
CMD     ["/bin/spinoza"]
