FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./main

FROM alpine:latest AS files-service
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]

FROM mysql AS files-db
ADD db-init.sql /docker-entrypoint-initdb.d