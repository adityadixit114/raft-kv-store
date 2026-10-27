FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o raft-server ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/raft-server .
CMD ["./raft-server"]
