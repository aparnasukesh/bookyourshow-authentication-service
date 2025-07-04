# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy Go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main cmd/main.go

# Stage 2: Create the final image with minimal size
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Install any required packages for running the app (if needed)
# RUN apk add --no-cache <package_name>.

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy environment variables file
COPY .env .

# Expose the port that the app listens on
EXPOSE 5052

# Command to run the application
CMD ["./main"]
