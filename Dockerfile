FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -v cmd/main.go

CMD ["./main"]