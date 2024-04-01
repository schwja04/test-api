FROM golang:1.22.1 as base

FROM base as dev

RUN go install github.com/cosmtrek/air@latest

RUN mkdir /air

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]

