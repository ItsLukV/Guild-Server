# Build stage
FROM golang:1.22.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy binary
COPY --from=builder /app/main .

# Copy required directories
COPY --from=builder /app/public ./public
COPY --from=builder /app/templates ./templates

# Install CA certificates (needed for HTTPS requests)
RUN apk add --no-cache ca-certificates

EXPOSE 8080
CMD ["./main"]
