# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Enable static linking by disabling CGO
ENV CGO_ENABLED=0

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /server ./src/main.go

# Stage 2: Create the final image with the binary
FROM scratch

# Copy the compiled Go binary from the builder stage
COPY --from=builder /server /server

# Expose the application port
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["/server"]
