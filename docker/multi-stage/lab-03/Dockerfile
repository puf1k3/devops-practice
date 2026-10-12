FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .

COPY main.go .

RUN go build -o server .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/server ./server

COPY assets/banner.txt ./assets/

USER 10001:10001

EXPOSE 8080

CMD ["./server"]
