# Stage 1: Build
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o status-checker ./cmd/server

# Stage 2: Runtime
From alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/status-checker .

COPY configs/services.yaml /app/configs/

EXPOSE 8080

ENV CONFIG_PATH=/app/configs/services.yaml

# Run
CMD ["./status-checker"]
