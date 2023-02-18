FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd

FROM alpine:3.14

RUN apk add --no-cache ca-certificates

WORKDIR /usr/src/app

COPY --from=builder ./ .

EXPOSE 3000

CMD ["./app"]