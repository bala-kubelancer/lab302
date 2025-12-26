# ---------- Stage 1: Build ----------
FROM golang:alpine AS builder

# Enable Go modules
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy Go source
COPY . .

# Build the Go executable
RUN go build -o server .

# ---------- Stage 2: Run ----------
FROM alpine:latest

# Create non-root user for security
RUN adduser -D appuser

WORKDIR /app
COPY --from=builder /app/server .

USER appuser

EXPOSE 8080

# Run the app
CMD ["./server"]
