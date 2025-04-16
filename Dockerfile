FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "./cmd/server/main.go"]
