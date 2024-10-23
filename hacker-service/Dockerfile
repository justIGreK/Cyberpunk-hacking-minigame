FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o hacker-service ./cmd/main.go

EXPOSE 8081

CMD ["./hacker-service"]