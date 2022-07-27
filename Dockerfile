FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .

RUN go get
RUN go install
RUN go mod tidy

COPY . .

EXPOSE 5000

CMD ["go", "run", "api/main.go"]