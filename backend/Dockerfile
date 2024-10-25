FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cmd/todolist-be/main ./cmd/todolist-be/main.go

CMD ["./cmd/todolist-be/main"]
