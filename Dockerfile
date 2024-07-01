FROM golang:1.22-bullseye AS builder
LABEL authors="developer"
ENV CGO_ENABLED=1

WORKDIR /app

# Download necessary Go modules
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o devcode-app .

# Stage 2: Run the Go application
FROM debian:bullseye-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/devcode-app .

# Expose port 3030 to the outside world
EXPOSE 3030

# Command to run the executable
CMD ["./devcode-app"]