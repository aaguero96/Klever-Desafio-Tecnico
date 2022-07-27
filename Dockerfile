FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

CMD ["go", "run", "api/main.go"]