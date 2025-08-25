FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o github-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/github-api .

EXPOSE 8080

CMD ["./github-api"]