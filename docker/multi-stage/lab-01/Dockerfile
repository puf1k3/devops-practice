# STAGE 1
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod
COPY main.go

RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /app/server .

# STAGE 2
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/server ./server

USER 10001:10001

EXPOSE 8080

CMD ["./server"]