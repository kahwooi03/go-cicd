FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server -a -installsuffix cgo .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

COPY --from=builder /app/. .

EXPOSE 8080

CMD ["./server"]