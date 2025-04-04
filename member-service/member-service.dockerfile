FROM golang:1.24.1-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o member-service ./cmd/api

EXPOSE 80 81

CMD ["./member-service"]
