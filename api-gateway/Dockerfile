# Use the official Golang image for building
FROM golang:1.23 as builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go app
RUN go mod download
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o ./main

# Use a minimal image for running
FROM alpine

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# it needs this certificate to work with auth0
RUN apk add ca-certificates

# Expose the port
EXPOSE 8080

# Command to run the service
CMD ["./main"]