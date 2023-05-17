# Build stage
FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o poker

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/poker .

CMD ["./poker"]
