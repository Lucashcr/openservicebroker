FROM golang:1.24.5

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build -o logger ./cmd/server/main.go

CMD ["./logger"]
