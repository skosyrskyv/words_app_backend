FROM golang:1.26.1-alpine AS builder

WORKDIR /app

# Copy protos module first (required by replace directive in gateway/go.mod)
COPY protos/ ./protos/

COPY gateway/go.mod gateway/go.sum ./gateway/
RUN cd gateway && go mod download

COPY gateway/ ./gateway/

RUN cd gateway && go build -o /app/bin/gateway ./cmd/api/main.go


FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/bin/gateway .
COPY --from=builder /app/gateway/static ./static

EXPOSE 8000

CMD ["./gateway"]
