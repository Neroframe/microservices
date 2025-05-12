FROM golang:1.23.6 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the stats binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o main ./cmd/main.go

# ─── Final Stage ───────────────────────────────────────────────────────────────
FROM alpine:latest
WORKDIR /root/

# (Optional) if you ever call HTTPS endpoints, you might need certs:
RUN apk add --no-cache ca-certificates

# Copy the statically-built binary
COPY --from=builder /app/main .

# Expose your gRPC port (inside-container 50051)
EXPOSE 50051

CMD ["./main"]
