FROM golang:1.18-alpine

EXPOSE 5000

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

CMD ["go", "run", "gRPC_server/main.go"]