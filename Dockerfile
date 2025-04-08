# --- Build stage ---
    FROM golang:1.23-alpine AS builder
    WORKDIR /app
    
    COPY go.mod ./
    RUN go mod download

    COPY . ./
    RUN go build -o gifbot
    
    # --- Runtime stage ---
    FROM alpine:latest
    WORKDIR /app
    
    COPY --from=builder /app/gifbot /app/gifbot
    
    ENTRYPOINT ["/app/gifbot"]
