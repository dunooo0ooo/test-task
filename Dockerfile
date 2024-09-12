FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/server

RUN go build -o /app/server/graphql-server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server/graphql-server .

ENV PORT=8080

EXPOSE 8080

CMD ["/app/graphql-server"]
