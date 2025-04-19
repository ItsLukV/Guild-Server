# Build stage
FROM golang:1.22.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# -ldflags="-s -w" removes debug symbols to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Final stage
# Use a minimal alpine image to run the application
FROM alpine:3.19

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy any other necessary files (like configs, templates, etc.)
# COPY --from=builder /app/config.yaml .
# COPY --from=builder /app/templates ./templates

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
