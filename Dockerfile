FROM golang:1.18 AS builder
RUN mkdir app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /app/main ./

CMD ["./main"]