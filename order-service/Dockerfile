FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 8082

CMD [ "./main" ]
