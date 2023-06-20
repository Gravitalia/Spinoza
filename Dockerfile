FROM golang:1.20 AS build

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN apt-get update && apt-get install -y libssl-dev pkg-config protobuf-compiler
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o spinoza

FROM alpine:3.18 AS runtime

COPY --from=build /app/spinoza /app/spinoza

EXPOSE 28717
CMD [ "/app/spinoza" ]
