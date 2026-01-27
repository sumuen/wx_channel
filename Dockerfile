# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev g++

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=1 is needed for some dependencies like sqlite (modernc.org/sqlite is pure go, but just in case)
# Using CGO_ENABLED=0 since we switched to modernc.org/sqlite which is pure Go
# Using CGO_ENABLED=1 because SunnyNet dependency appears to require it (compilation errors with CGO=0)
RUN CGO_ENABLED=1 GOOS=linux go build -o wx_channel .

# Runtime Stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates (required for HTTPS) and timezone data
RUN apk add --no-cache ca-certificates tzdata

# Copy binary from builder
COPY --from=builder /app/wx_channel .

# Copy web directory for Console UI
COPY --from=builder /app/web ./web

# Copy assets if needed (though they are embedded)
# COPY --from=builder /app/internal/assets ./internal/assets

# Create directories for logs and downloads
RUN mkdir -p logs downloads

# Expose ports
# 2025: Proxy Port
# 2026: Management/Console Port
EXPOSE 2025 2026

# Environment variables
ENV TZ=Asia/Shanghai

# Run the application
CMD ["./wx_channel"]
