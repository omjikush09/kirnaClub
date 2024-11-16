FROM golang:1.23.3 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .

RUN go build -o main .

#Server start step
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["./main"]
