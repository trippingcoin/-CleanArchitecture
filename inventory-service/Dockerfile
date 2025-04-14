FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o inventory-service ./cmd/main.go

EXPOSE 50051

CMD ["/app/inventory-service"]