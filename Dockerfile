FROM golang:1.23-alpine

WORKDIR /app
COPY . .
RUN go build -o app ./cmd/main.go

EXPOSE 8980

CMD ["./app"]
