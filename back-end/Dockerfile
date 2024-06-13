FROM golang:1.21.3-alpine3.18 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o web-login-service .

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/web-login-service .
ENV GIN_MODE=release
ENTRYPOINT ["./web-login-service"]
EXPOSE 2996