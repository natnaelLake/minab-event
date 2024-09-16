# Stage 1: Build the Go application
FROM golang:1.22.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a lightweight image with necessary files and certificates
FROM debian:bullseye-slim

# Install CA certificates to handle SSL verification
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Set permissions for the binary
RUN chmod +x ./main

# Expose the port your application will run on
EXPOSE 5050

# Command to run the Go binary
CMD ["./main"]
