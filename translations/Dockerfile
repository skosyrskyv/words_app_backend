FROM golang:1.26.1-alpine AS builder

WORKDIR /app

# Copy protos module first (required by replace directive in translations/go.mod)
COPY protos/ ./protos/

# Download dependencies separately for better layer caching
COPY translations/go.mod translations/go.sum ./translations/
RUN cd translations && go mod download

# Build
COPY translations/ ./translations/
RUN cd translations && go build -o /app/bin/translations ./cmd/api/main.go


FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/bin/translations .

EXPOSE 8080

CMD ["./translations"]
