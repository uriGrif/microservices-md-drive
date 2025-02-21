FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./bin/server ./cmd/server/main.go && go build -o ./bin/gateway ./cmd/gateway/main.go

FROM alpine:latest AS permissions-service-grpc
WORKDIR /app
COPY --from=builder /app/bin/server .
ENTRYPOINT ["./server"]

FROM alpine:latest AS permissions-service-gateway
WORKDIR /app
COPY --from=builder /app/bin/gateway .
ENTRYPOINT ["./gateway"]

FROM postgres AS permissions-db
COPY ./db-init.sql /docker-entrypoint-initdb.d/