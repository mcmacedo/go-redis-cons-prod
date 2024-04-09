FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./producer ./cmd/producer/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/producer /app/

CMD ["/app/producer"]