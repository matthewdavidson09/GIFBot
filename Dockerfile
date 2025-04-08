# --- Build stage ---
    FROM golang:1.21-alpine AS builder
    WORKDIR /app
    
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
    
    COPY . ./
    RUN go build -o gifbot
    
    # --- Runtime stage ---
    FROM alpine:latest
    WORKDIR /app
    
    COPY --from=builder /app/gifbot /app/gifbot
    
    ENTRYPOINT ["/app/gifbot"]
