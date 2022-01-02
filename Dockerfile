FROM golang:1.17-bullseye as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" .

FROM golang:1.17-bullseye

WORKDIR /app

COPY --from=builder /app/new-newt /usr/bin/

ENTRYPOINT ["new-newt"]