# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o bffApp ./cmd/api

RUN chmod +x /app/bffApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/bffApp /app

CMD [ "/app/bffApp" ]