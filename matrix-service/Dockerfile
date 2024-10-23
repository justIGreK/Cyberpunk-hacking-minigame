FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o matrix-service ./cmd/main.go

EXPOSE 8080

CMD ["./matrix-service"]