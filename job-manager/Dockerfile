# Use Go 1.22 for building the Go application
FROM golang:1.22-alpine as builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o job-manager main.go

# Final stage - smaller image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder image
COPY --from=builder /app/job-manager .

# Run the Go application
CMD ["./job-manager"]
