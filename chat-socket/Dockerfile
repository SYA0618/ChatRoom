FROM golang:1.21.3-alpine3.18 AS build
WORKDIR /app
COPY ./chat-socket/go.mod ./chat-socket/go.sum ./
RUN go mod download
COPY ./chat-socket .
RUN go build -o chat-service ./connection.go ./server.go

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/chat-service .
ENV GIN_MODE=release
ENTRYPOINT ["./chat-service"]
EXPOSE 8083