FROM golang:1.20-alpine AS build

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN apk update \
 && apk add openssl-dev pkgconf protobuf-dev
 
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o spinoza

FROM alpine:3.18 AS runtime

RUN apk update \
 && apk add --no-cache libssl1.1 musl-dev libgcc tini curl

COPY --from=build /app/spinoza /bin/spinoza

EXPOSE 28717
ENTRYPOINT ["tini", "--"]
CMD     /bin/spinoza
