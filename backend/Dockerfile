# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies (for SQLite CGO)
RUN apk add --no-cache gcc musl-dev

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application (CGO_ENABLED=1 required for mattn/go-sqlite3)
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/api ./cmd/api

# Run Stage
FROM alpine:latest

WORKDIR /app

# Install sqlite for runtime
RUN apk add --no-cache sqlite

# Copy binary from builder
COPY --from=builder /app/api .

# Expose port and set env
EXPOSE 8080
ENV PORT=8080
ENV DB_PATH=/data/local.db

# Run the binary
CMD ["./api"]
