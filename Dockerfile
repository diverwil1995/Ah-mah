# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /scrape

COPY go .mod go.sum ./

RUN go mod download

COPY * ./

RUN CGO_ENABLED=0 GOOS=linux go build

EXPOSE 8080

CMD [ "/scrape" ]