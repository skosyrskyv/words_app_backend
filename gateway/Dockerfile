FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/gateway ./cmd/api/main.go


FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/bin/gateway .
COPY static ./static

EXPOSE 8000

CMD ["./gateway"]
