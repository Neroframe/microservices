FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 🔥 Key: static binary + permissions
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/main.go

# ✅ Use lightweight secure base image (not Alpine)
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/main /
EXPOSE 8080

ENTRYPOINT ["/main"]
