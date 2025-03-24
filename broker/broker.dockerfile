FROM golang:1.24.1-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o broker ./cmd/api

EXPOSE 80

CMD ["./broker"]
