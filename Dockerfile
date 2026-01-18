FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat-api ./cmd

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/chat-api /usr/local/bin/chat-api
COPY migrations ./migrations

ENV HTTP_ADDR=:8080

CMD ["chat-api"]
