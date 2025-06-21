# Stage 1: Build
FROM golang:1.24.4 AS builder

# Install dependencies required for confluent-kafka-go
RUN apt-get update && apt-get install -y \
    librdkafka-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO must be enabled for confluent-kafka-go
ENV CGO_ENABLED=1
RUN go build -o app .

# Stage 2: Runtime
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y librdkafka1 && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/app /app/app
RUN chmod +x /app/app
CMD ["./app"]


