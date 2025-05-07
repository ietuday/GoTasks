# Use a more recent Go image that supports Go 1.24
FROM golang:1.24.2-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
