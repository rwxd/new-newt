FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" .

RUN cp /app/new-newt /usr/bin/

ENTRYPOINT ["new-newt"]